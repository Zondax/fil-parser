package actors

import (
	"fmt"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/actors/multisig"
	v11 "github.com/zondax/fil-parser/actors/v11"
	"github.com/zondax/fil-parser/parser"
)

var multisigParsers = map[string]multisig.MultisigParser{
	VersionV11: &v11.V11MultisigParser{},
	// Others versions
}

func (p *ActorParser) ParseMultisig(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	version := getVersionFromHeight(height)
	multisigParser, ok := multisigParsers[version]
	if !ok {
		return nil, fmt.Errorf("unsupported multisig version: %s", version)
	}
	return multisigParser.ParseMultisig(txType, msg, msgRct, height, key)
}
