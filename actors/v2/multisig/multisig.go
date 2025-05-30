package multisig

import (
	"bytes"
	"fmt"

	"github.com/filecoin-project/go-state-types/manifest"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	multisig10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	multisig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisig12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisig13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisig14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisig15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisig16 "github.com/filecoin-project/go-state-types/builtin/v16/multisig"
	multisig8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisig9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/multisig"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/multisig"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/multisig"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/multisig"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/multisig"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/multisig"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/multisig"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Msig) MsigConstructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V3)...):
		return parse(raw, &legacyv1.ConstructorParams{}, cborUnmarshaller[*legacyv1.ConstructorParams])
	case tools.AnyIsSupported(network, height, tools.V4, tools.V5, tools.V6, tools.V7, tools.V8, tools.V9):
		return parse(raw, &legacyv2.ConstructorParams{}, cborUnmarshaller[*legacyv2.ConstructorParams])
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parse(raw, &legacyv3.ConstructorParams{}, cborUnmarshaller[*legacyv3.ConstructorParams])
	case tools.V12.IsSupported(network, height):
		return parse(raw, &legacyv4.ConstructorParams{}, cborUnmarshaller[*legacyv4.ConstructorParams])
	case tools.V13.IsSupported(network, height):
		return parse(raw, &legacyv5.ConstructorParams{}, cborUnmarshaller[*legacyv5.ConstructorParams])
	case tools.V14.IsSupported(network, height):
		return parse(raw, &legacyv6.ConstructorParams{}, cborUnmarshaller[*legacyv6.ConstructorParams])
	case tools.V15.IsSupported(network, height):
		return parse(raw, &legacyv7.ConstructorParams{}, cborUnmarshaller[*legacyv7.ConstructorParams])

	case tools.V16.IsSupported(network, height):
		return parse(raw, &multisig8.ConstructorParams{}, cborUnmarshaller[*multisig8.ConstructorParams])
	case tools.V17.IsSupported(network, height):
		return parse(raw, &multisig9.ConstructorParams{}, cborUnmarshaller[*multisig9.ConstructorParams])
	case tools.V18.IsSupported(network, height):
		return parse(raw, &multisig10.ConstructorParams{}, cborUnmarshaller[*multisig10.ConstructorParams])
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parse(raw, &multisig11.ConstructorParams{}, cborUnmarshaller[*multisig11.ConstructorParams])
	case tools.V21.IsSupported(network, height):
		return parse(raw, &multisig12.ConstructorParams{}, cborUnmarshaller[*multisig12.ConstructorParams])
	case tools.V22.IsSupported(network, height):
		return parse(raw, &multisig13.ConstructorParams{}, cborUnmarshaller[*multisig13.ConstructorParams])
	case tools.V23.IsSupported(network, height):
		return parse(raw, &multisig14.ConstructorParams{}, cborUnmarshaller[*multisig14.ConstructorParams])
	case tools.V24.IsSupported(network, height):
		return parse(raw, &multisig15.ConstructorParams{}, cborUnmarshaller[*multisig15.ConstructorParams])
	case tools.V25.IsSupported(network, height):
		return parse(raw, &multisig16.ConstructorParams{}, cborUnmarshaller[*multisig16.ConstructorParams])
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (m *Msig) MsigParams(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, parser ParseFn) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V3)...):
		return parseWithMsigParser[*legacyv1.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*legacyv1.ConstructorParams], false, nil)
	case tools.AnyIsSupported(network, height, tools.V4, tools.V5, tools.V6, tools.V7, tools.V8, tools.V9):
		return parseWithMsigParser[*legacyv2.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*legacyv2.ConstructorParams], false, nil)
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parseWithMsigParser[*legacyv3.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*legacyv3.ConstructorParams], false, nil)
	case tools.V12.IsSupported(network, height):
		return parseWithMsigParser[*legacyv4.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*legacyv4.ConstructorParams], false, nil)
	case tools.V13.IsSupported(network, height):
		return parseWithMsigParser[*legacyv5.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*legacyv5.ConstructorParams], false, nil)
	case tools.V14.IsSupported(network, height):
		return parseWithMsigParser[*legacyv6.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*legacyv6.ConstructorParams], false, nil)
	case tools.V15.IsSupported(network, height):
		return parseWithMsigParser[*legacyv7.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*legacyv7.ConstructorParams], false, nil)
	case tools.V16.IsSupported(network, height):
		return parseWithMsigParser[*multisig8.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig8.ConstructorParams], false, nil)
	case tools.V17.IsSupported(network, height):
		return parseWithMsigParser[*multisig9.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig9.ConstructorParams], false, nil)
	case tools.V18.IsSupported(network, height):
		return parseWithMsigParser[*multisig10.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig10.ConstructorParams], false, nil)
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parseWithMsigParser[*multisig11.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig11.ConstructorParams], false, nil)
	case tools.V21.IsSupported(network, height):
		return parseWithMsigParser[*multisig12.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig12.ConstructorParams], false, nil)
	case tools.V22.IsSupported(network, height):
		return parseWithMsigParser[*multisig13.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig13.ConstructorParams], false, nil)
	case tools.V23.IsSupported(network, height):
		return parseWithMsigParser[*multisig14.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig14.ConstructorParams], false, nil)
	case tools.V24.IsSupported(network, height):
		return parseWithMsigParser[*multisig15.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig15.ConstructorParams], false, nil)
	case tools.V25.IsSupported(network, height):
		return parseWithMsigParser[*multisig16.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig16.ConstructorParams], false, nil)
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) Approve(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser ParseFn) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V3)...):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, jsonUnmarshaller[*legacyv1.ApproveReturn], true, &legacyv1.ApproveReturn{})
	case tools.AnyIsSupported(network, height, tools.V4, tools.V5, tools.V6, tools.V7, tools.V8, tools.V9):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, jsonUnmarshaller[*legacyv2.ApproveReturn], true, &legacyv2.ApproveReturn{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, jsonUnmarshaller[*legacyv3.ApproveReturn], true, &legacyv3.ApproveReturn{})
	case tools.V12.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, jsonUnmarshaller[*legacyv4.ApproveReturn], true, &legacyv4.ApproveReturn{})
	case tools.V13.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, jsonUnmarshaller[*legacyv5.ApproveReturn], true, &legacyv5.ApproveReturn{})
	case tools.V14.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, jsonUnmarshaller[*legacyv6.ApproveReturn], true, &legacyv6.ApproveReturn{})
	case tools.V15.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, jsonUnmarshaller[*legacyv7.ApproveReturn], true, &legacyv7.ApproveReturn{})

	case tools.V16.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig8.ApproveReturn], true, &multisig8.ApproveReturn{})
	case tools.V17.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig9.ApproveReturn], true, &multisig9.ApproveReturn{})
	case tools.V18.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig10.ApproveReturn], true, &multisig10.ApproveReturn{})
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig11.ApproveReturn], true, &multisig11.ApproveReturn{})
	case tools.V21.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig12.ApproveReturn], true, &multisig12.ApproveReturn{})
	case tools.V22.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig13.ApproveReturn], true, &multisig13.ApproveReturn{})
	case tools.V23.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig14.ApproveReturn], true, &multisig14.ApproveReturn{})
	case tools.V24.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig15.ApproveReturn], true, &multisig15.ApproveReturn{})
	case tools.V25.IsSupported(network, height):
		return parseWithMsigParser(msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig16.ApproveReturn], true, &multisig16.ApproveReturn{})
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (m *Msig) Propose(network string, msg *parser.LotusMessage, height int64, proposeKind string, key filTypes.TipSetKey, rawParams, rawReturn []byte, _ ParseFn) (map[string]interface{}, error) {
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

	metadata[parser.ParamsKey] = parser.MultisigPropose{
		To:     to.String(),
		Value:  value,
		Method: method,
		Params: innerParams,
	}

	r, err := proposeReturn(network, height)
	if err != nil {
		return map[string]interface{}{}, err
	}
	err = r.UnmarshalCBOR(bytes.NewReader(rawReturn))
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ReturnKey] = r

	return metadata, nil
}

