package evm

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/metrics"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"

	evmv10 "github.com/filecoin-project/go-state-types/builtin/v10/evm"
	evmv11 "github.com/filecoin-project/go-state-types/builtin/v11/evm"
	evmv12 "github.com/filecoin-project/go-state-types/builtin/v12/evm"
	evmv13 "github.com/filecoin-project/go-state-types/builtin/v13/evm"
	evmv14 "github.com/filecoin-project/go-state-types/builtin/v14/evm"
	evmv15 "github.com/filecoin-project/go-state-types/builtin/v15/evm"
	evmv16 "github.com/filecoin-project/go-state-types/builtin/v16/evm"

	"github.com/zondax/fil-parser/actors/v2/evm/types"
	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Evm struct {
	logger      *logger.Logger
	metrics     *metrics.ActorsMetricsClient
	helper      *helper.Helper
	actorParser actor_tools.ActorParserInterface
}

func New(helper *helper.Helper, logger *logger.Logger, metrics *metrics.ActorsMetricsClient, actorParser actor_tools.ActorParserInterface) *Evm {
	return &Evm{
		logger:      logger,
		metrics:     metrics,
		helper:      helper,
		actorParser: actorParser,
	}
}

func (e *Evm) Name() string {
	return manifest.EvmKey
}

func (*Evm) StartNetworkHeight() int64 {
	return tools.V18.Height()
}

func (e *Evm) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	var data map[abi.MethodNum]nonLegacyBuiltin.MethodMeta
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{}, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	case tools.V18.IsSupported(network, height):
		data = evmv10.Methods
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data = evmv11.Methods
	case tools.V21.IsSupported(network, height):
		data = evmv12.Methods
	case tools.V22.IsSupported(network, height):
		data = evmv13.Methods
	case tools.V23.IsSupported(network, height):
		data = evmv14.Methods
	case tools.V24.IsSupported(network, height):
		data = evmv15.Methods
	case tools.V25.IsSupported(network, height):
		data = evmv16.Methods
	default:
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
	// +1 because the comparison uses >
	// https://github.com/filecoin-project/builtin-actors/blob/8fdbdec5e3f46b60ba0132d90533783a44c5961f/actors/evm/src/lib.rs#L270
	data[abi.MethodNum(parser.EvmMaxReservedMethodNumber+1)] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodInvokeContractFilecoinHandler,
		Method: e.InvokeContractFilecoinHandler,
	}
	return data, nil
}

func (e *Evm) InvokeContract(network string, height int64, method string, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(rawParams)
	metadata[parser.ReturnKey] = parser.EthPrefix + hex.EncodeToString(rawReturn)

	var params abi.CborBytes
	if err := params.UnmarshalCBOR(reader); err != nil {
		_ = e.metrics.UpdateActorMethodErrorMetric(manifest.EvmKey, method)
		e.logger.Warnf("error deserializing rawParams: %s - hex data: %s", err.Error(), hex.EncodeToString(rawParams))
	}

	if reader.Len() == 0 { // This means that the reader has processed all the bytes
		metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(params)
	}

	reader = bytes.NewReader(rawReturn)
	var returnValue abi.CborBytes
	if err := returnValue.UnmarshalCBOR(reader); err != nil {
		_ = e.metrics.UpdateActorMethodErrorMetric(manifest.EvmKey, method)
		e.logger.Warnf("Error deserializing rawReturn: %s - hex data: %s", err.Error(), hex.EncodeToString(rawReturn))
	}

	if reader.Len() == 0 { // This means that the reader has processed all the bytes
		metadata[parser.ReturnKey] = parser.EthPrefix + hex.EncodeToString(returnValue)
	}
	return metadata, nil
}

func (*Evm) Resurrect(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := resurrectParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Evm) InvokeContractDelegate(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := delegateCallParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}

	return parse(rawParams, rawReturn, true, params(), &abi.CborBytes{}, parser.ParamsKey)
}

func (*Evm) GetBytecode(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := getBytecodeReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, returnValue(), &abi.CborBytes{}, parser.ReturnKey)
}

func (*Evm) GetBytecodeHash(network string, height int64, raw []byte) (map[string]interface{}, error) {
	return parse(raw, nil, false, &abi.CborBytes{}, &abi.CborBytes{}, parser.ReturnKey)
}

func (*Evm) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := constructorParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.CborBytes{}, parser.ParamsKey)
}

func (*Evm) GetStorageAt(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getStorageAtParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}

	return parse(rawParams, rawReturn, true, params(), &abi.CborBytes{}, parser.ParamsKey)
}

func (e *Evm) InvokeContractFilecoinHandler(network string, height int64, msgCid cid.Cid, msgFrom address.Address, tipsetKey filTypes.TipSetKey, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	fmt.Println("rawParams: ", hex.EncodeToString(rawParams))

	params := types.HandleFilecoinMethodParams{}
	returns := types.HandleFilecoinMethodReturn{}

	metadata, err := parse(rawParams, rawReturn, true, &params, &returns, parser.ParamsKey)
	if err != nil {
		return nil, err
	}

	fmt.Println("metadata: ", metadata)

	actorName, err := e.helper.GetActorNameFromAddress(msgFrom, height, tipsetKey)
	if err != nil {
		return nil, err
	}

	fmt.Println("actorName: ", actorName)

	actorMethod, err := actor_tools.GetMethodName(context.Background(), abi.MethodNum(params.Method), actorName, height, network, e.helper, e.logger, e.actorParser)
	if err != nil {
		return nil, err
	}
	actor, err := e.actorParser.GetActor(actorName, e.metrics)
	if err != nil {
		return nil, err
	}

	parsedInnerMetadata, _, err := actor.Parse(context.Background(), network, height, actorMethod, &parser.LotusMessage{
		Params: params.AbiBytes,
	}, &parser.LotusMessageReceipt{
		Return: params.AbiBytes,
	}, msgCid, tipsetKey)

	metadata[parser.ParamsKey] = map[string]interface{}{
		"Args":   parsedInnerMetadata,
		"Method": actorMethod,
	}

	return metadata, nil
}
