package actors

import (
	"bytes"
	"encoding/hex"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type ActorParserInterface interface {
	GetMetadata(txType string, msg *parser.LotusMessage, mainMsgCid cid.Cid, msgRct *parser.LotusMessageReceipt,
		height int64, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error)
}

func ParseSend(msg *parser.LotusMessage) map[string]interface{} {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = msg.Params
	return metadata
}

// ParseConstructor parse methods with format: *new(func(*address.Address) *abi.EmptyValue)
func ParseConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	return metadata, nil
}

func ParseUnknownMetadata(msgParams, msgReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if len(msgParams) > 0 {
		metadata[parser.ParamsKey] = hex.EncodeToString(msgParams)
	}
	if len(msgReturn) > 0 {
		metadata[parser.ReturnKey] = hex.EncodeToString(msgReturn)
	}
	return metadata, nil
}

func ParseEmptyParamsAndReturn() (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}
