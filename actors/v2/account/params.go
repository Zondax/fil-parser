package account

import (
	accountv10 "github.com/filecoin-project/go-state-types/builtin/v10/account"
	accountv11 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv12 "github.com/filecoin-project/go-state-types/builtin/v12/account"
	accountv13 "github.com/filecoin-project/go-state-types/builtin/v13/account"
	accountv14 "github.com/filecoin-project/go-state-types/builtin/v14/account"
	accountv15 "github.com/filecoin-project/go-state-types/builtin/v15/account"
	accountv16 "github.com/filecoin-project/go-state-types/builtin/v16/account"
	accountv9 "github.com/filecoin-project/go-state-types/builtin/v9/account"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
)

func authenticateMessageParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &accountv9.AuthenticateMessageParams{},
		tools.V18.String(): &accountv10.AuthenticateMessageParams{},
		tools.V19.String(): &accountv11.AuthenticateMessageParams{},
		tools.V20.String(): &accountv11.AuthenticateMessageParams{},
		tools.V21.String(): &accountv12.AuthenticateMessageParams{},
		tools.V22.String(): &accountv13.AuthenticateMessageParams{},
		tools.V23.String(): &accountv14.AuthenticateMessageParams{},
		tools.V24.String(): &accountv15.AuthenticateMessageParams{},
		tools.V25.String(): &accountv16.AuthenticateMessageParams{},
	}
}
