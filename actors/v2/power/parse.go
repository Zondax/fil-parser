package power

import (
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *Power) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, *types.AddressInfo, error) {
	var err error
	var addressInfo *types.AddressInfo
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodSend:
		// metadata = p.parseSend(msg)
	case parser.MethodConstructor:
		metadata, err = p.Constructor(network, height, msg, msg.Params)
	case parser.MethodCreateMiner, parser.MethodCreateMinerExported:
		metadata, err = p.CreateMinerExported(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodUpdateClaimedPower:
		metadata, err = p.UpdateClaimedPower(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodEnrollCronEvent:
		metadata, err = p.EnrollCronEvent(network, msg, height, msg.Params, msgRct.Return)
	case parser.MethodCronTick:
		// metadata, err = p.emptyParamsAndReturn()
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
	case parser.UnknownStr:
		// metadata, err = p.unknownMetadata(msg.Params, msgRct.Return)
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, addressInfo, err
}
