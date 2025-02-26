package actors

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-state-types/manifest"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v11/miner"
	"github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	multisig2 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/filecoin-project/go-state-types/exitcode"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
)

/*
Still needs to parse:

	Receive
*/
func (p *ActorParser) ParseMultisig(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodConstructor: // TODO: not tested
		return p.msigConstructor(msg.Params)
	case parser.MethodSend:
		return p.parseSend(msg), nil
	case parser.MethodPropose, parser.MethodProposeExported:
		return p.propose(msg.Params, msgRct.Return)
	case parser.MethodApprove, parser.MethodApproveExported:
		return p.approve(msg, msgRct.Return, txType, key)
	case parser.MethodCancel, parser.MethodCancelExported:
		return p.cancel(msg, txType, key)
	case parser.MethodAddSigner, parser.MethodAddSignerExported, parser.MethodSwapSigner, parser.MethodSwapSignerExported:
		return p.msigParams(msg, txType, key)
	case parser.MethodRemoveSigner, parser.MethodRemoveSignerExported:
		return p.removeSigner(msg, txType, key)
	case parser.MethodChangeNumApprovalsThreshold, parser.MethodChangeNumApprovalsThresholdExported:
		return p.changeNumApprovalsThreshold(msg.Params)
	case parser.MethodLockBalance, parser.MethodLockBalanceExported:
		return p.lockBalance(msg.Params)
	case parser.MethodMsigUniversalReceiverHook: // TODO: not tested
		return p.universalReceiverHook(msg.Params)
	case parser.UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func (p *ActorParser) msigConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var proposeParams multisig.ConstructorParams
	err := proposeParams.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = proposeParams
	return metadata, nil
}

