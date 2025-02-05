package multisig

import (
	"bytes"
	"fmt"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	multisig10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	multisig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisig12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisig13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisig14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisig15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisig8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisig9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Msig) MsigConstructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
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
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (m *Msig) MsigParams(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, parser ParseFn) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
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
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) Approve(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser ParseFn) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
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
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) Propose(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, _ ParseFn) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	rawParams, methodNum, to, value, _, err := getProposeParams(network, height)
	if err != nil {
		return nil, err
	}
	method, innerParams, err := innerProposeParams(network, height, methodNum, rawParams)
	if err != nil {
		return map[string]interface{}{}, err
	}

	metadata[parser.ParamsKey] = parser.Propose{
		To:     to,
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
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V17.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V18.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V21.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V22.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V23.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V24.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) RemoveSigner(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser ParseFn) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V17.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V18.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V21.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V22.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V23.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	case tools.V24.IsSupported(network, height):
		return parseWithMsigParser[metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false, nil)
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) ChangeNumApprovalsThreshold(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser ParseFn) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		return parse(rawReturn, &multisig8.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig8.ChangeNumApprovalsThresholdParams])
	case tools.V17.IsSupported(network, height):
		return parse(rawReturn, &multisig9.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig9.ChangeNumApprovalsThresholdParams])
	case tools.V18.IsSupported(network, height):
		return parse(rawReturn, &multisig10.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig10.ChangeNumApprovalsThresholdParams])
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parse(rawReturn, &multisig11.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig11.ChangeNumApprovalsThresholdParams])
	case tools.V21.IsSupported(network, height):
		return parse(rawReturn, &multisig12.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig12.ChangeNumApprovalsThresholdParams])
	case tools.V22.IsSupported(network, height):
		return parse(rawReturn, &multisig13.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig13.ChangeNumApprovalsThresholdParams])
	case tools.V23.IsSupported(network, height):
		return parse(rawReturn, &multisig14.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig14.ChangeNumApprovalsThresholdParams])
	case tools.V24.IsSupported(network, height):
		return parse(rawReturn, &multisig15.ChangeNumApprovalsThresholdParams{}, cborUnmarshaller[*multisig15.ChangeNumApprovalsThresholdParams])
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) LockBalance(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser ParseFn) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		return parse(rawReturn, &multisig8.LockBalanceParams{}, cborUnmarshaller[*multisig8.LockBalanceParams])
	case tools.V17.IsSupported(network, height):
		return parse(rawReturn, &multisig9.LockBalanceParams{}, cborUnmarshaller[*multisig9.LockBalanceParams])
	case tools.V18.IsSupported(network, height):
		return parse(rawReturn, &multisig10.LockBalanceParams{}, cborUnmarshaller[*multisig10.LockBalanceParams])
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parse(rawReturn, &multisig11.LockBalanceParams{}, cborUnmarshaller[*multisig11.LockBalanceParams])
	case tools.V21.IsSupported(network, height):
		return parse(rawReturn, &multisig12.LockBalanceParams{}, cborUnmarshaller[*multisig12.LockBalanceParams])
	case tools.V22.IsSupported(network, height):
		return parse(rawReturn, &multisig13.LockBalanceParams{}, cborUnmarshaller[*multisig13.LockBalanceParams])
	case tools.V23.IsSupported(network, height):
		return parse(rawReturn, &multisig14.LockBalanceParams{}, cborUnmarshaller[*multisig14.LockBalanceParams])
	case tools.V24.IsSupported(network, height):
		return parse(rawReturn, &multisig15.LockBalanceParams{}, cborUnmarshaller[*multisig15.LockBalanceParams])
	}
	return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) UniversalReceiverHook(network string, msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser ParseFn) (map[string]interface{}, error) {
	return parse(rawReturn, &abi.CborBytesTransparent{}, cborUnmarshaller[*abi.CborBytesTransparent])
}

func (m *Msig) parseMsigParams(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (string, error) {
	msgSerial, err := msg.MarshalJSON() // TODO: this may not work properly
	if err != nil {
		// m.helper.GetLogger().Sugar().Errorf("Could not parse params. Cannot serialize lotus message: %v", err)
		return "", err
	}

	actorCode, err := m.helper.GetActorsCache().GetActorCode(msg.To, key, false)
	if err != nil {
		return "", err
	}

	c, err := cid.Parse(actorCode)
	if err != nil {
		// m.helper.GetLogger().Sugar().Errorf("Could not parse params. Cannot cid.parse actor code: %v", err)
		return "", err
	}
	parsedParams, err := m.helper.GetFilecoinLib().ParseParamsMultisigTx(string(msgSerial), c)
	if err != nil {
		// m.helper.GetLogger().Sugar().Errorf("Could not parse params. ParseParamsMultisigTx returned with error: %v", err)
		return "", err
	}

	return parsedParams, nil
}
