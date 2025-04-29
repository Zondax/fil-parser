package actors

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/golem/pkg/logger"
)

func ConsolidateRobustAddress(address address.Address, actorCache *cache.ActorsCache, logger *logger.Logger, bestEffort bool) (string, error) {
	if isRobust, _ := common.IsRobustAddress(address); isRobust {
		return address.String(), nil
	}

	robustAddress, err := actorCache.GetRobustAddress(address)
	if err != nil && !bestEffort {
		logger.Warnf("Error converting address %s to robust format: %v", address, err)
		return "", fmt.Errorf("error converting address to robust format: %v", err) // Fallback
	}

	return robustAddress, nil
}
