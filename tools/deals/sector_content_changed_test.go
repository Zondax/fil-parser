package deals_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"os"
	"sort"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	miner16 "github.com/filecoin-project/go-state-types/builtin/v16/miner"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors/v2/market"
	parserContext "github.com/zondax/fil-parser/context"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
)

// sectorContentChangedFixture is a real market SectorContentChanged call, taken from the execution
// trace of a ProveCommitSectors3 / ProveReplicaUpdates3 message, together with the deal ids of the
// `deal-activated` events the node reports for that message.
type sectorContentChangedFixture struct {
	Name                string   `json:"name"`
	Network             string   `json:"network"`
	Height              uint64   `json:"height"`
	TxCid               string   `json:"txCid"`
	ParentMethod        uint64   `json:"parentMethod"`
	From                string   `json:"from"`
	Params              string   `json:"params"`
	Return              string   `json:"return"`
	DealActivatedEvents []uint64 `json:"dealActivatedEvents"`
}

func loadSectorContentChangedFixtures(t *testing.T) []sectorContentChangedFixture {
	raw, err := os.ReadFile("testdata/sector_content_changed_fixtures.json")
	require.NoError(t, err)
	var fixtures []sectorContentChangedFixture
	require.NoError(t, json.Unmarshal(raw, &fixtures))
	require.NotEmpty(t, fixtures)
	return fixtures
}

// sectorContentChangedTx parses raw SectorContentChanged params and return with the market actor
// parser, as the transaction parser does, and wraps the metadata in a transaction.
func sectorContentChangedTx(t *testing.T, network string, height uint64, txCid, from string, rawParams, rawReturn []byte) *types.Transaction {
	//#nosec G115
	metadata, err := market.New(logger.NewDevelopmentLogger()).SectorContentChanged(network, int64(height), rawParams, rawReturn)
	require.NoError(t, err)
	metadata[parser.MethodNumKey] = "2034386435"
	jsonMetadata, err := json.Marshal(metadata)
	require.NoError(t, err)

	return &types.Transaction{
		TxBasicBlockData: types.TxBasicBlockData{BasicBlockData: types.BasicBlockData{Height: height}},
		Id:               "tx-" + txCid,
		TxCid:            txCid,
		TxType:           parser.MethodSectorContentChanged,
		TxFrom:           from,
		TxTo:             txTo,
		TxMetadata:       string(jsonMetadata),
		Status:           tools.GetExitCodeOK(),
		SubcallStatus:    tools.GetExitCodeOK(),
	}
}

func dealIDs(activations []*types.DealsActivations) []uint64 {
	ids := make([]uint64, 0, len(activations))
	for _, activation := range activations {
		ids = append(ids, activation.DealID)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
}

func TestSectorContentChanged_RealFixtures(t *testing.T) {
	for _, fixture := range loadSectorContentChangedFixtures(t) {
		t.Run(fixture.Name, func(t *testing.T) {
			eg, node := setupTest(t, fixture.Network)
			ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})

			rawParams, err := base64.StdEncoding.DecodeString(fixture.Params)
			require.NoError(t, err)
			rawReturn, err := base64.StdEncoding.DecodeString(fixture.Return)
			require.NoError(t, err)

			// the reference decoding of the call, independent of the metadata path
			var sectors []miner16.SectorChanges
			reader := cbg.NewCborReader(bytes.NewReader(rawParams))
			maj, n, err := reader.ReadHeader()
			require.NoError(t, err)
			require.Equal(t, byte(cbg.MajArray), maj)
			for i := uint64(0); i < n; i++ {
				var sector miner16.SectorChanges
				require.NoError(t, sector.UnmarshalCBOR(reader))
				sectors = append(sectors, sector)
			}
			expiry := map[uint64]int64{}
			sectorOf := map[uint64]uint64{}
			size := map[uint64]uint64{}
			for _, sector := range sectors {
				for _, piece := range sector.Added {
					maj, dealID, err := cbg.CborReadHeader(bytes.NewReader(piece.Payload))
					require.NoError(t, err)
					require.Equal(t, byte(cbg.MajUnsignedInt), maj)
					expiry[dealID] = int64(sector.MinimumCommitmentEpoch)
					sectorOf[dealID] = uint64(sector.Sector)
					size[dealID] = uint64(piece.Size)
				}
			}

			tx := sectorContentChangedTx(t, fixture.Network, fixture.Height, fixture.TxCid, fixture.From, rawParams, rawReturn)
			events, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{tx}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)

			// cross-check: the activations are exactly the deal-activated events of the message
			require.Equal(t, fixture.DealActivatedEvents, dealIDs(events.DealsActivations))
			require.Len(t, events.DealsSpaceInfo, len(fixture.DealActivatedEvents))

			ids := map[string]bool{}
			for _, activation := range events.DealsActivations {
				require.False(t, ids[activation.ID], "duplicated id")
				ids[activation.ID] = true
				require.Equal(t, fixture.From, activation.ActorAddress)
				require.Equal(t, fixture.TxCid, activation.TxCid)
				require.Equal(t, fixture.Height, activation.Height)
				require.Equal(t, parser.MethodSectorContentChanged, activation.ActionType)
				require.Equal(t, expiry[activation.DealID], activation.SectorExpiry)
				require.NotNil(t, activation.SectorNumber)
				require.Equal(t, sectorOf[activation.DealID], *activation.SectorNumber)
			}
			for _, space := range events.DealsSpaceInfo {
				require.Equal(t, []uint64{space.DealID}, space.GroupDealIDs)
				require.Equal(t, new(big.Int).SetUint64(size[space.DealID]), space.NonVerifiedDealSpace)
				require.Equal(t, big.NewInt(0), space.VerifiedDealSpace)
				require.False(t, space.SpaceAsWeight)
			}
		})
	}
}

