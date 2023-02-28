package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v10/miner"
	"github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	"github.com/filecoin-project/go-state-types/cbor"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/database"
	"github.com/zondax/rosetta-filecoin-lib/actors"
	"go.uber.org/zap"
)

func (p *Parser) parseMultisig(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	switch txType {
	case MethodConstructor:
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodPropose:
		return p.propose(msg, msgRct)
	case MethodApprove:
		return p.approve(msg, msgRct, height, key)
	case MethodCancel:
		return p.cancel(msg, height, key)
	case MethodAddSigner, MethodSwapSigner:
		return p.msigParams(msg, height, key)
	case MethodRemoveSigner:
		return p.removeSigner(msg, height, key)
	case MethodChangeNumApprovalsThreshold:
		return p.changeNumApprovalsThreshold(msg.Params)
	case MethodAddVerifies:
	case MethodLockBalance:
		return p.lockBalance(msg.Params)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) msigParams(msg *filTypes.Message, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	params, err := p.parseMsigParams(msg, height, key)
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

func (p *Parser) propose(msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var proposeParams multisig.ProposeParams
	reader := bytes.NewReader(msg.Params)
	err := proposeParams.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	method, innerParams, err := p.innerProposeParams(proposeParams)
	if err != nil {
		zap.S().Errorf("could not decode multisig inner params. Method: %v. Err: %v", proposeParams.Method.String(), err)
	}
	metadata[ParamsKey] = propose{
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
	metadata[ReturnKey] = proposeReturn
	return metadata, nil
}

func (p *Parser) approve(msg *filTypes.Message, msgRct *filTypes.MessageReceipt, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := p.parseMsigParams(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[ParamsKey] = params
	reader := bytes.NewReader(msgRct.Return)
	var approveReturn multisig.ApproveReturn
	err = approveReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = approveReturn
	return metadata, nil
}

func (p *Parser) cancel(msg *filTypes.Message, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := p.parseMsigParams(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) removeSigner(msg *filTypes.Message, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := p.parseMsigParams(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) changeNumApprovalsThreshold(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params multisig.ChangeNumApprovalsThresholdParams
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) lockBalance(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params multisig.LockBalanceParams
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) parseMsigParams(msg *filTypes.Message, height int64, key filTypes.TipSetKey) (string, error) {
	msgSerial, err := msg.MarshalJSON()
	if err != nil {
		zap.S().Errorf("Could not parse params. Cannot serialize lotus message: %s", err.Error())
		return "", err
	}

	actorCode, err := database.ActorsDB.GetActorCode(msg.To, height, key)
	if err != nil {
		return "", err
	}

	if !p.lib.BuiltinActors.IsActor(actorCode, actors.ActorMultisigName) {
		return "", fmt.Errorf("this id doesn't correspond to a multisig actor")
	}

	parsedParams, err := p.lib.ParseParamsMultisigTx(string(msgSerial), actorCode)
	if err != nil {
		zap.S().Errorf("Could not parse params. ParseParamsMultisigTx returned with error: %s", err.Error())
		return "", err
	}

	return parsedParams, nil
}

func (p *Parser) innerProposeParams(propose multisig.ProposeParams) (string, cbor.Unmarshaler, error) {
	reader := bytes.NewReader(propose.Params)
	switch propose.Method {
	case builtin.MethodSend:
		var params multisig.ProposeParams
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return "", nil, err
		}
		return MethodSend, &params, nil
	case builtin.MethodsMultisig.Approve,
		builtin.MethodsMultisig.Cancel:
		var params multisig.TxnIDParams
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return "", nil, err
		}
		return MethodApprove, &params, nil
	case builtin.MethodsMultisig.AddSigner:
		var params multisig.AddSignerParams
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return "", nil, err
		}
		return MethodAddSigner, &params, nil
	case builtin.MethodsMultisig.RemoveSigner:
		var params multisig.RemoveSignerParams
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return "", nil, err
		}
		return MethodRemoveSigner, &params, nil
	case builtin.MethodsMultisig.SwapSigner:
		var params multisig.SwapSignerParams
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return "", nil, err
		}
		return MethodSwapSigner, &params, nil
	case builtin.MethodsMultisig.ChangeNumApprovalsThreshold:
		var params multisig.ChangeNumApprovalsThresholdParams
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return "", nil, err
		}
		return MethodChangeNumApprovalsThreshold, &params, nil
	case builtin.MethodsMultisig.LockBalance:
		var params multisig.LockBalanceParams
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return "", nil, err
		}
		return MethodLockBalance, &params, nil
	case builtin.MethodsMiner.WithdrawBalance:
		var params miner.WithdrawBalanceParams
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return "", nil, err
		}
		return MethodWithdrawBalance, &params, nil
	case builtin.MethodsVerifiedRegistry.AddVerifier:
		var params verifreg.AddVerifierParams
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return "", nil, err
		}
		return MethodAddVerifier, &params, nil
	}
	return "", nil, errUnknownMethod
}
