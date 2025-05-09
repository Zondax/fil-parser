package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/types"
)

const OnChainImpl = "on-chain"

// OnChain implementation
type OnChain struct {
	Node   api.FullNode
	logger *logger.Logger
}

func (m *OnChain) StoreAddressInfo(info types.AddressInfo) {
	// Not implemented
}

func (m *OnChain) BackFill() error {
	// Nothing to do
	return nil
}

func (m *OnChain) NewImpl(source common.DataSource, logger *logger.Logger) error {
	// Node datastore is required
	m.logger = logger2.GetSafeLogger(logger)
	if source.Node == nil {
		m.logger.Panic("[ActorsCache] - Node ptr is nil")
	}

	m.Node = source.Node
	return nil
}

func (m *OnChain) ImplementationType() string {
	return OnChainImpl
}

func (m *OnChain) GetActorCode(address address.Address, key filTypes.TipSetKey) (string, error) {
	actorCid, err := m.retrieveActorFromLotus(address, key)
	if err != nil {
		return cid.Undef.String(), err
	}

	return actorCid.String(), nil
}

func (m *OnChain) GetRobustAddress(address address.Address) (string, error) {
	isRobustAddress, err := common.IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if isRobustAddress {
		// Already a robust address
		return address.String(), nil
	}

	// Address is not in cache, get robust address from lotus
	robustAdd, err := m.retrieveActorPubKeyFromLotus(address, false)
	if err != nil {
		return "", err
	}

	return robustAdd, nil
}

func (m *OnChain) GetShortAddress(address address.Address) (string, error) {
	isRobustAddress, err := common.IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if !isRobustAddress {
		// Already a short address
		return address.String(), nil
	}

	shortAdd, err := m.retrieveActorPubKeyFromLotus(address, true)
	if err != nil {
		return "", common.ErrKeyNotFound
	}

	return shortAdd, nil
}

func (m *OnChain) retrieveActorFromLotus(add address.Address, key filTypes.TipSetKey) (cid.Cid, error) {
	actor, err := m.Node.StateGetActor(context.Background(), add, filTypes.EmptyTSK)
	if err != nil {
		// Try again but using the corresponding tipset Key
		actor, err = m.Node.StateGetActor(context.Background(), add, key)
		if err != nil {
			m.logger.Errorf("[ActorsCache] - retrieveActorFromLotus: %s", err.Error())
			return cid.Cid{}, err
		}
	}

	return actor.Code, nil
}

func (m *OnChain) retrieveActorPubKeyFromLotus(add address.Address, reverse bool) (string, error) {
	var maxAttempts = 3
	var maxWaitBeforeRetry = 5 * time.Second
	var key address.Address
	var err error

	if reverse {
		key, err = stateLookupWithRetry(maxAttempts, maxWaitBeforeRetry, func() (address.Address, error) {
			return m.Node.StateLookupID(context.Background(), add, filTypes.EmptyTSK)
		})
	} else {
		key, err = stateLookupWithRetry(maxAttempts, maxWaitBeforeRetry, func() (address.Address, error) {
			return m.Node.StateAccountKey(context.Background(), add, filTypes.EmptyTSK)
		})
	}

	if err != nil {
		if strings.Contains(err.Error(), "actor code is not account") {
			key, err = stateLookupWithRetry(maxAttempts, maxWaitBeforeRetry, func() (address.Address, error) {
				return m.Node.StateLookupRobustAddress(context.Background(), add, filTypes.EmptyTSK)
			})
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

func (m *OnChain) GetEVMSelectorSig(ctx context.Context, selectorHash string) (string, error) {
	return "", fmt.Errorf("unimplimented")
}

func (m *OnChain) StoreEVMSelectorSig(ctx context.Context, selectorHash, selectorSig string) error {
	return fmt.Errorf("unimplimented")
}
