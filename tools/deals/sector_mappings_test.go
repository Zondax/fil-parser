package deals_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/api"
	filApiTypes "github.com/filecoin-project/lotus/api/types"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	cid "github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	minerActor "github.com/zondax/fil-parser/actors/v2/miner"
	parserContext "github.com/zondax/fil-parser/context"
	filMetrics "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/deals"
	"github.com/zondax/fil-parser/tools/mocks"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

const minerAddr = "f01919423"

func identityCid(t *testing.T, data string) cid.Cid {
	c, err := cid.Prefix{Version: 1, Codec: cid.Raw, MhType: multihash.IDENTITY, MhLength: -1}.Sum([]byte(data))
	require.NoError(t, err)
	return c
}

// setupMinerTest is setupTest where f05 resolves to the market actor and every other address to a
// miner, so the generator sees both the miner's own methods and its calls to the market.
func setupMinerTest(t *testing.T, network string) (deals.EventGenerator, api.FullNode) {
	log := logger.NewDevelopmentLogger()
	metrics := filMetrics.NewMetricsClient(metrics2.NewNoopMetrics())
	marketCid := identityCid(t, "fil/1/storagemarket")
	minerCid := identityCid(t, "fil/1/storageminer")
	market, err := address.NewIDAddress(5)
	require.NoError(t, err)

	node := &mocks.FullNode{}
	node.On("StateNetworkName", mock.Anything).Return(dtypes.NetworkName(network), nil)
	node.On("StateNetworkVersion", mock.Anything, mock.Anything).Return(filApiTypes.NetworkVersion(16), nil)
	node.On("StateActorCodeCIDs", mock.Anything, mock.Anything).Return(map[string]cid.Cid{
		manifest.MarketKey: marketCid,
		manifest.MinerKey:  minerCid,
	}, nil)

	isMarket := func(a address.Address) bool { return a == market }
	cache := &mocks.IActorsCache{}
	cache.On("StoreAddressInfo", mock.Anything).Return(nil)
	cache.On("GetActorCode", mock.Anything, mock.MatchedBy(isMarket), mock.Anything, mock.Anything, mock.Anything).Return(marketCid.String(), nil)
	cache.On("GetActorCode", mock.Anything, mock.MatchedBy(func(a address.Address) bool { return !isMarket(a) }), mock.Anything, mock.Anything, mock.Anything).Return(minerCid.String(), nil)

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(node)
	h := helper.NewHelper(lib, cache, node, log, metrics)
	return deals.NewEventGenerator(h, log, metrics, network, parser.Config{}), node
}

func generateWithMiner(t *testing.T, network string, txs ...*types.Transaction) *types.DealsEvents {
	eg, node := setupMinerTest(t, network)
	ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})
	events, err := eg.GenerateDealsEvents(ctx, txs, tipsetCid, filTypes.EmptyTSK)
	require.NoError(t, err)
	return events
}

type update struct {
	sector uint64
	deals  []uint64
}

func replicaUpdatesParams(t *testing.T, updates []update, v2 bool) []byte {
	sealed := identityCid(t, "sealed")
	var buf bytes.Buffer
	if v2 {
		params := miner12.ProveReplicaUpdatesParams2{}
		for _, u := range updates {
			params.Updates = append(params.Updates, miner12.ReplicaUpdate2{
				SectorID: abi.SectorNumber(u.sector), NewSealedSectorCID: sealed, NewUnsealedSectorCID: sealed,
				Deals: toDealIDs(u.deals), ReplicaProof: []byte{1},
			})
		}
		require.NoError(t, params.MarshalCBOR(&buf))
		return buf.Bytes()
	}
	params := miner12.ProveReplicaUpdatesParams{}
	for _, u := range updates {
		params.Updates = append(params.Updates, miner12.ReplicaUpdate{
			SectorID: abi.SectorNumber(u.sector), NewSealedSectorCID: sealed,
			Deals: toDealIDs(u.deals), ReplicaProof: []byte{1},
		})
	}
	require.NoError(t, params.MarshalCBOR(&buf))
	return buf.Bytes()
}

func toDealIDs(ids []uint64) []abi.DealID {
	out := make([]abi.DealID, 0, len(ids))
	for _, id := range ids {
		out = append(out, abi.DealID(id))
	}
	return out
}

func bitfieldCBOR(t *testing.T, bits ...uint64) []byte {
	var buf bytes.Buffer
	bf := bitfield.NewFromSet(bits)
	require.NoError(t, bf.MarshalCBOR(&buf))
	return buf.Bytes()
}

