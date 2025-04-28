package actors

import (
	"context"
	"fmt"

	actormetrics "github.com/zondax/fil-parser/actors/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"

	"github.com/zondax/fil-parser/actors/v2/account"
	"github.com/zondax/fil-parser/actors/v2/cron"
	"github.com/zondax/fil-parser/actors/v2/datacap"
	"github.com/zondax/fil-parser/actors/v2/eam"
	"github.com/zondax/fil-parser/actors/v2/ethaccount"
	"github.com/zondax/fil-parser/actors/v2/evm"
	initActor "github.com/zondax/fil-parser/actors/v2/init"
	"github.com/zondax/fil-parser/actors/v2/market"
	"github.com/zondax/fil-parser/actors/v2/miner"
	"github.com/zondax/fil-parser/actors/v2/multisig"
	paymentchannel "github.com/zondax/fil-parser/actors/v2/paymentChannel"
	"github.com/zondax/fil-parser/actors/v2/placeholder"
	"github.com/zondax/fil-parser/actors/v2/power"
	"github.com/zondax/fil-parser/actors/v2/reward"
	"github.com/zondax/fil-parser/actors/v2/system"
	verifiedregistry "github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	"github.com/zondax/golem/pkg/logger"

	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
)

type Actor interface {
	Name() string
	Parse(ctx context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error)
	StartNetworkHeight() int64
	TransactionTypes() map[string]any
	Methods(ctx context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error)
}

var _ actor_tools.ActorParserInterface = &ActorParser{}

type ActorParser struct {
	network string
	helper  *helper.Helper
	logger  *logger.Logger
	metrics *actormetrics.ActorsMetricsClient
}

func NewActorParser(network string, helper *helper.Helper, logger *logger.Logger, metrics metrics2.MetricsClient) actor_tools.ActorParserInterface {
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
	actorParser, err := p.GetActor(actor, p.metrics)
	if err != nil {
		return actor, nil, nil, parser.ErrNotValidActor
	}
	if txType == parser.UnknownStr {
		// https: //github.com/filecoin-project/builtin-actors/blob/8fdbdec5e3f46b60ba0132d90533783a44c5961f/actors/account/src/lib.rs#L96
		if actor == manifest.AccountKey || actor == manifest.EthAccountKey {
			if msg.Method > parser.FirstExportedMethodNumber {
				fmt.Println("Using fallback for account actor")
				txType = parser.MethodFallback
			}
		} else {
			fromActor, err := p.helper.GetActorNameFromAddress(msg.From, height, key)
			if err != nil {
				return "", metadata, nil, fmt.Errorf("error getting actor name from address: %w", err)
			}
			fmt.Println("--------------------------------")
			fmt.Println("Method Num: ", msg.Method)
			fmt.Println("Actual Actor To:", actor)
			fmt.Println("Actual Actor From:", fromActor)
			fmt.Println("--------------------------------")
			// try all
			actors := []string{
				manifest.AccountKey,
				manifest.CronKey,
				manifest.DatacapKey,
				manifest.EamKey,
				manifest.EthAccountKey,
				manifest.InitKey,
				manifest.MarketKey,
				manifest.MinerKey,
				manifest.MultisigKey,
				manifest.PaychKey,
				manifest.PowerKey,
				manifest.RewardKey,
				manifest.VerifregKey,
				manifest.PlaceholderKey,
				manifest.SystemKey,
			}
			for _, actor := range actors {
				fmt.Println("trying actor", actor)
				actorParser, _ := p.GetActor(actor, p.metrics)
				methods, _ := actorParser.Methods(ctx, p.network, height)
				for _, method := range methods {
					fmt.Println("	trying method", method.Name)
					metadata, _, err := actorParser.Parse(ctx, p.network, height, method.Name, msg, msgRct, mainMsgCid, key)
					if err != nil {
						fmt.Println("		failed: ", err)
						continue
					}
					fmt.Println("	success: ", metadata)
					// return actor, metadata, addressInfo, nil
				}
			}
		}
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

func (p *ActorParser) GetActor(actor string, metrics *actormetrics.ActorsMetricsClient) (actor_tools.Actor, error) {
	switch actor {
	case manifest.AccountKey:
		return account.New(p.logger), nil
	case manifest.CronKey:
		return cron.New(p.logger), nil
	case manifest.DatacapKey:
		return datacap.New(p.logger), nil
	case manifest.EamKey:
		return eam.New(p.helper, p.logger), nil
	case manifest.EthAccountKey:
		return ethaccount.New(p.logger), nil
	case manifest.EvmKey:
		return evm.New(p.helper, p.logger, p.metrics, p), nil
	case manifest.InitKey:
		return initActor.New(p.helper, p.logger), nil
	case manifest.MarketKey:
		return market.New(p.logger), nil
	case manifest.MinerKey:
		return miner.New(p.logger), nil
	case manifest.MultisigKey:
		return multisig.New(p.helper, p.logger, metrics), nil
	case manifest.PaychKey:
		return paymentchannel.New(p.logger), nil
	case manifest.PowerKey:
		return power.New(p.helper, p.logger), nil
	case manifest.RewardKey:
		return reward.New(p.logger), nil
	case manifest.VerifregKey:
		return verifiedregistry.New(p.logger), nil
	case manifest.PlaceholderKey:
		return placeholder.New(p.logger), nil
	case manifest.SystemKey:
		return system.New(p.logger), nil
	default:
		return nil, fmt.Errorf("actor %s not found", actor)
	}
}

func (p *ActorParser) AllActors() []string {
	return []string{
		manifest.AccountKey,
		manifest.CronKey,
		manifest.DatacapKey,
		manifest.EamKey,
		manifest.EvmKey,
		manifest.EthAccountKey,
		manifest.InitKey,
		manifest.MarketKey,
		manifest.MinerKey,
		manifest.MultisigKey,
		manifest.PaychKey,
		manifest.PowerKey,
		manifest.RewardKey,
		manifest.VerifregKey,
		manifest.PlaceholderKey,
		manifest.SystemKey,
	}
}