// encodeSectorContentChanged builds raw SectorContentChanged params and return.
func encodeSectorContentChanged(t *testing.T, sectors []miner16.SectorChanges, accepted [][]bool) ([]byte, []byte) {
	params := new(bytes.Buffer)
	require.NoError(t, cbg.WriteMajorTypeHeader(params, cbg.MajArray, uint64(len(sectors))))
	for i := range sectors {
		require.NoError(t, sectors[i].MarshalCBOR(params))
	}

	ret := new(bytes.Buffer)
	require.NoError(t, cbg.WriteMajorTypeHeader(ret, cbg.MajArray, uint64(len(accepted))))
	for _, pieces := range accepted {
		require.NoError(t, cbg.WriteMajorTypeHeader(ret, cbg.MajArray, uint64(len(pieces))))
		for _, flag := range pieces {
			require.NoError(t, cbg.WriteBool(ret, flag))
		}
	}
	return params.Bytes(), ret.Bytes()
}

func dealPayload(t *testing.T, dealID uint64) []byte {
	buf := new(bytes.Buffer)
	require.NoError(t, cbg.WriteMajorTypeHeader(buf, cbg.MajUnsignedInt, dealID))
	return buf.Bytes()
}

func TestSectorContentChanged_Crafted(t *testing.T) {
	const (
		height  = 6559222
		network = "mainnet"
		miner   = "f03046733"
	)
	pieceCid, err := cid.Parse("baga6ea4seaqpol3htxs5b2iwsvzqssuuzvld6a3g445gn6axaumw4yzuvkmvgei")
	require.NoError(t, err)
	piece := func(payload []byte) miner16.PieceChange {
		return miner16.PieceChange{Data: pieceCid, Size: abi.PaddedPieceSize(34359738368), Payload: payload}
	}
	// a payload that is not a DealID: a CBOR byte string
	notADeal := []byte{0x43, 0x01, 0x02, 0x03}
	// a DealID followed by trailing bytes is not a valid DealID either
	trailing := append(dealPayload(t, 77), 0x00)

	tests := []struct {
		name          string
		sectors       []miner16.SectorChanges
		accepted      [][]bool
		parentFailed  bool
		expectedDeals []uint64
		expectedSize  int
	}{
		{
			name: "rejected piece is skipped, accepted sibling is kept",
			sectors: []miner16.SectorChanges{{Sector: 1, MinimumCommitmentEpoch: 7000000, Added: []miner16.PieceChange{
				piece(dealPayload(t, 10)), piece(dealPayload(t, 11)),
			}}},
			accepted:      [][]bool{{false, true}},
			expectedDeals: []uint64{11},
		},
		{
			name: "non-deal payloads are skipped whatever the flag",
			sectors: []miner16.SectorChanges{{Sector: 1, MinimumCommitmentEpoch: 7000000, Added: []miner16.PieceChange{
				piece(notADeal), piece(notADeal), piece(trailing), piece(dealPayload(t, 12)),
			}}},
			accepted:      [][]bool{{false, true, true, true}},
			expectedDeals: []uint64{12},
		},
		{
			name: "every sector uses its own accepted flags",
			sectors: []miner16.SectorChanges{
				{Sector: 1, MinimumCommitmentEpoch: 7000000, Added: []miner16.PieceChange{piece(dealPayload(t, 20))}},
				{Sector: 2, MinimumCommitmentEpoch: 7000001, Added: []miner16.PieceChange{piece(dealPayload(t, 21)), piece(dealPayload(t, 22))}},
				{Sector: 3, MinimumCommitmentEpoch: 7000002, Added: []miner16.PieceChange{}},
				{Sector: 4, MinimumCommitmentEpoch: 7000003, Added: []miner16.PieceChange{piece(dealPayload(t, 23))}},
			},
			accepted:      [][]bool{{false}, {true, false}, {}, {true}},
			expectedDeals: []uint64{21, 23},
		},
		{
			name: "all rejected",
			sectors: []miner16.SectorChanges{{Sector: 1, MinimumCommitmentEpoch: 7000000, Added: []miner16.PieceChange{
				piece(dealPayload(t, 30)),
			}}},
			accepted:      [][]bool{{false}},
			expectedDeals: []uint64{},
		},
		{
			// e.g. the miner aborts the message because require_notification_success is set and
			// another notification was rejected: the market call returned ok but was reverted
			name: "accepted piece under a failed parent is not emitted",
			sectors: []miner16.SectorChanges{{Sector: 1, MinimumCommitmentEpoch: 7000000, Added: []miner16.PieceChange{
				piece(dealPayload(t, 40)),
			}}},
			accepted:      [][]bool{{true}},
			parentFailed:  true,
			expectedDeals: []uint64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eg, node := setupTest(t, network)
			ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})

			rawParams, rawReturn := encodeSectorContentChanged(t, tt.sectors, tt.accepted)
			tx := sectorContentChangedTx(t, network, height, txCid, miner, rawParams, rawReturn)
			if tt.parentFailed {
				tx.Status = tools.GetExitCodeStatus("", exitcode.ErrForbidden)
			}

			events, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{tx}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.Equal(t, tt.expectedDeals, dealIDs(events.DealsActivations))
			require.Len(t, events.DealsSpaceInfo, len(tt.expectedDeals))
			if tt.parentFailed {
				require.Empty(t, events.DealsMessages)
			} else {
				require.Len(t, events.DealsMessages, 1)
			}

			minimumCommitment := map[uint64]int64{}
			sectorOf := map[uint64]uint64{}
			for _, sector := range tt.sectors {
				for _, p := range sector.Added {
					if maj, dealID, err := cbg.CborReadHeader(bytes.NewReader(p.Payload)); err == nil && maj == cbg.MajUnsignedInt {
						minimumCommitment[dealID] = int64(sector.MinimumCommitmentEpoch)
						sectorOf[dealID] = uint64(sector.Sector)
					}
				}
			}
			for _, activation := range events.DealsActivations {
				require.Equal(t, minimumCommitment[activation.DealID], activation.SectorExpiry)
				require.NotNil(t, activation.SectorNumber)
				require.Equal(t, sectorOf[activation.DealID], *activation.SectorNumber)
			}
		})
	}
}

