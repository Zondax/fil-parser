package evm_test

import (
	"context"
	_ "embed"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	filMetrics "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/evm"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	metrics2 "github.com/zondax/golem/pkg/metrics"
)

//go:embed test-data/FilecoinPayV1.abi.json
var filecoinPayV1ABI []byte
var filecoinPayV1Address = "f410feoy6aghqro4yenelcwug52jg527x6tnkxl36cyi"

func setupTest(t *testing.T) evm.EventGenerator {
	logger := logger.NewDevelopmentLogger()
	metrics := filMetrics.NewMetricsClient(metrics2.NewNoopMetrics())

	contracts := map[string][]byte{
		filecoinPayV1Address: filecoinPayV1ABI,
	}

	eg, err := evm.NewEventGenerator(logger, metrics, contracts)
	require.NoError(t, err)
	return eg
}

var (
	txFrom    = "f02"
	tipsetCid = "bafy2bzaceczpzd5k7u6hwaim7fdpwx2ujg7uhrdbpijf7q5ryvh7ogmawxupk"
	txCid     = "bafy2bzaceczpzd5k7u6hwaim7fdpwx2ujg7uhrdbpijf7q5ryvh7ogmawxupk"
)

func TestEvmEvents(t *testing.T) {
	eg := setupTest(t)
	tests := []struct {
		name      string
		txType    string
		actorName string
		txFrom    string
		txTo      string
		metadata  string
		want      []types.EvmEvent
	}{
		{
			name:      "modifyRailPayment",
			txType:    parser.MethodInvokeContract,
			actorName: manifest.EvmKey,
			txFrom:    txFrom,
			txTo:      filecoinPayV1Address,
			metadata:  `{"MethodNum":"3844450837","Params":"0x97d3ea34000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000a1b01d4b1c0000000000000000000000000000000000000000000000000000000000000000","Return":"0x"}`,
			want: []types.EvmEvent{
				{
					FunctionName: "modifyRailPayment",
					ParsedMetadata: map[string]any{
						"Params": map[string]any{
							"newRate":        big.NewInt(694444444444),
							"oneTimePayment": big.NewInt(0),
							"railId":         big.NewInt(12),
						},
						"Return": map[string]any{},
					},
				},
			},
		},
		{
			name:      "createRail",
			txType:    parser.MethodInvokeContract,
			actorName: manifest.EvmKey,
			txFrom:    txFrom,
			txTo:      filecoinPayV1Address,
			metadata:  `{"MethodNum":"3844450837","Params":"0xf9f78de800000000000000000000000080b98d3aa09ffff255c3ba4a241111ff1262f045000000000000000000000000d47f5823b1858df3c4928fd24aa0a3d48579e3c400000000000000000000000032c90c26bca6ed3945de9b29ba4e19d38314d6180000000000000000000000008408502033c418e1bbc97ce9ac48e5528f371a9f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000008408502033c418e1bbc97ce9ac48e5528f371a9f","Return":"0x0000000000000000000000000000000000000000000000000000000000000059"}`,
			want: []types.EvmEvent{
				{
					FunctionName: "modifyRailPayment",
					ParsedMetadata: map[string]any{
						"Params": map[string]any{
							"commissionRateBps":   big.NewInt(0),
							"from":                common.HexToAddress("0xD47f5823B1858DF3C4928FD24AA0A3D48579E3c4"),
							"serviceFeeRecipient": common.HexToAddress("0x8408502033C418E1bbC97cE9ac48E5528F371A9f"),
							"to":                  common.HexToAddress("0x32c90c26bCA6eD3945De9b29BA4e19D38314D618"),
							"token":               common.HexToAddress("0x80B98d3aa09ffff255c3ba4A241111Ff1262F045"),
							"validator":           common.HexToAddress("0x8408502033C418E1bbC97cE9ac48E5528F371A9f"),
						},
						"Return": map[string]any{
							"": big.NewInt(89),
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.txType, func(t *testing.T) {
			events, err := eg.GenerateEvmEvents(context.Background(), []*types.Transaction{
				{
					TxCid:         txCid,
					TxType:        test.txType,
					TxFrom:        test.txFrom,
					TxTo:          test.txTo,
					TxMetadata:    test.metadata,
					Status:        tools.GetExitCodeOK(),
					SubcallStatus: tools.GetExitCodeOK(),
				},
			}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)

			require.Len(t, events, 1)
			require.Len(t, test.want, 1)

			// Compare the structure manually to handle big.Int comparison properly
			compareMetadata := func(actual, expected map[string]any, section string) {
				for key, expectedValue := range expected {
					actualValue, ok := actual[key]
					require.True(t, ok, "missing key in %s: %s", section, key)

					// Special handling for big.Int
					if expectedBigInt, ok := expectedValue.(*big.Int); ok {
						actualBigInt, ok := actualValue.(*big.Int)
						require.True(t, ok, "expected *big.Int for key in %s: %s", section, key)
						assert.Equal(t, 0, expectedBigInt.Cmp(actualBigInt), "big.Int values not equal for key in %s: %s", section, key)
					} else {
						assert.Equal(t, expectedValue, actualValue, "values not equal for key in %s: %s", section, key)
					}
				}
			}

			// Compare params
			actualParams := events[0].ParsedMetadata["Params"].(map[string]any)
			expectedParams := test.want[0].ParsedMetadata["Params"].(map[string]any)
			compareMetadata(actualParams, expectedParams, "Params")

			// Compare return values
			actualReturn := events[0].ParsedMetadata["Return"].(map[string]any)
			expectedReturn := test.want[0].ParsedMetadata["Return"].(map[string]any)
			compareMetadata(actualReturn, expectedReturn, "Return")
		})
	}
}
