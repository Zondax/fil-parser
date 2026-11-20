package deals_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	v12Market "github.com/filecoin-project/go-state-types/builtin/v12/market"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/require"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors/v2/market"
	parserContext "github.com/zondax/fil-parser/context"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
)

// Last actors v12 and first actors v13 heights (NV22).
const (
	mainnetLastV12      = 3855359
	mainnetFirstV13     = 3855360
	calibrationLastV12  = 1427973
	calibrationFirstV13 = 1427974
)

// dealIDsTerminateParams encodes {epoch, deal_ids}, the params up to actors v12.
func dealIDsTerminateParams(t *testing.T, epoch int64, dealIDs []uint64) []byte {
	ids := make([]abi.DealID, 0, len(dealIDs))
	for _, id := range dealIDs {
		ids = append(ids, abi.DealID(id))
	}
	var buf bytes.Buffer
	require.NoError(t, (&v12Market.OnMinerSectorsTerminateParams{Epoch: abi.ChainEpoch(epoch), DealIDs: ids}).MarshalCBOR(&buf))
	return buf.Bytes()
}

// sectorsTerminateParams encodes {epoch, sectors}, the params from actors v13 on
// (builtin-actors actors/market/src/types.rs, a tuple of an epoch and a BitField).
func sectorsTerminateParams(t *testing.T, epoch int64, sectors bitfield.BitField) []byte {
	var buf bytes.Buffer
	w := cbg.NewCborWriter(&buf)
	require.NoError(t, w.WriteMajorTypeHeader(cbg.MajArray, 2))
	require.NoError(t, w.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(epoch)))
	require.NoError(t, sectors.MarshalCBOR(w))
	return buf.Bytes()
}

// terminateTx parses raw OnMinerSectorsTerminate params with the market actor parser, as the
// transaction parser does, and wraps the metadata in a successful sub-call from the miner.
// When the actor parser cannot decode the params, the metadata is the raw-only fallback the
// transaction parser stores (parser/v2/parser.go).
func terminateTx(t *testing.T, network string, height uint64, id string, rawParams []byte) *types.Transaction {
	//#nosec G115
	metadata, err := market.New(logger.NewDevelopmentLogger()).OnMinerSectorsTerminateExported(network, int64(height), rawParams)
	if err != nil {
		metadata = map[string]interface{}{parser.ParamsRawKey: rawParams}
	}
	metadata[parser.MethodNumKey] = "7"
	jsonMetadata, err := json.Marshal(metadata)
	require.NoError(t, err)
	tx := okTx(height, id, parser.MethodOnMinerSectorsTerminate, string(jsonMetadata))
	tx.TxCid = fmt.Sprintf("%s-%s", txCid, id)
	return tx
}

func generate(t *testing.T, network string, txs ...*types.Transaction) *types.DealsEvents {
	eg, node := setupTest(t, network)
	ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})
	events, err := eg.GenerateDealsEvents(ctx, txs, tipsetCid, filTypes.EmptyTSK)
	require.NoError(t, err)
	return events
}

type terminationRow struct {
	deal   *uint64
	sector *uint64
	epoch  int64
}

func terminationRows(events *types.DealsEvents) []terminationRow {
	rows := make([]terminationRow, 0, len(events.DealsTerminations))
	for _, r := range events.DealsTerminations {
		rows = append(rows, terminationRow{deal: r.DealID, sector: r.SectorNumber, epoch: r.TerminationEpoch})
	}
	return rows
}

