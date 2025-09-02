package verifreg

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
)

type EventGenerator interface {
	GenerateVerifregEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.VerifregEvents, error)
}

var _ EventGenerator = &eventGenerator{}

type eventGenerator struct {
	helper           *helper.Helper
	logger           *logger.Logger
	verifiedRegistry *verifiedRegistry.VerifiedRegistry
	config           parser.Config
	network          string
	metrics          *verifregMetricsClient
}

func NewEventGenerator(helper *helper.Helper, logger *logger.Logger, metrics metrics.MetricsClient, network string, config parser.Config) EventGenerator {
	return &eventGenerator{
		helper:           helper,
		logger:           logger,
		network:          network,
		verifiedRegistry: verifiedRegistry.New(logger),
		config:           config,
		metrics:          newClient(metrics, manifest.VerifregKey),
	}
}

func (eg *eventGenerator) GenerateVerifregEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.VerifregEvents, error) {
	events := &types.VerifregEvents{
		VerifierInfo: make([]*types.VerifregEvent, 0),
		ClientInfo:   make([]*types.VerifregClientInfo, 0),
		Deals:        make([]*types.VerifregDeal, 0),
	}

	for _, tx := range transactions {
		if !common.IsTxSuccess(tx) {
			eg.logger.Debug("failed tx found, skipping it")
			continue
		}

		addr, err := address.NewFromString(tx.TxTo)
		if err != nil {
			return nil, fmt.Errorf("could not parse address. err: %w", err)
		}

		// #nosec G115
		actorName, err := eg.helper.GetActorNameFromAddress(addr, int64(tx.Height), tipsetKey)
		if err != nil {
			_ = eg.metrics.UpdateActorNameFromAddressMetric()
			return nil, fmt.Errorf("could not get actor name from address. err: %w", err)
		}

		if !eg.isVerifregMessage(actorName, tx.TxType) {
			continue
		}

		events, err = eg.createVerifregInfo(tx, tipsetCid, events)
		if err != nil {
			return nil, fmt.Errorf("could not create verifreg info. err: %w", err)
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
			ID:           tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
			ActorAddress: addr,
			TxCid:        tx.TxCid,
			Height:       tx.Height,
			ActionType:   tx.TxType,
			Data:         tx.TxMetadata,
			TxTimestamp:  tx.TxTimestamp,
		}, &types.VerifregClientInfo{
			ID:            tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
			ActorAddress:  addr,
			ClientAddress: addr,
			TxCid:         tx.TxCid,
			Height:        tx.Height,
			ActionType:    tx.TxType,
			DataCap:       dataCapAmount,
			Data:          tx.TxMetadata,
			TxTimestamp:   tx.TxTimestamp,
		}, nil
}

func (eg *eventGenerator) removeVerifier(tx *types.Transaction, metadata map[string]interface{}, tipsetCid string) (*types.VerifregEvent, *types.VerifregClientInfo, error) {
	addr, err := parseRemoveVerifier(metadata)
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
			ID:           tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
			ActorAddress: addr,
			TxCid:        tx.TxCid,
			Height:       tx.Height,
			ActionType:   tx.TxType,
			Data:         tx.TxMetadata,
			TxTimestamp:  tx.TxTimestamp,
		}, &types.VerifregClientInfo{
			ID:            tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
			ActorAddress:  addr,
			ClientAddress: addr,
			TxCid:         tx.TxCid,
			Height:        tx.Height,
			ActionType:    tx.TxType,
			DataCap:       big.NewInt(0),
			Data:          tx.TxMetadata,
			TxTimestamp:   tx.TxTimestamp,
		}, nil
}

func (eg *eventGenerator) addVerifiedClient(tx *types.Transaction, metadata map[string]interface{}, tipsetCid string) (*types.VerifregEvent, *types.VerifregClientInfo, error) {
	addr, dataCapAmount, err := getAddressAllowance(metadata)
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
			ID:           tools.BuildId(tipsetCid, tx.TxCid, addr, tx.TxFrom, fmt.Sprint(tx.Height)),
			ActorAddress: tx.TxTo,
			TxCid:        tx.TxCid,
			Height:       tx.Height,
			ActionType:   tx.TxType,
			Data:         tx.TxMetadata,
			TxTimestamp:  tx.TxTimestamp,
		},
		&types.VerifregClientInfo{
			ID:            tools.BuildId(tipsetCid, tx.TxCid, addr, tx.TxFrom, fmt.Sprint(tx.Height)),
			ClientAddress: addr,
			ActorAddress:  tx.TxFrom,
			TxCid:         tx.TxCid,
			Height:        tx.Height,
			ActionType:    tx.TxType,
			DataCap:       dataCapAmount,
			Data:          tx.TxMetadata,
			TxTimestamp:   tx.TxTimestamp,
		},
		nil
}

func (eg *eventGenerator) removeVerifiedClient(tx *types.Transaction, metadata map[string]interface{}, tipsetCid string) (*types.VerifregEvent, *types.VerifregClientInfo, error) {
	// #nosec G115
	verifiedClientToRemove, _, removedDatacap, err := parseRemoveVerifiedClient(metadata, eg.network, int64(tx.Height))
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
			ID:           tools.BuildId(tipsetCid, tx.TxCid, verifiedClientToRemove, fmt.Sprint(tx.Height)),
			ActorAddress: tx.TxTo,
			TxCid:        tx.TxCid,
			Height:       tx.Height,
			ActionType:   tx.TxType,
			Data:         tx.TxMetadata,
			TxTimestamp:  tx.TxTimestamp,
		},
		&types.VerifregClientInfo{
			ID:            tools.BuildId(tipsetCid, tx.TxCid, verifiedClientToRemove, fmt.Sprint(tx.Height)),
			ActorAddress:  tx.TxTo,
			ClientAddress: verifiedClientToRemove,
			TxCid:         tx.TxCid,
			Height:        tx.Height,
			ActionType:    tx.TxType,
			DataCap:       removedDatacap,
			Data:          tx.TxMetadata,
			TxTimestamp:   tx.TxTimestamp,
		}, nil
}

func (eg *eventGenerator) universalReceiverHook(tx *types.Transaction, tipsetCid string) (*types.VerifregEvent, []*types.VerifregDeal, error) {
	clientAddress, clientValue, dealValue, err := eg.parserUniversalReceiverHook(tx, tipsetCid)
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, clientAddress, fmt.Sprint(tx.Height)),
		ActorAddress: clientAddress,
		TxCid:        tx.TxCid,
		Height:       tx.Height,
		ActionType:   tx.TxType,
		Data:         clientValue,
		TxTimestamp:  tx.TxTimestamp,
	}, dealValue, nil
}
