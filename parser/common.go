package parser

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/cache/impl"
	cacheMetrics "github.com/zondax/fil-parser/actors/cache/metrics"
	"github.com/zondax/fil-parser/types"
	golemBackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"
)

func parseMetadata(key string, metadata map[string]interface{}) string {
	params, ok := metadata[key].(string)
	if ok && params != "" {
		return params
	}
	jsonMetadata, err := json.Marshal(metadata[key])
	if err == nil && string(jsonMetadata) != "null" && string(jsonMetadata) != "\"\"" {
		return string(jsonMetadata)
	}
	return ""
}

func ParseParams(metadata map[string]interface{}) string {
	return parseMetadata(ParamsKey, metadata)
}

func ParseReturn(metadata map[string]interface{}) string {
	return parseMetadata(ReturnKey, metadata)
}

func GetTimestamp(timestamp uint64) time.Time {
	// #nosec G115
	blockTimeStamp := int64(timestamp) * 1000
	return time.Unix(blockTimeStamp/1000, blockTimeStamp%1000)
}

func AppendToAddressesMap(addressMap *types.AddressInfoMap, info ...*types.AddressInfo) {
	if addressMap == nil {
		return
	}

	for _, i := range info {
		switch i.ActorType {
		case manifest.EvmKey:
			// we store the address and later update it if it didn't have a creation_tx_cid or actor_cid
			cond := i.Robust != "" && i.Short != "" && i.Robust != i.Short
			if cond {
				prev, ok := addressMap.Get(i.Short)
				if ok {
					// this may happen because of direct storage from the evm parser
					if prev.CreationTxCid == "" || prev.ActorCid == "" {
						ok = false
					}
				}
				if !ok {
					addressMap.Set(i.Short, i)
				}
			}
		// miner & multisig: we skip any addressInfo without a CreationTxCid. The CreationTxCid is only obtained from parsing init.Exec Txs.
		case manifest.MultisigKey, manifest.MinerKey:
			// with multisig accounts we can skip checking for robust addresses because some
			// addresses do not have a robust address (genesis addresses)

			// we store the address and later update it if it didn't have a creation_tx_cid
			cond := i.Short != "" && i.Short != i.Robust && i.ActorCid != ""
			if i.IsSystemActor {
				cond = i.Short != "" && i.ActorCid != "" && i.ActorType != ""
			}
			if cond {
				prev, ok := addressMap.Get(i.Short)
				if ok {
					// this may happen because of direct storage from the miner/msig parser for diff. tx_types on the same address
					if prev.CreationTxCid == "" || prev.ActorCid == "" {
						ok = false
					}
				}
				if !ok {
					addressMap.Set(i.Short, i)
				}
			}
		default:
			cond := i.Robust != "" && i.Short != "" && i.Robust != i.Short
			if i.IsSystemActor {
				cond = i.Short != "" && i.ActorCid != "" && i.ActorType != ""
			}
			if cond {
				if _, ok := addressMap.Get(i.Short); !ok {
					addressMap.Set(i.Short, i)
				}
			}
		}
	}
}

func GetParentBaseFeeByHeight(tipset *types.ExtendedTipSet, logger *logger.Logger) (uint64, error) {
	defaultError := errors.New("could not find base fee")
	if tipset == nil {
		logger.Error("get-parent-base-fee: tipset is nil")
		return 0, defaultError
	}

	if len(tipset.TipSet.Blocks()) == 0 {
		logger.Error("get-parent-base-fee: no blocks found in the Tipset")
		return 0, defaultError
	}

	parentBaseFee := tipset.TipSet.Blocks()[0].ParentBaseFee
	return parentBaseFee.Uint64(), nil
}

func TranslateTxCidToTxHash(nodeClient api.FullNode, mainMsgCid cid.Cid, metrics *cacheMetrics.ActorsCacheMetricsClient, backoff *golemBackoff.BackOff) (string, error) {
	ctx := context.Background()

	nodeApiCallOptions := &impl.NodeApiCallWithRetryOptions[*ethtypes.EthHash]{
		RequestName: "EthGetTransactionHashByCid",
		BackOff:     *backoff,
		Request: func() (*ethtypes.EthHash, error) {
			ethHash, err := nodeClient.EthGetTransactionHashByCid(ctx, mainMsgCid)
			if err != nil || ethHash == nil {
				return nil, err
			}
			return ethHash, nil
		},
		RetryErrStrings: []string{"RPC client error"},
	}

	ethHash, err := impl.NodeApiCallWithRetry(nodeApiCallOptions, metrics)

	if err != nil || ethHash == nil {
		hash, err := ethtypes.EthHashFromCid(mainMsgCid)
		if err == nil {
			return hash.String(), nil
		}
		return "", err
	}

	return ethHash.String(), nil
}
