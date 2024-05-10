package parser

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/api"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/types"
)

func GetExitCodeStatus(exitCode exitcode.ExitCode) string {
	code := exitCode.String()
	status := strings.Split(code, "(")
	if len(status) == 2 {
		return status[0]
	}
	return CheckExitCodeCommonError(code)
}

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
	blockTimeStamp := int64(timestamp) * 1000
	return time.Unix(blockTimeStamp/1000, blockTimeStamp%1000)
}

func AppendToAddressesMap(addressMap *types.AddressInfoMap, info ...*types.AddressInfo) {
	if addressMap == nil {
		return
	}

	for _, i := range info {
		switch i.ActorType {
		case manifest.MultisigKey:
			// with multisig accounts we can skip checking for robust addresses because some
			// addresses do not have a robust address (genesis addresses)
			if i.Short != "" {
				if _, ok := addressMap.Get(i.Short); !ok {
					addressMap.Set(i.Short, i)
				}
			}
		default:
			if i.Robust != "" && i.Short != "" && i.Robust != i.Short {
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

func TranslateTxCidToTxHash(nodeClient api.FullNode, mainMsgCid cid.Cid) (string, error) {
	ctx := context.Background()
	ethHash, err := nodeClient.EthGetTransactionHashByCid(ctx, mainMsgCid)
	if err != nil || ethHash == nil {
		return "", nil
	}

	return ethHash.String(), nil
}
