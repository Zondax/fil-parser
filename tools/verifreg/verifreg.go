package verifreg

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
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
	case parser.MethodAddVerifiedClient:
		clientInfo, err := eg.parseAddVerifiedClient(tx, tipsetCid)
		if err != nil {
			return nil, err
		}
		events.ClientInfo = append(events.ClientInfo, clientInfo)
		return events, nil
	}

	return events, nil
}

func (eg *eventGenerator) parseAddVerifier(tx *types.Transaction, tipsetCid string) (*types.VerifierInfo, error) {
	addr, intAllowance, err := eg.getAddressAllowance(tx.TxMetadata)
	if err != nil {
		return nil, err
	}

	return &types.VerifierInfo{
		ID:          tools.BuildId(tipsetCid, tx.TxCid, addr, fmt.Sprint(tx.Height)),
		Address:     addr,
		Allowance:   intAllowance,
		TxCid:       tx.TxCid,
		Height:      tx.Height,
		TxTimestamp: tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) parseAddVerifiedClient(tx *types.Transaction, tipsetCid string) (*types.ClientInfo, error) {
	addr, intAllowance, err := eg.getAddressAllowance(tx.TxMetadata)
	if err != nil {
		return nil, err
	}

	return &types.ClientInfo{
		ID:              tools.BuildId(tipsetCid, tx.TxCid, addr, tx.TxFrom, fmt.Sprint(tx.Height)),
		Verifier: tx.TxFrom,
		Address:         addr,
		Allowance:       intAllowance,
		TxCid:           tx.TxCid,
		Height:          tx.Height,
		TxTimestamp:     tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) getAddressAllowance(metadata string) (string, uint64, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(metadata), &value)
	if err != nil {
		return "", 0, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, ok := value[parser.ParamsKey].(map[string]interface{})
	if !ok {
		return "", 0, fmt.Errorf("error parsing params: %w", err)
	}

	addr, ok := params["Address"].(string)
	if !ok {
		return "", 0, fmt.Errorf("error parsing address: %w", err)
	}

	allowance, ok := params["Allowance"].(string)
	if !ok {
		return "", 0, fmt.Errorf("error parsing allowance: %w", err)
	}

	intAllowance, err := strconv.ParseUint(allowance, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("error parsing allowance string '%s': %w", allowance, err)
	}

	return addr, intAllowance, nil
}
