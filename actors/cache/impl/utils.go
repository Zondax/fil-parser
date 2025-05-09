package impl

import (
	"strings"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/zondax/golem/pkg/zhttpclient/backoff"
)

const (
	Short2CidMapPrefix        = "short2Cid"
	Robust2ShortMapPrefix     = "robust2Short"
	Short2RobustMapPrefix     = "short2Robust"
	SelectorHash2SigMapPrefix = "hash2Sig"
)

func stateLookupWithRetry(maxAttempts int, maxWaitBeforeRetry time.Duration, request func() (address.Address, error)) (address.Address, error) {
	var address address.Address
	var err error

	// try without backoff
	address, err = request()
	if err != nil {
		if !strings.Contains(err.Error(), "RPC client error") {
			return address, err
		}
	}

	b := backoff.New().
		WithMaxAttempts(maxAttempts).
		WithInitialDuration(maxWaitBeforeRetry).
		WithMaxDuration(maxWaitBeforeRetry).
		Linear()

	err = backoff.Do(func() error {
		address, err = request()
		if err != nil {
			return err
		}
		return nil
	}, b)

	return address, err
}