func terminatedDeals(t *testing.T, events *types.DealsEvents) []uint64 {
	out := []uint64{}
	for _, r := range events.DealsTerminations {
		require.NotNil(t, r.DealID)
		require.Nil(t, r.SectorNumber)
		out = append(out, *r.DealID)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

func terminatedSectorNumbers(t *testing.T, events *types.DealsEvents) []uint64 {
	out := []uint64{}
	for _, r := range events.DealsTerminations {
		require.NotNil(t, r.SectorNumber)
		require.Nil(t, r.DealID)
		out = append(out, *r.SectorNumber)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

func requireUniqueTerminationIDs(t *testing.T, events *types.DealsEvents) {
	ids := map[string]bool{}
	for _, r := range events.DealsTerminations {
		require.NotEmpty(t, r.ID)
		require.Falsef(t, ids[r.ID], "duplicated termination id %s", r.ID)
		ids[r.ID] = true
	}
}

// Up to actors v12 the params are {epoch, deal_ids}: one row per deal id, no sector number, and
// termination_epoch is the epoch in the params (the epoch the sectors were terminated at).
func TestOnMinerSectorsTerminate_DealIDs(t *testing.T) {
	tests := []struct {
		network string
		height  uint64
	}{
		{tools.MainnetNetwork, 100000},         // actors v0 types
		{tools.MainnetNetwork, 1000000},        // actors v3/v4 types
		{tools.MainnetNetwork, 3000000},        // actors v11 types
		{tools.MainnetNetwork, mainnetLastV12}, // last actors v12 height
		{tools.CalibrationNetwork, calibrationLastV12},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %d", tt.network, tt.height), func(t *testing.T) {
			epoch := int64(tt.height) - 30 //nolint:gosec
			tx := terminateTx(t, tt.network, tt.height, "a", dealIDsTerminateParams(t, epoch, []uint64{7, 3, 1 << 40}))
			events := generate(t, tt.network, tx)

			require.Len(t, events.DealsMessages, 1)
			require.Equal(t, parser.MethodOnMinerSectorsTerminate, events.DealsMessages[0].ActionType)
			require.Equal(t, []uint64{3, 7, 1 << 40}, terminatedDeals(t, events))
			for _, r := range events.DealsTerminations {
				require.Equal(t, epoch, r.TerminationEpoch)
				require.Equal(t, tx.TxFrom, r.ActorAddress)
				require.Equal(t, tx.TxCid, r.TxCid)
				require.Equal(t, tt.height, r.Height)
				require.Equal(t, parser.MethodOnMinerSectorsTerminate, r.ActionType)
			}
			requireUniqueTerminationIDs(t, events)
			// terminations never produce activations or space rows
			require.Empty(t, events.DealsActivations)
			require.Empty(t, events.DealsSpaceInfo)
		})
	}
}

// From actors v13 the params are {epoch, sectors}: one row per sector, no deal id.
func TestOnMinerSectorsTerminate_Sectors(t *testing.T) {
	sectors := bitfield.NewFromSet([]uint64{0, 5, 6, 7, 1000, 1 << 40})
	tests := []struct {
		network string
		height  uint64
	}{
		{tools.MainnetNetwork, mainnetFirstV13},
		{tools.MainnetNetwork, 5000000},
		{tools.MainnetNetwork, 6539228},
		{tools.CalibrationNetwork, calibrationFirstV13},
		{tools.CalibrationNetwork, 4100979},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %d", tt.network, tt.height), func(t *testing.T) {
			tx := terminateTx(t, tt.network, tt.height, "a", sectorsTerminateParams(t, int64(tt.height), sectors)) //nolint:gosec
			events := generate(t, tt.network, tx)

			require.Len(t, events.DealsMessages, 1)
			// sector 0 is a valid sector number
			require.Equal(t, []uint64{0, 5, 6, 7, 1000, 1 << 40}, terminatedSectorNumbers(t, events))
			for _, r := range events.DealsTerminations {
				require.Equal(t, int64(tt.height), r.TerminationEpoch) //nolint:gosec
				require.Equal(t, tx.TxFrom, r.ActorAddress)
			}
			requireUniqueTerminationIDs(t, events)
		})
	}
}

// A real call: the miner f03644142 -> market OnMinerSectorsTerminate sub-call of the TerminateSectors
// message bafy2bzacea6ppodj4hxgwqru4fhplxxh4udycxt2f4phfumrtjxqkjqlbtqdk at mainnet 6,539,228.
func TestOnMinerSectorsTerminate_RealV13Params(t *testing.T) {
	raw, err := base64.StdEncoding.DecodeString("ghoAY8fcREDYXSA=")
	require.NoError(t, err)

	// the sectors, decoded independently from the CBOR
	var decoded struct {
		Epoch   abi.ChainEpoch
		Sectors bitfield.BitField
	}
	r := cbg.NewCborReader(bytes.NewReader(raw))
	_, _, err = r.ReadHeader()
	require.NoError(t, err)
	_, epoch, err := r.ReadHeader()
	require.NoError(t, err)
	decoded.Epoch = abi.ChainEpoch(epoch) //nolint:gosec
	require.NoError(t, decoded.Sectors.UnmarshalCBOR(r))
	want, err := decoded.Sectors.All(1 << 20)
	require.NoError(t, err)
	require.NotEmpty(t, want)

	tx := terminateTx(t, tools.MainnetNetwork, 6539228, "a", raw)
	tx.TxFrom = "f03644142"
	events := generate(t, tools.MainnetNetwork, tx)
	require.Equal(t, want, terminatedSectorNumbers(t, events))
	for _, row := range events.DealsTerminations {
		require.Equal(t, int64(6539228), row.TerminationEpoch)
		require.Equal(t, int64(decoded.Epoch), row.TerminationEpoch)
		require.Equal(t, "f03644142", row.ActorAddress)
	}
}

