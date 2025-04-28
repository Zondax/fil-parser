package account

import (
	"context"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Account struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Account {
	return &Account{
		logger: logger,
	}
}

func (a *Account) Name() string {
	return manifest.AccountKey
}

func (*Account) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

func (a *Account) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		resp := actor_tools.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodConstructor:
		resp, err := actor_tools.ParseConstructor(msg.Params)
		return resp, nil, err
	case parser.MethodPubkeyAddress:
		resp, err := a.PubkeyAddress(network, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodAuthenticateMessage:
		resp, err := a.AuthenticateMessage(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodUniversalReceiverHook, parser.MethodReceive:
		resp, err := a.UniversalReceiverHook(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodFallback:
		resp, err := a.Fallback(network, height, msg.Params)
		return resp, nil, err
	case parser.UnknownStr:
		resp, err := actor_tools.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (a *Account) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:                  actor_tools.ParseSend,
		parser.MethodConstructor:           actor_tools.ParseConstructor,
		parser.MethodPubkeyAddress:         a.PubkeyAddress,
		parser.MethodAuthenticateMessage:   a.AuthenticateMessage,
		parser.MethodUniversalReceiverHook: a.UniversalReceiverHook,
		parser.MethodReceive:               a.UniversalReceiverHook,
	}
}
