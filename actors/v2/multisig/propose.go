package multisig

import (
	"bytes"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/cbor"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"

	multisig10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	multisig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisig12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisig13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisig14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisig15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisig8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisig9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/multisig"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/multisig"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/multisig"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/multisig"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/multisig"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/multisig"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/multisig"

	verifregv10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	verifregv11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifregv12 "github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	verifregv13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifregv14 "github.com/filecoin-project/go-state-types/builtin/v14/verifreg"
	verifregv15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"
	verifregv8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	verifregv9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"

	legacyverifreg1 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	legacyverifreg2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/verifreg"
	legacyverifreg3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/verifreg"
	legacyverifreg4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/verifreg"
	legacyverifreg5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/verifreg"
	legacyverifreg6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/verifreg"
	legacyverifreg7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/verifreg"

	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"

	legacyminer1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	legacyminer2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyminer3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyminer4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyminer5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyminer6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyminer7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"
)

func innerProposeParams(network string, height int64, method abi.MethodNum, proposeParams []byte) (string, cbor.Unmarshaler, error) {
	var params multisigParams
	var err error
	var methodName string
	reader := bytes.NewReader(proposeParams)
	switch method {
	case builtin.MethodSend:
		if proposeParams == nil {
			return parser.MethodSend, nil, nil
		}
		_, _, _, _, params, err = getProposeParams(network, height, proposeParams)
		return parser.MethodSend, params, err
	case builtin.MethodsMultisig.Approve:
		methodName = parser.MethodApprove
		params, err = txnIDParams(network, height)
	case builtin.MethodsMultisig.Cancel:
		methodName = parser.MethodCancel
		params, err = txnIDParams(network, height)
	case builtin.MethodsMultisig.AddSigner:
		methodName = parser.MethodAddSigner
		params, err = addSignerParams(network, height)
	case builtin.MethodsMultisig.RemoveSigner:
		methodName = parser.MethodRemoveSigner
		params, err = removeSignerParams(network, height)
	case builtin.MethodsMultisig.SwapSigner:
		methodName = parser.MethodSwapSigner
		params, err = swapSignerParams(network, height)
	case builtin.MethodsMultisig.ChangeNumApprovalsThreshold:
		methodName = parser.MethodChangeNumApprovalsThreshold
		params, err = changeNumApprovalsThresholdParams(network, height)
	case builtin.MethodsMultisig.LockBalance:
		methodName = parser.MethodLockBalance
		params, err = lockBalanceParams(network, height)
	case builtin.MethodsMiner.WithdrawBalance:
		methodName = parser.MethodWithdrawBalance
		params, err = withdrawBalanceParams(network, height)
	case builtin.MethodsVerifiedRegistry.AddVerifier:
		methodName = parser.MethodAddVerifier
		params, err = verifierParams(network, height)
	default:
		err = parser.ErrUnknownMethod
	}
	if err == nil {
		err := params.UnmarshalCBOR(reader)
		return methodName, params, err
	}
	return "", nil, err
}

func getProposeParams(network string, height int64, rawParams []byte) (raw []byte, methodNum abi.MethodNum, to, value string, params multisigParams, err error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		tmp := &legacyv1.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil

	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		tmp := &legacyv2.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		tmp := &legacyv3.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.V12.IsSupported(network, height):
		tmp := &legacyv4.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.V13.IsSupported(network, height):
		tmp := &legacyv5.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.V14.IsSupported(network, height):
		tmp := &legacyv6.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.V15.IsSupported(network, height):
		tmp := &legacyv7.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil

	case tools.V16.IsSupported(network, height):
		tmp := &multisig8.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.V17.IsSupported(network, height):
		tmp := &multisig9.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.V18.IsSupported(network, height):
		tmp := &multisig10.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		tmp := &multisig11.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.V21.IsSupported(network, height):
		tmp := &multisig12.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.V22.IsSupported(network, height):
		tmp := &multisig13.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.V23.IsSupported(network, height):
		tmp := &multisig14.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	case tools.V24.IsSupported(network, height):
		tmp := &multisig15.ProposeParams{}
		err = tmp.UnmarshalCBOR(bytes.NewReader(rawParams))
		if err != nil {
			break
		}
		return tmp.Params, tmp.Method, tmp.To.String(), tmp.Value.String(), tmp, nil
	default:
		return nil, 0, "", "", nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return nil, 0, "", "", nil, err
}

