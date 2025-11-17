package actors

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/logger"
)

// ConsolidateToRobustAddress consolidates an address to a robust address
// if the address is a zero address account actor, it returns the robust address of the zero address account actor
// if the address is already a robust address, it returns the address
// if the address is f2 evm, we consolidate f2 -> f0 -> f4
func ConsolidateToRobustAddress(addr address.Address, h *helper.Helper, logger *logger.Logger, bestEffort bool, canonical bool) (string, error) {
	actorCache := h.GetActorsCache()
	if ok, _, _ := h.IsZeroAddressAccountActor(addr); ok {
		return helper.ZeroAddressAccountActorRobust, nil
	}

	if isRobust, _ := common.IsRobustAddress(addr); isRobust {
		// we need to handle cases where a f2 address for evm actors is used
		// f2 -> f0 -> f4, as we want to consolidate the address to f4 style
		shortAddressStr, err := actorCache.GetShortAddress(addr, canonical)
		if err == nil {
			shortAddress, _ := address.NewFromString(shortAddressStr)
			addrStr, err := actorCache.GetRobustAddress(shortAddress, canonical)
			if err == nil {
				addr, _ = address.NewFromString(addrStr)
			}
		}
		return addr.String(), nil
	}

	robustAddress, err := actorCache.GetRobustAddress(addr, canonical)
	if err != nil && !bestEffort {
		logger.Warnf("Error converting address %s to robust format: %v", addr, err)
		return "", fmt.Errorf("error converting address to robust format: %v", err) // Fallback
	}

	return robustAddress, nil
}
