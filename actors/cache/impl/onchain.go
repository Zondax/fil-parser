package impl

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	parserContext "github.com/zondax/fil-parser/context"

	golemBackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"

	cacheMetrics "github.com/zondax/fil-parser/actors/cache/metrics"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/types"
)

const OnChainImpl = "on-chain"

// OnChain implementation
type OnChain struct {
	logger     *logger.Logger
	backoff    *golemBackoff.BackOff
	metrics    *cacheMetrics.ActorsCacheMetricsClient
	httpClient *resty.Client
}

func (m *OnChain) StoreAddressInfo(info types.AddressInfo) {
	// Not implemented
}

func (m *OnChain) BackFill() error {
	// Nothing to do
	return nil
}

func (m *OnChain) NewImpl(source common.DataSource, logger *logger.Logger, metrics *cacheMetrics.ActorsCacheMetricsClient, backoff *golemBackoff.BackOff) error {
	// Node datastore is required
	m.logger = logger2.GetSafeLogger(logger)

	m.metrics = metrics
	m.backoff = backoff
	m.httpClient = resty.New().SetTimeout(30 * time.Second)

	return nil
}

func (m *OnChain) ImplementationType() string {
	return OnChainImpl
}

func (m *OnChain) GetActorCode(ctx context.Context, address address.Address, key filTypes.TipSetKey, _, _ bool) (string, error) {
	var actorCid cid.Cid
	var err error

	nodes, err := parserContext.GetNodes(ctx)
	if err != nil {
		return "", err
	}
	for _, node := range nodes {
		actorCid, err = m.retrieveActorFromLotus(node, address, key)
		if err != nil && !IsRetriableError(err) {
			return cid.Undef.String(), err
		}
	}

	return actorCid.String(), nil
}

func (m *OnChain) GetRobustAddress(ctx context.Context, address address.Address, _ bool) (string, error) {
	isRobustAddress, err := common.IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if isRobustAddress {
		// Already a robust address
		return address.String(), nil
	}

	var robustAdd string
	nodes, err := parserContext.GetNodes(ctx)
	if err != nil {
		return "", err
	}
	// Address is not in cache, get robust address from lotus
	for _, node := range nodes {
		robustAdd, err = m.retrieveActorPubKeyFromLotus(node, address, false)
		if err != nil && !IsRetriableError(err) {
			return "", err
		}
	}

	return robustAdd, nil
}

func (m *OnChain) GetShortAddress(ctx context.Context, address address.Address, _ bool) (string, error) {
	isRobustAddress, err := common.IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if !isRobustAddress {
		// Already a short address
		return address.String(), nil
	}

	var shortAdd string
	nodes, err := parserContext.GetNodes(ctx)
	if err != nil {
		return "", err
	}
	for _, node := range nodes {
		shortAdd, err = m.retrieveActorPubKeyFromLotus(node, address, true)
		if err != nil && !IsRetriableError(err) {
			return "", common.ErrKeyNotFound
		}
	}

	return shortAdd, nil
}

// IsSystemActor returns false for all OnChain implementations as the system actors list is maintained by the helper.
// Use the ActorsCache directly.
// Only required to satisfy IActorsCache.
func (m *OnChain) IsSystemActor(_ string) bool {
	return false
}

// IsGenesisActor returns false for all OnChain implementations as the genesis actors list is maintained by the helper.
// Use the ActorsCache directly.
// Only required to satisfy IActorsCache.
func (m *OnChain) IsGenesisActor(_ string) bool {
	return false
}

