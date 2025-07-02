package verifreg

import (
	"context"
	"fmt"
	"strings"

	"github.com/zondax/fil-parser/tools"

	"github.com/zondax/golem/pkg/logger"

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
		ClientInfo:   make([]*types.VerifregEvent, 0),
		Deals:        make([]*types.VerifregDeal, 0),
	}

	for _, tx := range transactions {
		if !strings.EqualFold(tx.Status, "ok") {
			eg.logger.Debug("failed tx found, skipping it")
			continue
		}

		/*
			TODO: Improve this logic
			addr, err := address.NewFromString(tx.TxTo)
			if err != nil {
				eg.logger.Errorf("could not parse address. Err: %s", err)
			}

			_, actorName, err := eg.helper.GetActorNameFromAddress(addr, int64(tx.Height), tipsetKey)
			if err != nil {
				// _ = eg.metrics.UpdateActorNameFromAddressMetric()
				eg.logger.Errorf("could not get actor name from address. Err: %s", err)
				continue
			}

			if !eg.isVerifregMessage("", tx.TxType) {
				continue
			}
		*/
		var err error
		events, err = eg.createVerifregInfo(tx, tipsetCid, events)
		if err != nil {
			eg.logger.Errorf("could not create verifreg info. Err: %s", err)
			continue
		}

	}

	return events, nil
}

func (eg *eventGenerator) isVerifregMessage(actorName, txType string) bool {
	switch {
	case strings.EqualFold(actorName, manifest.VerifregKey):
		return true
	case txType == parser.MethodAddVerifier:
		return true
	case txType == parser.MethodAddVerifiedClient:
		return true
	}

	return false
}

func (eg *eventGenerator) createVerifregInfo(tx *types.Transaction, tipsetCid string, events *types.VerifregEvents) (*types.VerifregEvents, error) {
	switch tx.TxType {
	case parser.MethodAddVerifier:
		addVerifier, err := eg.parseAddVerifier(tx, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.VerifierInfo = append(events.VerifierInfo, addVerifier)
	case parser.MethodRemoveVerifier:
		removeVerifier, err := eg.removeVerifier(tx, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.VerifierInfo = append(events.VerifierInfo, removeVerifier)
	case parser.MethodAddVerifiedClient:
		verifierInfo, clientInfo, err := eg.addVerifiedClient(tx, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.VerifierInfo = append(events.VerifierInfo, verifierInfo)
		events.ClientInfo = append(events.ClientInfo, clientInfo)
	case parser.MethodRemoveVerifiedClientDataCap:
		verifierInfo, clientInfo, err := eg.removeVerifiedClient(tx, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.VerifierInfo = append(events.VerifierInfo, verifierInfo)
		events.ClientInfo = append(events.ClientInfo, clientInfo)
		//case parser.MethodTransferExported:

	}

	return events, nil
}

func (eg *eventGenerator) parseAddVerifier(tx *types.Transaction, tipsetCid string) (*types.VerifregEvent, error) {
	addr, _, err := getAddressAllowance(tx.TxMetadata)
	if err != nil {
		return nil, err
	}

	return &types.VerifregEvent{
		ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
		Address:     addr,
		TxCid:       tx.TxCid,
		Height:      tx.Height,
		ActionType:  tx.TxType,
		Value:       tx.TxMetadata,
		TxTimestamp: tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) removeVerifier(tx *types.Transaction, tipsetCid string) (*types.VerifregEvent, error) {
	addr, err := parseRemoveVerifier(tx.TxMetadata)
	if err != nil {
		return nil, err
	}

	return &types.VerifregEvent{
		ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
		Address:     addr,
		TxCid:       tx.TxCid,
		Height:      tx.Height,
		ActionType:  tx.TxType,
		Value:       tx.TxMetadata,
		TxTimestamp: tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) addVerifiedClient(tx *types.Transaction, tipsetCid string) (*types.VerifregEvent, *types.VerifregEvent, error) {
	addr, _, err := getAddressAllowance(tx.TxMetadata)
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, tx.TxFrom, fmt.Sprint(tx.Height)),
			Address:     tx.TxFrom,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		},
		&types.VerifregEvent{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, tx.TxFrom, fmt.Sprint(tx.Height)),
			Address:     addr,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		},
		nil
}

func (eg *eventGenerator) removeVerifiedClient(tx *types.Transaction, tipsetCid string) (*types.VerifregEvent, *types.VerifregEvent, error) {
	verifiedClientToRemove, verifier, _, err := parseRemoveVerifiedClient(tx.TxMetadata, eg.network, int64(tx.Height))
	if err != nil {
		return nil, nil, err
	}

	return &types.VerifregEvent{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, verifiedClientToRemove, verifier, fmt.Sprint(tx.Height)),
			Address:     tx.TxFrom,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		},
		&types.VerifregEvent{
			ID:          tools.BuildId(tipsetCid, tx.TxCid, verifiedClientToRemove, verifier, fmt.Sprint(tx.Height)),
			Address:     verifiedClientToRemove,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			ActionType:  tx.TxType,
			Value:       tx.TxMetadata,
			TxTimestamp: tx.TxTimestamp,
		}, nil
}

//func (eg *eventGenerator) transferExported(tx *types.Transaction, tipsetCid string) (*types.VerifregEvent, error) {
//
//}
