package multisig

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/manifest"

	filTypes "github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Msig) MsigConstructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := constructorParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("constructorParams: %s not found", version.String())
	}

	metadata, err := parseCBOR(raw, nil, params(), nil)
	if err != nil {
		versions := tools.GetSupportedVersions(network)
		for _, version := range versions {
			params, ok := constructorParams[version.String()]
			if !ok {
				continue
			}
			metadata, err = parseCBOR(raw, nil, params(), nil)
			if err == nil {
				break
			}
		}
	}
	return metadata, err
}

func (*Msig) Approve(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := txnIDParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("txnIDParams: %s not found", version.String())
	}

	ret, ok := approveReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("approveReturn: %s not found", version.String())
	}
	return parseCBOR(rawParams, rawReturn, params(), ret())
}

func (*Msig) Cancel(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := txnIDParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("txnIDParams: %s not found", version.String())
	}

	return parseCBOR(rawParams, nil, params(), nil)
}

func (m *Msig) Propose(network string, msg *parser.LotusMessage, height int64, proposeKind string, key filTypes.TipSetKey, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	innerParamsRaw, methodNum, to, value, _, err := getProposeParams(network, height, rawParams)
	if err != nil {
		return nil, err
	}
	applied, innerReturnRaw, innerReturnParsed, err := getProposeReturn(network, height, rawReturn)
	if err != nil {
		return nil, err
	}

	method, innerMsg, err := m.parseInnerProposeMsg(msg, to, network, height, methodNum, innerParamsRaw, innerReturnRaw, key)
	if err != nil {
		_ = m.metrics.UpdateMultisigProposeMetric(manifest.MultisigKey, proposeKind, fmt.Sprint(methodNum))
		m.logger.Errorf("could not decode multisig inner params. Method: %v. Err: %v", methodNum.String(), err)
	}

	proposalData := parser.MultisigPropose{
		To:     to.String(),
		Value:  value,
		Method: method,
		Params: innerMsg,
	}

	// this is the params of the multisig proposal execution (always present)
	if innerMsg != nil && innerMsg[parser.ParamsKey] != nil {
		parsedCBORParams, err := m.paramsToMap(innerMsg[parser.ParamsKey])
		if err != nil {
			return nil, err
		}
		proposalData.Params = parsedCBORParams
	}

	// this is the return data of the multisig proposal execution, only present if applied=true
	if applied && innerMsg != nil && innerMsg[parser.ReturnKey] != nil {
		parsedCBORReturn, err := m.paramsToMap(innerMsg[parser.ReturnKey])
		if err != nil {
			return nil, err
		}
		proposalData.Return = parsedCBORReturn
	}

	metadata[parser.ParamsKey] = proposalData

	// this is the return status of the multisig proposal that indicates the TxnID and if the proposal was applied
	metadata[parser.ReturnKey] = innerReturnParsed

	return metadata, nil
}

// paramsToMap converts the parameters to a map from a generic CBORUnmarshaler interface type.
func (*Msig) paramsToMap(params any) (map[string]any, error) {
	dataAsMap := make(map[string]any)

	tmp, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(tmp, &dataAsMap)
	if err != nil {
		var dataAsAny any
		err = json.Unmarshal(tmp, &dataAsAny)
		if err != nil {
			return nil, err
		}
		dataAsMap[parser.ValueKey] = dataAsAny
	}

	return dataAsMap, nil
}

func (*Msig) RemoveSigner(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := removeSignerParams2[version.String()]
	if !ok {
		return nil, fmt.Errorf("removeSignerParams: %s not found", version.String())
	}

	return parseCBOR(rawParams, nil, params(), nil)
}

func (*Msig) ChangeNumApprovalsThreshold(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := changeNumApprovalsThresholdParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("changeNumApprovalsThresholdParams: %s not found", version.String())
	}

	return parseCBOR(rawParams, nil, params(), nil)
}

func (*Msig) LockBalance(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := lockBalanceParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("lockBalanceParams: %s not found", version.String())
	}

	return parseCBOR(rawParams, nil, params(), nil)
}

func (m *Msig) WithdrawBalance(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte, rawReturn []byte) (map[string]interface{}, error) {
	return m.miner.WithdrawBalanceExported(network, height, rawParams, rawReturn)
}

func (m *Msig) InvokeContract(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = hex.EncodeToString(rawParams)
	metadata[parser.ReturnKey] = hex.EncodeToString(rawReturn)
	return metadata, nil
}

func (m *Msig) AddSigner(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := addSignerParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("lockBalanceParams: %s not found", version.String())
	}

	return parseCBOR(rawParams, nil, params(), nil)
}
func (m *Msig) SwapSigner(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := swapSignerParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("lockBalanceParams: %s not found", version.String())
	}

	return parseCBOR(rawParams, nil, params(), nil)
}

func (m *Msig) AddVerifier(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte, rawReturn []byte) (map[string]interface{}, error) {
	return m.verifreg.AddVerifier(network, height, rawParams)
}

func (m *Msig) ChangeOwnerAddress(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	return metadata, nil
}

func (*Msig) UniversalReceiverHook(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte) (map[string]interface{}, error) {
	return parseCBOR(rawParams, nil, &abi.CborBytesTransparent{}, nil)
}
