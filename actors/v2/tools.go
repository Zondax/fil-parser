package v2

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"

	"github.com/zondax/fil-parser/actors/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/logger"
)

func GetMethodName(ctx context.Context, methodNum abi.MethodNum, actorName string, height int64, network string, helper *helper.Helper, logger *logger.Logger) (string, error) {
	actorMethods, err := ActorMethods(ctx, actorName, height, network, helper, logger)
	if err != nil {
		return "", err
	}

	method, ok := actorMethods[methodNum]
	if !ok {

		if (actorName == manifest.AccountKey || actorName == manifest.EthAccountKey) && methodNum >= abi.MethodNum(parser.FirstExportedMethodNumber) {
			return parser.MethodFallback, nil
		}
		if actorName == manifest.EvmKey && methodNum > abi.MethodNum(parser.EvmMaxReservedMethodNumber) {
			return parser.MethodHandleFilecoinMethod, nil
		}

		return parser.UnknownStr, nil

	}
	return method.Name, nil
}

// EthAccount and Placeholder can receive tokens with Send and InvokeEVM methods
// We set evm.Methods instead of empty array of methods. Therefore, we will be able to understand
// this specific method (3844450837) - tx cid example: bafy2bzacedgmcvsp56ieciutvgwza2qpvz7pvbhhu4l5y5tdl35rwfnjn5buk
func ActorMethods(ctx context.Context, actorName string, height int64, network string, helper *helper.Helper, logger *logger.Logger) (actorMethods map[abi.MethodNum]builtin.MethodMeta, err error) {
	metricsClient := &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()}
	mActorName := actorName
	actorParser := &ActorParser{network, helper, logger, metricsClient}
	if actorName == manifest.EthAccountKey || actorName == manifest.PlaceholderKey {
		mActorName = manifest.EvmKey
	}

	actor, err := actorParser.GetActor(mActorName)
	if err != nil {
		return nil, err
	}

	actorMethods, err = actor.Methods(ctx, network, height)
	if err != nil {
		return nil, err
	}

	return actorMethods, nil
}
