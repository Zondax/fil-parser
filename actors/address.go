package actors

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"go.uber.org/zap"
)

func ConsolidateRobustAddress(address address.Address, actorCache *cache.ActorsCache, logger *zap.Logger) (string, error) {
	if isRobust, _ := common.IsRobustAddress(address); isRobust {
		return address.String(), nil
	}

	robustAddress, err := actorCache.GetRobustAddress(address)
	if err != nil {
		logger.Sugar().Warnf("Error converting address %s to robust format: %v", address, err)
		return address.String(), fmt.Errorf("error converting address to robust format: %v", err) // Fallback
	}

	return robustAddress, nil
}
