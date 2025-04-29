package evm

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/zondax/fil-parser/actors/metrics"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"

	evmv10 "github.com/filecoin-project/go-state-types/builtin/v10/evm"
	evmv11 "github.com/filecoin-project/go-state-types/builtin/v11/evm"
	evmv12 "github.com/filecoin-project/go-state-types/builtin/v12/evm"
	evmv13 "github.com/filecoin-project/go-state-types/builtin/v13/evm"
	evmv14 "github.com/filecoin-project/go-state-types/builtin/v14/evm"
	evmv15 "github.com/filecoin-project/go-state-types/builtin/v15/evm"
	evmv16 "github.com/filecoin-project/go-state-types/builtin/v16/evm"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Evm struct {
	logger  *logger.Logger
	metrics *metrics.ActorsMetricsClient
}

func New(logger *logger.Logger, metrics *metrics.ActorsMetricsClient) *Evm {
	return &Evm{
		logger:  logger,
		metrics: metrics,
	}
}

func (e *Evm) Name() string {
	return manifest.EvmKey
}

func (*Evm) StartNetworkHeight() int64 {
	return tools.V18.Height()
}

func (e *Evm) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	case tools.V18.IsSupported(network, height):
		return evmv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return evmv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return evmv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return evmv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return evmv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return evmv15.Methods, nil
	case tools.V25.IsSupported(network, height):
		return evmv16.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
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
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Evm) InvokeContractDelegate(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := delegateCallParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(rawParams, rawReturn, true, params(), &abi.CborBytes{}, parser.ParamsKey)
}

func (*Evm) GetBytecode(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := getBytecodeReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
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
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.CborBytes{}, parser.ParamsKey)
}

func (*Evm) GetStorageAt(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getStorageAtParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(rawParams, rawReturn, true, params(), &abi.CborBytes{}, parser.ParamsKey)
}
