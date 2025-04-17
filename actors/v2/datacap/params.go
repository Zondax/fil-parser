package datacap

import (
	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv16 "github.com/filecoin-project/go-state-types/builtin/v16/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
)

func increaseAllowanceParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.IncreaseAllowanceParams{},
		tools.V18.String(): &datacapv10.IncreaseAllowanceParams{},
		tools.V19.String(): &datacapv11.IncreaseAllowanceParams{},
		tools.V20.String(): &datacapv11.IncreaseAllowanceParams{},
		tools.V21.String(): &datacapv12.IncreaseAllowanceParams{},
		tools.V22.String(): &datacapv13.IncreaseAllowanceParams{},
		tools.V23.String(): &datacapv14.IncreaseAllowanceParams{},
		tools.V24.String(): &datacapv15.IncreaseAllowanceParams{},
		tools.V25.String(): &datacapv16.IncreaseAllowanceParams{},
	}
}

func decreaseAllowanceParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.DecreaseAllowanceParams{},
		tools.V18.String(): &datacapv10.DecreaseAllowanceParams{},
		tools.V19.String(): &datacapv11.DecreaseAllowanceParams{},
		tools.V20.String(): &datacapv11.DecreaseAllowanceParams{},
		tools.V21.String(): &datacapv12.DecreaseAllowanceParams{},
		tools.V22.String(): &datacapv13.DecreaseAllowanceParams{},
		tools.V23.String(): &datacapv14.DecreaseAllowanceParams{},
		tools.V24.String(): &datacapv15.DecreaseAllowanceParams{},
		tools.V25.String(): &datacapv16.DecreaseAllowanceParams{},
	}
}

func revokeAllowanceParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.RevokeAllowanceParams{},
		tools.V18.String(): &datacapv10.RevokeAllowanceParams{},
		tools.V19.String(): &datacapv11.RevokeAllowanceParams{},
		tools.V20.String(): &datacapv11.RevokeAllowanceParams{},
		tools.V21.String(): &datacapv12.RevokeAllowanceParams{},
		tools.V22.String(): &datacapv13.RevokeAllowanceParams{},
		tools.V23.String(): &datacapv14.RevokeAllowanceParams{},
		tools.V24.String(): &datacapv15.RevokeAllowanceParams{},
		tools.V25.String(): &datacapv16.RevokeAllowanceParams{},
	}
}

func allowanceParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.GetAllowanceParams{},
		tools.V18.String(): &datacapv10.GetAllowanceParams{},
		tools.V19.String(): &datacapv11.GetAllowanceParams{},
		tools.V20.String(): &datacapv11.GetAllowanceParams{},
		tools.V21.String(): &datacapv12.GetAllowanceParams{},
		tools.V22.String(): &datacapv13.GetAllowanceParams{},
		tools.V23.String(): &datacapv14.GetAllowanceParams{},
		tools.V24.String(): &datacapv15.GetAllowanceParams{},
		tools.V25.String(): &datacapv16.GetAllowanceParams{},
	}
}

func burnParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.BurnParams{},
		tools.V18.String(): &datacapv10.BurnParams{},
		tools.V19.String(): &datacapv11.BurnParams{},
		tools.V20.String(): &datacapv11.BurnParams{},
		tools.V21.String(): &datacapv12.BurnParams{},
		tools.V22.String(): &datacapv13.BurnParams{},
		tools.V23.String(): &datacapv14.BurnParams{},
		tools.V24.String(): &datacapv15.BurnParams{},
		tools.V25.String(): &datacapv16.BurnParams{},
	}
}

func burnReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.BurnReturn{},
		tools.V18.String(): &datacapv10.BurnReturn{},
		tools.V19.String(): &datacapv11.BurnReturn{},
		tools.V20.String(): &datacapv11.BurnReturn{},
		tools.V21.String(): &datacapv12.BurnReturn{},
		tools.V22.String(): &datacapv13.BurnReturn{},
		tools.V23.String(): &datacapv14.BurnReturn{},
		tools.V24.String(): &datacapv15.BurnReturn{},
		tools.V25.String(): &datacapv16.BurnReturn{},
	}
}

func burnFromParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.BurnFromParams{},
		tools.V18.String(): &datacapv10.BurnFromParams{},
		tools.V19.String(): &datacapv11.BurnFromParams{},
		tools.V20.String(): &datacapv11.BurnFromParams{},
		tools.V21.String(): &datacapv12.BurnFromParams{},
		tools.V22.String(): &datacapv13.BurnFromParams{},
		tools.V23.String(): &datacapv14.BurnFromParams{},
		tools.V24.String(): &datacapv15.BurnFromParams{},
		tools.V25.String(): &datacapv16.BurnFromParams{},
	}
}

func burnFromReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.BurnFromReturn{},
		tools.V18.String(): &datacapv10.BurnFromReturn{},
		tools.V19.String(): &datacapv11.BurnFromReturn{},
		tools.V20.String(): &datacapv11.BurnFromReturn{},
		tools.V21.String(): &datacapv12.BurnFromReturn{},
		tools.V22.String(): &datacapv13.BurnFromReturn{},
		tools.V23.String(): &datacapv14.BurnFromReturn{},
		tools.V24.String(): &datacapv15.BurnFromReturn{},
		tools.V25.String(): &datacapv16.BurnFromReturn{},
	}
}

func destroyParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.DestroyParams{},
		tools.V18.String(): &datacapv10.DestroyParams{},
		tools.V19.String(): &datacapv11.DestroyParams{},
		tools.V20.String(): &datacapv11.DestroyParams{},
		tools.V21.String(): &datacapv12.DestroyParams{},
		tools.V22.String(): &datacapv13.DestroyParams{},
		tools.V23.String(): &datacapv14.DestroyParams{},
		tools.V24.String(): &datacapv15.DestroyParams{},
		tools.V25.String(): &datacapv16.DestroyParams{},
	}
}

func granuralityReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): new(datacapv10.GranularityReturn),
		tools.V19.String(): new(datacapv11.GranularityReturn),
		tools.V20.String(): new(datacapv11.GranularityReturn),
		tools.V21.String(): new(datacapv12.GranularityReturn),
		tools.V22.String(): new(datacapv13.GranularityReturn),
		tools.V23.String(): new(datacapv14.GranularityReturn),
		tools.V24.String(): new(datacapv15.GranularityReturn),
		tools.V25.String(): new(datacapv16.GranularityReturn),
	}
}

func mintParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.MintParams{},
		tools.V18.String(): &datacapv10.MintParams{},
		tools.V19.String(): &datacapv11.MintParams{},
		tools.V20.String(): &datacapv11.MintParams{},
		tools.V21.String(): &datacapv12.MintParams{},
		tools.V22.String(): &datacapv13.MintParams{},
		tools.V23.String(): &datacapv14.MintParams{},
		tools.V24.String(): &datacapv15.MintParams{},
		tools.V25.String(): &datacapv16.MintParams{},
	}
}

func mintReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.MintReturn{},
		tools.V18.String(): &datacapv10.MintReturn{},
		tools.V19.String(): &datacapv11.MintReturn{},
		tools.V20.String(): &datacapv11.MintReturn{},
		tools.V21.String(): &datacapv12.MintReturn{},
		tools.V22.String(): &datacapv13.MintReturn{},
		tools.V23.String(): &datacapv14.MintReturn{},
		tools.V24.String(): &datacapv15.MintReturn{},
		tools.V25.String(): &datacapv16.MintReturn{},
	}
}

func transferParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.TransferParams{},
		tools.V18.String(): &datacapv10.TransferParams{},
		tools.V19.String(): &datacapv11.TransferParams{},
		tools.V20.String(): &datacapv11.TransferParams{},
		tools.V21.String(): &datacapv12.TransferParams{},
		tools.V22.String(): &datacapv13.TransferParams{},
		tools.V23.String(): &datacapv14.TransferParams{},
		tools.V24.String(): &datacapv15.TransferParams{},
		tools.V25.String(): &datacapv16.TransferParams{},
	}
}

func transferReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.TransferReturn{},
		tools.V18.String(): &datacapv10.TransferReturn{},
		tools.V19.String(): &datacapv11.TransferReturn{},
		tools.V20.String(): &datacapv11.TransferReturn{},
		tools.V21.String(): &datacapv12.TransferReturn{},
		tools.V22.String(): &datacapv13.TransferReturn{},
		tools.V23.String(): &datacapv14.TransferReturn{},
		tools.V24.String(): &datacapv15.TransferReturn{},
		tools.V25.String(): &datacapv16.TransferReturn{},
	}
}

func transferFromParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.TransferFromParams{},
		tools.V18.String(): &datacapv10.TransferFromParams{},
		tools.V19.String(): &datacapv11.TransferFromParams{},
		tools.V20.String(): &datacapv11.TransferFromParams{},
		tools.V21.String(): &datacapv12.TransferFromParams{},
		tools.V22.String(): &datacapv13.TransferFromParams{},
		tools.V23.String(): &datacapv14.TransferFromParams{},
		tools.V24.String(): &datacapv15.TransferFromParams{},
		tools.V25.String(): &datacapv16.TransferFromParams{},
	}
}

func transferFromReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.TransferFromReturn{},
		tools.V18.String(): &datacapv10.TransferFromReturn{},
		tools.V19.String(): &datacapv11.TransferFromReturn{},
		tools.V20.String(): &datacapv11.TransferFromReturn{},
		tools.V21.String(): &datacapv12.TransferFromReturn{},
		tools.V22.String(): &datacapv13.TransferFromReturn{},
		tools.V23.String(): &datacapv14.TransferFromReturn{},
		tools.V24.String(): &datacapv15.TransferFromReturn{},
		tools.V25.String(): &datacapv16.TransferFromReturn{},
	}
}
