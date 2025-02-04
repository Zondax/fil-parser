package v2

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/v2/account"
	"github.com/zondax/fil-parser/actors/v2/cron"
	"github.com/zondax/fil-parser/actors/v2/datacap"
	"github.com/zondax/fil-parser/actors/v2/eam"
	"github.com/zondax/fil-parser/actors/v2/evm"
	initActor "github.com/zondax/fil-parser/actors/v2/init"
	"github.com/zondax/fil-parser/actors/v2/market"
	"github.com/zondax/fil-parser/actors/v2/miner"
	"github.com/zondax/fil-parser/actors/v2/multisig"
	paymentchannel "github.com/zondax/fil-parser/actors/v2/paymentChannel"
	"github.com/zondax/fil-parser/actors/v2/power"
	"github.com/zondax/fil-parser/actors/v2/reward"
	verifiedregistry "github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

type Actor interface {
	Name() string
	Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error)
	TransactionTypes() map[string]any
}

type ActorParser struct {
	helper *helper.Helper
	logger *zap.Logger
}

func NewActorParser(helper *helper.Helper, logger *zap.Logger) *ActorParser {
	return &ActorParser{
		helper: helper,
		logger: logger2.GetSafeLogger(logger),
	}
}

func (p *ActorParser) GetMetadata(txType string, msg *parser.LotusMessage, mainMsgCid cid.Cid, msgRct *parser.LotusMessageReceipt,
	height int64, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	if msg == nil {
		return metadata, nil, nil
	}

	actor, err := p.helper.GetActorNameFromAddress(msg.To, height, key)
	if err != nil {
		return metadata, nil, err
	}
	network := ""

	var addressInfo *types.AddressInfo
	switch actor {
	case manifest.InitKey:
		initActor := &initActor.Init{}
		metadata, addressInfo, err = initActor.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.CronKey:
		cron := &cron.Cron{}
		metadata, addressInfo, err = cron.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.AccountKey:
		account := &account.Account{}
		metadata, addressInfo, err = account.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.PowerKey:
		power := &power.Power{}
		metadata, addressInfo, err = power.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.MinerKey:
		miner := &miner.Miner{}
		metadata, addressInfo, err = miner.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.MarketKey:
		market := &market.Market{}
		metadata, addressInfo, err = market.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.PaychKey:
		paymentChannel := &paymentchannel.PaymentChannel{}
		metadata, addressInfo, err = paymentChannel.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.MultisigKey:
		multisig := &multisig.Msig{}
		metadata, err = multisig.Parse(network, height, txType, msg, msgRct, key)
	case manifest.RewardKey:
		reward := &reward.Reward{}
		metadata, addressInfo, err = reward.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.VerifregKey:
		verifiedRegistry := &verifiedregistry.VerifiedRegistry{}
		metadata, addressInfo, err = verifiedRegistry.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.EvmKey:
		evm := &evm.Evm{}
		metadata, addressInfo, err = evm.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.EamKey:
		eam := &eam.Eam{}
		metadata, addressInfo, err = eam.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.DatacapKey:
		datacap := &datacap.Datacap{}
		metadata, addressInfo, err = datacap.Parse(network, height, txType, msg, msgRct, mainMsgCid)
	case manifest.EthAccountKey:
		// metadata, err = p.ParseEthAccount(txType, msg, msgRct)
	case manifest.PlaceholderKey:
		// metadata, err = p.ParsePlaceholder(txType, msg, msgRct)
	default:
		err = parser.ErrNotValidActor
	}
	return metadata, addressInfo, err
}
