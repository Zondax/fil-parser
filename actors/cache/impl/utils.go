package impl

import (
	"strings"
	"time"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/golem/pkg/zhttpclient/backoff"
)

const (
	Short2CidMapPrefix        = "short2Cid"
	Robust2ShortMapPrefix     = "robust2Short"
	Short2RobustMapPrefix     = "short2Robust"
	SelectorHash2SigMapPrefix = "hash2Sig"
)

func stateLookupWithRetry[T address.Address | *filTypes.Actor](maxAttempts int, maxWaitBeforeRetry time.Duration, request func() (T, error)) (T, error) {
	// try without backoff
	result, err := request()
	if err != nil {
		if !strings.Contains(err.Error(), "RPC client error") {
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
