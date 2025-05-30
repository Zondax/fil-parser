package multisig

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/manifest"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

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

	return parseCBOR(raw, nil, params(), nil)
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

	method, innerParams, err := m.innerProposeParams(msg, to, network, height, methodNum, innerParamsRaw, key)
	if err != nil {
		_ = m.metrics.UpdateMultisigProposeMetric(manifest.MultisigKey, proposeKind, fmt.Sprint(methodNum))
		m.logger.Errorf("could not decode multisig inner params. Method: %v. Err: %v", methodNum.String(), err)
	}

	params := innerParams
	// get ParamsKey for innerParams if possible
	if innerParams != nil && innerParams[parser.ParamsKey] != nil {
		if inner, ok := innerParams[parser.ParamsKey].(map[string]any); ok {
			params = inner
		}
	}

	metadata[parser.ParamsKey] = parser.MultisigPropose{
		To:     to.String(),
		Value:  value,
		Method: method,
		Params: params,
	}

	version := tools.VersionFromHeight(network, height)
	r, ok := proposeReturn[version.String()]
	if !ok {
		return map[string]interface{}{}, fmt.Errorf("proposeReturn: %s not found", version.String())
	}
	val := r()
	err = val.UnmarshalCBOR(bytes.NewReader(rawReturn))
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ReturnKey] = val

	return metadata, nil
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

func (m *Msig) parseMsigParams(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (string, error) {
	msgSerial, err := msg.MarshalJSON() // TODO: this may not work properly
	if err != nil {
		// m.helper.GetLogger().Errorf("Could not parse params. Cannot serialize lotus message: %v", err)
		return "", err
	}

	actorCode, err := m.helper.GetActorsCache().GetActorCode(msg.To, key, false)
	if err != nil {
		return "", err
	}

	c, err := cid.Parse(actorCode)
	if err != nil {
		// m.helper.GetLogger().Errorf("Could not parse params. Cannot cid.parse actor code: %v", err)
		return "", err
	}
	parsedParams, err := m.helper.GetFilecoinLib().ParseParamsMultisigTx(string(msgSerial), c)
	if err != nil {
		// m.helper.GetLogger().Errorf("Could not parse params. ParseParamsMultisigTx returned with error: %v", err)
		return "", err
	}

	return parsedParams, nil
}
