package multisig

import (
	"bytes"
	"fmt"

	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"

	"github.com/filecoin-project/go-address"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/multisig"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/multisig"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/multisig"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/multisig"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/multisig"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/multisig"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/multisig"

	"github.com/filecoin-project/go-state-types/abi"
	multisig10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	multisig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisig12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisig13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisig14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisig15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisig16 "github.com/filecoin-project/go-state-types/builtin/v16/multisig"
	multisig17 "github.com/filecoin-project/go-state-types/builtin/v17/multisig"
	multisig8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisig9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"
	"github.com/filecoin-project/go-state-types/exitcode"
)

var removeSignerParams2 = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return &legacyv1.RemoveSignerParams{} },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return &legacyv1.RemoveSignerParams{} },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return &legacyv1.RemoveSignerParams{} },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return &legacyv1.RemoveSignerParams{} },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return &legacyv2.RemoveSignerParams{} },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return &legacyv2.RemoveSignerParams{} },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return &legacyv2.RemoveSignerParams{} },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return &legacyv2.RemoveSignerParams{} },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return &legacyv2.RemoveSignerParams{} },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return &legacyv2.RemoveSignerParams{} },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return &legacyv3.RemoveSignerParams{} },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return &legacyv3.RemoveSignerParams{} },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return &legacyv4.RemoveSignerParams{} },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return &legacyv5.RemoveSignerParams{} },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return &legacyv6.RemoveSignerParams{} },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return &legacyv7.RemoveSignerParams{} },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return &multisig8.RemoveSignerParams{} },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return &multisig9.RemoveSignerParams{} },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return &multisig10.RemoveSignerParams{} },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return &multisig11.RemoveSignerParams{} },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return &multisig11.RemoveSignerParams{} },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return &multisig12.RemoveSignerParams{} },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return &multisig13.RemoveSignerParams{} },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return &multisig14.RemoveSignerParams{} },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return &multisig15.RemoveSignerParams{} },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return &multisig16.RemoveSignerParams{} },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return &multisig16.RemoveSignerParams{} },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return &multisig17.RemoveSignerParams{} },
}

var changeNumApprovalsThresholdParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ChangeNumApprovalsThresholdParams{} },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ChangeNumApprovalsThresholdParams{} },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ChangeNumApprovalsThresholdParams{} },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ChangeNumApprovalsThresholdParams{} },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ChangeNumApprovalsThresholdParams{} },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ChangeNumApprovalsThresholdParams{} },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ChangeNumApprovalsThresholdParams{} },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ChangeNumApprovalsThresholdParams{} },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ChangeNumApprovalsThresholdParams{} },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ChangeNumApprovalsThresholdParams{} },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return &legacyv3.ChangeNumApprovalsThresholdParams{} },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return &legacyv3.ChangeNumApprovalsThresholdParams{} },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return &legacyv4.ChangeNumApprovalsThresholdParams{} },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return &legacyv5.ChangeNumApprovalsThresholdParams{} },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return &legacyv6.ChangeNumApprovalsThresholdParams{} },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return &legacyv7.ChangeNumApprovalsThresholdParams{} },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return &multisig8.ChangeNumApprovalsThresholdParams{} },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return &multisig9.ChangeNumApprovalsThresholdParams{} },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return &multisig10.ChangeNumApprovalsThresholdParams{} },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return &multisig11.ChangeNumApprovalsThresholdParams{} },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return &multisig11.ChangeNumApprovalsThresholdParams{} },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return &multisig12.ChangeNumApprovalsThresholdParams{} },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return &multisig13.ChangeNumApprovalsThresholdParams{} },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return &multisig14.ChangeNumApprovalsThresholdParams{} },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return &multisig15.ChangeNumApprovalsThresholdParams{} },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return &multisig16.ChangeNumApprovalsThresholdParams{} },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return &multisig16.ChangeNumApprovalsThresholdParams{} },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return &multisig17.ChangeNumApprovalsThresholdParams{} },
}

var lockBalanceParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return &legacyv1.LockBalanceParams{} },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return &legacyv1.LockBalanceParams{} },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return &legacyv1.LockBalanceParams{} },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return &legacyv1.LockBalanceParams{} },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return &legacyv2.LockBalanceParams{} },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return &legacyv2.LockBalanceParams{} },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return &legacyv2.LockBalanceParams{} },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return &legacyv2.LockBalanceParams{} },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return &legacyv2.LockBalanceParams{} },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return &legacyv2.LockBalanceParams{} },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return &legacyv3.LockBalanceParams{} },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return &legacyv3.LockBalanceParams{} },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return &legacyv4.LockBalanceParams{} },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return &legacyv5.LockBalanceParams{} },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return &legacyv6.LockBalanceParams{} },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return &legacyv7.LockBalanceParams{} },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return &multisig8.LockBalanceParams{} },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return &multisig9.LockBalanceParams{} },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return &multisig10.LockBalanceParams{} },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return &multisig11.LockBalanceParams{} },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return &multisig11.LockBalanceParams{} },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return &multisig12.LockBalanceParams{} },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return &multisig13.LockBalanceParams{} },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return &multisig14.LockBalanceParams{} },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return &multisig15.LockBalanceParams{} },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return &multisig16.LockBalanceParams{} },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return &multisig16.LockBalanceParams{} },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return &multisig17.LockBalanceParams{} },
}

var approveReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ApproveReturn{} },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ApproveReturn{} },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ApproveReturn{} },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ApproveReturn{} },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ApproveReturn{} },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ApproveReturn{} },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ApproveReturn{} },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ApproveReturn{} },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ApproveReturn{} },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ApproveReturn{} },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return &legacyv3.ApproveReturn{} },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return &legacyv3.ApproveReturn{} },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return &legacyv4.ApproveReturn{} },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return &legacyv5.ApproveReturn{} },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return &legacyv6.ApproveReturn{} },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return &legacyv7.ApproveReturn{} },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return &multisig8.ApproveReturn{} },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return &multisig9.ApproveReturn{} },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return &multisig10.ApproveReturn{} },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return &multisig11.ApproveReturn{} },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return &multisig11.ApproveReturn{} },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return &multisig12.ApproveReturn{} },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return &multisig13.ApproveReturn{} },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return &multisig14.ApproveReturn{} },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return &multisig15.ApproveReturn{} },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return &multisig16.ApproveReturn{} },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return &multisig16.ApproveReturn{} },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return &multisig17.ApproveReturn{} },
}

var constructorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ConstructorParams{} },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ConstructorParams{} },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ConstructorParams{} },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ConstructorParams{} },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ConstructorParams{} },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ConstructorParams{} },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ConstructorParams{} },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ConstructorParams{} },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ConstructorParams{} },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ConstructorParams{} },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return &legacyv3.ConstructorParams{} },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return &legacyv3.ConstructorParams{} },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return &legacyv4.ConstructorParams{} },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return &legacyv5.ConstructorParams{} },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return &legacyv6.ConstructorParams{} },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return &legacyv7.ConstructorParams{} },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return &multisig8.ConstructorParams{} },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return &multisig9.ConstructorParams{} },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return &multisig10.ConstructorParams{} },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return &multisig11.ConstructorParams{} },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return &multisig11.ConstructorParams{} },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return &multisig12.ConstructorParams{} },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return &multisig13.ConstructorParams{} },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return &multisig14.ConstructorParams{} },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return &multisig15.ConstructorParams{} },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return &multisig16.ConstructorParams{} },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return &multisig16.ConstructorParams{} },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return &multisig17.ConstructorParams{} },
}

var addSignerParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return &legacyv1.AddSignerParams{} },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return &legacyv1.AddSignerParams{} },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return &legacyv1.AddSignerParams{} },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return &legacyv1.AddSignerParams{} },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return &legacyv2.AddSignerParams{} },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return &legacyv2.AddSignerParams{} },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return &legacyv2.AddSignerParams{} },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return &legacyv2.AddSignerParams{} },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return &legacyv2.AddSignerParams{} },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return &legacyv2.AddSignerParams{} },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return &legacyv3.AddSignerParams{} },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return &legacyv3.AddSignerParams{} },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return &legacyv4.AddSignerParams{} },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return &legacyv5.AddSignerParams{} },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return &legacyv6.AddSignerParams{} },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return &legacyv7.AddSignerParams{} },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return &multisig8.AddSignerParams{} },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return &multisig9.AddSignerParams{} },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return &multisig10.AddSignerParams{} },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return &multisig11.AddSignerParams{} },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return &multisig11.AddSignerParams{} },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return &multisig12.AddSignerParams{} },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return &multisig13.AddSignerParams{} },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return &multisig14.AddSignerParams{} },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return &multisig15.AddSignerParams{} },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return &multisig16.AddSignerParams{} },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return &multisig16.AddSignerParams{} },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return &multisig17.AddSignerParams{} },
}

var swapSignerParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return &legacyv1.SwapSignerParams{} },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return &legacyv1.SwapSignerParams{} },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return &legacyv1.SwapSignerParams{} },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return &legacyv1.SwapSignerParams{} },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return &legacyv2.SwapSignerParams{} },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return &legacyv2.SwapSignerParams{} },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return &legacyv2.SwapSignerParams{} },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return &legacyv2.SwapSignerParams{} },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return &legacyv2.SwapSignerParams{} },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return &legacyv2.SwapSignerParams{} },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return &legacyv3.SwapSignerParams{} },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return &legacyv3.SwapSignerParams{} },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return &legacyv4.SwapSignerParams{} },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return &legacyv5.SwapSignerParams{} },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return &legacyv6.SwapSignerParams{} },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return &legacyv7.SwapSignerParams{} },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return &multisig8.SwapSignerParams{} },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return &multisig9.SwapSignerParams{} },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return &multisig10.SwapSignerParams{} },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return &multisig11.SwapSignerParams{} },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return &multisig11.SwapSignerParams{} },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return &multisig12.SwapSignerParams{} },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return &multisig13.SwapSignerParams{} },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return &multisig14.SwapSignerParams{} },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return &multisig15.SwapSignerParams{} },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return &multisig16.SwapSignerParams{} },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return &multisig16.SwapSignerParams{} },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return &multisig17.SwapSignerParams{} },
}

var txnIDParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return &legacyv1.TxnIDParams{} },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return &legacyv1.TxnIDParams{} },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return &legacyv1.TxnIDParams{} },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return &legacyv1.TxnIDParams{} },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return &legacyv2.TxnIDParams{} },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return &legacyv2.TxnIDParams{} },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return &legacyv2.TxnIDParams{} },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return &legacyv2.TxnIDParams{} },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return &legacyv2.TxnIDParams{} },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return &legacyv2.TxnIDParams{} },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return &legacyv3.TxnIDParams{} },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return &legacyv3.TxnIDParams{} },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return &legacyv4.TxnIDParams{} },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return &legacyv5.TxnIDParams{} },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return &legacyv6.TxnIDParams{} },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return &legacyv7.TxnIDParams{} },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return &multisig8.TxnIDParams{} },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return &multisig9.TxnIDParams{} },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return &multisig10.TxnIDParams{} },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return &multisig11.TxnIDParams{} },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return &multisig11.TxnIDParams{} },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return &multisig12.TxnIDParams{} },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return &multisig13.TxnIDParams{} },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return &multisig14.TxnIDParams{} },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return &multisig15.TxnIDParams{} },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return &multisig16.TxnIDParams{} },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return &multisig16.TxnIDParams{} },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return &multisig17.TxnIDParams{} },
}

var proposeReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ProposeReturn{} },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ProposeReturn{} },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ProposeReturn{} },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ProposeReturn{} },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeReturn{} },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeReturn{} },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeReturn{} },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeReturn{} },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeReturn{} },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeReturn{} },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return &legacyv3.ProposeReturn{} },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return &legacyv3.ProposeReturn{} },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return &legacyv4.ProposeReturn{} },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return &legacyv5.ProposeReturn{} },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return &legacyv6.ProposeReturn{} },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return &legacyv7.ProposeReturn{} },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return &multisig8.ProposeReturn{} },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return &multisig9.ProposeReturn{} },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return &multisig10.ProposeReturn{} },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return &multisig11.ProposeReturn{} },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return &multisig11.ProposeReturn{} },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return &multisig12.ProposeReturn{} },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return &multisig13.ProposeReturn{} },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return &multisig14.ProposeReturn{} },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return &multisig15.ProposeReturn{} },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return &multisig16.ProposeReturn{} },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return &multisig16.ProposeReturn{} },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return &multisig17.ProposeReturn{} },
}

var proposeParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ProposeParams{} },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ProposeParams{} },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ProposeParams{} },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return &legacyv1.ProposeParams{} },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeParams{} },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeParams{} },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeParams{} },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeParams{} },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeParams{} },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return &legacyv2.ProposeParams{} },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return &legacyv3.ProposeParams{} },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return &legacyv3.ProposeParams{} },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return &legacyv4.ProposeParams{} },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return &legacyv5.ProposeParams{} },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return &legacyv6.ProposeParams{} },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return &legacyv7.ProposeParams{} },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return &multisig8.ProposeParams{} },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return &multisig9.ProposeParams{} },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return &multisig10.ProposeParams{} },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return &multisig11.ProposeParams{} },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return &multisig11.ProposeParams{} },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return &multisig12.ProposeParams{} },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return &multisig13.ProposeParams{} },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return &multisig14.ProposeParams{} },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return &multisig15.ProposeParams{} },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return &multisig16.ProposeParams{} },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return &multisig16.ProposeParams{} },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return &multisig17.ProposeParams{} },
}

func getProposeParams(network string, height int64, rawParams []byte) (raw []byte, methodNum abi.MethodNum, to address.Address, value string, params cbg.CBORUnmarshaler, err error) {
	version := tools.VersionFromHeight(network, height)
	tmp, ok := proposeParams[version.String()]
	if !ok {
		return nil, 0, address.Address{}, "", nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	val := tmp()
	err = val.UnmarshalCBOR(bytes.NewReader(rawParams))
	if err != nil {
		return nil, 0, address.Address{}, "", nil, err
	}

	switch parsedParams := val.(type) {
	case *legacyv1.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	// exact same type, commented out due to compiler error.
	// case *legacyv2.ProposeParams:
	// 	return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	// case *legacyv3.ProposeParams:
	// 	return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	// case *legacyv4.ProposeParams:
	// 	return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	// case *legacyv5.ProposeParams:
	// 	return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	// case *legacyv6.ProposeParams:
	// 	return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	// case *legacyv7.ProposeParams:
	// 	return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	case *multisig8.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	case *multisig9.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	case *multisig10.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	case *multisig11.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	case *multisig12.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	case *multisig13.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	case *multisig14.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	case *multisig15.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	case *multisig16.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	case *multisig17.ProposeParams:
		return parsedParams.Params, parsedParams.Method, parsedParams.To, parsedParams.Value.String(), parsedParams, nil
	default:
		return nil, 0, address.Address{}, "", nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func getProposeReturn(network string, height int64, rawReturn []byte) (applied bool, exitCode exitcode.ExitCode, raw []byte, retValue cbg.CBORUnmarshaler, err error) {
	version := tools.VersionFromHeight(network, height)
	tmp, ok := proposeReturn[version.String()]
	if !ok {
		return false, 0, rawReturn, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	val := tmp()
	err = val.UnmarshalCBOR(bytes.NewReader(rawReturn))
	if err != nil {
		return false, 0, rawReturn, nil, err
	}

	switch parsedReturn := val.(type) {
	case *legacyv1.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	// exact same type, commented out due to compiler error.
	// case *legacyv2.ProposeParams:
	// 	return parsedReturn.Applied, parsedReturn.Ret, parsedReturn, nil
	// case *legacyv3.ProposeParams:
	// 	return parsedReturn.Applied, parsedReturn.Ret, parsedReturn, nil
	// case *legacyv4.ProposeParams:
	// 	return parsedReturn.Applied, parsedReturn.Ret, parsedReturn, nil
	// case *legacyv5.ProposeParams:
	// 	return parsedReturn.Applied, parsedReturn.Ret, parsedReturn, nil
	// case *legacyv6.ProposeParams:
	// 	return parsedReturn.Applied, parsedReturn.Ret, parsedReturn, nil
	// case *legacyv7.ProposeParams:
	// 	return parsedReturn.Applied, parsedReturn.Ret, parsedReturn, nil
	case *multisig8.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	case *multisig9.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	case *multisig10.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	case *multisig11.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	case *multisig12.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	case *multisig13.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	case *multisig14.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	case *multisig15.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	case *multisig16.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	case *multisig17.ProposeReturn:
		return parsedReturn.Applied, parsedReturn.Code, parsedReturn.Ret, parsedReturn, nil
	default:
		return false, 0, rawReturn, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}
