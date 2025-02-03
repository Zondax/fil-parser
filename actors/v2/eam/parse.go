package eam

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type Eam struct{}

func (p *Eam) Name() string {
	return manifest.EamKey
}

func (p *Eam) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	var err error

	return metadata, nil, err
}
