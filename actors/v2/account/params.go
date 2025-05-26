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

var authenticateMessageParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(accountv9.AuthenticateMessageParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(accountv10.AuthenticateMessageParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(accountv11.AuthenticateMessageParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(accountv11.AuthenticateMessageParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(accountv12.AuthenticateMessageParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(accountv13.AuthenticateMessageParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(accountv14.AuthenticateMessageParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(accountv15.AuthenticateMessageParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(accountv16.AuthenticateMessageParams) },
}
