package verifreg_test

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

// The allocation Provider is an interface{} field: with the exact decoding an ActorID provider is a
// json.Number and must still be consolidated to an ID address, with the same deal data as before.
func TestUniversalReceiverHook_ActorIDProvider(t *testing.T) {
	eg, node := setupTest(t, tools.MainnetNetwork)
	ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})
	events, err := eg.GenerateVerifregEvents(ctx, []*types.Transaction{{
		TxBasicBlockData: types.TxBasicBlockData{BasicBlockData: types.BasicBlockData{Height: 2790000}},
		TxCid:            txCid,
		TxType:           parser.MethodUniversalReceiverHook,
		TxFrom:           txFrom,
		TxTo:             txTo,
		TxMetadata:       `{"MethodNum":"3726118371","Params":{"Type_":2233613279,"Payload":"hhoAIA0lBgVNAN4Lazp2QAAAAAAAAFhNgoGGGgARICPYKlgoAAGB4gOSICDR58jNGlgp2quDxIKIVbj0CP07ogD/LURzu8p5DTeXNhsAAAAQAAAAABoACTqAGgANLwAaACrKroBA"},"Return":{"AllocationResults":{"SuccessCount":1,"FailCodes":null},"ExtensionResults":{"SuccessCount":0,"FailCodes":null},"NewAllocations":[15773663]}}`,
		Status:           tools.GetExitCodeOK(),
		SubcallStatus:    tools.GetExitCodeOK(),
	}}, tipsetCid, filTypes.EmptyTSK)
	require.NoError(t, err)
	require.Len(t, events.Deals, 1)
	require.Equal(t, "f01122339", events.Deals[0].ProviderAddress)
	require.Equal(t, `{"Provider":1122339,"Data":{"/":"baga6ea4seaqndz6izunfqko2vob4jauikw4pich5horab7znirz3xstzbu3zonq"},"Size":68719476736,"TermMin":604800,"TermMax":864000,"Expiration":2804398,"DealID":"15773663"}`, events.Deals[0].Data)
}
