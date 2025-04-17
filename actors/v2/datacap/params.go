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

var increaseAllowanceParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.IncreaseAllowanceParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.IncreaseAllowanceParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.IncreaseAllowanceParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.IncreaseAllowanceParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.IncreaseAllowanceParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.IncreaseAllowanceParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.IncreaseAllowanceParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.IncreaseAllowanceParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.IncreaseAllowanceParams) },
}

var decreaseAllowanceParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.DecreaseAllowanceParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.DecreaseAllowanceParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.DecreaseAllowanceParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.DecreaseAllowanceParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.DecreaseAllowanceParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.DecreaseAllowanceParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.DecreaseAllowanceParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.DecreaseAllowanceParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.DecreaseAllowanceParams) },
}

var revokeAllowanceParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.RevokeAllowanceParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.RevokeAllowanceParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.RevokeAllowanceParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.RevokeAllowanceParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.RevokeAllowanceParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.RevokeAllowanceParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.RevokeAllowanceParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.RevokeAllowanceParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.RevokeAllowanceParams) },
}

var allowanceParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.GetAllowanceParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.GetAllowanceParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.GetAllowanceParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.GetAllowanceParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.GetAllowanceParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.GetAllowanceParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.GetAllowanceParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.GetAllowanceParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.GetAllowanceParams) },
}

var burnParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.BurnParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.BurnParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.BurnParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.BurnParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.BurnParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.BurnParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.BurnParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.BurnParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.BurnParams) },
}

var burnReturn = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.BurnReturn) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.BurnReturn) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.BurnReturn) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.BurnReturn) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.BurnReturn) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.BurnReturn) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.BurnReturn) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.BurnReturn) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.BurnReturn) },
}

var burnFromParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.BurnFromParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.BurnFromParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.BurnFromParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.BurnFromParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.BurnFromParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.BurnFromParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.BurnFromParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.BurnFromParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.BurnFromParams) },
}

var burnFromReturn = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.BurnFromReturn) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.BurnFromReturn) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.BurnFromReturn) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.BurnFromReturn) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.BurnFromReturn) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.BurnFromReturn) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.BurnFromReturn) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.BurnFromReturn) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.BurnFromReturn) },
}

var destroyParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.DestroyParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.DestroyParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.DestroyParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.DestroyParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.DestroyParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.DestroyParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.DestroyParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.DestroyParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.DestroyParams) },
}

var granularityReturn = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.GranularityReturn) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.GranularityReturn) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.GranularityReturn) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.GranularityReturn) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.GranularityReturn) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.GranularityReturn) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.GranularityReturn) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.GranularityReturn) },
}

var mintParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.MintParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.MintParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.MintParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.MintParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.MintParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.MintParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.MintParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.MintParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.MintParams) },
}

var mintReturn = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.MintReturn) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.MintReturn) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.MintReturn) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.MintReturn) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.MintReturn) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.MintReturn) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.MintReturn) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.MintReturn) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.MintReturn) },
}

var transferParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.TransferParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.TransferParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.TransferParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.TransferParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.TransferParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.TransferParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.TransferParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.TransferParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.TransferParams) },
}

var transferReturn = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.TransferReturn) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.TransferReturn) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.TransferReturn) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.TransferReturn) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.TransferReturn) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.TransferReturn) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.TransferReturn) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.TransferReturn) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.TransferReturn) },
}

var transferFromParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.TransferFromParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.TransferFromParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.TransferFromParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.TransferFromParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.TransferFromParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.TransferFromParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.TransferFromParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.TransferFromParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.TransferFromParams) },
}

var transferFromReturn = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(datacapv9.TransferFromReturn) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(datacapv10.TransferFromReturn) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.TransferFromReturn) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(datacapv11.TransferFromReturn) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(datacapv12.TransferFromReturn) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(datacapv13.TransferFromReturn) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(datacapv14.TransferFromReturn) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(datacapv15.TransferFromReturn) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(datacapv16.TransferFromReturn) },
}
