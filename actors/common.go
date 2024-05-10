package actors

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/lotus/api"
	"github.com/zondax/fil-parser/types"

	"github.com/filecoin-project/go-address"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/parser"
)

func (p *ActorParser) parseSend(msg *parser.LotusMessage) map[string]interface{} {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = msg.Params
	return metadata
}

// parseConstructor parse methods with format: *new(func(*address.Address) *abi.EmptyValue)
func (p *ActorParser) parseConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	return metadata, nil
}

func (p *ActorParser) unknownMetadata(msgParams, msgReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if len(msgParams) > 0 {
		metadata[parser.ParamsKey] = hex.EncodeToString(msgParams)
	}
	if len(msgReturn) > 0 {
		metadata[parser.ReturnKey] = hex.EncodeToString(msgReturn)
	}
	return metadata, nil
}

func (p *ActorParser) emptyParamsAndReturn() (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}

func ConsolidateRobustAddress(address address.Address, actorCache *cache.ActorsCache, logger *logger.Logger, config *parser.ConsolidateAddressesToRobust) (string, error) {
	var err error
	addressStr := address.String()
	if config != nil && config.Enable {
		if addressStr, err = EnsureRobustAddress(address, actorCache, logger); err != nil && !config.BestEffort {
			return "", err
		}
	}

	return addressStr, nil
}

func EnsureRobustAddress(address address.Address, actorCache *cache.ActorsCache, logger *logger.Logger) (string, error) {
	if isRobust, _ := common.IsRobustAddress(address); isRobust {
		return address.String(), nil
	}

	robustAddress, err := actorCache.GetRobustAddress(address)
	if err != nil {
		logger.Warnf("Error converting address to robust format: %v", err)
		return address.String(), fmt.Errorf("error converting address to robust format: %v", err) // Fallback
	}
	return robustAddress, nil
}

func CalculateTransactionFees(gasCost api.MsgGasCost, tipset *types.ExtendedTipSet, blockCid string, actorCache *cache.ActorsCache, logger *logger.Logger, config *parser.FilecoinParserConfig) []byte {
	minerAddressStr, err := tipset.GetBlockMiner(blockCid)
	if err == nil {
		minerAddress, err := address.NewFromString(minerAddressStr)
		if err != nil {
			logger.Errorf("Error when trying to parse miner address: %v", err)
		}

		minerAddressStr, err = ConsolidateRobustAddress(minerAddress, actorCache, logger, &config.ConsolidateAddressesToRobust)
		if err != nil {
			logger.Errorf("Error when trying to consolidate miner address to robust: %v", err)
		}
	} else {
		logger.Errorf("Error when trying to get miner address from block cid '%s': %v", blockCid, err)
	}

	feeData := parser.FeeData{
		FeesMetadata: parser.FeesMetadata{
			MinerFee: parser.MinerFee{
				MinerAddress: minerAddressStr,
				Amount:       gasCost.MinerTip.String(),
			},
			OverEstimationBurnFee: parser.OverEstimationBurnFee{
				BurnAddress: parser.BurnAddress,
				Amount:      gasCost.OverEstimationBurn.String(),
			},
			BurnFee: parser.BurnFee{
				BurnAddress: parser.BurnAddress,
				Amount:      gasCost.BaseFeeBurn.String(),
			},
		},
		Amount: gasCost.TotalCost.Int.String(),
	}

	data, err := json.Marshal(feeData)
	if err != nil {
		logger.Errorf("Error when trying to marshal fees data: %v", err)
	}

	return data
}