// replicaUpdateTx parses raw ProveReplicaUpdates(2) params and return with the miner actor parser,
// as the transaction parser does, and wraps the metadata in a successful message to the miner.
func replicaUpdateTx(t *testing.T, network string, height uint64, id string, v2 bool, rawParams, rawReturn []byte) *types.Transaction {
	m := minerActor.New(logger.NewDevelopmentLogger())
	var (
		metadata map[string]interface{}
		err      error
		txType   = parser.MethodProveReplicaUpdates
	)
	if v2 {
		txType = parser.MethodProveReplicaUpdates2
		metadata, err = m.ProveReplicaUpdates2(network, int64(height), rawParams, rawReturn) //nolint:gosec
	} else {
		metadata, err = m.ProveReplicaUpdates(network, int64(height), rawParams, rawReturn) //nolint:gosec
	}
	require.NoError(t, err)
	jsonMetadata, err := json.Marshal(metadata)
	require.NoError(t, err)
	tx := okTx(height, id, txType, string(jsonMetadata))
	tx.TxCid = fmt.Sprintf("%s-%s", txCid, id)
	tx.TxFrom = "f3worker"
	tx.TxTo = minerAddr
	return tx
}

func minerTx(height uint64, id, txType, metadata string) *types.Transaction {
	tx := okTx(height, id, txType, metadata)
	tx.TxCid = fmt.Sprintf("%s-%s", txCid, id)
	tx.TxFrom = "f3worker"
	tx.TxTo = minerAddr
	return tx
}

type mapping struct {
	sector, deal uint64
}

func mappings(t *testing.T, events *types.DealsEvents) []mapping {
	out := []mapping{}
	ids := map[string]bool{}
	for _, m := range events.DealsSectorMappings {
		require.Falsef(t, ids[m.ID], "duplicated mapping id %s", m.ID)
		ids[m.ID] = true
		out = append(out, mapping{m.SectorNumber, m.DealID})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].sector != out[j].sector {
			return out[i].sector < out[j].sector
		}
		return out[i].deal < out[j].deal
	})
	return out
}

// mainnet heights per network version with replica updates (actors v7 to v12)
var replicaUpdateHeights = []struct {
	name   string
	height uint64
	v2     bool // ProveReplicaUpdates2 exists (actors v9 / NV17 on)
}{
	{"NV15 actors v7", 1700000, false},
	{"NV16 actors v8", 2000000, false},
	{"NV17 actors v9", 2400000, true},
	{"NV18 actors v10", 2700000, true},
	{"NV19 actors v11", 2820000, true},
	{"NV20 actors v11", 3000000, true},
	{"NV21 actors v12", 3500000, true},
	{"last NV21 height", 3855359, true},
}

// Only the updates the return reports as successful map their deals; a skipped update in the same
// batch maps nothing.
func TestReplicaUpdateMappings_PartialSuccess(t *testing.T) {
	updates := []update{{10, []uint64{1, 2}}, {11, []uint64{3}}, {12, []uint64{4}}}
	for _, tt := range replicaUpdateHeights {
		for _, v2 := range []bool{false, true} {
			if v2 && !tt.v2 {
				continue
			}
			t.Run(fmt.Sprintf("%s v2=%t", tt.name, v2), func(t *testing.T) {
				tx := replicaUpdateTx(t, tools.MainnetNetwork, tt.height, "a", v2, replicaUpdatesParams(t, updates, v2), bitfieldCBOR(t, 10, 12))
				events := generateWithMiner(t, tools.MainnetNetwork, tx)
				require.Equal(t, []mapping{{10, 1}, {10, 2}, {12, 4}}, mappings(t, events))
				for _, m := range events.DealsSectorMappings {
					require.Equal(t, minerAddr, m.ActorAddress)
					require.Equal(t, tx.TxType, m.ActionType)
					require.Equal(t, tx.TxCid, m.TxCid)
					require.Equal(t, tt.height, m.Height)
				}
				// a miner method is not a deals message, and maps nothing else
				require.Empty(t, events.DealsMessages)
				require.Empty(t, events.DealsActivations)
				require.Empty(t, events.DealsTerminations)
			})
		}
	}
}

