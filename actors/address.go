package actors

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/logger"
)

func ConsolidateRobustAddress(addr address.Address, h *helper.Helper, logger *logger.Logger, bestEffort bool) (string, error) {
	actorCache := h.GetActorsCache()
	isZeroAddressAccountActor, _, _ := h.IsZeroAddressAccountActor(addr)

	if isRobust, _ := common.IsRobustAddress(addr); isRobust {
		if isZeroAddressAccountActor {
			return helper.ZeroAddressAccountActorShort, nil
		}
		// we need to handle cases where a f2 address for evm actors is used
		// f2 -> f0 -> f4, as we want to consolidate the address to f4 style
		shortAddressStr, err := actorCache.GetShortAddress(addr)
		if err == nil {
			shortAddress, _ := address.NewFromString(shortAddressStr)
			addrStr, err := actorCache.GetRobustAddress(shortAddress)
			if err == nil {
				addr, _ = address.NewFromString(addrStr)
			}
		}
		return addr.String(), nil
	}

	if isZeroAddressAccountActor {
		return helper.ZeroAddressAccountActorRobust, nil
	}
	robustAddress, err := actorCache.GetRobustAddress(addr)
	if err != nil && !bestEffort {
		logger.Warnf("Error converting address %s to robust format: %v", addr, err)
		return "", fmt.Errorf("error converting address to robust format: %v", err) // Fallback
	}

	return robustAddress, nil
}
