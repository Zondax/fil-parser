package deals_test

import (
	"testing"

	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/require"
	parserContext "github.com/zondax/fil-parser/context"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

const (
	flatActivateDealsMetadata    = `{"MethodNum":"6","Params":{"DealIDs":[1596649,1596648],"SectorExpiry":1079530}}`
	sectorsActivateDealsMetadata = `{"MethodNum":"6","Params":{"Sectors":[{"SectorNumber":37656,"SectorType":8,"SectorExpiry":4920235,"DealIDs":[78950968]},{"SectorNumber":37888,"SectorType":8,"SectorExpiry":4920235,"DealIDs":[78951195]},{"SectorNumber":37549,"SectorType":8,"SectorExpiry":4920235,"DealIDs":[79166044]}],"ComputeCID":false},"Return":{"ActivationResults":{"SuccessCount":2,"FailCodes":[{"Idx":2,"Code":16}]},"Activations":[{"NonVerifiedDealSpace":"0","VerifiedInfos":[{"Client":3061409,"AllocationId":61239862,"Data":{"/":"baga6ea4seaqgvrjfj65lawcocwvrpgq7h53oghvto6akrys6wllhbbckchfgefy"},"Size":34359738368}],"UnsealedCid":{}},{"NonVerifiedDealSpace":"0","VerifiedInfos":[{"Client":3061409,"AllocationId":61240089,"Data":{"/":"baga6ea4seaqbuieim7slc3wu7kms436xpnorao5jxr6tqftnqsysfxcp5dnduia"},"Size":34359738368}],"UnsealedCid":{}}]}}`
	sectorsParamsOnly            = `{"Sectors":[{"SectorNumber":37656,"SectorType":8,"SectorExpiry":4920235,"DealIDs":[78950968]},{"SectorNumber":37888,"SectorType":8,"SectorExpiry":4920235,"DealIDs":[78951195]}],"ComputeCID":false}`
)

func okTx(height uint64, id, txType, metadata string) *types.Transaction {
	return &types.Transaction{
		TxBasicBlockData: types.TxBasicBlockData{BasicBlockData: types.BasicBlockData{Height: height}},
		Id:               id,
		TxCid:            txCid,
		TxType:           txType,
		TxFrom:           txFrom,
		TxTo:             txTo,
		TxMetadata:       metadata,
		Status:           tools.GetExitCodeOK(),
		SubcallStatus:    tools.GetExitCodeOK(),
	}
}

// The ids of rows that existed before SectorContentChanged support must not change for the same
// input. The expected ids were produced by origin/dev at e67c606 (#583).
func TestDealsEventIdsUnchanged(t *testing.T) {
	eg, node := setupTest(t, "mainnet")
	ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})

	txs := []*types.Transaction{
		okTx(1594681, "a", parser.MethodVerifyDealsForActivation, `{"MethodNum":"5","Params":{"Sectors":[{"SectorExpiry":3149856,"DealIDs":[4137277]}]},"Return":{"Sectors":[{"DealSpace":34359738368,"DealWeight":"0","VerifiedDealWeight":"52941484877414400"}]}}`),
		okTx(550367, "b", parser.MethodActivateDeals, flatActivateDealsMetadata),
		okTx(3857557, "c", parser.MethodActivateDeals, sectorsActivateDealsMetadata),
	}
	events, err := eg.GenerateDealsEvents(ctx, txs, tipsetCid, filTypes.EmptyTSK)
	require.NoError(t, err)

	messages := []string{}
	for _, m := range events.DealsMessages {
		messages = append(messages, m.ActionType+" "+m.ID)
	}
	require.Equal(t, []string{
		"VerifyDealsForActivation 918fd49a-af25-5b8c-9484-f9913c4f4cd0",
		"ActivateDeals e0abf3e8-e369-5f79-b29a-e344dfbfa604",
		"BatchActivateDeals 31a1e4ee-7ae8-5192-b35d-f975fa0ae648",
	}, messages)

	activations := map[uint64]string{}
	for _, a := range events.DealsActivations {
		activations[a.DealID] = a.ID
	}
	require.Equal(t, map[uint64]string{
		1596649:  "f9faaa90-52bd-575d-a9d9-6bf5235656df",
		1596648:  "ab528652-8dc2-5e75-b5b0-036d52311657",
		78950968: "aed172f2-795b-58b1-8324-01edef8bee63",
		78951195: "60b1c405-7f55-5f5a-8d8e-509a8ce4f1bd",
	}, activations)

	space := map[uint64]string{}
	for _, s := range events.DealsSpaceInfo {
		space[s.DealID] = s.ID
	}
	require.Equal(t, map[uint64]string{
		4137277:  "51f051fd-7ae6-5410-9fa4-32812b4be75f",
		78950968: "aed172f2-795b-58b1-8324-01edef8bee63",
		78951195: "60b1c405-7f55-5f5a-8d8e-509a8ce4f1bd",
	}, space)
}

