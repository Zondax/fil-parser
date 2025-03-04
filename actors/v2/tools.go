package v2

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"go.uber.org/zap"
)

func GetMethodName(methodNum abi.MethodNum, actorName string, height int64, network string, helper *helper.Helper, logger *zap.Logger) (string, error) {
	actorMethods, err := ActorMethods(actorName, height, network, helper, logger)
	if err != nil {
		return "", err
	}

	method, ok := actorMethods[methodNum]
	if !ok {
		return parser.UnknownStr, nil
	}
	return method.Name, nil
}

// EthAccount and Placeholder can receive tokens with Send and InvokeEVM methods
// We set evm.Methods instead of empty array of methods. Therefore, we will be able to understand
// this specific method (3844450837) - tx cid example: bafy2bzacedgmcvsp56ieciutvgwza2qpvz7pvbhhu4l5y5tdl35rwfnjn5buk
func ActorMethods(actorName string, height int64, network string, helper *helper.Helper, logger *zap.Logger) (actorMethods map[abi.MethodNum]builtin.MethodMeta, err error) {
	mActorName := actorName
	actorParser := &ActorParser{network, helper, logger}
	if actorName == manifest.EthAccountKey || actorName == manifest.PlaceholderKey {
		mActorName = manifest.EvmKey
	}
	actor, err := actorParser.GetActor(mActorName)
	if err != nil {
		return nil, err
	}

	actorMethods, err = actor.Methods(network, height)
	if err != nil {
		return nil, err
	}
	return actorMethods, nil
}