func TestSectorContentChanged_ReturnShapes(t *testing.T) {
	eg, node := setupTest(t, "mainnet")
	ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})

	params := func(sectors int) string {
		out := "["
		for i := 0; i < sectors; i++ {
			if i > 0 {
				out += ","
			}
			// payload 0x1864 is the CBOR encoding of DealID 100
			out += `{"Sector":1,"MinimumCommitmentEpoch":7000000,"Added":[{"Data":{"/":"baga6ea4seaqpol3htxs5b2iwsvzqssuuzvld6a3g445gn6axaumw4yzuvkmvgei"},"Size":2048,"Payload":"GGQ="}]}`
		}
		return out + "]"
	}
	tx := func(metadata string) *types.Transaction {
		return &types.Transaction{
			TxBasicBlockData: types.TxBasicBlockData{BasicBlockData: types.BasicBlockData{Height: 6559222}},
			TxCid:            txCid,
			TxType:           parser.MethodSectorContentChanged,
			TxFrom:           "f03046733",
			TxTo:             txTo,
			TxMetadata:       metadata,
			Status:           tools.GetExitCodeOK(),
			SubcallStatus:    tools.GetExitCodeOK(),
		}
	}

	t.Run("legacy flat return with one sector", func(t *testing.T) {
		events, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{tx(`{"MethodNum":"2034386435","Params":` + params(1) + `,"Return":[true]}`)}, tipsetCid, filTypes.EmptyTSK)
		require.NoError(t, err)
		require.Equal(t, []uint64{100}, dealIDs(events.DealsActivations))
	})
	t.Run("legacy flat return with several sectors is ambiguous", func(t *testing.T) {
		_, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{tx(`{"MethodNum":"2034386435","Params":` + params(2) + `,"Return":[true]}`)}, tipsetCid, filTypes.EmptyTSK)
		require.Error(t, err)
	})
	t.Run("return with fewer sectors than params", func(t *testing.T) {
		_, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{tx(`{"MethodNum":"2034386435","Params":` + params(2) + `,"Return":[[true]]}`)}, tipsetCid, filTypes.EmptyTSK)
		require.Error(t, err)
	})
	t.Run("return with fewer pieces than params", func(t *testing.T) {
		_, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{tx(`{"MethodNum":"2034386435","Params":` + params(1) + `,"Return":[[]]}`)}, tipsetCid, filTypes.EmptyTSK)
		require.Error(t, err)
	})
}