func (m *OnChain) retrieveActorFromLotus(node api.FullNode, add address.Address, key filTypes.TipSetKey) (cid.Cid, error) {
	nodeApiCallOptions := &NodeApiCallWithRetryOptions[*filTypes.Actor]{
		RequestName: "StateGetActorWithTipSetKey",
		BackOff:     *m.backoff,
		Request: func() (*filTypes.Actor, error) {
			return node.StateGetActor(context.Background(), add, key)
		},
		RetryErrStrings: []string{"ipld: could not find", "RPC client error", "503"},
	}

	actor, err := NodeApiCallWithRetry(nodeApiCallOptions, m.metrics)
	if err != nil {
		// Try again but with an empty tipset Key
		nodeApiCallOptions.RequestName = "StateGetActor"
		nodeApiCallOptions.Request = func() (*filTypes.Actor, error) {
			return node.StateGetActor(context.Background(), add, filTypes.EmptyTSK)
		}
		actor, err = NodeApiCallWithRetry(nodeApiCallOptions, m.metrics)
		if err != nil {
			m.logger.Errorf("[ActorsCache] - retrieveActorFromLotus(%s): %s", add.String(), err.Error())
			return cid.Cid{}, err
		}
	}

	return actor.Code, nil
}

func (m *OnChain) retrieveActorPubKeyFromLotus(node api.FullNode, add address.Address, reverse bool) (string, error) {
	var key address.Address
	var err error

	nodeApiCallOptions := &NodeApiCallWithRetryOptions[address.Address]{
		BackOff:         *m.backoff,
		RetryErrStrings: []string{"RPC client error"},
	}

	if reverse {
		nodeApiCallOptions.RequestName = "StateLookupID"
		nodeApiCallOptions.Request = func() (address.Address, error) {
			return node.StateLookupID(context.Background(), add, filTypes.EmptyTSK)
		}
		key, err = NodeApiCallWithRetry(nodeApiCallOptions, m.metrics)
	} else {
		nodeApiCallOptions.RequestName = "StateAccountKey"
		nodeApiCallOptions.Request = func() (address.Address, error) {
			return node.StateAccountKey(context.Background(), add, filTypes.EmptyTSK)
		}
		key, err = NodeApiCallWithRetry(nodeApiCallOptions, m.metrics)
	}

	if err != nil {
		if strings.Contains(err.Error(), "actor code is not account") {
			nodeApiCallOptions.RequestName = "StateLookupRobustAddress"
			nodeApiCallOptions.Request = func() (address.Address, error) {
				return node.StateLookupRobustAddress(context.Background(), add, filTypes.EmptyTSK)
			}
			key, err = NodeApiCallWithRetry(nodeApiCallOptions, m.metrics)
			if err != nil {
				m.logger.Errorf("[ActorsCache] - retrieveActorPubKeyFromLotus(StateLookupRobustAddress): %s", err.Error())
				return "", common.ErrKeyNotFound
			}
		} else {
			m.logger.Errorf("[ActorsCache] - retrieveActorPubKeyFromLotus: %s", err.Error())
			return "", common.ErrKeyNotFound
		}
	}

	// Must check here because if lotus cannot find the pair, it will return the same address as result
	if key.String() == add.String() {
		return "", common.ErrKeyNotFound
	}

	return key.String(), nil
}

func (m *OnChain) GetEVMSelectorSig(ctx context.Context, selectorID string, _ bool) (string, error) {
	resp, err := m.httpClient.NewRequest().
		SetQueryParam("hex_signature", selectorID).
		SetContext(ctx).
		SetResult(&FourBytesSignatureResult{}).
		Get(SignatureDBURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("error from 4bytes: %v", resp.Error())
	}

	signatureData, ok := resp.Result().(*FourBytesSignatureResult)
	if !ok {
		return "", errors.New("error asserting result to SignatureResult")
	}

	if len(signatureData.Results) == 0 {
		return "", fmt.Errorf("signature not found: %s", selectorID)
	}

	sig := signatureData.Results[0].TextSignature
	return sig, nil
}

func (m *OnChain) StoreEVMSelectorSig(ctx context.Context, selectorHash, selectorSig string, _ bool) error {
	return fmt.Errorf("unimplimented")
}

func (m *OnChain) ClearBadAddressCache() {
	// Nothing to do
}
