package deals_test

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
	"github.com/zondax/fil-parser/tools/deals"
	"github.com/zondax/fil-parser/tools/mocks"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

var (
	tipsetCid   = "bafy2bzaceaacubfrkq2wng5vcm53r7u5v5pjjs72s6kgqbm44dyfyik4y2azc"
	txCid       = "bafy2bzaceadjahvfnskpsz33ansyymo7uoqm7iyapx5cjpja3327pffvml7je"
	actorCidStr = "bafkqae3gnfwc6mjpon2g64tbm5sw2ylsnnsxi"
	txFrom      = "f2ncucfhaemzclocsau6o4u5b5bigjvpii2ql4rhi"
	txTo        = "f05"
	actorCid    cid.Cid
)

func init() {
	var err error
	actorCid, err = cid.Parse(actorCidStr)
	if err != nil {
		panic(err)
	}
}

func setupTest(_ *testing.T, network string) deals.EventGenerator {
	logger := logger.NewDevelopmentLogger()
	metrics := filMetrics.NewMetricsClient(metrics2.NewNoopMetrics())

	node := &mocks.FullNode{}
	node.On("StateNetworkName", mock.Anything).Return(dtypes.NetworkName("calibrationnet"), nil)
	node.On("StateNetworkVersion", mock.Anything, mock.Anything).Return(filApiTypes.NetworkVersion(16), nil)
	node.On("StateActorCodeCIDs", mock.Anything, mock.Anything).Return(map[string]cid.Cid{
		manifest.MarketKey: actorCid,
	}, nil)

	cache := &mocks.IActorsCache{}
	cache.On("StoreAddressInfo", mock.Anything).Return(nil)
	cache.On("GetActorCode", mock.Anything, mock.Anything, mock.Anything).Return(actorCidStr, nil)
	cache.On("GetActorNameFromAddress", mock.Anything, mock.Anything, mock.Anything).Return(manifest.MinerKey, nil)

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(node)
	helper := helper.NewHelper(lib, cache, node, logger, metrics)

	return deals.NewEventGenerator(helper, logger, metrics, network, parser.Config{})
}

func TestParseVerifyDealsForActivation(t *testing.T) {
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
			name:      "NV3",
			txType:    parser.MethodVerifyDealsForActivation,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"5","Params":{"DealIDs":null,"SectorExpiry":1644975,"SectorStart":94001},"Return":{"DealWeight":"0","VerifiedDealWeight":"0"}}`,
			height:    94000,
		},
		{
			name:      "NV10",
			txType:    parser.MethodVerifyDealsForActivation,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"5","Params":{"Sectors":[{"SectorExpiry":2097712,"DealIDs":[1596539]}]},"Return":{"Sectors":[{"DealSpace":34359738368,"DealWeight":"0","VerifiedDealWeight":"52446704644915200"}]}}`,
			height:    550350,
		},
		{
			name:      "NV15",
			txType:    parser.MethodVerifyDealsForActivation,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"5","Params":{"Sectors":[{"SectorExpiry":3149856,"DealIDs":[4137277]}]},"Return":{"Sectors":[{"DealSpace":34359738368,"DealWeight":"0","VerifiedDealWeight":"52941484877414400"}]}}`,
			height:    1594681,
		},
		{
			name:      "NV17",
			txType:    parser.MethodVerifyDealsForActivation,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"5","Params":{"Sectors":[{"SectorType":8,"SectorExpiry":3938176,"DealIDs":[17514444]}]},"Return":{"Sectors":[{"CommD":{"/":"baga6ea4seaqcvxnumewlmsy3dgd6fzqzkmx6gwg2hrlea3mfsgsp342hmqausea"}}]}}`,
			height:    2383682,
		},
		{
			name:      "NV18",
			txType:    parser.MethodVerifyDealsForActivation,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"5","Params":{"Sectors":[{"SectorType":9,"SectorExpiry":3752099,"DealIDs":[28436872,28436293]}]},"Return":{"Sectors":[{"CommD":{"/":"baga6ea4seaqezng67jc7t2mowage34tt4emcd4w7d5jnpusv5obmpexsq335ihy"}}]}}`,
			height:    2683349,
		},
		{
			name:      "NV19",
			txType:    parser.MethodVerifyDealsForActivation,
			actorName: manifest.VerifregKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"5","Params":{"Sectors":[{"SectorType":8,"SectorExpiry":4364609,"DealIDs":[34223220]}]},"Return":{"Sectors":[{"CommD":{"/":"baga6ea4seaqbro6cy6d6hp7ywf3ktxt3hygi376cannpbint5sqeih6nintqioq"}}]}}`,
			height:    2809801,
		},
	}

	for _, test := range tests {
		t.Run(test.txType, func(t *testing.T) {
			_, err := eg.GenerateDealsEvents(context.Background(), []*types.Transaction{
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
