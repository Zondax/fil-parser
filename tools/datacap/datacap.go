package datacap

import (
	"context"
	"fmt"
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

type EventGenerator interface {
	GenerateDataCapEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.DataCapEvents, error)
}

var _ EventGenerator = &eventGenerator{}

type eventGenerator struct {
	helper  *helper.Helper
	logger  *logger.Logger
	metrics *datacapMetricsClient
	config  parser.Config
}

func NewEventGenerator(helper *helper.Helper, logger *logger.Logger, metrics metrics.MetricsClient, config parser.Config) EventGenerator {
	return &eventGenerator{
		helper:  helper,
		logger:  logger,
		metrics: newClient(metrics, manifest.DatacapKey),
		config:  config,
	}
}

func (eg *eventGenerator) GenerateDataCapEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.DataCapEvents, error) {
	events := &types.DataCapEvents{
		DataCapInfo:           []*types.DataCapInfo{},
		DataCapTokenEvent:     []*types.DataCapTokenEvent{},
		DataCapAllowanceEvent: []*types.DataCapAllowanceEvent{},
	}

	for _, tx := range transactions {
		if !common.IsTxSuccess(tx) {
			eg.logger.Debug("failed tx found, skipping it")
			continue
		}

		actorAddress := tx.TxTo

		addr, err := address.NewFromString(actorAddress)
		if err != nil {
			return nil, fmt.Errorf("could not parse address. err: %w", err)
		}

		// #nosec G115
		_, actorName, err := eg.helper.GetActorInfoFromAddress(addr, int64(tx.Height), tipsetKey)
		if err != nil {
			_ = eg.metrics.UpdateActorNameFromAddressMetric()
			return nil, fmt.Errorf("could not get actor name from address. err: %w", err)
		}

		if !eg.isDataCapStateMessage(actorName) {
			continue
		}

		dataCapInfo, err := eg.createDataCapInfo(tx, tipsetCid, actorAddress)
		if err != nil {
			return nil, fmt.Errorf("could not create datacap info. err: %w", err)
		}

		events.DataCapInfo = append(events.DataCapInfo, dataCapInfo)
		if isDatacapTokenMessage(tx.TxType) {
			tokenEvents, allowanceEvents, err := eg.createDataCapTokenEvents(ctx, tx, tipsetCid)
			if err != nil {
				return nil, fmt.Errorf("could not create datacap token events. err: %w", err)
			}
			events.DataCapTokenEvent = append(events.DataCapTokenEvent, tokenEvents...)
			events.DataCapAllowanceEvent = append(events.DataCapAllowanceEvent, allowanceEvents...)
		}
	}

	return events, nil
}

func (eg *eventGenerator) isDataCapStateMessage(actorName string) bool {
	return strings.Contains(actorName, manifest.DatacapKey)
}