// BatchActivateDeals is the name of market method 6 from actors v12 (NV21) on; it is still called by
// ProveCommitAggregate and ProveReplicaUpdates until actors v16, alongside SectorContentChanged.
func TestBatchActivateDealsIsProcessed(t *testing.T) {
	eg, node := setupTest(t, "mainnet")
	ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})

	metadata := `{"MethodNum":"6","Params":{"Sectors":[{"SectorNumber":5926,"SectorType":8,"SectorExpiry":5388210,"DealIDs":[79173041]},{"SectorNumber":5839,"SectorType":8,"SectorExpiry":5388210,"DealIDs":[79174788]}],"ComputeCID":false},"Return":{"ActivationResults":{"SuccessCount":1,"FailCodes":[{"Idx":0,"Code":16}]},"Activations":[{"NonVerifiedDealSpace":"0","VerifiedInfos":[{"Client":2026357,"AllocationId":61463679,"Data":{"/":"baga6ea4seaqpol3htxs5b2iwsvzqssuuzvld6a3g445gn6axaumw4yzuvkmvgei"},"Size":34359738368}],"UnsealedCid":{}}]}}`
	events, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{{
		TxBasicBlockData: types.TxBasicBlockData{BasicBlockData: types.BasicBlockData{Height: 3857651}},
		TxCid:            txCid,
		TxType:           parser.MethodBatchActivateDeals,
		TxFrom:           txFrom,
		TxTo:             txTo,
		TxMetadata:       metadata,
		Status:           tools.GetExitCodeOK(),
		SubcallStatus:    tools.GetExitCodeOK(),
	}}, tipsetCid, filTypes.EmptyTSK)
	require.NoError(t, err)
	require.Len(t, events.DealsMessages, 1)
	require.Equal(t, []uint64{79174788}, dealIDs(events.DealsActivations))
	require.NotNil(t, events.DealsActivations[0].SectorNumber)
	require.Equal(t, uint64(5839), *events.DealsActivations[0].SectorNumber)
	require.Equal(t, int64(5388210), events.DealsActivations[0].SectorExpiry)
	require.Len(t, events.DealsSpaceInfo, 1)
	require.Equal(t, big.NewInt(34359738368), events.DealsSpaceInfo[0].VerifiedDealSpace)
}

// ActivateDeals params before actors v12 carry no sector number.
func TestActivateDealsWithoutSectorNumber(t *testing.T) {
	eg, node := setupTest(t, "mainnet")
	ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})

	events, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{{
		TxBasicBlockData: types.TxBasicBlockData{BasicBlockData: types.BasicBlockData{Height: 2383680}},
		TxCid:            txCid,
		TxType:           parser.MethodActivateDeals,
		TxFrom:           txFrom,
		TxTo:             txTo,
		TxMetadata:       `{"MethodNum":"6","Params":{"DealIDs":[17507589],"SectorExpiry":3937912}}`,
		Status:           tools.GetExitCodeOK(),
		SubcallStatus:    tools.GetExitCodeOK(),
	}}, tipsetCid, filTypes.EmptyTSK)
	require.NoError(t, err)
	require.Equal(t, []uint64{17507589}, dealIDs(events.DealsActivations))
	require.Nil(t, events.DealsActivations[0].SectorNumber)
}