func TestReplicaUpdateMappings_AllAndNone(t *testing.T) {
	updates := []update{{10, []uint64{1}}, {11, []uint64{2}}}
	events := generateWithMiner(t, tools.MainnetNetwork,
		replicaUpdateTx(t, tools.MainnetNetwork, 3500000, "a", false, replicaUpdatesParams(t, updates, false), bitfieldCBOR(t, 10, 11)),
	)
	require.Equal(t, []mapping{{10, 1}, {11, 2}}, mappings(t, events))

	// an empty bitfield: nothing updated
	events = generateWithMiner(t, tools.MainnetNetwork,
		replicaUpdateTx(t, tools.MainnetNetwork, 3500000, "a", false, replicaUpdatesParams(t, updates, false), bitfieldCBOR(t)),
	)
	require.Empty(t, events.DealsSectorMappings)
}

// A sector listed twice is skipped as a duplicate from its second occurrence on: its deals are the
// first occurrence's.
func TestReplicaUpdateMappings_DuplicateSector(t *testing.T) {
	updates := []update{{20, []uint64{5}}, {20, []uint64{6}}}
	events := generateWithMiner(t, tools.MainnetNetwork,
		replicaUpdateTx(t, tools.MainnetNetwork, 3500000, "a", true, replicaUpdatesParams(t, updates, true), bitfieldCBOR(t, 20)),
	)
	require.Equal(t, []mapping{{20, 5}}, mappings(t, events))
}

// From actors v13 (NV22) on the mappings come from the activations, which carry the sector
// number; replica updates and pre-commits sent to the miner map nothing. The tipset at the upgrade
// height still runs actors v12 (its deals are activated by v12 BatchActivateDeals, without a
// sector), so its replica updates and pre-commits do map.
func TestSectorMappings_NV22Boundary(t *testing.T) {
	updates := []update{{10, []uint64{1}}}
	preCommit := `{"MethodNum":"6","Params":{"SealProof":8,"SectorNumber":30,"DealIDs":[7],"Expiration":1}}`
	for _, tt := range []struct {
		network                   string
		before, upgrade, firstV13 uint64
	}{
		{tools.MainnetNetwork, 3855359, 3855360, 3855361},
		{tools.CalibrationNetwork, 1427973, 1427974, 1427975},
	} {
		t.Run(tt.network, func(t *testing.T) {
			for _, height := range []uint64{tt.before, tt.upgrade} {
				events := generateWithMiner(t, tt.network,
					replicaUpdateTx(t, tt.network, height, "a", true, replicaUpdatesParams(t, updates, true), bitfieldCBOR(t, 10)),
					minerTx(height, "b", parser.MethodPreCommitSector, preCommit),
				)
				require.Equalf(t, []mapping{{10, 1}, {30, 7}}, mappings(t, events), "height %d", height)
			}

			events := generateWithMiner(t, tt.network,
				replicaUpdateTx(t, tt.network, tt.firstV13, "a", true, replicaUpdatesParams(t, updates, true), bitfieldCBOR(t, 10)),
				minerTx(tt.firstV13, "b", parser.MethodPreCommitSector, preCommit),
			)
			require.Empty(t, events.DealsSectorMappings)
		})
	}
}

// A (sector, deal) repeated in one call gives one row (it would carry the same id).
func TestSectorMappings_RepeatedPairInOneCall(t *testing.T) {
	events := generateWithMiner(t, tools.MainnetNetwork,
		minerTx(2907480, "a", parser.MethodPreCommitSectorBatch, `{"MethodNum":"25","Params":{"Sectors":[{"SectorNumber":40,"DealIDs":[1,1,2]}]}}`),
		minerTx(2907480, "b", parser.MethodPreCommitSectorBatch, `{"MethodNum":"25","Params":{"Sectors":[{"SectorNumber":40,"DealIDs":[1]}]}}`),
	)
	// the second tx keeps its own row, with a distinct id
	require.Equal(t, []mapping{{40, 1}, {40, 1}, {40, 2}}, mappings(t, events))
}

