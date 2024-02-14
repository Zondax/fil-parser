package actors

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/parser"
	"go.uber.org/zap"
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

func ConsolidateRobustAddresses(from, to address.Address, actorCache *cache.ActorsCache, logger *zap.Logger, config *parser.ConsolidateAddressesToRobust) (string, string, error) {
	var err error
	txFrom := from.String()
	txTo := to.String()
	if config != nil && config.Enable {
		if txFrom, err = EnsureRobustAddress(from, actorCache, logger); err != nil && !config.BestEffort {
			return "", "", err
		}
		if txTo, err = EnsureRobustAddress(to, actorCache, logger); err != nil && !config.BestEffort {
			return "", "", err
		}
	}

	return txFrom, txTo, nil
}

func EnsureRobustAddress(address address.Address, actorCache *cache.ActorsCache, logger *zap.Logger) (string, error) {
	if isRobust, _ := common.IsRobustAddress(address); isRobust {
		return address.String(), nil
	}

	robustAddress, err := actorCache.GetRobustAddress(address)
	if err != nil {
		logger.Sugar().Warnf("Error converting address to robust format: %v", err)
		return address.String(), fmt.Errorf("error converting address to robust format: %v", err) // Fallback
	}
	return robustAddress, nil
}
