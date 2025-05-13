package impl

import (
	"strings"
	"time"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/zondax/golem/pkg/zhttpclient/backoff"
)

const (
	Short2CidMapPrefix        = "short2Cid"
	Robust2ShortMapPrefix     = "robust2Short"
	Short2RobustMapPrefix     = "short2Robust"
	SelectorHash2SigMapPrefix = "hash2Sig"
)

// NodeApiCallWithRetry makes an API call with automatic retries for specific errors.
// Parameters:
//   - errStrings: list of error strings that SHOULD trigger a retry
//   - maxAttempts: maximum number of retry attempts
//   - maxWaitBeforeRetry: maximum duration to wait before retrying
//   - request: the function that makes the actual API call
//
// Returns the result of the API call and any error encountered.
func NodeApiCallWithRetry[T address.Address | *filTypes.Actor | *ethtypes.EthHash](errStrings []string, maxAttempts int, maxWaitBeforeRetry time.Duration, request func() (T, error)) (T, error) {
	// try without backoff
	result, err := request()
	if err != nil {
		shouldRetry := false
		for _, errString := range errStrings {
			if strings.Contains(err.Error(), errString) {
				shouldRetry = true
				break
			}
		}
		if !shouldRetry {
			return result, err
		}
	} else {
		return result, nil
	}

	b := backoff.New().
		WithMaxAttempts(maxAttempts).
		WithInitialDuration(maxWaitBeforeRetry).
		WithMaxDuration(maxWaitBeforeRetry).
		Linear()

	err = backoff.Do(func() error {
		result, err = request()
		if err != nil {
			return err
		}
		return nil
	}, b)

	return result, err
}
