package evm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
)

const (
	NumSelectorBytes = 4
)

type EventGenerator interface {
	GenerateEvmEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) ([]types.EvmEvent, error)
}

var _ EventGenerator = &eventGenerator{}

type eventGenerator struct {
	logger    *logger.Logger
	metrics   *evmMetricsClient
	contracts map[string]*abi.ABI
}

func NewEventGenerator(logger *logger.Logger, metrics metrics.MetricsClient, contracts map[string][]byte) (EventGenerator, error) {
	eg := &eventGenerator{
		logger:    logger,
		metrics:   newClient(metrics, manifest.EvmKey),
		contracts: map[string]*abi.ABI{},
	}

	for contractName, contractData := range contracts {
		abi, err := abi.JSON(bytes.NewReader(contractData))
		if err != nil {
			return nil, fmt.Errorf("error parsing contract: %s: %w", contractName, err)
		}
		eg.contracts[contractName] = &abi
	}
	return eg, nil
}

func (eg *eventGenerator) GenerateEvmEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) ([]types.EvmEvent, error) {
	evmEvents := []types.EvmEvent{}
	for _, tx := range transactions {
		if !isEVMContractCall(tx) {
			continue
		}
		decoder, ok := eg.contracts[tx.TxTo]
		if !ok {
			continue
		}

		var value map[string]interface{}
		err := json.Unmarshal([]byte(tx.TxMetadata), &value)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
		}

		params, err := common.GetItem[string](value, parser.ParamsKey, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing params: %w", err)
		}
		ret, err := common.GetItem[string](value, parser.ReturnKey, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing params: %w", err)
		}

		paramData, err := hexutil.Decode(params)
		if err != nil {
			_ = eg.metrics.UpdateDecodeParamMetric()
			return nil, fmt.Errorf("error decoding params: %w", err)
		}
		returnData, err := hexutil.Decode(ret)
		if err != nil {
			_ = eg.metrics.UpdateDecodeReturnMetric()
			return nil, fmt.Errorf("error decoding return data: %w", err)
		}

		method, err := decoder.MethodById(paramData[:NumSelectorBytes])
		if err != nil {
			_ = eg.metrics.UpdateDecodeSelectorMetric()
			return nil, fmt.Errorf("error decoding function selector: %w", err)
		}

		if method.StateMutability == "view" {
			// ignore evm get calls
			continue
		}
		parsedInput := map[string]any{}
		err = method.Inputs.UnpackIntoMap(parsedInput, paramData[4:])
		if err != nil {
			_ = eg.metrics.UpdateUnpackInputsMetric()
			return nil, fmt.Errorf("error unpacking inputs: %w", err)
		}

		parsedOutput := map[string]any{}
		err = method.Outputs.UnpackIntoMap(parsedOutput, returnData)
		if err != nil {
			_ = eg.metrics.UpdateUnpackOutputsMetric()
			return nil, fmt.Errorf("error unpacking outputs: %w", err)
		}
		evmEvents = append(evmEvents, types.EvmEvent{
			Height:          tx.Height,
			TxCid:           tx.TxCid,
			ContractAddress: tx.TxTo,
			FunctionName:    method.Name,
			ParsedMetadata: map[string]any{
				parser.ParamsKey: parsedInput,
				parser.ReturnKey: parsedOutput,
			},
		})
	}

	return evmEvents, nil
}

func isEVMContractCall(tx *types.Transaction) bool {
	switch tx.TxType {
	case parser.MethodInvokeContract, parser.MethodInvokeContractDelegate:
		return true
	}
	return false
}
