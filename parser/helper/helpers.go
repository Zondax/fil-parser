package helper

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v12/account"
	"github.com/filecoin-project/go-state-types/builtin/v12/cron"
	"github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	"github.com/filecoin-project/go-state-types/builtin/v12/eam"
	"github.com/filecoin-project/go-state-types/builtin/v12/evm"
	filInit "github.com/filecoin-project/go-state-types/builtin/v12/init"
	"github.com/filecoin-project/go-state-types/builtin/v12/market"
	"github.com/filecoin-project/go-state-types/builtin/v12/miner"
	"github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v12/paych"
	"github.com/filecoin-project/go-state-types/builtin/v12/power"
	"github.com/filecoin-project/go-state-types/builtin/v12/reward"
	"github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/go-resty/resty/v2"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/cache"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/golem/pkg/zcache"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"github.com/zondax/rosetta-filecoin-lib/actors"
	"go.uber.org/zap"

	"github.com/zondax/fil-parser/types"
)

var allMethods = map[string]map[abi.MethodNum]builtin.MethodMeta{
	manifest.InitKey:     filInit.Methods,
	manifest.CronKey:     cron.Methods,
	manifest.AccountKey:  account.Methods,
	manifest.PowerKey:    power.Methods,
	manifest.MinerKey:    miner.Methods,
	manifest.MarketKey:   market.Methods,
	manifest.PaychKey:    paych.Methods,
	manifest.MultisigKey: multisig.Methods,
	manifest.RewardKey:   reward.Methods,
	manifest.VerifregKey: verifreg.Methods,
	manifest.EvmKey:      evm.Methods,
	manifest.EamKey:      eam.Methods,
	manifest.DatacapKey:  datacap.Methods,

	// EthAccount and Placeholder can receive tokens with Send and InvokeEVM methods
	// We set evm.Methods instead of empty array of methods. Therefore, we will be able to understand
	// this specific method (3844450837) - tx cid example: bafy2bzacedgmcvsp56ieciutvgwza2qpvz7pvbhhu4l5y5tdl35rwfnjn5buk
	manifest.PlaceholderKey: evm.Methods,
	manifest.EthAccountKey:  evm.Methods,
}

type Helper struct {
	lib        *rosettaFilecoinLib.RosettaConstructionFilecoin
	node       api.FullNode
	actorCache *cache.ActorsCache
	logger     *zap.Logger
	httpClient *resty.Client
}

func NewHelper(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, actorsCache *cache.ActorsCache, node api.FullNode, logger *zap.Logger) *Helper {
	return &Helper{lib: lib, actorCache: actorsCache, node: node, logger: logger2.GetSafeLogger(logger), httpClient: resty.New().SetTimeout(30 * time.Second)}
}

func (h *Helper) GetActorsCache() *cache.ActorsCache {
	return h.actorCache
}

func (h *Helper) GetFilecoinLib() *rosettaFilecoinLib.RosettaConstructionFilecoin {
	return h.lib
}

func (h *Helper) GetFilecoinNodeClient() api.FullNode {
	return h.node
}

func (h *Helper) GetActorAddressInfo(add address.Address, key filTypes.TipSetKey) *types.AddressInfo {
	var err error
	addInfo := &types.AddressInfo{}

	addInfo.ActorCid, err = h.actorCache.GetActorCode(add, key, false)
	if err != nil {
		h.logger.Sugar().Errorf("could not get actor code from address. Err: %s", err)
	} else {
		c, err := cid.Parse(addInfo.ActorCid)
		if err != nil {
			h.logger.Sugar().Errorf("Could not parse params. Cannot cid.parse actor code: %v", err)
		}
		addInfo.ActorType, _ = h.lib.BuiltinActors.GetActorNameFromCid(c)
	}

	addInfo.Short, err = h.actorCache.GetShortAddress(add)
	if err != nil {
		h.logger.Sugar().Errorf("could not get short address for %s. Err: %v", add.String(), err)
	}

	// Ignore searching robust addresses for Msig and miners
	if addInfo.ActorType == manifest.MinerKey || addInfo.ActorType == manifest.MultisigKey {
		return addInfo
	}

	addInfo.Robust, err = h.actorCache.GetRobustAddress(add)
	if err != nil {
		h.logger.Sugar().Errorf("could not get robust address for %s. Err: %v", add.String(), err)
	}

	return addInfo
}

func (h *Helper) GetActorNameFromAddress(address address.Address, height int64, key filTypes.TipSetKey) (string, error) {
	onChainOnly := false
	for {
		// Search for actor in cache
		actorCode, err := h.actorCache.GetActorCode(address, key, onChainOnly)
		if err != nil {
			return actors.UnknownStr, err
		}

		c, err := cid.Parse(actorCode)
		if err != nil {
			h.logger.Sugar().Errorf("Could not parse params. Cannot cid.parse actor code: %v", err)
			return actors.UnknownStr, err
		}

		actorName, err := h.lib.BuiltinActors.GetActorNameFromCid(c)
		if err != nil {
			return actors.UnknownStr, err
		}

		if actorName == manifest.PlaceholderKey && !onChainOnly {
			onChainOnly = true
		} else {
			return actorName, nil
		}
	}
}

func (h *Helper) GetMethodName(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (string, error) {

	if msg == nil {
		return "", errors.New("malformed value")
	}

	// Shortcut 1 - Method "0" corresponds to "MethodSend"
	if msg.Method == 0 {
		return parser.MethodSend, nil
	}

	// Shortcut 2 - Method "1" corresponds to "MethodConstructor"
	if msg.Method == 1 {
		return parser.MethodConstructor, nil
	}

	actorName, _ := h.GetActorNameFromAddress(msg.To, height, key)

	actorMethods, ok := allMethods[actorName]
	if !ok {
		return "", parser.ErrNotKnownActor
	}
	method, ok := actorMethods[msg.Method]
	if !ok {
		return parser.UnknownStr, nil
	}
	return method.Name, nil
}

// GetEVMSelectorSig attempts to get the selector_signature from SignatureDBURL and utilizes a combined cache.
func (h *Helper) GetEVMSelectorSig(selectorID string, cache zcache.ZCache) (string, error) {
	ctx := context.Background()

	var selectorSig string
	if err := cache.Get(ctx, selectorID, &selectorSig); err != nil {
		if !cache.IsNotFoundError(err) {
			return "", err
		}
	}

	if selectorSig != "" {
		return selectorSig, nil
	}

	// not found in cache
	resp, err := h.httpClient.NewRequest().
		SetQueryParam("hex_signature", selectorID).
		SetContext(ctx).
		SetResult(&SignatureResult{}).
		Get(SignatureDBURL)
	if err != nil {
		return selectorSig, err
	}

	if resp.StatusCode() != http.StatusOK {
		return selectorSig, fmt.Errorf("error from 4bytes: %v", resp.Error())
	}

	signatureData, ok := resp.Result().(*SignatureResult)
	if !ok {
		return selectorSig, errors.New("error asserting result to SignatureResult")
	}

	if len(signatureData.Results) == 0 {
		return selectorSig, fmt.Errorf("signature not found: %s", selectorID)
	}

	sig := signatureData.Results[0].TextSignature

	if err := cache.Set(ctx, selectorID, sig, 0); err != nil {
		return selectorSig, fmt.Errorf("error adding selector_sig to cache: %w", err)
	}
	return sig, nil
}
