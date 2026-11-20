package deals

import (
	"context"
	"fmt"
	"github.com/zondax/fil-parser/tools"
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

		DealsTerminations:   []*types.DealsTerminations{},
		DealsSectorMappings: []*types.DealsSectorMappings{},
	}

	dealsActivationsIds := tools.NewEventIds()
	dealsMessagesIds := tools.NewEventIds()
	dealsProposalsIds := tools.NewEventIds()
	dealsSpaceInfoIds := tools.NewEventIds()
	dealsTerminationsIds := tools.NewEventIds()
	dealsSectorMappingsIds := tools.NewEventIds()
	// appendMappings drops a (miner, sector, deal) repeated within one tx (e.g. a deal listed twice
	// in a pre-commit): the row would carry the same id.
	appendMappings := func(tx *types.Transaction, mappings []*types.DealsSectorMappings) {
		seen := make(map[string]bool, len(mappings))
		for _, mapping := range mappings {
			if seen[mapping.ID] {
				continue
			}
			seen[mapping.ID] = true
			mapping.ID = dealsSectorMappingsIds.Unique(mapping.ID, tx.Id)
			events.DealsSectorMappings = append(events.DealsSectorMappings, mapping)
		}
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
		actorName, err := common.GetActorNameFromAddress(ctx, eg.helper, addr, int64(tx.Height), tipsetKey, true)
		if err != nil {
			_ = eg.metrics.UpdateActorNameFromAddressMetric()
			return nil, err
		}

		// pre-NV22 pre-commits and replica updates sent to a miner: sector -> deal mappings only
		if eg.isSectorMappingMessage(actorName, tx) {
			mappings, err := eg.createSectorMappings(tx)
			if err != nil {
				// Same rule as activations: unusable metadata (e.g. raw-only, or a
				// ProveReplicaUpdates stored without its return) skips this tx only.
				eg.logger.Warnf("skipping sector deal mappings of tx %s (tx type %s, height %d): %s", tx.TxCid, tx.TxType, tx.Height, err.Error())
				_ = eg.metrics.UpdateSkippedSectorMappingsMetric(tx.TxType)
				continue
			}
			appendMappings(tx, mappings)
			continue
		}

		if !eg.isDealsStateMessage(actorName, tx.TxType) {
			continue
		}

		// the address that calls the market actor
		actorAddress := tx.TxFrom
		dealMessage, err := eg.createDealMessage(tx, tipsetCid, actorAddress)
		if err != nil {
			return nil, fmt.Errorf("could not create deal message. err: %w", err)
		}

		dealMessage.ID = dealsMessagesIds.Unique(dealMessage.ID, tx.Id)
		events.DealsMessages = append(events.DealsMessages, dealMessage)

		if eg.isPublishStorageDeals(tx.TxType) {
			dealsInfo, err := eg.createDealsInfo(ctx, tx)
			if err != nil {
				return nil, fmt.Errorf("could not create deal proposal. err: %w", err)
			}
			for i := range dealsInfo {
				dealsInfo[i].ID = dealsProposalsIds.Unique(dealsInfo[i].ID, tx.Id)
			}
			events.DealsProposals = append(events.DealsProposals, dealsInfo...)
		}
		if eg.isDealActivation(tx) {
			dealActivations, dealSpaceInfo, err := eg.createDealActivations(ctx, tx)
			if err != nil {
				// Metadata whose params or return are missing or unusable (e.g. raw-only metadata
				// after a decode failure, or SectorContentChanged metadata stored by a parser
				// that lost the flags of all sectors but the first) cannot tell which deals were
				// activated. Skip the activations of this tx only; its message row is kept and
				// the rest of the height is processed.
				eg.logger.Warnf("skipping deal activations of tx %s (tx type %s, height %d): %s", tx.TxCid, tx.TxType, tx.Height, err.Error())
				_ = eg.metrics.UpdateSkippedActivationsMetric(tx.TxType)
				continue
			}
			for i := range dealActivations {
				dealActivations[i].ID = dealsActivationsIds.Unique(dealActivations[i].ID, tx.Id)
			}
			events.DealsActivations = append(events.DealsActivations, dealActivations...)
			appendMappings(tx, activationMappings(tx, dealActivations))
			for i := range dealSpaceInfo {
				dealSpaceInfo[i].ID = dealsSpaceInfoIds.Unique(dealSpaceInfo[i].ID, tx.Id)
			}
			events.DealsSpaceInfo = append(events.DealsSpaceInfo, dealSpaceInfo...)
		}
		if eg.isDealTermination(tx) {
			dealTerminations, err := eg.createDealTerminations(tx)
			if err != nil {
				// Same rule as activations: unusable metadata (e.g. raw-only after a decode
				// failure) skips the terminations of this tx only, never the height.
				eg.logger.Warnf("skipping deal terminations of tx %s (tx type %s, height %d): %s", tx.TxCid, tx.TxType, tx.Height, err.Error())
				_ = eg.metrics.UpdateSkippedTerminationsMetric(tx.TxType)
				continue
			}
			for i := range dealTerminations {
				dealTerminations[i].ID = dealsTerminationsIds.Unique(dealTerminations[i].ID, tx.Id)
			}
			events.DealsTerminations = append(events.DealsTerminations, dealTerminations...)
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
		// market method 6 is named BatchActivateDeals from actors v12 (NV21) on
		strings.EqualFold(txType, parser.MethodBatchActivateDeals),
		strings.EqualFold(txType, parser.MethodSettleDealPaymentsExported),
		strings.EqualFold(txType, parser.MethodSectorContentChanged),
		strings.EqualFold(txType, parser.MethodOnMinerSectorsTerminate):
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