// Across the v12 -> v13 boundary each height uses its own shape. Params of the other shape do not
// decode with the height's type (the market would have rejected them too); the tx is skipped,
// keeps its message row, and does not affect the other txs of the height.
func TestOnMinerSectorsTerminate_UpgradeBoundary(t *testing.T) {
	sectors := bitfield.NewFromSet([]uint64{11, 12})
	for _, tt := range []struct {
		network           string
		lastV12, firstV13 uint64
	}{
		{tools.MainnetNetwork, mainnetLastV12, mainnetFirstV13},
		{tools.CalibrationNetwork, calibrationLastV12, calibrationFirstV13},
	} {
		t.Run(tt.network, func(t *testing.T) {
			// right shape on each side
			events := generate(t, tt.network,
				terminateTx(t, tt.network, tt.lastV12, "a", dealIDsTerminateParams(t, 100, []uint64{42})),
			)
			require.Equal(t, []uint64{42}, terminatedDeals(t, events))
			events = generate(t, tt.network,
				terminateTx(t, tt.network, tt.firstV13, "a", sectorsTerminateParams(t, int64(tt.firstV13), sectors)), //nolint:gosec
			)
			require.Equal(t, []uint64{11, 12}, terminatedSectorNumbers(t, events))

			// a bitfield at the last v12 height does not decode as {epoch, deal_ids}: raw-only metadata
			events = generate(t, tt.network,
				terminateTx(t, tt.network, tt.lastV12, "a", sectorsTerminateParams(t, int64(tt.lastV12), sectors)), //nolint:gosec
				terminateTx(t, tt.network, tt.lastV12, "b", dealIDsTerminateParams(t, 100, []uint64{43})),
			)
			require.Len(t, events.DealsMessages, 2)
			require.Equal(t, []uint64{43}, terminatedDeals(t, events))

			// deal ids at the first v13 height decode as an array (`SectorNumbers`), which no v13
			// market accepts: never read as sectors
			events = generate(t, tt.network,
				terminateTx(t, tt.network, tt.firstV13, "a", dealIDsTerminateParams(t, 100, []uint64{44})),
				terminateTx(t, tt.network, tt.firstV13, "b", sectorsTerminateParams(t, int64(tt.firstV13), sectors)), //nolint:gosec
			)
			require.Len(t, events.DealsMessages, 2)
			require.Equal(t, []uint64{11, 12}, terminatedSectorNumbers(t, events))
		})
	}
}

// Unusable metadata never fails the height: the tx's terminations are skipped, its message row is
// kept, and the other txs are processed.
func TestOnMinerSectorsTerminate_UnusableMetadata(t *testing.T) {
	good := terminateTx(t, tools.MainnetNetwork, 6539228, "good", sectorsTerminateParams(t, 6539228, bitfield.NewFromSet([]uint64{9})))
	for name, metadata := range map[string]string{
		"raw only":              `{"MethodNum":"7","ParamsRaw":"gg=="}`,
		"no params":             `{"MethodNum":"7"}`,
		"params not a map":      `{"MethodNum":"7","Params":[1,2]}`,
		"no epoch":              `{"MethodNum":"7","Params":{"DealIDs":[1]}}`,
		"negative deal id":      `{"MethodNum":"7","Params":{"Epoch":1,"DealIDs":[-1]}}`,
		"fractional deal id":    `{"MethodNum":"7","Params":{"Epoch":1,"DealIDs":[1.5]}}`,
		"deal ids not an array": `{"MethodNum":"7","Params":{"Epoch":1,"DealIDs":{"a":1}}}`,
		"no sectors":            `{"MethodNum":"7","Params":{"Epoch":1}}`,
		"null bitfield":         `{"MethodNum":"7","Params":{"Epoch":1,"SectorBitField":null,"SectorNumbers":null}}`,
		"bad bitfield":          `{"MethodNum":"7","Params":{"Epoch":1,"SectorBitField":[1,0,2]}}`,
		"huge bitfield":         `{"MethodNum":"7","Params":{"Epoch":1,"SectorBitField":[0,1099511627776]}}`,
		"sector numbers array":  `{"MethodNum":"7","Params":{"Epoch":1,"SectorBitField":[0],"SectorNumbers":[5]}}`,
		"not json":              `{`,
	} {
		t.Run(name, func(t *testing.T) {
			bad := okTx(6539228, "bad", parser.MethodOnMinerSectorsTerminate, metadata)
			events := generate(t, tools.MainnetNetwork, bad, good)
			// the message row is built from the raw metadata and kept
			require.Len(t, events.DealsMessages, 2)
			require.Equal(t, []uint64{9}, terminatedSectorNumbers(t, events))
		})
	}
}