// Unusable metadata never fails the height: the tx's mappings are skipped and the other txs are
// processed.
func TestReplicaUpdateMappings_UnusableMetadata(t *testing.T) {
	good := replicaUpdateTx(t, tools.MainnetNetwork, 3500000, "good", false, replicaUpdatesParams(t, []update{{9, []uint64{90}}}, false), bitfieldCBOR(t, 9))
	params := `{"Updates":[{"SectorID":10,"Deals":[1]},{"SectorID":11,"Deals":[2]}]}`
	for name, tt := range map[string]struct{ txType, metadata string }{
		// ProveReplicaUpdates metadata stored before the return was parsed
		"no return":              {parser.MethodProveReplicaUpdates, `{"MethodNum":"27","Params":` + params + `}`},
		"null return":            {parser.MethodProveReplicaUpdates2, `{"MethodNum":"29","Params":` + params + `,"Return":null}`},
		"sector not in params":   {parser.MethodProveReplicaUpdates, `{"MethodNum":"27","Params":` + params + `,"Return":[12,1]}`},
		"more bits than updates": {parser.MethodProveReplicaUpdates, `{"MethodNum":"27","Params":` + params + `,"Return":[10,3]}`},
		"bad bitfield":           {parser.MethodProveReplicaUpdates, `{"MethodNum":"27","Params":` + params + `,"Return":[1,0,2]}`},
		"bad sector id":          {parser.MethodProveReplicaUpdates, `{"MethodNum":"27","Params":{"Updates":[{"SectorID":-1,"Deals":[1]}]},"Return":[0]}`},
		"bad deal id":            {parser.MethodProveReplicaUpdates, `{"MethodNum":"27","Params":{"Updates":[{"SectorID":1,"Deals":[1.5]}]},"Return":[1,1]}`},
		"no updates":             {parser.MethodProveReplicaUpdates, `{"MethodNum":"27","Params":{},"Return":[0]}`},
		"raw only":               {parser.MethodProveReplicaUpdates, `{"MethodNum":"27","ParamsRaw":"gA=="}`},
		"bad pre-commit":         {parser.MethodPreCommitSector, `{"MethodNum":"6","Params":{"SectorNumber":"x","DealIDs":[1]}}`},
		"bad pre-commit batch":   {parser.MethodPreCommitSectorBatch2, `{"MethodNum":"28","Params":{"Sectors":{}}}`},
		"not json":               {parser.MethodPreCommitSector, `{`},
	} {
		t.Run(name, func(t *testing.T) {
			events := generateWithMiner(t, tools.MainnetNetwork, minerTx(3500000, "bad", tt.txType, tt.metadata), good)
			require.Equal(t, []mapping{{9, 90}}, mappings(t, events))
		})
	}
}

// Pre-commits before NV22: one row per deal id of each sector; sectors without deals map nothing.
func TestPreCommitMappings(t *testing.T) {
	events := generateWithMiner(t, tools.MainnetNetwork,
		minerTx(2907480, "a", parser.MethodPreCommitSector, `{"MethodNum":"6","Params":{"SealProof":8,"SectorNumber":30,"SealedCID":{"/":"x"},"DealIDs":[7,8],"Expiration":1}}`),
		minerTx(2907480, "b", parser.MethodPreCommitSector, `{"MethodNum":"6","Params":{"SealProof":8,"SectorNumber":31,"DealIDs":null,"Expiration":1}}`),
		minerTx(1700000, "c", parser.MethodPreCommitSectorBatch, `{"MethodNum":"25","Params":{"Sectors":[{"SectorNumber":40,"DealIDs":[1]},{"SectorNumber":41,"DealIDs":[]},{"SectorNumber":0,"DealIDs":[2,3]}]}}`),
		minerTx(3500000, "d", parser.MethodPreCommitSectorBatch2, `{"MethodNum":"28","Params":{"Sectors":[{"SectorNumber":9007199254740993,"DealIDs":[9007199254740995]}]}}`),
	)
	require.Equal(t, []mapping{{0, 2}, {0, 3}, {30, 7}, {30, 8}, {40, 1}, {9007199254740993, 9007199254740995}}, mappings(t, events))
	require.Empty(t, events.DealsMessages)
}

// Only successful txs map deals.
func TestSectorMappings_OnlySuccessful(t *testing.T) {
	failed := replicaUpdateTx(t, tools.MainnetNetwork, 3500000, "a", false, replicaUpdatesParams(t, []update{{9, []uint64{90}}}, false), bitfieldCBOR(t, 9))
	failed.Status = "ErrIllegalArgument"
	failedPreCommit := minerTx(2907480, "b", parser.MethodPreCommitSector, `{"MethodNum":"6","Params":{"SectorNumber":30,"DealIDs":[7]}}`)
	failedPreCommit.SubcallStatus = "ErrIllegalArgument"
	events := generateWithMiner(t, tools.MainnetNetwork, failed, failedPreCommit)
	require.Empty(t, events.DealsSectorMappings)
}

