package evm

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/zondax/fil-parser/actors/metrics"
	"github.com/zondax/fil-parser/actors/v2/evm/types"
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
	"github.com/zondax/fil-parser/actors/v2/miner"
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

func customMethods(e *Evm) map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		// This is a miner method verified from testing with CID: bafy2bzacealgb5zr5g2cc5emi7yc2mpragoufvt5lm54xzdkhdorpfgjhbshi on calibration.
		abi.MethodNum(23): {
			Name:   parser.MethodChangeOwnerAddressExported,
			Method: e.ChangeOwnerAddress,
		},

		// This is a miner method verified from testing with CID: f3vmqpcytevkwn6fktjd2zelo4lftq6xzsb2vnmp2r3qarbr4vnso7c7y3nqi5gmxifp22m2pbqdctfxrwkmga on calibration.
		abi.MethodNum(18): {
			Name:   parser.MethodChangeMultiaddrsExported,
			Method: e.ChangeMultiAddrs,
		},
		// https://github.com/filecoin-project/ref-fvm/issues/835#issuecomment-1236096270
		abi.MethodNum(0): {
			Name:   parser.MethodValueTransfer,
			Method: e.ValueTransfer,
		},
	}
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V18.String(): actors.CopyMethods(evmv10.Methods, customMethods(&Evm{})),
	tools.V19.String(): actors.CopyMethods(evmv11.Methods, customMethods(&Evm{})),
	tools.V20.String(): actors.CopyMethods(evmv11.Methods, customMethods(&Evm{})),
	tools.V21.String(): actors.CopyMethods(evmv12.Methods, customMethods(&Evm{})),
	tools.V22.String(): actors.CopyMethods(evmv13.Methods, customMethods(&Evm{})),
	tools.V23.String(): actors.CopyMethods(evmv14.Methods, customMethods(&Evm{})),
	tools.V24.String(): actors.CopyMethods(evmv15.Methods, customMethods(&Evm{})),
	tools.V25.String(): actors.CopyMethods(evmv16.Methods, customMethods(&Evm{})),
}

func (e *Evm) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return methods, nil
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

func (*Evm) HandleFilecoinMethod(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata, err := parse(rawParams, rawReturn, true, &types.HandleFilecoinMethodParams{}, &types.HandleFilecoinMethodReturn{}, parser.ParamsKey)
	if err != nil {
		if metadata == nil {
			metadata = make(map[string]interface{})
		}
		metadata[parser.ParamsRawKey] = hex.EncodeToString(rawParams)
		metadata[parser.ReturnRawKey] = hex.EncodeToString(rawReturn)
		return metadata, nil
	}
	return metadata, nil

}

func (e *Evm) ChangeOwnerAddress(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	m := miner.New(e.logger)
	return m.ChangeOwnerAddressExported(network, height, rawParams)
}

func (e *Evm) ChangeMultiAddrs(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	m := miner.New(e.logger)
	return m.ChangeMultiaddrsExported(network, height, rawParams)
}

func (e *Evm) ValueTransfer(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
