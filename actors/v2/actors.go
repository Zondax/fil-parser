package v2

import (
	"context"
	"fmt"

	"github.com/zondax/fil-parser/actors/v2/internal"

	actormetrics "github.com/zondax/fil-parser/actors/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/actors/v2/multisig"
	"github.com/zondax/golem/pkg/logger"

	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
)

type Actor = actors.Actor

var _ actors.ActorParserInterface = &ActorParser{}

type ActorParser struct {
	network string
	helper  *helper.Helper
	logger  *logger.Logger
	metrics *actormetrics.ActorsMetricsClient
}

func NewActorParser(network string, helper *helper.Helper, logger *logger.Logger, metrics metrics2.MetricsClient) actors.ActorParserInterface {
	return &ActorParser{
		network: network,
		helper:  helper,
		logger:  logger2.GetSafeLogger(logger),
		metrics: actormetrics.NewClient(metrics, "actorV2"),
	}
}

func (p *ActorParser) GetMetadata(ctx context.Context, txType string, msg *parser.LotusMessage, mainMsgCid cid.Cid, msgRct *parser.LotusMessageReceipt,
	height int64, key filTypes.TipSetKey) (string, map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	if msg == nil {
		return "", metadata, nil, nil
	}

	actor, err := p.helper.GetActorNameFromAddress(msg.To, height, key)
	if err != nil {
		return "", metadata, nil, fmt.Errorf("error getting actor name from address: %w", err)
	}
	actorParser, err := p.GetActor(actor)
	if err != nil {
		return actor, nil, nil, parser.ErrNotValidActor
	}
	metadata, addressInfo, err := actorParser.Parse(ctx, p.network, height, txType, msg, msgRct, mainMsgCid, key)
	return actor, metadata, addressInfo, err
}

func (p *ActorParser) LatestSupportedVersion(actor string) (uint64, error) {
	keys := manifest.GetBuiltinActorsKeys(10)

	for _, key := range keys {
		if key == actor {
			return 10, nil
		}
	}
	return 0, nil
}

func (p *ActorParser) GetActor(actor string) (Actor, error) {
	if actor == manifest.MultisigKey {
		return multisig.New(p.helper, p.logger, p.metrics), nil
	}

	return internal.GetActor(actor, p.logger, p.helper, p.metrics)
}
