package verifreg

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/zondax/fil-parser/tools"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
)

type EventGenerator interface {
	GenerateVerifregEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.VerifregEvents, error)
}

var _ EventGenerator = &eventGenerator{}

type eventGenerator struct {
	helper  *helper.Helper
	logger  *logger.Logger
	network string
}

func NewEventGenerator(helper *helper.Helper, logger *logger.Logger, network string) EventGenerator {
	return &eventGenerator{
		helper:  helper,
		logger:  logger,
		network: network,
	}
}

func (eg *eventGenerator) GenerateVerifregEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.VerifregEvents, error) {
	events := &types.VerifregEvents{
		VerifierInfo: make([]*types.VerifregEvent, 0),
		ClientInfo:   make([]*types.VerifregClientInfo, 0),
		Deals:        make([]*types.VerifregDeal, 0),
	}

	for _, tx := range transactions {
		if !strings.EqualFold(tx.Status, "ok") {
			eg.logger.Debug("failed tx found, skipping it")
			continue
		}

		addr, err := address.NewFromString(tx.TxTo)
		if err != nil {
			return nil, fmt.Errorf("could not parse address. Err: %s", err)
		}

		actorName, err := eg.helper.GetActorNameFromAddress(addr, int64(tx.Height), tipsetKey)
		if err != nil {
			return nil, fmt.Errorf("could not get actor name from address. Err: %s", err)
		}

		if !eg.isVerifregMessage(actorName, tx.TxType) {
			continue
		}

		events, err = eg.createVerifregInfo(tx, tipsetCid, events)
		if err != nil {
			return nil, fmt.Errorf("could not create verifreg info. Err: %s", err)
		}

	}

	return events, nil
}

func (eg *eventGenerator) isVerifregMessage(actorName, txType string) bool {
	return strings.EqualFold(actorName, manifest.VerifregKey)
}

func (eg *eventGenerator) createVerifregInfo(tx *types.Transaction, tipsetCid string, events *types.VerifregEvents) (*types.VerifregEvents, error) {

	metadata := map[string]interface{}{}
	err := json.Unmarshal([]byte(tx.TxMetadata), &metadata)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}
	switch tx.TxType {
	case parser.MethodAddVerifier:
		addVerifier, clientInfo, err := eg.parseAddVerifier(tx, metadata, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.VerifierInfo = append(events.VerifierInfo, addVerifier)
		events.ClientInfo = append(events.ClientInfo, clientInfo)
	case parser.MethodRemoveVerifier:
		removeVerifier, clientInfo, err := eg.removeVerifier(tx, metadata, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.VerifierInfo = append(events.VerifierInfo, removeVerifier)
		events.ClientInfo = append(events.ClientInfo, clientInfo)
	case parser.MethodAddVerifiedClient:
		verifierInfo, clientInfo, err := eg.addVerifiedClient(tx, metadata, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.VerifierInfo = append(events.VerifierInfo, verifierInfo)
		events.ClientInfo = append(events.ClientInfo, clientInfo)
	case parser.MethodRemoveVerifiedClientDataCap:
		verifierInfo, clientInfo, err := eg.removeVerifiedClient(tx, metadata, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.VerifierInfo = append(events.VerifierInfo, verifierInfo)
		events.ClientInfo = append(events.ClientInfo, clientInfo)
	case parser.MethodUniversalReceiverHook:
		clientInfo, dealInfo, err := eg.universalReceiverHook(tx, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.VerifierInfo = append(events.VerifierInfo, clientInfo)
		events.Deals = append(events.Deals, dealInfo...)
	}

	return events, nil
}

/*
{"MethodNum":"2","Params":{"Address":"f1arlxnqbq2bpyw7wzcz7stnsr4v24xuwaj7p2uhq","Allowance":"100000000000000000000000000000000000000000"}}
*/

func (eg *eventGenerator) parseAddVerifier(tx *types.Transaction, metadata map[string]interface{}, tipsetCid string) (*types.VerifregEvent, *types.VerifregClientInfo, error) {
	addr, dataCapAmount, err := getAddressAllowance(metadata)
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
			Address:     addr,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		}, &types.VerifregClientInfo{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
			Address:     addr,
			Client:      addr,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			DataCap:     dataCapAmount,
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		}, nil
}

func (eg *eventGenerator) removeVerifier(tx *types.Transaction, metadata map[string]interface{}, tipsetCid string) (*types.VerifregEvent, *types.VerifregClientInfo, error) {
	addr, err := parseRemoveVerifier(metadata)
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
			Address:     addr,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		}, &types.VerifregClientInfo{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
			Address:     addr,
			Client:      addr,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			DataCap:     big.NewInt(0),
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		}, nil
}

func (eg *eventGenerator) addVerifiedClient(tx *types.Transaction, metadata map[string]interface{}, tipsetCid string) (*types.VerifregEvent, *types.VerifregClientInfo, error) {
	addr, dataCapAmount, err := getAddressAllowance(metadata)
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, tx.TxFrom, fmt.Sprint(tx.Height)),
			Address:     tx.TxTo,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		},
		&types.VerifregClientInfo{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, tx.TxFrom, fmt.Sprint(tx.Height)),
			Client:      addr,
			Address:     tx.TxFrom,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			DataCap:     dataCapAmount,
			Verifiers:   []string{tx.TxFrom},
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		},
		nil
}

func (eg *eventGenerator) removeVerifiedClient(tx *types.Transaction, metadata map[string]interface{}, tipsetCid string) (*types.VerifregEvent, *types.VerifregClientInfo, error) {
	verifiedClientToRemove, verifiers, removedDatacap, err := parseRemoveVerifiedClient(metadata, eg.network, int64(tx.Height))
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, verifiedClientToRemove, fmt.Sprint(tx.Height)),
			Address:     tx.TxTo,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		},
		&types.VerifregClientInfo{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, verifiedClientToRemove, fmt.Sprint(tx.Height)),
			Address:     tx.TxTo,
			Client:      verifiedClientToRemove,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			DataCap:     removedDatacap,
			Verifiers:   verifiers,
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		}, nil
}

func (eg *eventGenerator) universalReceiverHook(tx *types.Transaction, tipsetCid string) (*types.VerifregEvent, []*types.VerifregDeal, error) {
	clientValue, dealValue, err := eg.parserUniversalReceiverHook(tx, tipsetCid)
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
		ID:          tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, fmt.Sprint(tx.Height)),
		Address:     tx.TxFrom, // This is the datacap actor. Should be the client.
		TxCid:       tx.TxCid,
		Height:      tx.Height,
		ActionType:  tx.TxType,
		Value:       clientValue,
		TxTimestamp: tx.TxTimestamp,
	}, dealValue, nil
}
