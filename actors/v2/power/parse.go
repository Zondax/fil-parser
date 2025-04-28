package power

import (
	"context"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *Power) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	var err error
	var addressInfo *types.AddressInfo
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodSend:
		resp := actor_tools.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodConstructor:
		metadata, err = p.Constructor(network, height, msg, msg.Params)
	case parser.MethodCreateMiner, parser.MethodCreateMinerExported:
		metadata, addressInfo, err = p.CreateMinerExported(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodUpdateClaimedPower:
		metadata, err = p.UpdateClaimedPower(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodEnrollCronEvent:
		metadata, err = p.EnrollCronEvent(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodCronTick:
		resp, err := actor_tools.ParseEmptyParamsAndReturn()
		return resp, nil, err
	case parser.MethodUpdatePledgeTotal:
		metadata, err = p.UpdatePledgeTotal(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodSubmitPoRepForBulkVerify:
		metadata, err = p.SubmitPoRepForBulkVerify(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodCurrentTotalPower:
		metadata, err = p.CurrentTotalPower(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodNetworkRawPowerExported:
		metadata, err = p.NetworkRawPowerExported(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodMinerRawPowerExported:
		metadata, err = p.MinerRawPowerExported(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodMinerCountExported:
		metadata, err = p.MinerCountExported(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodMinerConsensusCountExported:
		metadata, err = p.MinerConsensusCountExported(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodOnEpochTickEnd:
		resp, err := actor_tools.ParseEmptyParamsAndReturn()
		return resp, nil, err
	case parser.MethodOnConsensusFault:
		metadata, err = p.OnConsensusFault(network, height, msg, msg.Params)
		return metadata, nil, err
	case parser.UnknownStr:
		resp, err := actor_tools.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, addressInfo, err
}

func (p *Power) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:                        actor_tools.ParseSend,
		parser.MethodConstructor:                 p.Constructor,
		parser.MethodCreateMiner:                 p.CreateMinerExported,
		parser.MethodCreateMinerExported:         p.CreateMinerExported,
		parser.MethodUpdateClaimedPower:          p.UpdateClaimedPower,
		parser.MethodEnrollCronEvent:             p.EnrollCronEvent,
		parser.MethodCronTick:                    actor_tools.ParseEmptyParamsAndReturn,
		parser.MethodUpdatePledgeTotal:           p.UpdatePledgeTotal,
		parser.MethodSubmitPoRepForBulkVerify:    p.SubmitPoRepForBulkVerify,
		parser.MethodCurrentTotalPower:           p.CurrentTotalPower,
		parser.MethodNetworkRawPowerExported:     p.NetworkRawPowerExported,
		parser.MethodMinerRawPowerExported:       p.MinerRawPowerExported,
		parser.MethodMinerCountExported:          p.MinerCountExported,
		parser.MethodMinerConsensusCountExported: p.MinerConsensusCountExported,
		parser.MethodOnEpochTickEnd:              actor_tools.ParseEmptyParamsAndReturn,
		parser.MethodOnConsensusFault:            p.OnConsensusFault,
	}
}
