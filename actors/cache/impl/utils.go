package impl

import (
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	cacheMetrics "github.com/zondax/fil-parser/actors/cache/metrics"
	zbackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"
)

const (
	Short2CidMapPrefix        = "short2Cid"
	Robust2ShortMapPrefix     = "robust2Short"
	Short2RobustMapPrefix     = "short2Robust"
	SelectorHash2SigMapPrefix = "hash2Sig"
)

const (
	isRetry        = true
	isNotRetry     = false
	isRetriable    = true
	isNotRetriable = false
	isSuccess      = true
	isNotSuccess   = false
)

type NodeApiResponse interface {
	address.Address | *filTypes.Actor | *ethtypes.EthHash
}

type NodeApiCallWithRetryOptions[T NodeApiResponse] struct {
	BackOff         backoff.BackOff
	RequestName     string
	Request         func() (T, error)
	RetryErrStrings []string
}

// NodeApiCallWithRetry makes an API call with automatic retries for specific errors.
// Parameters:
//   - errStrings: list of error strings that SHOULD trigger a retry
//   - maxAttempts: maximum number of retry attempts
//   - maxWaitBeforeRetry: maximum duration to wait before retrying
//   - request: the function that makes the actual API call
//
// Returns the result of the API call and any error encountered.
func NodeApiCallWithRetry[T NodeApiResponse](options *NodeApiCallWithRetryOptions[T], metrics *cacheMetrics.ActorsCacheMetricsClient) (T, error) {
	errStrings := options.RetryErrStrings

	// time the request
	request := func() (T, error) {
		start := time.Now()
		result, err := options.Request()
		latency := time.Since(start)
		_ = metrics.UpdateNodeApiCallLatencyMetric(options.RequestName, err == nil, latency)
		return result, err
	}

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
			// update failure without a retry
			_ = metrics.UpdateNodeApiCallMetric(options.RequestName, isNotSuccess, isNotRetry, isNotRetriable)
			return result, err
		}
	} else {
		// update successful call without a retry
		_ = metrics.UpdateNodeApiCallMetric(options.RequestName, isSuccess, isNotRetry, isNotRetriable)
		return result, nil
	}

	_ = metrics.UpdateNodeApiCallMetric(options.RequestName, isNotSuccess, isNotRetry, isRetriable)

	err = zbackoff.Do(func() error {
		result, err = request()
		if err != nil {
			// update failed retries
			_ = metrics.UpdateNodeApiCallMetric(options.RequestName, isNotSuccess, isRetry, isRetriable)
			return err
		}
		return nil
	}, options.BackOff)

	if err != nil {
		// update failure after retry attempts
		_ = metrics.UpdateNodeApiCallMetric(options.RequestName, isNotSuccess, isRetry, isNotRetriable)
	} else {
		// update success after retry attempts
		_ = metrics.UpdateNodeApiCallMetric(options.RequestName, isSuccess, isRetry, isNotRetriable)
	}

	return result, err
}