// ActivateDeals params are dispatched on their shape, not on the network version: NV21 (actors
// v12) already uses the sectors shape, so the boundary is the NV21 upgrade height on both networks.
func TestActivateDealsShapeAtUpgradeBoundary(t *testing.T) {
	tests := []struct {
		network  string
		height   uint64
		txType   string
		metadata string
		expected []uint64
	}{
		// last flat height (NV20) and first sectors height (NV21, Watermelon)
		{"mainnet", 3469379, parser.MethodActivateDeals, flatActivateDealsMetadata, []uint64{1596648, 1596649}},
		{"mainnet", 3469380, parser.MethodBatchActivateDeals, sectorsActivateDealsMetadata, []uint64{78950968, 78951195}},
		{"calibration", 1108173, parser.MethodActivateDeals, flatActivateDealsMetadata, []uint64{1596648, 1596649}},
		{"calibration", 1108174, parser.MethodBatchActivateDeals, sectorsActivateDealsMetadata, []uint64{78950968, 78951195}},
	}
	for _, tt := range tests {
		t.Run(tt.network+" "+tt.txType, func(t *testing.T) {
			eg, node := setupTest(t, tt.network)
			ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})
			events, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{okTx(tt.height, "x", tt.txType, tt.metadata)}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.Equal(t, tt.expected, dealIDs(events.DealsActivations))
			require.Len(t, events.DealsMessages, 1)
		})
	}
}

// Missing or inconsistent params/return never panic or fail the height: the activations of that
// tx are skipped and its message row is kept.
func TestBatchActivateDealsUnusableMetadata(t *testing.T) {
	eg, node := setupTest(t, "mainnet")
	ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})

	tests := []struct {
		name     string
		metadata string
	}{
		{"params without return", `{"MethodNum":"6","Params":` + sectorsParamsOnly + `}`},
		{"return with no activations", `{"MethodNum":"6","Params":` + sectorsParamsOnly + `,"Return":{"ActivationResults":{"SuccessCount":2,"FailCodes":null},"Activations":null}}`},
		{"return shorter than the successful sectors", `{"MethodNum":"6","Params":` + sectorsParamsOnly + `,"Return":{"ActivationResults":{"SuccessCount":2,"FailCodes":null},"Activations":[{"NonVerifiedDealSpace":"0","VerifiedInfos":[],"UnsealedCid":{}}]}}`},
		{"raw-only metadata", `{"MethodNum":"6","ParamsRaw":"gg==","ReturnRaw":"gg=="}`},
		{"params not a map", `{"MethodNum":"6","Params":[1,2]}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanics(t, func() {
				bad := okTx(3857651, "bad", parser.MethodBatchActivateDeals, tt.metadata)
				good := okTx(3857651, "good", parser.MethodBatchActivateDeals, sectorsActivateDealsMetadata)
				events, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{bad, good}, tipsetCid, filTypes.EmptyTSK)
				require.NoError(t, err)
				require.Len(t, events.DealsMessages, 2)
				require.Equal(t, []uint64{78950968, 78951195}, dealIDs(events.DealsActivations))
				require.Len(t, events.DealsSpaceInfo, 2)
			})
		})
	}
}

// SectorDeals carries the sector number only from actors v13 (NV22) on. For actors v12 (NV21) the
// decoded metadata shows the Go struct's zero value, which must not be recorded.
func TestBatchActivateDealsSectorNumberFromActorsV13(t *testing.T) {
	// as decoded for actors v12: SectorNumber is the zero value
	v12Metadata := `{"MethodNum":"6","Params":{"Sectors":[{"SectorNumber":0,"SectorType":8,"SectorExpiry":4920235,"DealIDs":[78950968]},{"SectorNumber":0,"SectorType":8,"SectorExpiry":4920235,"DealIDs":[78951195]}],"ComputeCID":false},"Return":{"ActivationResults":{"SuccessCount":2,"FailCodes":null},"Activations":[{"NonVerifiedDealSpace":"0","VerifiedInfos":[],"UnsealedCid":{}},{"NonVerifiedDealSpace":"0","VerifiedInfos":[],"UnsealedCid":{}}]}}`

	tests := []struct {
		name     string
		network  string
		height   uint64
		metadata string
		// expected sector number per deal, nil when it must not be recorded
		expected map[uint64]*uint64
	}{
		{"mainnet last v12 height, zero value", "mainnet", 3855359, v12Metadata, map[uint64]*uint64{78950968: nil, 78951195: nil}},
		{"mainnet last v12 height, non-zero value", "mainnet", 3855359, sectorsActivateDealsMetadata, map[uint64]*uint64{78950968: nil, 78951195: nil}},
		{"mainnet first v13 height", "mainnet", 3855360, sectorsActivateDealsMetadata, map[uint64]*uint64{78950968: ptr(37656), 78951195: ptr(37888)}},
		{"mainnet first v13 height, sector 0", "mainnet", 3855360, v12Metadata, map[uint64]*uint64{78950968: ptr(0), 78951195: ptr(0)}},
		{"mainnet v12 fixture height", "mainnet", 3573062, v12Metadata, map[uint64]*uint64{78950968: nil, 78951195: nil}},
		{"calibration last v12 height", "calibration", 1427973, v12Metadata, map[uint64]*uint64{78950968: nil, 78951195: nil}},
		{"calibration first v13 height", "calibration", 1427974, sectorsActivateDealsMetadata, map[uint64]*uint64{78950968: ptr(37656), 78951195: ptr(37888)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eg, node := setupTest(t, tt.network)
			ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})
			events, err := eg.GenerateDealsEvents(ctx, []*types.Transaction{okTx(tt.height, "x", parser.MethodBatchActivateDeals, tt.metadata)}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)
			got := map[uint64]*uint64{}
			for _, a := range events.DealsActivations {
				got[a.DealID] = a.SectorNumber
			}
			require.Equal(t, tt.expected, got)
		})
	}
}

func ptr(v uint64) *uint64 { return &v }