func TestOnMinerSectorsTerminate_EmptyParams(t *testing.T) {
	// null deal ids (empty list) and an empty bitfield: no rows, no error
	events := generate(t, tools.MainnetNetwork,
		okTx(3000000, "a", parser.MethodOnMinerSectorsTerminate, `{"MethodNum":"7","Params":{"Epoch":1,"DealIDs":null}}`),
		okTx(6539228, "b", parser.MethodOnMinerSectorsTerminate, `{"MethodNum":"7","Params":{"Epoch":1,"SectorBitField":[0],"SectorNumbers":null}}`),
	)
	require.Len(t, events.DealsMessages, 2)
	require.Empty(t, events.DealsTerminations)
}

// Only successful calls count: a failed call, or a successful call under a reverted parent (the
// cron path swallows the market's error, so the parent succeeds while the sub-call fails), emits
// nothing.
func TestOnMinerSectorsTerminate_OnlySuccessfulCalls(t *testing.T) {
	raw := sectorsTerminateParams(t, 6539228, bitfield.NewFromSet([]uint64{1}))
	failedCall := terminateTx(t, tools.MainnetNetwork, 6539228, "a", raw)
	failedCall.SubcallStatus = "SysErrOutOfGas"
	revertedParent := terminateTx(t, tools.MainnetNetwork, 6539228, "b", raw)
	revertedParent.Status = "ErrIllegalState"
	events := generate(t, tools.MainnetNetwork, failedCall, revertedParent)
	require.Empty(t, events.DealsMessages)
	require.Empty(t, events.DealsTerminations)
}

// Deal ids and sector numbers above 2^53 keep their exact value (float64 would round them).
func TestOnMinerSectorsTerminate_Precision(t *testing.T) {
	big := uint64(1<<53 + 1)
	events := generate(t, tools.MainnetNetwork,
		terminateTx(t, tools.MainnetNetwork, 3000000, "a", dealIDsTerminateParams(t, 2999999, []uint64{big, math.MaxInt64})),
	)
	require.Equal(t, []uint64{big, math.MaxInt64}, terminatedDeals(t, events))

	events = generate(t, tools.MainnetNetwork,
		terminateTx(t, tools.MainnetNetwork, 6539228, "a", sectorsTerminateParams(t, 6539228, bitfield.NewFromSet([]uint64{big, math.MaxInt64}))),
	)
	require.Equal(t, []uint64{big, math.MaxInt64}, terminatedSectorNumbers(t, events))
}

// Ids are deterministic, and the same deal id or sector number in two calls of the same tx (the v12
// miner sends deal ids in chunks of 8192, one call per termination epoch) gets distinct ids.
func TestOnMinerSectorsTerminate_Ids(t *testing.T) {
	raw := dealIDsTerminateParams(t, 100, []uint64{1, 2})
	first := terminateTx(t, tools.MainnetNetwork, 3000000, "a", raw)
	second := terminateTx(t, tools.MainnetNetwork, 3000000, "b", raw)
	second.TxCid = first.TxCid

	events := generate(t, tools.MainnetNetwork, first, second)
	require.Len(t, events.DealsTerminations, 4)
	requireUniqueTerminationIDs(t, events)

	again := generate(t, tools.MainnetNetwork, first, second)
	for i := range events.DealsTerminations {
		require.Equal(t, events.DealsTerminations[i].ID, again.DealsTerminations[i].ID)
	}
	require.Equal(t, terminationRows(events), terminationRows(again))
}
