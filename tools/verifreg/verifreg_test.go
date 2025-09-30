package verifreg_test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-state-types/manifest"
	filApiTypes "github.com/filecoin-project/lotus/api/types"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	cid "github.com/ipfs/go-cid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	filMetrics "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/mocks"
	"github.com/zondax/fil-parser/tools/verifreg"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

var (
	tipsetCid   = "bafy2bzaceazaznpb47y6ljwmqrtkwcogpkuaslm2345w6khxdjq7r3cemurdw"
	txCid       = "bafy2bzaceacl3ywfoqdznurtpzwmhdgua6ak753ooxhvo4i64ekby2icvusgs"
	actorCidStr = "bafkqaftgnfwc6mjpozsxe2lgnfswi4tfm5uxg5dspe"
	txFrom      = "f07"
	txTo        = "f06"
	actorCid    cid.Cid
)

func init() {
	var err error
	actorCid, err = cid.Parse(actorCidStr)
	if err != nil {
		panic(err)
	}
}

func setupTest(_ *testing.T, network string) verifreg.EventGenerator {
	logger := logger.NewDevelopmentLogger()
	metrics := filMetrics.NewMetricsClient(metrics2.NewNoopMetrics())

	node := &mocks.FullNode{}
	node.On("StateNetworkName", mock.Anything).Return(dtypes.NetworkName("calibrationnet"), nil)
	node.On("StateNetworkVersion", mock.Anything, mock.Anything).Return(filApiTypes.NetworkVersion(16), nil)
	node.On("StateActorCodeCIDs", mock.Anything, mock.Anything).Return(map[string]cid.Cid{
		manifest.VerifregKey: actorCid,
	}, nil)

	cache := &mocks.IActorsCache{}
	cache.On("StoreAddressInfo", mock.Anything).Return(nil)
	cache.On("GetActorCode", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(actorCidStr, nil)
	cache.On("GetActorNameFromAddress", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(manifest.MinerKey, nil)

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(node)
	helper := helper.NewHelper(lib, cache, node, logger, metrics)

	return verifreg.NewEventGenerator(helper, logger, metrics, network, parser.Config{})
}

func TestParseUniversalReceiverHook(t *testing.T) {
	eg := setupTest(t, "mainnet")

	tests := []struct {
		name      string
		txType    string
		actorName string
		txFrom    string
		txTo      string
		metadata  string
		height    uint64
	}{

		{
			name:      "NV18",
			txType:    parser.MethodUniversalReceiverHook,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"3726118371","Params":{"Type_":2233613279,"Payload":"hhoAHKF4BgVNAG8FtZ07IAAAAAAAAFhNgoGGRADWqXfYKlgoAAGB4gOSICBKxq8wxilsMCvhalmlZIWNbI26TBugLVasm9mGBz8NCBsAAAAIAAAAABoAF4LAGgAbd0AaACkp04BA"},"Return":{"AllocationResults":{"SuccessCount":1,"FailCodes":null},"ExtensionResults":{"SuccessCount":0,"FailCodes":null},"NewAllocations":[10789415]}}`,
			height:    2683348,
		},
		{
			name:      "NV19",
			txType:    parser.MethodUniversalReceiverHook,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"3726118371","Params":{"Type_":2233613279,"Payload":"hhoAHw8tBgVOAAG8FtZ07IAAAAAAAABZASuChIYaACBt7dgqWCgAAYHiA5IgIC8nrg1pZodWnyqq81c3a9a1RWYpIDfPJIOkWDWrD942GwAAAAgAAAAAGgAXgsAaABt3QBoAKxf4hhoAIG3t2CpYKAABgeIDkiAg/Rz0HnyAisvyihsxL27uTE2CMx/m2Ag6Kxg2RpYWSyAbAAAACAAAAAAaABeCwBoAG3dAGgArF/eGGgAgbe3YKlgoAAGB4gOSICBlhlPm5grkPkh4+d0DXJ6K1ZosYMGk55sZ99v1bD2/DRsAAAAIAAAAABoAF4LAGgAbd0AaACsYBIYaACBt7dgqWCgAAYHiA5IgICe7t+VR372dFxW1vzFJBKCIE4Q0M8xJosnuQuOMXmo3GwAAAAgAAAAAGgAXgsAaABt3QBoAKxgEgEA="},"Return":{"AllocationResults":{"SuccessCount":4,"FailCodes":null},"ExtensionResults":{"SuccessCount":0,"FailCodes":null},"NewAllocations":[16545922,16545923,16545924,16545925]}}`,
			height:    2809800,
		},
		{
			name:      "NV20",
			txType:    parser.MethodUniversalReceiverHook,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"3726118371","Params":{"Type_":2233613279,"Payload":"hhoAIHyuBgVNAG8FtZ07IAAAAAAAAFhNgoGGGgAcYBPYKlgoAAGB4gOSICCwqXDFYFPgWWdT22G2HhrUn0jWcltsfb9CXyFvAHXCPhsAAAAIAAAAABoAF0qAGgAbPwAaACwOEYBA"},"Return":{"AllocationResults":{"SuccessCount":1,"FailCodes":null},"ExtensionResults":{"SuccessCount":0,"FailCodes":null},"NewAllocations":[19490992]}}`,
			height:    2870280,
		},
		{
			name:      "NV21",
			txType:    parser.MethodUniversalReceiverHook,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"3726118371","Params":{"Type_":2233613279,"Payload":"hhoAK7UABgVOAAN4Lazp2QAAAAAAAABZAlOCiIYaACs7rdgqWCgAAYHiA5IgIBYW1GKyYhQkOfCgL54UZN2XmQ4ij9Fn/X+mlKT8mC8eGwAAAAgAAAAAGgAH6WQaAAvd5BoANUpEhhoAKzut2CpYKAABgeIDkiAg80Fxtm3Qpc7RevJk71IeAXZkxhP4gyYL9V+USoSC5gYbAAAACAAAAAAaAAfpZBoAC93kGgA1SkSGGgArO63YKlgoAAGB4gOSICADfGvNgbjHQyEucYw20L6j1p+UraWHGAvcTL1N7SOXBBsAAAAIAAAAABoAB+lkGgAL3eQaADVKRIYaACs7rdgqWCgAAYHiA5IgIBhgflxIBly9SHNXHoXGqMBgVhz8QevuxRKbk1UhLF0SGwAAAAgAAAAAGgAH6WQaAAvd5BoANUpEhhoAKzut2CpYKAABgeIDkiAglSWOfNnVP9yADh5z8epCHQEs/pscxI+uYVtmE2QQcDgbAAAACAAAAAAaAAfpZBoAC93kGgA1SkSGGgArO63YKlgoAAGB4gOSICDL3x5gOf1r++/shWW8b6TnLoGA8pr6+tXum/uC2kOgFBsAAAAIAAAAABoAB+lkGgAL3eQaADVKRIYaACs7rdgqWCgAAYHiA5IgIDahWkgTJpk7R3mIZmrBiddqmYbaMdbUHRrOqI3ITJk5GwAAAAgAAAAAGgAH6WQaAAvd5BoANUpEhhoAKzut2CpYKAABgeIDkiAg/fZNjXgcdnflTMxkDb8BiRaqFMKrwH7QrExtR8Q5zw8bAAAACAAAAAAaAAfpZBoAC93kGgA1SkSAQA=="},"Return":{"AllocationResults":{"SuccessCount":8,"FailCodes":null},"ExtensionResults":{"SuccessCount":0,"FailCodes":null},"NewAllocations":[48504969,48504970,48504971,48504972,48504973,48504974,48504975,48504976]}}`,
			height:    3469380,
		},
		{
			name:      "NV22",
			txType:    parser.MethodUniversalReceiverHook,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"3726118371","Params":{"Type_":2233613279,"Payload":"hhoALsBLBgVOAAIrHIwSJ6AAAAAAAABZAXWChYYaACtK6tgqWCgAAYHiA5IgIEEQVzaOEKxdy3L5SXOYdCzKEzFLqpnvv2GOSzv8OeQMGwAAAAgAAAAAGgAJOoAaAA0vABoAOwEAhhoAK0rq2CpYKAABgeIDkiAgLrpweGLqVXRA72lH7bchNIb0mFo45+7xiX4k93BnwxAbAAAACAAAAAAaAAk6gBoADS8AGgA7AQCGGgArSurYKlgoAAGB4gOSICCt8qw8IWn29heqgUjkM8ReXUzIy03q8a+SqSRP7NOvCxsAAAAIAAAAABoACTqAGgANLwAaADsBAYYaACtK6tgqWCgAAYHiA5IgIAvVtBZ0Y4LdCcEvzh3yidgj1Y5oeJqfWbeIOjRtGcMHGwAAAAgAAAAAGgAJOoAaAA0vABoAOwD/hhoAK0rq2CpYKAABgeIDkiAgkwImfW+tmnsQnwX+I0kd5q0tgwOfJtknxgU+wZ7U4QkbAAAACAAAAAAaAAk6gBoADS8AGgA7AP+AQA=="},"Return":{"AllocationResults":{"SuccessCount":5,"FailCodes":null},"ExtensionResults":{"SuccessCount":0,"FailCodes":null},"NewAllocations":[61537644,61537645,61537646,61537647,61537648]}}`,
			height:    3855360,
		},
		{
			name:      "NV23",
			txType:    parser.MethodUniversalReceiverHook,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"3726118371","Params":{"Type_":2233613279,"Payload":"hhoAL/9FBgVOAAN4Lazp2QAAAAAAAABZAlOCiIYaAC8etNgqWCgAAYHiA5IgILAY/I3lGAA76/0hMSqyJ+m62Er5KM4L5Jufgtq+9ME6GwAAAAgAAAAAGgAH6QAaAAvdgBoAP51QhhoALx602CpYKAABgeIDkiAg3DU3Y1akJ70pWS7qCk8U7BSuC+TdrH8Jof4Tii4hSBcbAAAACAAAAAAaAAfpABoAC92AGgA/nVCGGgAvHrTYKlgoAAGB4gOSICB7h6VNkGO8TwPnFVvoN/P7JAwG6xjBc1yc7yS5443lGBsAAAAIAAAAABoAB+kAGgAL3YAaAD+dUIYaAC8etNgqWCgAAYHiA5IgIJoxaU2Vm/wD9ap6RA+MRlmW/d1o1HeHrkpmGu0gn7EUGwAAAAgAAAAAGgAH6QAaAAvdgBoAP51QhhoALx602CpYKAABgeIDkiAglaNQgZ04nqvrYZjxNa9j9DMUG6canxzZo41i8CQfPiAbAAAACAAAAAAaAAfpABoAC92AGgA/nVCGGgAvHrTYKlgoAAGB4gOSICBX5+HkPu2PvxsaGtUuFQp9m/vFd/5V1b+U6q/APAQ4PRsAAAAIAAAAABoAB+kAGgAL3YAaAD+dUIYaAC8etNgqWCgAAYHiA5IgIHuKW0NgymMT/uY/vZABQyN2QSkQxR93a2ZDCM3PxW0nGwAAAAgAAAAAGgAH6QAaAAvdgBoAP51QhhoALx602CpYKAABgeIDkiAgpAFzWgxfgEqlyKai4vIaRoFHMMGNn8TNlTWQ/XaffSsbAAAACAAAAAAaAAfpABoAC92AGgA/nVCAQA=="},"Return":{"AllocationResults":{"SuccessCount":8,"FailCodes":null},"ExtensionResults":{"SuccessCount":0,"FailCodes":null},"NewAllocations":[70800098,70800099,70800100,70800101,70800102,70800103,70800104,70800105]}}`,
			height:    4154640,
		},
		{
			name:      "NV24",
			txType:    parser.MethodUniversalReceiverHook,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"3726118371","Params":{"Type_":2233613279,"Payload":"hhoAMaLiBgVOAAN4Lazp2QAAAAAAAABZAlOCiIYaADEUZdgqWCgAAYHiA5IgIHZiT+qvxPQRQr2UkFSSLfM6cyHpz9yMkDVPobKtc9Y9GwAAAAgAAAAAGgAPmcAaABOOQBoARErxhhoAMRRl2CpYKAABgeIDkiAgteAL0KLTl6hWOgZLHGIf0y1SB67D3WueJtoG3rueixcbAAAACAAAAAAaAA+ZwBoAE45AGgBESvGGGgAxFGXYKlgoAAGB4gOSICC3GaWvnv3ZZN+ryLglzzvzzUZFAQQZOT5HIpFM7tf8CBsAAAAIAAAAABoAD5nAGgATjkAaAERK8YYaADEUZdgqWCgAAYHiA5IgIG4DmNA0QdtS1oBYwbQ8weg48N/EyIZN6Au6WK4L0EIaGwAAAAgAAAAAGgAPmcAaABOOQBoARErwhhoAMRRl2CpYKAABgeIDkiAgKE6J1ifDFbhUZ8pnxzhItwF2atg2Aft/oy4SCekR/RsbAAAACAAAAAAaAA+ZwBoAE45AGgBESvCGGgAxFGXYKlgoAAGB4gOSICBwzu+eDAHLRaZsLEFWgN7IrEpKcigKMHe0MUuVH3inABsAAAAIAAAAABoAD5nAGgATjkAaAERK8IYaADEUZdgqWCgAAYHiA5IgILV9hAHH1wNNWH4WchL14RecS6tkW+NYe45ALQrEEZITGwAAAAgAAAAAGgAPmcAaABOOQBoARErxhhoAMRRl2CpYKAABgeIDkiAgiaph02bu32fgUIv6ep2Ef0IgRAx1Q+G3zYzaBXgE3R8bAAAACAAAAAAaAA+ZwBoAE45AGgBESvGAQA=="},"Return":{"AllocationResults":{"SuccessCount":8,"FailCodes":null},"ExtensionResults":{"SuccessCount":0,"FailCodes":null},"NewAllocations":[82182485,82182486,82182487,82182488,82182489,82182490,82182491,82182492]}}`,
			height:    4461240,
		},
		{
			name:      "NV25",
			txType:    parser.MethodUniversalReceiverHook,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"3726118371","Params":{"Type_":2233613279,"Payload":"hhoANeG+BgVOAARWORgkT0AAAAAAAABZAueCioYaADVn4NgqWCgAAYHiA5IgIJmpTRSCyzJfKcab7PBXD5yzCCmMAMxsqmFNTxsyvLEJGwAAAAgAAAAAGgAXSoAaABs/ABoASsCvhhoANWfg2CpYKAABgeIDkiAg2zMf+FSrgjjpw0YRacs+LA9rOqqMpjiQRXdOyC2ehA4bAAAACAAAAAAaABdKgBoAGz8AGgBKwLKGGgA1Z+DYKlgoAAGB4gOSICAqn2/FovemDlEax+ltVZOZg+BiqgMhj5X4s/z/RpcWLBsAAAAIAAAAABoAF0qAGgAbPwAaAErAsoYaADVn4NgqWCgAAYHiA5IgIEK/tYwwGhjx4crcyvL8sfhAyp0i4dIvertuR/oPsEc5GwAAAAgAAAAAGgAXSoAaABs/ABoASsC2hhoANWfg2CpYKAABgeIDkiAgu3hmhC2A/Qf3SAzD7YnMdHB5zXolhniefvkfacFbEzsbAAAACAAAAAAaABdKgBoAGz8AGgBKwK6GGgA1Z+DYKlgoAAGB4gOSICB9ienI6YsYCKBftIytlXGSCbeYtdPUGa0oShKJex7qGRsAAAAIAAAAABoAF0qAGgAbPwAaAErAr4YaADVn4NgqWCgAAYHiA5IgIF4JkdWBapOqZbvdwGw6JMVBAg/BLoHhrCG+OSKOCOEpGwAAAAgAAAAAGgAXSoAaABs/ABoASsCvhhoANWfg2CpYKAABgeIDkiAg9zbQSJtHVbdAHun2ty1zOD+4pMqM6j3EfNaFiXMJmicbAAAACAAAAAAaABdKgBoAGz8AGgBKwK+GGgA1Z+DYKlgoAAGB4gOSICBZs6rBTEi0P0/hEoomCWmmlotoyMaHZHqKjK676KvrFxsAAAAIAAAAABoAF0qAGgAbPwAaAErArYYaADVn4NgqWCgAAYHiA5IgIADN/HhSlubFty1GuizWeLHCGhv+h/URkQ8Mc6u3Jz8nGwAAAAgAAAAAGgAXSoAaABs/ABoASsCzgEA="},"Return":{"AllocationResults":{"SuccessCount":10,"FailCodes":null},"ExtensionResults":{"SuccessCount":0,"FailCodes":null},"NewAllocations":[96263896,96263897,96263898,96263899,96263900,96263901,96263902,96263903,96263904,96263905]}}`,
			height:    4878840,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := eg.GenerateVerifregEvents(context.Background(), []*types.Transaction{
				{
					TxBasicBlockData: types.TxBasicBlockData{
						BasicBlockData: types.BasicBlockData{
							Height: test.height,
						},
					},
					TxCid:         txCid,
					TxType:        test.txType,
					TxFrom:        test.txFrom,
					TxTo:          test.txTo,
					TxMetadata:    test.metadata,
					Status:        tools.GetExitCodeStatus(exitcode.Ok),
					SubcallStatus: tools.GetExitCodeStatus(exitcode.Ok),
				},
			}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)

		})
	}
}
