package miner_test

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

// Integers above 2^53 in the tx metadata must reach the miner events unchanged: float64 would
// turn NewExpiration 9223372036854775807 into 9223372036854776000 and round sector numbers.
func TestMinerEvents_KeepLargeIntegers(t *testing.T) {
	eg, node := setupTest(t)

	generate := func(t *testing.T, txType, metadata string) *types.MinerEvents {
		t.Helper()
		ctx := parserContext.SetNodes(t.Context(), []api.FullNode{node})
		events, err := eg.GenerateMinerEvents(ctx, []*types.Transaction{{
			TxCid:         txCid,
			TxType:        txType,
			TxFrom:        txFrom,
			TxTo:          txTo,
			TxMetadata:    metadata,
			Status:        tools.GetExitCodeOK(),
			SubcallStatus: tools.GetExitCodeOK(),
		}}, tipsetCid, filTypes.EmptyTSK)
		require.NoError(t, err)
		return events
	}

	t.Run("ChangeBeneficiary NewExpiration MaxInt64", func(t *testing.T) {
		// the inner ChangeBeneficiary params of mainnet message
		// bafy2bzacedyg4uiezhkabfgvbbtfluie5xiukxdiivnjlfqsn7omti3r63rxy (keys in encoding order,
		// as json.Marshal writes a map)
		metadata := `{"MethodNum":"30","Params":{"NewBeneficiary":"f02117153","NewExpiration":9223372036854775807,"NewQuota":"340282366920938463463374607431768211455"}}`
		events := generate(t, parser.MethodChangeBeneficiary, metadata)
		require.Len(t, events.MinerInfo, 1)
		require.Equal(t, metadata, events.MinerInfo[0].Data)
	})

	t.Run("small values and token amounts unchanged", func(t *testing.T) {
		metadata := `{"MethodNum":"30","Params":{"NewBeneficiary":"f02117153","NewExpiration":4000000,"NewQuota":"1000000000000000000"}}`
		events := generate(t, parser.MethodChangeBeneficiary, metadata)
		require.Len(t, events.MinerInfo, 1)
		require.Equal(t, metadata, events.MinerInfo[0].Data)
	})

	t.Run("PreCommitSector numbers above 2^53", func(t *testing.T) {
		metadata := `{"MethodNum":"6","Params":{"SealProof":8,"SectorNumber":9007199254740993,"SealedCID":{"/":"bagboea4b5abcapw3cbg4xe4cq7enh7uouusqsd3c3oiki34vjyjinp25vxpj3hlk"},"SealRandEpoch":517402,"DealIDs":[18446744073709551615,9007199254740993,67798],"Expiration":9223372036854775807,"ReplaceCapacity":false,"ReplaceSectorDeadline":0,"ReplaceSectorPartition":0,"ReplaceSectorNumber":0}}`
		events := generate(t, parser.MethodPreCommitSector, metadata)
		require.Len(t, events.MinerSectors, 1)
		require.Equal(t, uint64(9007199254740993), events.MinerSectors[0].SectorNumber)
		require.JSONEq(t, `{"DealIDs":[18446744073709551615,9007199254740993,67798],"Expiration":9223372036854775807,"SectorNumber":9007199254740993,"SectorSize":34359738368}`, events.MinerSectors[0].Data)
	})
}
