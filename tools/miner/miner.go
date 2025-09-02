package miner

import (
	"context"
	"strings"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
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
	logger  *logger.Logger
	metrics *minerMetricsClient
	config  parser.Config
}

func NewEventGenerator(helper *helper.Helper, logger *logger.Logger, metrics metrics.MetricsClient, config parser.Config) EventGenerator {
	return &eventGenerator{
		helper:  helper,
		logger:  logger,
		metrics: newClient(metrics, "miner"),
		config:  config,
	}
}

func (eg *eventGenerator) GenerateMinerEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MinerEvents, error) {
	events := &types.MinerEvents{
		MinerInfo:    []*types.MinerInfo{},
		MinerSectors: []*types.MinerSectorEvent{},
	}

	for _, tx := range transactions {
		if !strings.EqualFold(tx.SubcallStatus, common.TxStatusOk) {
			eg.logger.Debug("failed tx found, skipping it")
			continue
		}

		actorAddress := tx.TxTo
		// this is executed by(from) the miner actor
		if tx.TxType == parser.MethodUpdateClaimedPower || tx.TxType == parser.TotalFeeOp {
			actorAddress = tx.TxFrom
		}

		addr, err := address.NewFromString(actorAddress)
		if err != nil {
			eg.logger.Errorf("could not parse address. Err: %s", err)
		}

		// #nosec G115
		_, actorName, err := eg.helper.GetActorInfoFromAddress(addr, int64(tx.Height), tipsetKey)
		if err != nil {
			_ = eg.metrics.UpdateActorNameFromAddressMetric()
			eg.logger.Errorf("could not get actor name from address. Err: %s", err)
			continue
		}

		if !eg.isMinerStateMessage(actorName, tx.TxType) {
			continue
		}

		minerInfo, err := eg.createMinerInfo(tx, tipsetCid, actorAddress)
		if err != nil {
			eg.logger.Errorf("could not create miner info. Err: %s", err)
			continue
		}

		events.MinerInfo = append(events.MinerInfo, minerInfo)

		if eg.isMinerSectorMessage(actorName, tx.TxType) {
			minerSectors, err := eg.createSectorEvents(ctx, tx, tipsetCid)
			if err != nil {
				eg.logger.Errorf("could not create miner sector. Err: %s", err)
				continue
			}
			events.MinerSectors = append(events.MinerSectors, minerSectors...)
		}
	}

	return events, nil
}

func (eg *eventGenerator) isMinerStateMessage(actorName, txType string) bool {
	switch {
	case strings.Contains(actorName, manifest.MinerKey):
		return !strings.EqualFold(txType, parser.MethodOnDeferredCronEvent)
	case strings.EqualFold(txType, parser.MethodAwardBlockReward):
		return true
	case strings.EqualFold(txType, parser.MethodUpdateClaimedPower):
		return true
	case strings.EqualFold(txType, parser.MethodOnMinerSectorsTerminate),
		strings.EqualFold(txType, parser.MethodPublishStorageDeals),
		strings.EqualFold(txType, parser.MethodPublishStorageDealsExported),
		strings.EqualFold(txType, parser.MethodActivateDeals):
		return true
	case strings.EqualFold(txType, parser.MethodCurrentTotalPower):
		return true
	case strings.EqualFold(txType, parser.MethodThisEpochReward):
		return true
	}

	return false
}
