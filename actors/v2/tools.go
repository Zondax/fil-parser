package v2

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/actors/metrics"
	"github.com/zondax/fil-parser/actors/v2/reward"
	metrics2 "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
)

func GetMethodName(ctx context.Context, methodNum abi.MethodNum, actorName string, height int64, network string, helper *helper.Helper, logger *logger.Logger) (string, error) {
	actorMethods, err := ActorMethods(ctx, actorName, height, network, helper, logger)
	if err != nil {
		return "", err
	}

	method, ok := actorMethods[methodNum]
	if !ok {
		version := tools.VersionFromHeight(network, height)

		if strings.Contains(actorName, manifest.PlaceholderKey) {
			return parser.MethodFallback, nil
		}

		if strings.Contains(actorName, manifest.AccountKey) {
			if version.NodeVersion() == tools.V17.NodeVersion() {
				// https://github.com/filecoin-project/builtin-actors/blob/0c3720c05da4733c3a5ed39c124bc8027c143aa8/actors/account/src/lib.rs#L107
				return parser.MethodUniversalReceiverHook, nil
			} else if version.NodeVersion() > tools.V17.NodeVersion() {
				// https://github.com/filecoin-project/builtin-actors/blob/8fdbdec5e3f46b60ba0132d90533783a44c5961f/actors/account/src/lib.rs#L96
				if methodNum >= abi.MethodNum(parser.FirstExportedMethodNumber) {
					return parser.MethodFallback, nil
				}
			}
		}

		// handle fallback methods
		if strings.Contains(actorName, manifest.EthAccountKey) {
			//https://github.com/filecoin-project/builtin-actors/blob/8fdbdec5e3f46b60ba0132d90533783a44c5961f/actors/ethaccount/src/lib.rs#L54
			if methodNum >= abi.MethodNum(parser.FirstExportedMethodNumber) {
				return parser.MethodFallback, nil
			}
		}
		if strings.Contains(actorName, manifest.EvmKey) {
			// https://github.com/filecoin-project/builtin-actors/blob/8fdbdec5e3f46b60ba0132d90533783a44c5961f/actors/evm/src/lib.rs#L266
			if methodNum > abi.MethodNum(parser.EvmMaxReservedMethodNumber) {
				return parser.MethodHandleFilecoinMethod, nil
			}
		}

		if strings.Contains(actorName, manifest.MultisigKey) {
			// https://github.com/filecoin-project/builtin-actors/blob/b86938e410daebf27f9397fd622370a16b24f58b/actors/multisig/src/lib.rs#L439
			if methodNum >= abi.MethodNum(parser.FirstExportedMethodNumber) {
				return parser.MethodFallback, nil
			}
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

func GetBlockCidFromMsgCid(msgCid, txType string, txMetadata map[string]interface{}, tipset *types.ExtendedTipSet, logger *logger.Logger) (string, error) {
	// Default value
	blockCid := ""

	// Process the special cases first were this kind of txs are not explicitly included in a block
	switch txType {
	case parser.MethodAwardBlockReward:
		if txMetadata == nil {
			return blockCid, fmt.Errorf("received tx of type '%s' with nil metadata", txType)
		}
		// Get the miner that received the reward
		params, ok := txMetadata["Params"]
		if !ok {
			logger.Errorf("Could no get paramater 'Params' inside tx '%s'", txType)
			return blockCid, nil
		}
		miner := reward.GetMinerFromAwardBlockRewardParams(params)
		if miner == "" {
			logger.Errorf("Could not parse parameters for tx '%s', param type: %T", txType, params)
			return blockCid, nil
		}
		// Get the block that this miner mined
		c, err := tipset.GetBlockMinedByMiner(miner)
		if err != nil {
			return blockCid, err
		}
		return c, nil
	case parser.MethodApplyRewards, parser.MethodUpdatePledgeTotal, parser.MethodCronTick,
		parser.MethodEpochTick, parser.MethodThisEpochReward, parser.MethodConfirmSectorProofsValid,
		parser.MethodActivateDeals, parser.MethodClaimAllocations, parser.MethodBurnExported,
		parser.MethodEnrollCronEvent, parser.MethodOnDeferredCronEvent, parser.MethodUpdateNetworkKPI,
		parser.MethodFallback, parser.MethodHandleFilecoinMethod:
		// These txs are not included in a block
		return blockCid, nil
	}

	blockCids, ok := tipset.BlockMessages[msgCid]
	if !ok {
		return blockCid, errors.New("could not find block hash for message cid")
	}

	if len(blockCids) == 0 {
		return blockCid, errors.New("could not find block hash for message cid. Slice is empty")
	} else {
		blockCid = blockCids[0].Cid
	}

	return blockCid, nil
}
