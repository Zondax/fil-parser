package parser

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	parser "github.com/zondax/fil-parser/types"
)

func (p *Parser) searchForActorCreation(msg *filTypes.Message, receipt *filTypes.MessageReceipt,
	height int64, key filTypes.TipSetKey) (*parser.AddressInfo, error) {

	toAddressInfo := p.getActorAddressInfo(msg.To, height, key)
	actorName, err := p.lib.BuiltinActors.GetActorNameFromCid(toAddressInfo.ActorCid)
	if err != nil {
		return nil, err
	}

	switch actorName {
	case manifest.InitKey:
		{
			params, err := ParseInitActorExecParams(msg.Params)
			if err != nil {
				return nil, err
			}
			createdActorName, err := p.lib.BuiltinActors.GetActorNameFromCid(params.CodeCID)
			if err != nil {
				return nil, err
			}
			switch createdActorName {
			case manifest.MultisigKey, manifest.PaychKey:
				{
					execReturn, err := ParseExecReturn(receipt.Return)
					if err != nil {
						return nil, err
					}

					return &parser.AddressInfo{
						Short:     execReturn.IDAddress.String(),
						Robust:    execReturn.RobustAddress.String(),
						ActorCid:  params.CodeCID,
						ActorType: createdActorName,
					}, nil
				}
			default:
				return nil, nil
			}
		}
	case manifest.PowerKey:
		{
			execReturn, err := ParseExecReturn(receipt.Return)
			if err != nil {
				return nil, err
			}
			return &parser.AddressInfo{
				Short:     execReturn.IDAddress.String(),
				Robust:    execReturn.RobustAddress.String(),
				ActorType: "miner",
			}, nil
		}
	default:
		return nil, nil
	}
}

func IsOpSupported(op string) bool {
	supported, ok := SupportedOperations[op]
	if ok && supported {
		return true
	}

	return false
}

func SetupSupportedOperations(ops []string) {
	for s := range SupportedOperations {
		for _, op := range ops {
			found := false
			if s == op {
				found = true
			}
			SupportedOperations[s] = found
			if found {
				break
			}
		}
	}
}
