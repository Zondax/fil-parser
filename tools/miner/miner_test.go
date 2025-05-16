package miner_test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	filApiTypes "github.com/filecoin-project/lotus/api/types"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	filMetrics "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools/miner"
	"github.com/zondax/fil-parser/tools/mocks"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

var (
	tipsetCid   = "bafy2bzaceczpzd5k7u6hwaim7fdpwx2ujg7uhrdbpijf7q5ryvh7ogmawxupk"
	txCid       = "bafy2bzaceczpzd5k7u6hwaim7fdpwx2ujg7uhrdbpijf7q5ryvh7ogmawxupk"
	actorCidStr = "bafk2bzacea6rabflc7kpwr6y4lzcqsnuahr4zblyq3rhzrrsfceeiw2lufrb4"
	txFrom      = "f02"
	txTo        = "f2ddsjma6hfwcqhdp4vv6z4t5fighlhrjrqyxcekq"
	actorCid    cid.Cid
)

func setupTest(t *testing.T) miner.EventGenerator {
	logger := logger.NewDevelopmentLogger()
	metrics := filMetrics.NewMetricsClient(metrics2.NewNoopMetrics())

	var err error
	actorCid, err = cid.Parse(actorCidStr)
	require.NoError(t, err)

	node := &mocks.FullNode{}
	node.On("StateNetworkName", mock.Anything).Return(dtypes.NetworkName("calibrationnet"), nil)
	node.On("StateNetworkVersion", mock.Anything, mock.Anything).Return(filApiTypes.NetworkVersion(16), nil)
	node.On("StateActorCodeCIDs", mock.Anything, mock.Anything).Return(map[string]cid.Cid{
		manifest.MinerKey: actorCid,
	}, nil)

	cache := &mocks.IActorsCache{}
	cache.On("StoreAddressInfo", mock.Anything).Return(nil)
	cache.On("GetActorCode", mock.Anything, mock.Anything, mock.Anything).Return(actorCidStr, nil)
	cache.On("GetActorNameFromAddress", mock.Anything, mock.Anything, mock.Anything).Return(manifest.MinerKey, nil)

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(node)
	helper := helper.NewHelper(lib, cache, node, logger, metrics)

	return miner.NewEventGenerator(helper, logger, metrics)
}

func assertSectorEvents(t *testing.T, want []*types.MinerSectorEvent, got []*types.MinerSectorEvent) {
	require.Equal(t, len(want), len(got))

	for i, sector := range got {
		assert.Equal(t, want[i].ActionType, sector.ActionType)
		assert.Equalf(t, want[i].SectorNumber, sector.SectorNumber, "sector number mismatch for index %d, want: %d, got: %d", i, want[i].SectorNumber, sector.SectorNumber)
		assert.Equal(t, want[i].MinerAddress, sector.MinerAddress)
		assert.Equal(t, want[i].TxCid, sector.TxCid)
	}
}

func getSectorEvents(_ *testing.T, txType, minerAddress, txCid string, sectorNumbers ...uint64) []*types.MinerSectorEvent {
	sectorEvents := []*types.MinerSectorEvent{}
	for _, sectorNumber := range sectorNumbers {
		sectorEvents = append(sectorEvents, &types.MinerSectorEvent{
			SectorNumber: sectorNumber,
			MinerAddress: minerAddress,
			TxCid:        txCid,
			ActionType:   txType,
		})
	}
	return sectorEvents
}

func gen(from, to uint64) []uint64 {
	arr := []uint64{}
	for i := from; i <= to; i++ {
		arr = append(arr, i)
	}
	return arr
}

