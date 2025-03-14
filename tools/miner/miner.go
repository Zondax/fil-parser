package miner

import (
	"context"
	"strings"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

const (
	txStatusOk = "ok"
)

type EventGenerator interface {
	GenerateMinerEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MinerEvents, error)
}

var _ EventGenerator = &eventGenerator{}

type eventGenerator struct {
	helper  *helper.Helper
	logger  *zap.Logger
	metrics *minerMetricsClient
}

func NewEventGenerator(helper *helper.Helper, logger *zap.Logger, metrics metrics.MetricsClient) EventGenerator {
	return &eventGenerator{
		helper:  helper,
		logger:  logger,
		metrics: newClient(metrics, "miner"),
	}
}

func (eg *eventGenerator) GenerateMinerEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MinerEvents, error) {
	events := &types.MinerEvents{
		MinerInfo:    []*types.MinerInfo{},
		MinerSectors: []*types.MinerSectorEvent{},
	}

	for _, tx := range transactions {
		if !strings.EqualFold(tx.Status, txStatusOk) {
			eg.logger.Sugar().Debug("failed tx found, skipping it")
			continue
		}

		addrTo, err := address.NewFromString(tx.TxTo)
		if err != nil {
			eg.logger.Sugar().Errorf("could not parse address. Err: %s", err)
			continue
		}

		actorName, err := eg.helper.GetActorNameFromAddress(addrTo, int64(tx.Height), tipsetKey)
		if err != nil {
			_ = eg.metrics.UpdateActorNameFromAddressMetric()
			eg.logger.Sugar().Errorf("could not get actor name from address. Err: %s", err)
			continue
		}

		actorAddress := tx.TxTo
		if !eg.isMinerStateMessage(actorName, tx.TxType) {
			continue
		}

		// this is executed by(from) the miner actor
		if tx.TxType == parser.MethodUpdateClaimedPower {
			actorAddress = tx.TxFrom
		}

		minerInfo, err := eg.createMinerInfo(tx, tipsetCid, actorAddress)
		if err != nil {
			eg.logger.Sugar().Errorf("could not create miner info. Err: %s", err)
			continue
		}

		events.MinerInfo = append(events.MinerInfo, minerInfo)

		if eg.isMinerSectorMessage(actorName, tx.TxType) {
			minerSectors, err := eg.createSectorEvents(ctx, tx, tipsetCid)
			if err != nil {
				eg.logger.Sugar().Errorf("could not create miner sector. Err: %s", err)
				continue
			}
			events.MinerSectors = append(events.MinerSectors, minerSectors...)
		}

	}
	return events, nil
}

func (eg *eventGenerator) isMinerStateMessage(actorName, txType string) bool {
	switch {
	case strings.EqualFold(actorName, manifest.MinerKey):
		return (!strings.EqualFold(txType, parser.MethodOnDeferredCronEvent))
	case strings.EqualFold(txType, parser.MethodUpdateClaimedPower):
		return true
	case strings.EqualFold(actorName, manifest.MarketKey):
		return (strings.EqualFold(txType, parser.MethodOnMinerSectorsTerminate) ||
			strings.EqualFold(txType, parser.MethodPublishStorageDeals) ||
			strings.EqualFold(txType, parser.MethodPublishStorageDealsExported) ||
			strings.EqualFold(txType, parser.MethodActivateDeals))

	case strings.EqualFold(actorName, manifest.PowerKey):
		return strings.EqualFold(txType, parser.MethodCurrentTotalPower)

	case strings.EqualFold(actorName, manifest.RewardKey):
		return (strings.EqualFold(txType, parser.MethodThisEpochReward) ||
			strings.EqualFold(txType, parser.MethodAwardBlockReward))
	}

	return false
}