func (*Msig) Cancel(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser ParseFn) (map[string]interface{}, error) {
	return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
}

func (*Msig) RemoveSigner(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser ParseFn) (map[string]interface{}, error) {
	return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
}

func (*Msig) ChangeNumApprovalsThreshold(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte, parser ParseFn) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V3)...):
		return parseWithMsigParser[*legacyv1.ChangeNumApprovalsThresholdParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv1.ChangeNumApprovalsThresholdParams], false, nil)
	case tools.AnyIsSupported(network, height, tools.V4, tools.V5, tools.V6, tools.V7, tools.V8, tools.V9):
		return parseWithMsigParser[*legacyv2.ChangeNumApprovalsThresholdParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv2.ChangeNumApprovalsThresholdParams], false, nil)
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parseWithMsigParser[*legacyv3.ChangeNumApprovalsThresholdParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv3.ChangeNumApprovalsThresholdParams], false, nil)
	case tools.V12.IsSupported(network, height):
		return parseWithMsigParser[*legacyv4.ChangeNumApprovalsThresholdParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv4.ChangeNumApprovalsThresholdParams], false, nil)
	case tools.V13.IsSupported(network, height):
		return parseWithMsigParser[*legacyv5.ChangeNumApprovalsThresholdParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv5.ChangeNumApprovalsThresholdParams], false, nil)
	case tools.V14.IsSupported(network, height):
		return parseWithMsigParser[*legacyv6.ChangeNumApprovalsThresholdParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv6.ChangeNumApprovalsThresholdParams], false, nil)
	case tools.V15.IsSupported(network, height):
		return parseWithMsigParser[*legacyv7.ChangeNumApprovalsThresholdParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv7.ChangeNumApprovalsThresholdParams], false, nil)
	case tools.V16.IsSupported(network, height):
		return parse(rawParams, &multisig8.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig8.ChangeNumApprovalsThresholdParams])
	case tools.V17.IsSupported(network, height):
		return parse(rawParams, &multisig9.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig9.ChangeNumApprovalsThresholdParams])
	case tools.V18.IsSupported(network, height):
		return parse(rawParams, &multisig10.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig10.ChangeNumApprovalsThresholdParams])
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parse(rawParams, &multisig11.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig11.ChangeNumApprovalsThresholdParams])
	case tools.V21.IsSupported(network, height):
		return parse(rawParams, &multisig12.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig12.ChangeNumApprovalsThresholdParams])
	case tools.V22.IsSupported(network, height):
		return parse(rawParams, &multisig13.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig13.ChangeNumApprovalsThresholdParams])
	case tools.V23.IsSupported(network, height):
		return parse(rawParams, &multisig14.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig14.ChangeNumApprovalsThresholdParams])
	case tools.V24.IsSupported(network, height):
		return parse(rawParams, &multisig15.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig15.ChangeNumApprovalsThresholdParams])
	case tools.V25.IsSupported(network, height):
		return parse(rawParams, &multisig16.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig16.ChangeNumApprovalsThresholdParams])
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) LockBalance(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawParams []byte, parser ParseFn) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V3)...):
		return parseWithMsigParser[*legacyv1.LockBalanceParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv1.LockBalanceParams], false, nil)
	case tools.AnyIsSupported(network, height, tools.V4, tools.V5, tools.V6, tools.V7, tools.V8, tools.V9):
		return parseWithMsigParser[*legacyv2.LockBalanceParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv2.LockBalanceParams], false, nil)
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parseWithMsigParser[*legacyv3.LockBalanceParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv3.LockBalanceParams], false, nil)
	case tools.V12.IsSupported(network, height):
		return parseWithMsigParser[*legacyv4.LockBalanceParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv4.LockBalanceParams], false, nil)
	case tools.V13.IsSupported(network, height):
		return parseWithMsigParser[*legacyv5.LockBalanceParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv5.LockBalanceParams], false, nil)
	case tools.V14.IsSupported(network, height):
		return parseWithMsigParser[*legacyv6.LockBalanceParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv6.LockBalanceParams], false, nil)
	case tools.V15.IsSupported(network, height):
		return parseWithMsigParser[*legacyv7.LockBalanceParams](msg, height, key, parser, rawParams, cborUnmarshaller[*legacyv7.LockBalanceParams], false, nil)
	case tools.V16.IsSupported(network, height):
		return parse(rawParams, &multisig8.LockBalanceParams{}, cborUnmarshaller[*multisig8.LockBalanceParams])
	case tools.V17.IsSupported(network, height):
		return parse(rawParams, &multisig9.LockBalanceParams{}, cborUnmarshaller[*multisig9.LockBalanceParams])
	case tools.V18.IsSupported(network, height):
		return parse(rawParams, &multisig10.LockBalanceParams{}, cborUnmarshaller[*multisig10.LockBalanceParams])
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parse(rawParams, &multisig11.LockBalanceParams{}, cborUnmarshaller[*multisig11.LockBalanceParams])
	case tools.V21.IsSupported(network, height):
		return parse(rawParams, &multisig12.LockBalanceParams{}, cborUnmarshaller[*multisig12.LockBalanceParams])
	case tools.V22.IsSupported(network, height):
		return parse(rawParams, &multisig13.LockBalanceParams{}, cborUnmarshaller[*multisig13.LockBalanceParams])
	case tools.V23.IsSupported(network, height):
		return parse(rawParams, &multisig14.LockBalanceParams{}, cborUnmarshaller[*multisig14.LockBalanceParams])
	case tools.V24.IsSupported(network, height):
		return parse(rawParams, &multisig15.LockBalanceParams{}, cborUnmarshaller[*multisig15.LockBalanceParams])
	case tools.V25.IsSupported(network, height):
		return parse(rawParams, &multisig16.LockBalanceParams{}, cborUnmarshaller[*multisig16.LockBalanceParams])
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) UniversalReceiverHook(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser ParseFn) (map[string]interface{}, error) {
	return parse(rawReturn, &abi.CborBytesTransparent{}, cborUnmarshaller[*abi.CborBytesTransparent])
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