func TestMinerInfo_AwardBlockReward(t *testing.T) {
	eg := setupTest(t)

	tests := []struct {
		name      string
		txType    string
		actorName string
		txFrom    string
		txTo      string
		metadata  string
		want      *types.MinerEvents
	}{

		{
			name:      "Award Block Reward",
			txType:    parser.MethodAwardBlockReward,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"2","Params":{"Miner":"f01000","Penalty":"0","GasReward":"0","WinCount":3}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodAwardBlockReward, txTo, txCid, 1389),
				MinerInfo: []*types.MinerInfo{
					{
						ActorAddress: "f01000",
						TxCid:        txCid,
						ActionType:   parser.MethodAwardBlockReward,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.txType, func(t *testing.T) {
			events, err := eg.GenerateMinerEvents(context.Background(), []*types.Transaction{
				{
					TxCid:      txCid,
					TxType:     test.txType,
					TxFrom:     test.txFrom,
					TxTo:       test.txTo,
					TxMetadata: test.metadata,
					Status:     "Ok",
				},
			}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)

			for i, event := range events.MinerInfo {
				assert.Equal(t, test.want.MinerInfo[i].ActorAddress, event.ActorAddress)
				assert.Equal(t, test.want.MinerInfo[i].TxCid, event.TxCid)
				assert.Equal(t, test.want.MinerInfo[i].ActionType, event.ActionType)

			}
		})
	}
}

func TestMinerSectors_PreCommitStage(t *testing.T) {
	eg := setupTest(t)

	tests := []struct {
		name      string
		txType    string
		actorName string
		txFrom    string
		txTo      string
		metadata  string
		want      *types.MinerEvents
	}{

		{
			name:      "Pre Commit Sector",
			txType:    parser.MethodPreCommitSector,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"6","Params":{"SealProof":8,"SectorNumber":1389,"SealedCID":{"/":"bagboea4b5abcapw3cbg4xe4cq7enh7uouusqsd3c3oiki34vjyjinp25vxpj3hlk"},"SealRandEpoch":517402,"DealIDs":[67798],"Expiration":1126132,"ReplaceCapacity":false,"ReplaceSectorDeadline":0,"ReplaceSectorPartition":0,"ReplaceSectorNumber":0}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodPreCommitSector, txTo, txCid, 1389),
			},
		},
		{
			name:      "Pre Commit Sector Batch",
			txType:    parser.MethodPreCommitSectorBatch,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"25","Params":{"Sectors":[{"SealProof":8,"SectorNumber":1110,"SealedCID":{"/":"bagboea4b5abcbg5ftov3qjqvocmymtw6fpih4tbek2q7as46lbmk6n2xbhsswjsx"},"SealRandEpoch":260089,"DealIDs":[22023],"Expiration":1138519,"ReplaceCapacity":false,"ReplaceSectorDeadline":0,"ReplaceSectorPartition":0,"ReplaceSectorNumber":0},{"SealProof":8,"SectorNumber":1111,"SealedCID":{"/":"bagboea4b5abcbhhxyskvadcful3qgiuhjacibww5dtv7earyi47gahze45oy4vkw"},"SealRandEpoch":260089,"DealIDs":[22015],"Expiration":1138515,"ReplaceCapacity":false,"ReplaceSectorDeadline":0,"ReplaceSectorPartition":0,"ReplaceSectorNumber":0}]}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodPreCommitSectorBatch, txTo, txCid, 1110, 1111),
			},
		},
		{
			name:      "Pre Commit Sector Batch 2",
			txType:    parser.MethodPreCommitSectorBatch2,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"28","Params":{"Sectors":[{"SealProof":8,"SectorNumber":28,"SealedCID":{"/":"bagboea4b5abcbrhwah4llrepf3nvptyiqoh7s57bszcjfpqt3u3ddm7rwavv3lzx"},"SealRandEpoch":1416388,"DealIDs":null,"Expiration":2025118,"UnsealedCid":null}]}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodPreCommitSectorBatch2, txTo, txCid, 28),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.txType, func(t *testing.T) {
			events, err := eg.GenerateMinerEvents(context.Background(), []*types.Transaction{
				{
					TxCid:      txCid,
					TxType:     test.txType,
					TxFrom:     test.txFrom,
					TxTo:       test.txTo,
					TxMetadata: test.metadata,
					Status:     "Ok",
				},
			}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)

			assertSectorEvents(t, test.want.MinerSectors, events.MinerSectors)
		})
	}
}

func TestMinerSectors_ProveCommitStage(t *testing.T) {
	eg := setupTest(t)

	tests := []struct {
		name      string
		txType    string
		actorName string
		txFrom    string
		txTo      string
		metadata  string
		want      *types.MinerEvents
	}{
		{
			name:      "Prove Commit Aggregate",
			txType:    parser.MethodProveCommitAggregate,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"26","Params":{"SectorNumbers":[0,4],"AggregateProof":""}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodProveCommitAggregate, txTo, txCid, 0, 1, 2, 3),
			},
		},
		{
			name:      "Prove Commit Sector",
			txType:    parser.MethodProveCommitSector,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"Params":{"SectorNumber":428,"Proof":""}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodProveCommitSector, txTo, txCid, 428),
			},
		},
		{
			name:      "Confirm Sector Proofs Valid",
			txType:    parser.MethodConfirmSectorProofsValid,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"Params":{"Sectors":[435],"RewardSmoothed":{"PositionEstimate":"10356936834322900687769035838261409832064265323070766597037","VelocityEstimate":"-1137632398727324746087094268522211807413829360717816"},"RewardBaselinePower":"8371919687512492146","QualityAdjPowerSmoothed":{"PositionEstimate":"1046503062216955445726624465573726475809608663785146418","VelocityEstimate":"-31812310398129809759780491270823543363211012668"}}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodConfirmSectorProofsValid, txTo, txCid, 435),
			},
		},
		{
			name:      "Prove Commit Sectors NI",
			txType:    parser.MethodProveCommitSectorsNI,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"36","Params":{"Sectors":[{"SealingNumber":1,"SealerID":132414,"SealedCID":{"/":"bagboea4b5abcayrsf5tv5ea7nq6o3jjesdrqddzy3u5jbxul6azth4ndlvsj3qqe"},"SectorNumber":1,"SealRandEpoch":1784699,"Expiration":3342710},{"SealingNumber":2,"SealerID":132414,"SealedCID":{"/":"bagboea4b5abcakafd36kxd4yea75jkou34ajvszgkk6vmw7ja7c23us3p3iacjks"},"SectorNumber":2,"SealRandEpoch":1784699,"Expiration":3342710},{"SealingNumber":3,"SealerID":132414,"SealedCID":{"/":"bagboea4b5abcafgjxnnzispjss2hcp4nij5zu4vur3tpj7oou7mztxk4dgz6ivik"},"SectorNumber":3,"SealRandEpoch":1784699,"Expiration":3342710},{"SealingNumber":4,"SealerID":132414,"SealedCID":{"/":"bagboea4b5abcasvbeiz2bvv4lv5zu3hpazu6yf6wfacwr2qso3mw5sinnrcxtdqg"},"SectorNumber":4,"SealRandEpoch":1784699,"Expiration":3342710}],"AggregateProof":"","SealProofType":18,"AggregateProofType":1,"ProvingDeadline":27,"RequireActivationSuccess":true}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodProveCommitSectorsNI, txTo, txCid, 1, 2, 3, 4),
			},
		},
		{
			name:      "Prove Commit Sectors 3",
			txType:    parser.MethodProveCommitSectors3,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"34","Params":{"SectorActivations":[{"SectorNumber":3282,"Pieces":null}],"SectorProofs":[""],"AggregateProof":null,"AggregateProofType":null,"RequireActivationSuccess":false,"RequireNotificationSuccess":false},"Return":{"SuccessCount":1,"FailCodes":null}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodProveCommitSectors3, txTo, txCid, 3282),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.txType, func(t *testing.T) {
			events, err := eg.GenerateMinerEvents(context.Background(), []*types.Transaction{
				{
					TxCid:      txCid,
					TxType:     test.txType,
					TxFrom:     test.txFrom,
					TxTo:       test.txTo,
					TxMetadata: test.metadata,
					Status:     "Ok",
				},
			}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)

			assertSectorEvents(t, test.want.MinerSectors, events.MinerSectors)
		})
	}
}

