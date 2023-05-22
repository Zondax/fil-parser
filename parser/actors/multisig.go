package actors

import (
	"bytes"
	"encoding/json"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v11/miner"
	"github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	"github.com/filecoin-project/go-state-types/cbor"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
	"go.uber.org/zap"

	"github.com/zondax/fil-parser/database"
)

/*
Still needs to parse:

	Receive
*/
func ParseMultisig(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodConstructor: // TODO: not tested
		return msigConstructor(msg.Params)
	case parser.MethodSend:
		return parseSend(msg), nil
	case parser.MethodPropose, parser.MethodProposeExported:
		return propose(msg, msgRct)
	case parser.MethodApprove, parser.MethodApproveExported:
		return approve(msg, msgRct, height, key)
	case parser.MethodCancel, parser.MethodCancelExported:
		return cancel(msg, height, key)
	case parser.MethodAddSigner, parser.MethodAddSignerExported, parser.MethodSwapSigner, parser.MethodSwapSignerExported:
		return msigParams(msg, height, key)
	case parser.MethodRemoveSigner, parser.MethodRemoveSignerExported:
		return removeSigner(msg, height, key)
	case parser.MethodChangeNumApprovalsThreshold, parser.MethodChangeNumApprovalsThresholdExported:
		return changeNumApprovalsThreshold(msg.Params)
	case parser.MethodAddVerifies: // ?
	case parser.MethodLockBalance, parser.MethodLockBalanceExported:
		return lockBalance(msg.Params)
	case parser.MethodMsigUniversalReceiverHook: // TODO: not tested
		return universalReceiverHook(msg.Params)
	case parser.UnknownStr:
		return unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func msigConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var proposeParams multisig.ConstructorParams
	err := proposeParams.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	return metadata, nil
}

func msigParams(msg *filTypes.Message, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	params, err := parseMsigParams(msg, height, key)
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

func propose(msg *parser.LotusMessage, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var proposeParams multisig.ProposeParams
	reader := bytes.NewReader(msg.Params)
	err := proposeParams.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	method, innerParams, err := innerProposeParams(proposeParams)
	if err != nil {
		zap.S().Errorf("could not decode multisig inner params. Method: %v. Err: %v", proposeParams.Method.String(), err)
	}
	metadata[parser.ParamsKey] = parser.Propose{
		To:     proposeParams.To.String(),
		Value:  proposeParams.Value.String(),
		Method: method,
		Params: innerParams,
	}
	var proposeReturn multisig.ProposeReturn
	reader = bytes.NewReader(msgRct.Return)
	err = proposeReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = proposeReturn
	return metadata, nil
}

func approve(msg *parser.LotusMessage, msgRct *filTypes.MessageReceipt, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := parseMsigParams(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	reader := bytes.NewReader(msgRct.Return)
	var approveReturn multisig.ApproveReturn
	err = approveReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = approveReturn
	return metadata, nil
}

func cancel(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := parseMsigParams(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func removeSigner(msg *filTypes.Message, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := parseMsigParams(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func changeNumApprovalsThreshold(raw []byte) (map[string]interface{}, error) {
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

func lockBalance(raw []byte) (map[string]interface{}, error) {
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

func parseMsigParams(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (string, error) {
	msgSerial, err := msg.MarshalJSON()
	if err != nil {
		zap.S().Errorf("Could not parse params. Cannot serialize lotus message: %s", err.Error())
		return "", err
	}

	actorCode, err := database.ActorsDB.GetActorCode(msg.To, height, key)
	if err != nil {
		return "", err
	}

	parsedParams, err := Lib.ParseParamsMultisigTx(string(msgSerial), actorCode)
	if err != nil {
		zap.S().Errorf("Could not parse params. ParseParamsMultisigTx returned with error: %s", err.Error())
		return "", err
	}

	return parsedParams, nil
}

func universalReceiverHook(raw []byte) (map[string]interface{}, error) {
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

func innerProposeParams(propose multisig.ProposeParams) (string, cbor.Unmarshaler, error) {
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