// Every activation that carries its sector (SectorContentChanged, BatchActivateDeals from actors
// v13) is copied as a mapping, so downstream has a single mapping source. Activations without a
// sector (ActivateDeals before v13) map nothing; their sectors come from the pre-commit or the
// replica update.
func TestActivationMappings(t *testing.T) {
	batch := okTx(3857557, "a", parser.MethodActivateDeals, sectorsActivateDealsMetadata)
	batch.TxFrom = minerAddr
	flat := okTx(550367, "b", parser.MethodActivateDeals, flatActivateDealsMetadata)
	flat.TxFrom = minerAddr
	events := generateWithMiner(t, tools.MainnetNetwork, batch, flat)

	require.Len(t, events.DealsActivations, 4)
	require.Equal(t, []mapping{{37656, 78950968}, {37888, 78951195}}, mappings(t, events))
	for _, m := range events.DealsSectorMappings {
		require.Equal(t, minerAddr, m.ActorAddress)
		require.Equal(t, parser.MethodActivateDeals, m.ActionType)
	}
}

// joinTerminatedDeals is the downstream join, on the parser's own rows: a termination with a deal
// id (actors v12 and the NV22 upgrade tipset) is that deal; a termination with a sector (actors
// v13+) is every deal mapped to (actor_address, sector_number). Distinct deal ids.
func joinTerminatedDeals(events *types.DealsEvents) []uint64 {
	type key struct {
		miner  string
		sector uint64
	}
	mapped := map[key][]uint64{}
	for _, m := range events.DealsSectorMappings {
		k := key{m.ActorAddress, m.SectorNumber}
		mapped[k] = append(mapped[k], m.DealID)
	}
	deals := map[uint64]bool{}
	for _, t := range events.DealsTerminations {
		if t.DealID != nil {
			deals[*t.DealID] = true
			continue
		}
		for _, d := range mapped[key{t.ActorAddress, *t.SectorNumber}] {
			deals[d] = true
		}
	}
	out := []uint64{}
	for d := range deals {
		out = append(out, d)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

// The parser's own mappings join its own terminations, across eras: deals mapped before NV22 by a
// SnapDeals update and by a pre-commit, and a deal activated after NV22 with its sector, are found
// from one post-NV22 sector termination of the same miner. Synthetic: no real pair of a pre-NV22
// mapping and a later termination of the same sector is in the fixtures. The values of the SnapDeals
// update are the real f01919423 update at mainnet 3,573,062 (sector 17923, deal 69334621).
func TestSectorMappings_JoinTerminations(t *testing.T) {
	snap := replicaUpdateTx(t, tools.MainnetNetwork, 3573062, "snap", false,
		replicaUpdatesParams(t, []update{{17923, []uint64{69334621}}, {17924, []uint64{69334622}}}, false), bitfieldCBOR(t, 17923))
	preCommit := minerTx(2907480, "precommit", parser.MethodPreCommitSector, `{"MethodNum":"6","Params":{"SealProof":8,"SectorNumber":40,"DealIDs":[1,2],"Expiration":1}}`)
	// actors v13 BatchActivateDeals: sectors 37656 and 37888 (the third sector failed)
	activation := okTx(3857557, "activation", parser.MethodBatchActivateDeals, sectorsActivateDealsMetadata)
	activation.TxFrom = minerAddr
	// a v12 termination with deal ids, for another miner
	v12 := terminateTx(t, tools.MainnetNetwork, 3000000, "v12", dealIDsTerminateParams(t, 2999990, []uint64{500}))
	v12.TxFrom = "f01000"
	// one post-NV22 termination of: the SnapDeals sector, the pre-commit sector, the v13 activation
	// sector, the sector whose update failed, and a sector that never had deals
	v13 := terminateTx(t, tools.MainnetNetwork, 6539228, "v13", sectorsTerminateParams(t, 6539228, bitfield.NewFromSet([]uint64{17923, 40, 37656, 17924, 999})))
	v13.TxFrom = minerAddr
	// the same sectors terminated by another miner map nothing
	other := terminateTx(t, tools.MainnetNetwork, 6539228, "other", sectorsTerminateParams(t, 6539228, bitfield.NewFromSet([]uint64{17923, 40})))
	other.TxFrom = "f09999"

	events := generateWithMiner(t, tools.MainnetNetwork, snap, preCommit, activation, v12, v13, other)
	require.Equal(t, []uint64{1, 2, 500, 69334621, 78950968}, joinTerminatedDeals(events))
}