func (p *ActorParser) msigParams(msg *parser.LotusMessage, method string, key filTypes.TipSetKey) (map[string]interface{}, error) {
	params, err := p.parseMsigParams(msg, method, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	var paramsMap map[string]interface{}
	err = json.Unmarshal([]byte(params), &paramsMap)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return paramsMap, nil
}

func (p *ActorParser) propose(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var proposeParams multisig.ProposeParams
	reader := bytes.NewReader(rawParams)
	err := proposeParams.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	method, innerParams, err := p.innerProposeParams(proposeParams)
	if err != nil {
		_ = p.metrics.UpdateActorMethodErrorMetric(manifest.MultisigKey, parser.MethodPropose, err)
		p.logger.Sugar().Errorf("could not decode multisig inner params. Method: %v. Err: %v", proposeParams.Method.String(), err)
	}
	metadata[parser.ParamsKey] = parser.Propose{
		To:     proposeParams.To.String(),
		Value:  proposeParams.Value.String(),
		Method: method,
		Params: innerParams,
	}
	var proposeReturn multisig.ProposeReturn
	reader = bytes.NewReader(rawReturn)
	err = proposeReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = proposeReturn
	return metadata, nil
}

func (p *ActorParser) approve(msg *parser.LotusMessage, rawReturn []byte, method string, key filTypes.TipSetKey) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := p.parseMsigParams(msg, method, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	reader := bytes.NewReader(rawReturn)
	var approveReturn multisig.ApproveReturn
	err = approveReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = approveReturn
	return metadata, nil
}

func (p *ActorParser) cancel(msg *parser.LotusMessage, method string, key filTypes.TipSetKey) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := p.parseMsigParams(msg, method, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func (p *ActorParser) removeSigner(msg *parser.LotusMessage, method string, key filTypes.TipSetKey) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := p.parseMsigParams(msg, method, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func (p *ActorParser) changeNumApprovalsThreshold(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params multisig.ChangeNumApprovalsThresholdParams
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func (p *ActorParser) lockBalance(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params multisig.LockBalanceParams
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func (p *ActorParser) parseMsigParams(msg *parser.LotusMessage, method string, key filTypes.TipSetKey) (string, error) {
	msgSerial, err := msg.MarshalJSON() // TODO: this may not work properly
	if err != nil {
		p.logger.Sugar().Errorf("Could not parse params. Cannot serialize lotus message: %v", err)
		return "", err
	}

	actorCode, err := p.helper.GetActorsCache().GetActorCode(msg.To, key, false)
	if err != nil {
		return "", err
	}

	c, err := cid.Parse(actorCode)
	if err != nil {
		p.logger.Sugar().Errorf("Could not parse params. Cannot cid.parse actor code: %v", err)
		return "", err
	}
	parsedParams, err := p.helper.GetFilecoinLib().ParseParamsMultisigTx(string(msgSerial), c)
	if err != nil {
		p.logger.Sugar().Errorf("Could not parse params. ParseParamsMultisigTx returned with error: %v", err)
		return "", err
	}

	return parsedParams, nil
}

func (p *ActorParser) universalReceiverHook(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params abi.CborBytesTransparent
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func (p *ActorParser) innerProposeParams(propose multisig.ProposeParams) (string, cbor.Unmarshaler, error) {
	reader := bytes.NewReader(propose.Params)
	switch propose.Method {
	case builtin.MethodSend:
		if propose.Params == nil {
			return parser.MethodSend, nil, nil
		}
		var params multisig.ProposeParams // TODO: is this correct?
		err := params.UnmarshalCBOR(reader)
		return parser.MethodSend, &params, err
	case builtin.MethodsMultisig.Approve,
		builtin.MethodsMultisig.Cancel:
		var params multisig.TxnIDParams
		err := params.UnmarshalCBOR(reader)
		return parser.MethodApprove, &params, err
	case builtin.MethodsMultisig.AddSigner:
		var params multisig.AddSignerParams
		err := params.UnmarshalCBOR(reader)
		return parser.MethodAddSigner, &params, err
	case builtin.MethodsMultisig.RemoveSigner:
		var params multisig.RemoveSignerParams
		err := params.UnmarshalCBOR(reader)
		return parser.MethodRemoveSigner, &params, err
	case builtin.MethodsMultisig.SwapSigner:
		var params multisig.SwapSignerParams
		err := params.UnmarshalCBOR(reader)
		return parser.MethodSwapSigner, &params, err
	case builtin.MethodsMultisig.ChangeNumApprovalsThreshold:
		var params multisig.ChangeNumApprovalsThresholdParams
		err := params.UnmarshalCBOR(reader)
		return parser.MethodChangeNumApprovalsThreshold, &params, err
	case builtin.MethodsMultisig.LockBalance:
		var params multisig.LockBalanceParams
		err := params.UnmarshalCBOR(reader)
		return parser.MethodLockBalance, &params, err
	case builtin.MethodsMiner.WithdrawBalance:
		var params miner.WithdrawBalanceParams
		err := params.UnmarshalCBOR(reader)
		return parser.MethodWithdrawBalance, &params, err
	case builtin.MethodsVerifiedRegistry.AddVerifier:
		var params verifreg.AddVerifierParams
		err := params.UnmarshalCBOR(reader)
		return parser.MethodAddVerifier, &params, err
	}
	return "", nil, parser.ErrUnknownMethod
}

// TODO: Definitely we need refactor the pkg actors

func ParseMultisigMetadata(txType string, txMetadata string) (interface{}, error) {
	deserializationFuncs := map[string]func(string) (interface{}, error){
		parser.MethodAddSigner:                           parseAddSignerValue,
		parser.MethodApprove:                             parseApproveValue,
		parser.MethodCancel:                              parseCancelValue,
		parser.MethodChangeNumApprovalsThreshold:         parseChangeNumApprovalsThresholdValue,
		parser.MethodConstructor:                         parseConstructorValue,
		parser.MethodLockBalance:                         parseLockBalanceValue,
		parser.MethodRemoveSigner:                        parseRemoveSignerValue,
		parser.MethodSend:                                parseSendValue,
		parser.MethodSwapSigner:                          parseSwapSignerValue,
		parser.MethodAddVerifier:                         parseAddVerifierValue,
		parser.MethodChangeOwnerAddress:                  parseChangeOwnerAddressValue,
		parser.MethodWithdrawBalance:                     parseWithdrawBalanceValue,
		parser.MethodInvokeContract:                      parseInvokeContractValue,
		parser.MethodApproveExported:                     parseApproveValue,
		parser.MethodCancelExported:                      parseCancelValue,
		parser.MethodAddSignerExported:                   parseAddSignerValue,
		parser.MethodSwapSignerExported:                  parseSwapSignerValue,
		parser.MethodRemoveSignerExported:                parseRemoveSignerValue,
		parser.MethodChangeNumApprovalsThresholdExported: parseChangeNumApprovalsThresholdValue,
		parser.MethodLockBalanceExported:                 parseLockBalanceValue,
		parser.MethodMsigUniversalReceiverHook:           parseUniversalReceiverHookValue,
		parser.MethodChangeOwnerAddressExported:          parseChangeOwnerAddressValue,
		parser.MethodWithdrawBalanceExported:             parseWithdrawBalanceValue,
	}

	if parseFunc, found := deserializationFuncs[txType]; found {
		return parseFunc(txMetadata)
	}

	return nil, fmt.Errorf("unknown tx type: %s", txType)
}

func parseAddVerifierValue(txMetadata string) (interface{}, error) {
	var v verifreg.AddVerifierParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseChangeOwnerAddressValue(txMetadata string) (interface{}, error) {
	var v ChangeOwnerAddressParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseWithdrawBalanceValue(txMetadata string) (interface{}, error) {
	var v miner.WithdrawBalanceParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseInvokeContractValue(txMetadata string) (interface{}, error) {
	var v InvokeContractParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseAddSignerValue(txMetadata string) (interface{}, error) {
	var v multisig2.AddSignerParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseApproveValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsRaw, ok := raw["Params"].(string)
	if !ok {
		return nil, fmt.Errorf("Params not found or not a string")
	}

	var params ApproveValue
	err = json.Unmarshal([]byte(paramsRaw), &params)
	if err != nil {
		return nil, err
	}

	returnRaw, ok := raw["Return"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Return not found or not a map")
	}

	applied, ok := returnRaw["Applied"].(bool)
	if !ok {
		return nil, fmt.Errorf("Applied not found or not a bool")
	}

	code, ok := returnRaw["Code"].(float64)
	if !ok {
		return nil, fmt.Errorf("Code not found or not a float64")
	}

	ret, ok := returnRaw["Ret"].(string)
	if !ok {
		return nil, fmt.Errorf("Ret not found or not a string")
	}

	params.Return = multisig2.ApproveReturn{
		Applied: applied,
		Code:    exitcode.ExitCode(code),
		Ret:     []byte(ret),
	}

	return params, nil
}

func parseCancelValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsStr, ok := raw["Params"].(string)
	if !ok {
		return nil, fmt.Errorf("Params not found or not a string")
	}

	var paramsRaw map[string]interface{}
	err = json.Unmarshal([]byte(paramsStr), &paramsRaw)
	if err != nil {
		return nil, err
	}

	var v CancelValue
	err = mapToStruct(paramsRaw, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func parseChangeNumApprovalsThresholdValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsRaw, ok := raw["Params"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("params not found or not a map")
	}

	var v multisig2.ChangeNumApprovalsThresholdParams
	if newThreshold, ok := paramsRaw["NewThreshold"].(float64); ok {
		v.NewThreshold = uint64(newThreshold)
	} else {
		return nil, fmt.Errorf("NewThreshold not found or not a number")
	}

	return v, nil
}

func parseConstructorValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsRaw, ok := raw["Params"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("params not found or not a map")
	}

	var v multisig2.ConstructorParams
	err = mapToStruct(paramsRaw, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func parseLockBalanceValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsRaw, ok := raw["Params"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("params not found or not a map")
	}

	var v multisig2.LockBalanceParams
	err = mapToStruct(paramsRaw, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func parseRemoveSignerValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}
	paramsRaw, ok := raw["Params"].(string)
	if !ok {
		return nil, fmt.Errorf("Params not found or not a string")
	}
	var params multisig2.RemoveSignerParams
	err = json.Unmarshal([]byte(paramsRaw), &params)
	if err != nil {
		return nil, err
	}
	return params, nil
}

func parseSendValue(txMetadata string) (interface{}, error) {
	var v SendValue
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseSwapSignerValue(txMetadata string) (interface{}, error) {
	var v multisig2.SwapSignerParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseUniversalReceiverHookValue(txMetadata string) (interface{}, error) {
	var tx TransactionUniversalReceiverHookMetadata
	err := json.Unmarshal([]byte(txMetadata), &tx)
	if err != nil {
		return nil, err
	}

	var params UniversalReceiverHookParams
	err = json.Unmarshal([]byte(tx.Params), &params)
	if err != nil {
		return nil, err
	}

	result := UniversalReceiverHookValue{
		Type:    uint64(params.Type_),
		Payload: params.Payload,
		Return:  tx.Return,
	}

	return result, nil
}

func mapToStruct(m map[string]interface{}, v interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
