package account

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Account struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Account {
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

func (a *Account) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		resp := actors.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodConstructor:
		resp, err := actors.ParseConstructor(msg.Params)
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
	case parser.UnknownStr:
		resp, err := actors.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (a *Account) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:                  actors.ParseSend,
		parser.MethodConstructor:           actors.ParseConstructor,
		parser.MethodPubkeyAddress:         a.PubkeyAddress,
		parser.MethodAuthenticateMessage:   a.AuthenticateMessage,
		parser.MethodUniversalReceiverHook: a.UniversalReceiverHook,
		parser.MethodReceive:               a.UniversalReceiverHook,
	}
}
