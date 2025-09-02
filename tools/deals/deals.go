package deals

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

type EventGenerator interface {
	GenerateDealsEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.DealsEvents, error)
}

var _ EventGenerator = &eventGenerator{}

type eventGenerator struct {
	helper  *helper.Helper
	logger  *logger.Logger
	metrics *dealsMetricsClient
	network string
	config  parser.Config
}

func NewEventGenerator(helper *helper.Helper, logger *logger.Logger, metrics metrics.MetricsClient, network string, config parser.Config) EventGenerator {
	return &eventGenerator{
		helper:  helper,
		logger:  logger,
		metrics: newClient(metrics, "deals"),
		network: network,
		config:  config,
	}
}

func (eg *eventGenerator) GenerateDealsEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.DealsEvents, error) {
	events := &types.DealsEvents{
		DealsMessages:    []*types.DealsMessages{},
		DealsProposals:   []*types.DealsProposals{},
		DealsActivations: []*types.DealsActivations{},
		DealsSpaceInfo:   []*types.DealsSpaceInfo{},
	}

	for _, tx := range transactions {
		if !(strings.EqualFold(tx.Status, common.TxStatusOk) && strings.EqualFold(tx.SubcallStatus, common.TxStatusOk)) {
			eg.logger.Debug("failed tx found, skipping it")
			continue
		}

		addr, err := address.NewFromString(tx.TxTo)
		if err != nil {
			eg.logger.Errorf("could not parse address. Err: %s", err)
		}

		// #nosec G115
		actorName, err := eg.helper.GetActorNameFromAddress(addr, int64(tx.Height), tipsetKey)
		if err != nil {
			_ = eg.metrics.UpdateActorNameFromAddressMetric()
			eg.logger.Errorf("could not get actor name from address. Err: %s", err)
			continue
		}

		if !eg.isDealsStateMessage(actorName, tx.TxType) {
			continue
		}

		// the address that calls the market actor
		actorAddress := tx.TxFrom
		dealMessage, err := eg.createDealMessage(tx, tipsetCid, actorAddress)
		if err != nil {
			eg.logger.Errorf("could not create deal message. Err: %s", err)
			continue
		}

		events.DealsMessages = append(events.DealsMessages, dealMessage)

		if eg.isPublishStorageDeals(tx.TxType) {
			dealsInfo, err := eg.createDealsInfo(ctx, tx)
			if err != nil {
				eg.logger.Errorf("could not create deal proposal. Err: %s", err)
				continue
			}
			events.DealsProposals = append(events.DealsProposals, dealsInfo...)
		}
		if eg.isDealActivation(tx) {
			dealActivations, dealSpaceInfo, err := eg.createDealActivations(ctx, tx)
			if err != nil {
				eg.logger.Errorf("could not create deal activations. Err: %s", err)
				continue
			}
			events.DealsActivations = append(events.DealsActivations, dealActivations...)
			events.DealsSpaceInfo = append(events.DealsSpaceInfo, dealSpaceInfo...)
		}
	}

	return events, nil
}

func (eg *eventGenerator) isDealsStateMessage(actorName, txType string) bool {
	if !strings.Contains(actorName, manifest.MarketKey) {
		return false
	}

	switch {
	case strings.EqualFold(txType, parser.MethodPublishStorageDeals),
		strings.EqualFold(txType, parser.MethodPublishStorageDealsExported),
		strings.EqualFold(txType, parser.MethodVerifyDealsForActivation),
		strings.EqualFold(txType, parser.MethodActivateDeals),
		strings.EqualFold(txType, parser.MethodSettleDealPaymentsExported),
		strings.EqualFold(txType, parser.MethodSectorContentChanged):
		return true
	}

	return false
}

func (eg *eventGenerator) isPublishStorageDeals(txType string) bool {
	switch {
	case strings.EqualFold(txType, parser.MethodPublishStorageDeals),
		strings.EqualFold(txType, parser.MethodPublishStorageDealsExported):
		return true
	}

	return false
}
