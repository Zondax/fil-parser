package verifreg

import (
	"context"
	"errors"
	"strings"

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
	helper *helper.Helper
	logger *logger.Logger
}

func NewEventGenerator(helper *helper.Helper, logger *logger.Logger) EventGenerator {
	return &eventGenerator{
		helper: helper,
		logger: logger,
	}
}

func (eg *eventGenerator) GenerateVerifregEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.VerifregEvents, error) {
	events := &types.VerifregEvents{
		VerifierInfo: make([]*types.VerifierInfo, 0),
		ClientInfo:   make([]*types.ClientInfo, 0),
		Deals:        make([]*types.VerifregDeal, 0),
	}

	for _, tx := range transactions {
		if !strings.EqualFold(tx.Status, "ok") {
			eg.logger.Debug("failed tx found, skipping it")
			continue
		}

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

		if !eg.isVerifregMessage(actorName, tx.TxType) {
			continue
		}

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
	}

	return false
}

func (eg *eventGenerator) createVerifregInfo(tx *types.Transaction, tipsetCid string, events *types.VerifregEvents) (*types.VerifregEvents, error) {
	switch tx.TxType {
	case parser.MethodAddVerifier:
		verifierInfo, err := eg.parseAddVerifier(tx, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.VerifierInfo = append(events.VerifierInfo, verifierInfo)
		return events, nil
	}

	return nil, errors.New("unknown verifreg message")
}

func (eg *eventGenerator) parseAddVerifier(tx *types.Transaction, tipsetCid string) (*types.VerifierInfo, error) {

	return nil, nil
}
