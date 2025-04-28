package actor_tools

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/ipfs/go-cid"
	actormetrics "github.com/zondax/fil-parser/actors/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/logger"

	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/actors/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
)

var ErrUnsupportedHeight = errors.New("unsupported height")
var ErrInvalidHeightForMethod = errors.New("invalid height for method")

type ActorParserInterface interface {
	GetMetadata(ctx context.Context, txType string, msg *parser.LotusMessage, mainMsgCid cid.Cid, msgRct *parser.LotusMessageReceipt,
		height int64, key filTypes.TipSetKey) (actor string, metadata map[string]interface{}, addressInfo *types.AddressInfo, err error)
	GetActor(actor string, metrics *actormetrics.ActorsMetricsClient) (Actor, error)
	AllActors() []string
}

type Actor interface {
	Name() string
	Parse(ctx context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error)
	StartNetworkHeight() int64
	TransactionTypes() map[string]any
	Methods(ctx context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error)
}

type MethodNamesFn func(ctx context.Context, actorName, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error)

func GetMethodName(ctx context.Context, methodNum abi.MethodNum, actorName string, height int64, network string, helper *helper.Helper, logger *logger.Logger, actorParser ActorParserInterface) (string, error) {
	actorMethods, err := ActorMethods(ctx, actorName, height, network, helper, logger, actorParser)
	if err != nil {
		return "", err
	}

	method, ok := actorMethods[methodNum]
	if !ok {
		if (actorName == manifest.AccountKey || actorName == manifest.EthAccountKey) && methodNum >= abi.MethodNum(parser.FirstExportedMethodNumber) {
			return parser.MethodFallback, nil
		}
		if actorName == manifest.EvmKey && methodNum > abi.MethodNum(parser.EvmMaxReservedMethodNumber) {
			return parser.MethodInvokeContractFilecoinHandler, nil
		}

		return parser.UnknownStr, nil
	}
	return method.Name, nil
}

// EthAccount and Placeholder can receive tokens with Send and InvokeEVM methods
// We set evm.Methods instead of empty array of methods. Therefore, we will be able to understand
// this specific method (3844450837) - tx cid example: bafy2bzacedgmcvsp56ieciutvgwza2qpvz7pvbhhu4l5y5tdl35rwfnjn5buk
func ActorMethods(ctx context.Context, actorName string, height int64, network string, helper *helper.Helper, logger *logger.Logger, actorParser ActorParserInterface) (actorMethods map[abi.MethodNum]builtin.MethodMeta, err error) {
	metricsClient := &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()}
	mActorName := actorName
	if actorName == manifest.EthAccountKey || actorName == manifest.PlaceholderKey {
		mActorName = manifest.EvmKey
	}
	actor, err := actorParser.GetActor(mActorName, metricsClient)
	if err != nil {
		return nil, err
	}

	actorMethods, err = actor.Methods(ctx, network, height)
	if err != nil {
		return nil, err
	}
	return actorMethods, nil
}

func ParseSend(msg *parser.LotusMessage) map[string]interface{} {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = msg.Params
	return metadata
}

// ParseConstructor parse methods with format: *new(func(*address.Address) *abi.EmptyValue)
func ParseConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	return metadata, nil
}

func ParseUnknownMetadata(msgParams, msgReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if len(msgParams) > 0 {
		metadata[parser.ParamsKey] = hex.EncodeToString(msgParams)
	}
	if len(msgReturn) > 0 {
		metadata[parser.ReturnKey] = hex.EncodeToString(msgReturn)
	}
	return metadata, nil
}

func ParseEmptyParamsAndReturn() (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}