func TestMinerSectors_TerminationFaultAndRecoveries(t *testing.T) {
	eg := setupTest(t)

	tests := []struct {
		name      string
		txType    string
		actorName string
		txFrom    string
		txTo      string
		metadata  string
		want      *types.MinerEvents
	}{

		{
			name:      "Terminate Sectors",
			txType:    parser.MethodTerminateSectors,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"9","Params":{"Terminations":[{"Deadline":0,"Partition":0,"Sectors":[36,1]}]},"Return":{"Done":true}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodTerminateSectors, txTo, txCid, 36),
			},
		},
		{
			name:      "Declare Faults",
			txType:    parser.MethodDeclareFaults,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"10","Params":{"Faults":[{"Deadline":0,"Partition":0,"Sectors":[1,17]}]}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodDeclareFaults, txTo, txCid, gen(1, 17)...),
			},
		},
		{
			name:      "Declare Faults Recovered",
			txType:    parser.MethodDeclareFaultsRecovered,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"11","Params":{"Recoveries":[{"Deadline":2,"Partition":0,"Sectors":[9,1,3,2]}]}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodDeclareFaultsRecovered, txTo, txCid, 9, 13, 14),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.txType, func(t *testing.T) {
			events, err := eg.GenerateMinerEvents(context.Background(), []*types.Transaction{
				{
					TxCid:      txCid,
					TxType:     test.txType,
					TxFrom:     test.txFrom,
					TxTo:       test.txTo,
					TxMetadata: test.metadata,
					Status:     "Ok",
				},
			}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)

			assertSectorEvents(t, test.want.MinerSectors, events.MinerSectors)
		})
	}
}

func TestMinerSectors_ExpiryExtension(t *testing.T) {
	eg := setupTest(t)

	tests := []struct {
		name      string
		txType    string
		actorName string
		txFrom    string
		txTo      string
		metadata  string
		want      *types.MinerEvents
	}{

		{
			name:      "Extend Sector Expiration",
			txType:    parser.MethodExtendSectorExpiration,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"8","Params":{"Extensions":[{"Deadline":0,"Partition":0,"Sectors":[12,1],"NewExpiration":1581493}]}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodExtendSectorExpiration, txTo, txCid, 12),
			},
		},
		{
			name:      "Extend Sector Expiration 2",
			txType:    parser.MethodExtendSectorExpiration2,
			actorName: manifest.MinerKey,
			txFrom:    txFrom,
			txTo:      txTo,
			metadata:  `{"MethodNum":"32","Params":{"Extensions":[{"Deadline":0,"Partition":0,"Sectors":[10,1],"SectorsWithClaims":null,"NewExpiration":1162719}]}}`,
			want: &types.MinerEvents{
				MinerSectors: getSectorEvents(t, parser.MethodExtendSectorExpiration2, txTo, txCid, 10),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.txType, func(t *testing.T) {
			events, err := eg.GenerateMinerEvents(context.Background(), []*types.Transaction{
				{
					TxCid:      txCid,
					TxType:     test.txType,
					TxFrom:     test.txFrom,
					TxTo:       test.txTo,
					TxMetadata: test.metadata,
					Status:     "Ok",
				},
			}, tipsetCid, filTypes.EmptyTSK)
			require.NoError(t, err)

			assertSectorEvents(t, test.want.MinerSectors, events.MinerSectors)
		})
	}
}
