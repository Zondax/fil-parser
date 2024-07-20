package multisig

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v8/miner"
	"github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	"github.com/filecoin-project/go-state-types/cbor"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/parser"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

type MultisigMethods interface {
	ParseParams(raw []byte, paramsType cbor.Unmarshaler) (map[string]interface{}, error)
	ParseParamsWithReturn(rawParams, rawReturn []byte, paramsType, returnType cbor.Unmarshaler) (map[string]interface{}, error)
	ParseMsigConstructor(raw []byte, paramsType cbor.Unmarshaler) (map[string]interface{}, error)
	ParsePropose(rawParams, rawReturn []byte, paramsType, returnType cbor.Unmarshaler) (map[string]interface{}, error)
	ParseApprove(msg *parser.LotusMessage, rawReturn []byte, height int64, key filTypes.TipSetKey, paramsType, returnType cbor.Unmarshaler) (map[string]interface{}, error)
	ParseCancel(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, paramsType cbor.Unmarshaler) (map[string]interface{}, error)
	ParseRemoveSigner(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, paramsType cbor.Unmarshaler) (map[string]interface{}, error)
	ParseChangeNumApprovalsThreshold(raw []byte, paramsType cbor.Unmarshaler) (map[string]interface{}, error)
	ParseLockBalance(raw []byte, paramsType cbor.Unmarshaler) (map[string]interface{}, error)
	ParseUniversalReceiverHook(raw []byte, paramsType cbor.Unmarshaler) (map[string]interface{}, error)
	ParseInnerProposeParams(propose multisig.ProposeParams) (string, cbor.Unmarshaler, error)
}

type multisigMethods struct {
	actorsCache *cache.ActorsCache
	rosettaLib  *rosettaFilecoinLib.RosettaConstructionFilecoin
}

func NewMultisigMethods(actorsCache *cache.ActorsCache, rosettaLib *rosettaFilecoinLib.RosettaConstructionFilecoin) MultisigMethods {
	return &multisigMethods{actorsCache: actorsCache, rosettaLib: rosettaLib}
}

func (s *multisigMethods) ParseParams(raw []byte, paramsType cbor.Unmarshaler) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := paramsType.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = paramsType
	return metadata, nil
}

func (s *multisigMethods) ParseParamsWithReturn(rawParams, rawReturn []byte, paramsType, returnType cbor.Unmarshaler) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	err := paramsType.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = paramsType

	reader = bytes.NewReader(rawReturn)
	err = returnType.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = returnType

	return metadata, nil
}

func (s *multisigMethods) ParseMsigConstructor(raw []byte, paramsType cbor.Unmarshaler) (map[string]interface{}, error) {
	return s.ParseParams(raw, paramsType)
}

func (s *multisigMethods) ParsePropose(rawParams, rawReturn []byte, paramsType, returnType cbor.Unmarshaler) (map[string]interface{}, error) {
	return s.ParseParamsWithReturn(rawParams, rawReturn, paramsType, returnType)
}

func (s *multisigMethods) ParseApprove(msg *parser.LotusMessage, rawReturn []byte, height int64, key filTypes.TipSetKey, paramsType, returnType cbor.Unmarshaler) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := s.msigParams(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	reader := bytes.NewReader(rawReturn)
	err = returnType.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = returnType
	return metadata, nil
}

func (s *multisigMethods) ParseCancel(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, paramsType cbor.Unmarshaler) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := s.msigParams(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func (s *multisigMethods) ParseRemoveSigner(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, paramsType cbor.Unmarshaler) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	params, err := s.msigParams(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func (s *multisigMethods) ParseChangeNumApprovalsThreshold(raw []byte, paramsType cbor.Unmarshaler) (map[string]interface{}, error) {
	return s.ParseParams(raw, paramsType)
}

func (s *multisigMethods) ParseLockBalance(raw []byte, paramsType cbor.Unmarshaler) (map[string]interface{}, error) {
	return s.ParseParams(raw, paramsType)
}

func (s *multisigMethods) ParseUniversalReceiverHook(raw []byte, paramsType cbor.Unmarshaler) (map[string]interface{}, error) {
	return s.ParseParams(raw, paramsType)
}

func (s *multisigMethods) msigParams(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (string, error) {
	msgSerial, err := msg.MarshalJSON() // TODO: this may not work properly
	if err != nil {
		return "", err
	}

	actorCode, err := s.actorsCache.GetActorCode(msg.To, key, false)
	if err != nil {
		return "", err
	}

	c, err := cid.Parse(actorCode)
	if err != nil {
		return "", err
	}
	parsedParams, err := s.rosettaLib.ParseParamsMultisigTx(string(msgSerial), c)
	if err != nil {
		return "", err
	}

	return parsedParams, nil
}

func (s *multisigMethods) ParseInnerProposeParams(propose multisig.ProposeParams) (string, cbor.Unmarshaler, error) {
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
