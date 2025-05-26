package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

var l = logger.NewDevelopmentLogger()

type RPCClient struct {
	url          string
	token        string
	client       api.FullNode
	clientCloser *jsonrpc.ClientCloser
	ctx          context.Context
	rosettaLib   *rosettaFilecoinLib.RosettaConstructionFilecoin
	nodeInfo     types.NodeInfo
}

type RawData struct {
	Tipset         *types.ExtendedTipSet
	Trace          *api.ComputeStateOutput
	EthLogs        []types.EthLog
	NativeLogs     []*lotusChainTypes.ActorEvent
	TipsetMetadata types.BlockMetadata
}

// NewFilecoinRPCClient creates a new blockchain RPC remoteNode
func newFilecoinRPCClient(url string, token string) (*RPCClient, error) {
	ctx := context.Background()

	headers := http.Header{}
	if len(token) > 0 {
		headers.Add("Authorization", "Bearer "+token)
	}

	lotusAPI, closer, err := client.NewFullNodeRPCV1(ctx, url, headers)

	if err != nil {
		return nil, err
	}

	// Setup rosetta lib
	r := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusAPI)
	if r == nil {
		return nil, fmt.Errorf("could not create instance of rosetta filecoin-lib")
	}

	// Get node version
	nodeFullVersion, err := lotusAPI.Version(ctx)
	if err != nil {
		l.Error(fmt.Sprintf("Error getting node version: %s", err))
		return nil, err
	}
	nodeInfo, err := processNodeVersion(nodeFullVersion.Version)
	if err != nil {
		l.Error(fmt.Sprintf("Error processing node version: %s", err))
		return nil, err
	}

	return &RPCClient{
		url:          url,
		token:        token,
		client:       lotusAPI,
		clientCloser: &closer,
		ctx:          ctx,
		rosettaLib:   r,
		nodeInfo:     *nodeInfo,
	}, nil

}

func getTraceFileByHeight(height uint64, lotusClient api.FullNode) (*api.ComputeStateOutput, error) {
	// #nosec G115
	tipset, err := lotusClient.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(height), lotusChainTypes.EmptyTSK)
	if err != nil {
		return nil, err
	}

	// Check that the retrieved tipset is not empty nor invalid
	// #nosec G115
	if tipset == nil || uint64(tipset.Height()) != height {
		l.Infof("no tipset data received for the specified height: %d", height)
		return nil, nil
	}

	// #nosec G115
	traces, err := lotusClient.StateCompute(context.Background(), abi.ChainEpoch(height), nil, tipset.Key())
	if err != nil {
		return nil, fmt.Errorf("error retrieving traces for tipset %d: %+v", height, err)
	}

	if traces == nil {
		return nil, fmt.Errorf("nil trace received for tipset height: %d", height)
	}

	return traces, nil
}

func getTipsetFileByHeight(height uint64, key lotusChainTypes.TipSetKey, lotusClient api.FullNode) (*types.ExtendedTipSet, error) {
	// #nosec G115
	chainEpoch := abi.ChainEpoch(height)
	tipset, err := lotusClient.ChainGetTipSetByHeight(context.Background(), chainEpoch, key)
	if err != nil {
		// Try using empty key
		time.Sleep(time.Second * 3)
		tipset, err = lotusClient.ChainGetTipSetByHeight(context.Background(), chainEpoch, lotusChainTypes.EmptyTSK)
		if err != nil {
			return nil, err
		}
	}

	// Check that the retrieved tipset is valid
	if tipset == nil || tipset.Height() != chainEpoch {
		l.Infof("no tipset data received for the specified height: %d", height)
		return nil, nil
	}

	// Get messages CIDs stored on each block
	extendedTipset, err := fetchBlockMessagesCids(tipset, lotusClient)
	if err != nil {
		l.Error(err.Error())
		return nil, err
	}

	return extendedTipset, nil
}

func getNativeLogsByHeight(height uint64, lotusClient api.FullNode) ([]*lotusChainTypes.ActorEvent, error) {
	// #nosec G115
	currentChainEpoch := abi.ChainEpoch(height)
	res, err := lotusClient.GetActorEventsRaw(context.Background(), &lotusChainTypes.ActorEventFilter{
		FromHeight: &currentChainEpoch,
		ToHeight:   &currentChainEpoch,
	})

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return res, nil
}

func getEthLogsByHeight(height uint64, lotusClient api.FullNode) ([]types.EthLog, error) {
	fromBlockHex := "0x" + strconv.FormatUint(height, 16)
	res, err := lotusClient.EthGetLogs(context.Background(), &ethtypes.EthFilterSpec{
		FromBlock: &fromBlockHex,
		ToBlock:   &fromBlockHex,
	})

	if err != nil {
		return nil, err
	}

	if len(res.Results) == 0 {
		return nil, nil
	}

	logs := make([]types.EthLog, 0, len(res.Results))
	for _, result := range res.Results {
		var log types.EthLog
		resultJson, _ := json.Marshal(result)
		err = json.Unmarshal(resultJson, &log)
		if err != nil {
			return nil, fmt.Errorf("eth logs are not of the expected type 'EthLogs'")
		}
		// Get the ethHash <-> filCID mapping
		txCid, err := getCIDFromEthHashMappings(log.TransactionHash.String(), lotusClient)
		if err != nil {
			l.Errorf("Could not get filCid from ethHash. Height %d, hash %s", height, log.TransactionHash.String())
		} else {
			log.TransactionCid = txCid
		}

		logs = append(logs, log)
	}

	return logs, nil
}

func getMetadata(rpcClient *RPCClient) (types.BlockMetadata, error) {
	return types.BlockMetadata{
		NodeInfo: types.NodeInfo{
			NodeFullVersion:       rpcClient.nodeInfo.NodeFullVersion,
			NodeMajorMinorVersion: rpcClient.nodeInfo.NodeMajorMinorVersion,
		},
	}, nil
}

func fetchBlockMessagesCids(tipset *lotusChainTypes.TipSet, lotusClient api.FullNode) (*types.ExtendedTipSet, error) {
	extendedTipset := types.ExtendedTipSet{
		TipSet:        *tipset,
		BlockMessages: make(types.BlockMessages),
	}

	for _, header := range tipset.Blocks() {
		blockMessages, err := lotusClient.ChainGetBlockMessages(context.Background(), header.Cid())
		if err != nil {
			l.Errorf("error while calling lotus 'ChainGetBlockMessages': %v", err)
			continue
		}

		for _, msgCid := range blockMessages.Cids {
			if _, ok := extendedTipset.BlockMessages[msgCid.String()]; !ok {
				extendedTipset.BlockMessages[msgCid.String()] = make([]types.LightBlockHeader, 0, len(tipset.Blocks()))
			}
			extendedTipset.BlockMessages[msgCid.String()] = append(extendedTipset.BlockMessages[msgCid.String()], types.LightBlockHeader{
				Cid:        header.Cid().String(),
				BlockMiner: header.Miner.String(),
			})
		}
	}

	return &extendedTipset, nil
}

func getCIDFromEthHashMappings(hash string, lotusClient api.FullNode) (string, error) {
	ethHash, err := ethtypes.ParseEthHash(hash)
	if err != nil {
		l.Errorf("Error while trying to parse ethHash '%v'", hash)
		return "", err
	}

	txCid, err := lotusClient.EthGetMessageCidByTransactionHash(context.Background(), &ethHash)
	if err != nil || txCid == nil {
		l.Errorf("Error while trying to get hash mapping from node: %s", err)
		return "", err
	}

	return txCid.String(), nil
}