func proposeReturn(network string, height int64) (multisigParams, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return &legacyv1.ProposeReturn{}, nil
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return &legacyv2.ProposeReturn{}, nil
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return &legacyv3.ProposeReturn{}, nil
	case tools.V12.IsSupported(network, height):
		return &legacyv4.ProposeReturn{}, nil
	case tools.V13.IsSupported(network, height):
		return &legacyv5.ProposeReturn{}, nil
	case tools.V14.IsSupported(network, height):
		return &legacyv6.ProposeReturn{}, nil
	case tools.V15.IsSupported(network, height):
		return &legacyv7.ProposeReturn{}, nil
	case tools.V16.IsSupported(network, height):
		return &multisig8.ProposeReturn{}, nil
	case tools.V17.IsSupported(network, height):
		return &multisig9.ProposeReturn{}, nil
	case tools.V18.IsSupported(network, height):
		return &multisig10.ProposeReturn{}, nil
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return &multisig11.ProposeReturn{}, nil
	case tools.V21.IsSupported(network, height):
		return &multisig12.ProposeReturn{}, nil
	case tools.V22.IsSupported(network, height):
		return &multisig13.ProposeReturn{}, nil
	case tools.V23.IsSupported(network, height):
		return &multisig14.ProposeReturn{}, nil
	case tools.V24.IsSupported(network, height):
		return &multisig15.ProposeReturn{}, nil
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func txnIDParams(network string, height int64) (multisigParams, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return &legacyv1.TxnIDParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return &legacyv2.TxnIDParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return &legacyv3.TxnIDParams{}, nil
	case tools.V12.IsSupported(network, height):
		return &legacyv4.TxnIDParams{}, nil
	case tools.V13.IsSupported(network, height):
		return &legacyv5.TxnIDParams{}, nil
	case tools.V14.IsSupported(network, height):
		return &legacyv6.TxnIDParams{}, nil
	case tools.V15.IsSupported(network, height):
		return &legacyv7.TxnIDParams{}, nil

	case tools.V16.IsSupported(network, height):
		return &multisig8.TxnIDParams{}, nil
	case tools.V17.IsSupported(network, height):
		return &multisig9.TxnIDParams{}, nil
	case tools.V18.IsSupported(network, height):
		return &multisig10.TxnIDParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return &multisig11.TxnIDParams{}, nil
	case tools.V21.IsSupported(network, height):
		return &multisig12.TxnIDParams{}, nil
	case tools.V22.IsSupported(network, height):
		return &multisig13.TxnIDParams{}, nil
	case tools.V23.IsSupported(network, height):
		return &multisig14.TxnIDParams{}, nil
	case tools.V24.IsSupported(network, height):
		return &multisig15.TxnIDParams{}, nil
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func addSignerParams(network string, height int64) (multisigParams, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return &legacyv1.AddSignerParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return &legacyv2.AddSignerParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return &legacyv3.AddSignerParams{}, nil
	case tools.V12.IsSupported(network, height):
		return &legacyv4.AddSignerParams{}, nil
	case tools.V13.IsSupported(network, height):
		return &legacyv5.AddSignerParams{}, nil
	case tools.V14.IsSupported(network, height):
		return &legacyv6.AddSignerParams{}, nil
	case tools.V15.IsSupported(network, height):
		return &legacyv7.AddSignerParams{}, nil
	case tools.V16.IsSupported(network, height):
		return &multisig8.AddSignerParams{}, nil
	case tools.V17.IsSupported(network, height):
		return &multisig9.AddSignerParams{}, nil
	case tools.V18.IsSupported(network, height):
		return &multisig10.AddSignerParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return &multisig11.AddSignerParams{}, nil
	case tools.V21.IsSupported(network, height):
		return &multisig12.AddSignerParams{}, nil
	case tools.V22.IsSupported(network, height):
		return &multisig13.AddSignerParams{}, nil
	case tools.V23.IsSupported(network, height):
		return &multisig14.AddSignerParams{}, nil
	case tools.V24.IsSupported(network, height):
		return &multisig15.AddSignerParams{}, nil
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func removeSignerParams(network string, height int64) (multisigParams, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return &legacyv1.RemoveSignerParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return &legacyv2.RemoveSignerParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return &legacyv3.RemoveSignerParams{}, nil
	case tools.V12.IsSupported(network, height):
		return &legacyv4.RemoveSignerParams{}, nil
	case tools.V13.IsSupported(network, height):
		return &legacyv5.RemoveSignerParams{}, nil
	case tools.V14.IsSupported(network, height):
		return &legacyv6.RemoveSignerParams{}, nil
	case tools.V15.IsSupported(network, height):
		return &legacyv7.RemoveSignerParams{}, nil
	case tools.V16.IsSupported(network, height):
		return &multisig8.RemoveSignerParams{}, nil
	case tools.V17.IsSupported(network, height):
		return &multisig9.RemoveSignerParams{}, nil
	case tools.V18.IsSupported(network, height):
		return &multisig10.RemoveSignerParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return &multisig11.RemoveSignerParams{}, nil
	case tools.V21.IsSupported(network, height):
		return &multisig12.RemoveSignerParams{}, nil
	case tools.V22.IsSupported(network, height):
		return &multisig13.RemoveSignerParams{}, nil
	case tools.V23.IsSupported(network, height):
		return &multisig14.RemoveSignerParams{}, nil
	case tools.V24.IsSupported(network, height):
		return &multisig15.RemoveSignerParams{}, nil
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func swapSignerParams(network string, height int64) (multisigParams, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return &legacyv1.SwapSignerParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return &legacyv2.SwapSignerParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return &legacyv3.SwapSignerParams{}, nil
	case tools.V12.IsSupported(network, height):
		return &legacyv4.SwapSignerParams{}, nil
	case tools.V13.IsSupported(network, height):
		return &legacyv5.SwapSignerParams{}, nil
	case tools.V14.IsSupported(network, height):
		return &legacyv6.SwapSignerParams{}, nil
	case tools.V15.IsSupported(network, height):
		return &legacyv7.SwapSignerParams{}, nil
	case tools.V16.IsSupported(network, height):
		return &multisig8.SwapSignerParams{}, nil
	case tools.V17.IsSupported(network, height):
		return &multisig9.SwapSignerParams{}, nil
	case tools.V18.IsSupported(network, height):
		return &multisig10.SwapSignerParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return &multisig11.SwapSignerParams{}, nil
	case tools.V21.IsSupported(network, height):
		return &multisig12.SwapSignerParams{}, nil
	case tools.V22.IsSupported(network, height):
		return &multisig13.SwapSignerParams{}, nil
	case tools.V23.IsSupported(network, height):
		return &multisig14.SwapSignerParams{}, nil
	case tools.V24.IsSupported(network, height):
		return &multisig15.SwapSignerParams{}, nil
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func changeNumApprovalsThresholdParams(network string, height int64) (multisigParams, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return &legacyv1.ChangeNumApprovalsThresholdParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return &legacyv2.ChangeNumApprovalsThresholdParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return &legacyv3.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V12.IsSupported(network, height):
		return &legacyv4.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V13.IsSupported(network, height):
		return &legacyv5.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V14.IsSupported(network, height):
		return &legacyv6.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V15.IsSupported(network, height):
		return &legacyv7.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V16.IsSupported(network, height):
		return &multisig8.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V17.IsSupported(network, height):
		return &multisig9.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V18.IsSupported(network, height):
		return &multisig10.ChangeNumApprovalsThresholdParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return &multisig11.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V21.IsSupported(network, height):
		return &multisig12.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V22.IsSupported(network, height):
		return &multisig13.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V23.IsSupported(network, height):
		return &multisig14.ChangeNumApprovalsThresholdParams{}, nil
	case tools.V24.IsSupported(network, height):
		return &multisig15.ChangeNumApprovalsThresholdParams{}, nil
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
func lockBalanceParams(network string, height int64) (multisigParams, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return &legacyv1.LockBalanceParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return &legacyv2.LockBalanceParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return &legacyv3.LockBalanceParams{}, nil
	case tools.V12.IsSupported(network, height):
		return &legacyv4.LockBalanceParams{}, nil
	case tools.V13.IsSupported(network, height):
		return &legacyv5.LockBalanceParams{}, nil
	case tools.V14.IsSupported(network, height):
		return &legacyv6.LockBalanceParams{}, nil
	case tools.V15.IsSupported(network, height):
		return &legacyv7.LockBalanceParams{}, nil
	case tools.V16.IsSupported(network, height):
		return &multisig8.LockBalanceParams{}, nil
	case tools.V17.IsSupported(network, height):
		return &multisig9.LockBalanceParams{}, nil
	case tools.V18.IsSupported(network, height):
		return &multisig10.LockBalanceParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return &multisig11.LockBalanceParams{}, nil
	case tools.V21.IsSupported(network, height):
		return &multisig12.LockBalanceParams{}, nil
	case tools.V22.IsSupported(network, height):
		return &multisig13.LockBalanceParams{}, nil
	case tools.V23.IsSupported(network, height):
		return &multisig14.LockBalanceParams{}, nil
	case tools.V24.IsSupported(network, height):
		return &multisig15.LockBalanceParams{}, nil
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
func withdrawBalanceParams(network string, height int64) (multisigParams, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return &legacyminer1.WithdrawBalanceParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return &legacyminer2.WithdrawBalanceParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return &legacyminer3.WithdrawBalanceParams{}, nil
	case tools.V12.IsSupported(network, height):
		return &legacyminer4.WithdrawBalanceParams{}, nil
	case tools.V13.IsSupported(network, height):
		return &legacyminer5.WithdrawBalanceParams{}, nil
	case tools.V14.IsSupported(network, height):
		return &legacyminer6.WithdrawBalanceParams{}, nil
	case tools.V15.IsSupported(network, height):
		return &legacyminer7.WithdrawBalanceParams{}, nil
	case tools.V16.IsSupported(network, height):
		return &miner8.WithdrawBalanceParams{}, nil
	case tools.V17.IsSupported(network, height):
		return &miner9.WithdrawBalanceParams{}, nil
	case tools.V18.IsSupported(network, height):
		return &miner10.WithdrawBalanceParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return &miner11.WithdrawBalanceParams{}, nil
	case tools.V21.IsSupported(network, height):
		return &miner12.WithdrawBalanceParams{}, nil
	case tools.V22.IsSupported(network, height):
		return &miner13.WithdrawBalanceParams{}, nil
	case tools.V23.IsSupported(network, height):
		return &miner14.WithdrawBalanceParams{}, nil
	case tools.V24.IsSupported(network, height):
		return &miner15.WithdrawBalanceParams{}, nil
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
func verifierParams(network string, height int64) (multisigParams, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return &legacyverifreg1.AddVerifierParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return &legacyverifreg2.AddVerifierParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return &legacyverifreg3.AddVerifierParams{}, nil
	case tools.V12.IsSupported(network, height):
		return &legacyverifreg4.AddVerifierParams{}, nil
	case tools.V13.IsSupported(network, height):
		return &legacyverifreg5.AddVerifierParams{}, nil
	case tools.V14.IsSupported(network, height):
		return &legacyverifreg6.AddVerifierParams{}, nil
	case tools.V15.IsSupported(network, height):
		return &legacyverifreg7.AddVerifierParams{}, nil
	case tools.V16.IsSupported(network, height):
		return &verifregv8.AddVerifierParams{}, nil
	case tools.V17.IsSupported(network, height):
		return &verifregv9.AddVerifierParams{}, nil
	case tools.V18.IsSupported(network, height):
		return &verifregv10.AddVerifierParams{}, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return &verifregv11.AddVerifierParams{}, nil
	case tools.V21.IsSupported(network, height):
		return &verifregv12.AddVerifierParams{}, nil
	case tools.V22.IsSupported(network, height):
		return &verifregv13.AddVerifierParams{}, nil
	case tools.V23.IsSupported(network, height):
		return &verifregv14.AddVerifierParams{}, nil
	case tools.V24.IsSupported(network, height):
		return &verifregv15.AddVerifierParams{}, nil
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
