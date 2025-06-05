package helper

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/zondax/golem/pkg/logger"

	"github.com/ipfs/go-cid"
	// The following import is necessary to ensure that the init() function
	// from the lotus build package is invoked.
	// In a recent refactor (v1.30.0), some build packages were modularized to reduce
	// unnecessary dependencies. As a result, if this package is not explicitly
	// imported, its init() will not be triggered, potentially causing issues
	// with initialization, such as errors when searching for actorNameByCid.
	_ "github.com/filecoin-project/lotus/build"

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

	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/tools"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"github.com/zondax/rosetta-filecoin-lib/actors"

	"github.com/zondax/fil-parser/actors/cache"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	parsermetrics "github.com/zondax/fil-parser/parser/metrics"
	"github.com/zondax/fil-parser/types"
)

const (
	// keylessAccountActor f090 was a multisig actor until V23 where it was converted to an account actor
	// https://github.com/filecoin-project/lotus/releases/tag/v1.28.1
	// https://github.com/filecoin-project/FIPs/blob/master/FIPS/fip-0085.md
	keylessAccountActor = "f090"
	// ZeroAddressAccountActor f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a is a zero address actor that existed until V10.
	// Created: https://github.com/filecoin-project/lotus/blob/5750f49834deee9dfce752ff840630ae402a8b51/build/buildconstants/params_shared_vals.go#L56
	// Terminated: https://github.com/filecoin-project/lotus/blob/5750f49834deee9dfce752ff840630ae402a8b51/chain/consensus/filcns/upgrades.go#L1054
	ZeroAddressAccountActor = "f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a"
	// multisig actorcode for nv22
	msigCidStr = "bafk2bzacedef4sqdsfebspu7dqnk7naj27ac4lyho4zmvjrei5qnf2wn6v64u"
	// account actorcode for nv23
	accountCidStr = "bafk2bzacedbgei6jkx36fwdgvoohce4aghvpohqdhoco7p4thszgssms7olv2"
)

// Deprecated: Use v2/tools.GetMethodName instead
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

var (
	msigCid    = cid.MustParse(msigCidStr)
	accountCid = cid.MustParse(accountCidStr)

	// specialLegacyActors is a list of actors not included in the lotus manifest but appear on the network.
	// StateGetActor(f067253) returns "bafkqadlgnfwc6mrpmfrwg33vnz2a" which is not included in the Lotus Manifest along with the following actor cids.
	// https://github.com/filecoin-project/statediff/blob/3e676285574e7bdb4ae0b9e28e6f23cfc86dd089/transform.go#L164
	specialLegacyActors = map[string]string{
		// v1
		"bafkqaddgnfwc6mjpon4xg5dfnu":                 manifest.SystemKey,
		"bafkqactgnfwc6mjpnfxgs5a":                    manifest.InitKey,
		"bafkqaddgnfwc6mjpojsxoylsmq":                 manifest.RewardKey,
		"bafkqactgnfwc6mjpmnzg63q":                    manifest.CronKey,
		"bafkqaetgnfwc6mjpon2g64tbm5sxa33xmvza":       manifest.PowerKey,
		"bafkqae3gnfwc6mjpon2g64tbm5sw2ylsnnsxi":      manifest.MarketKey,
		"bafkqaftgnfwc6mjpozsxe2lgnfswi4tfm5uxg5dspe": manifest.VerifregKey,
		"bafkqadlgnfwc6mjpmfrwg33vnz2a":               manifest.AccountKey,
		"bafkqadtgnfwc6mjpnv2wy5djonuwo":              manifest.MultisigKey,
		"bafkqafdgnfwc6mjpobqxs3lfnz2gg2dbnzxgk3a":    manifest.PaychKey,
		"bafkqaetgnfwc6mjpon2g64tbm5sw22lomvza":       manifest.MinerKey,

		// v2
		"bafkqaddgnfwc6mrpon4xg5dfnu":                 manifest.SystemKey,
		"bafkqactgnfwc6mrpnfxgs5a":                    manifest.InitKey,
		"bafkqaddgnfwc6mrpojsxoylsmq":                 manifest.RewardKey,
		"bafkqactgnfwc6mrpmnzg63q":                    manifest.CronKey,
		"bafkqaetgnfwc6mrpon2g64tbm5sxa33xmvza":       manifest.PowerKey,
		"bafkqae3gnfwc6mrpon2g64tbm5sw2ylsnnsxi":      manifest.MarketKey,
		"bafkqaftgnfwc6mrpozsxe2lgnfswi4tfm5uxg5dspe": manifest.VerifregKey,
		"bafkqadlgnfwc6mrpmfrwg33vnz2a":               manifest.AccountKey,
		"bafkqadtgnfwc6mrpnv2wy5djonuwo":              manifest.MultisigKey,
		"bafkqafdgnfwc6mrpobqxs3lfnz2gg2dbnzxgk3a":    manifest.PaychKey,
		"bafkqaetgnfwc6mrpon2g64tbm5sw22lomvza":       manifest.MinerKey,
	}

	// https://github.com/filecoin-project/lotus/blob/58c1ed844b2424a66008728e4c135fa2f6097b60/build/builtin_actors.go#L59
	calibrationBuggyActors = map[string]string{
		"bafk2bzacecnh2ouohmonvebq7uughh4h3ppmg4cjsk74dzxlbbtlcij4xbzxq": manifest.MinerKey,
		"bafk2bzaced7emkbbnrewv5uvrokxpf5tlm4jslu2jsv77ofw2yqdglg657uie": manifest.MinerKey,
		"bafk2bzacednskl3bykz5qpo54z2j2p4q44t5of4ktd6vs6ymmg2zebsbxazkm": manifest.VerifregKey,
	}
)

type Helper struct {
	lib        *rosettaFilecoinLib.RosettaConstructionFilecoin
	node       api.FullNode
	actorCache cache.IActorsCache
	logger     *logger.Logger
	metrics    *parsermetrics.ParserMetricsClient
	network    string
}

func NewHelper(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, actorsCache cache.IActorsCache, node api.FullNode, logger *logger.Logger, metrics metrics.MetricsClient) *Helper {
	h := &Helper{
		lib:        lib,
		actorCache: actorsCache,
		node:       node,
		logger:     logger2.GetSafeLogger(logger),
		metrics:    parsermetrics.NewClient(metrics, "helper"),
	}
	network, err := h.node.StateNetworkName(context.Background())
	if err != nil {
		h.logger.Errorf("could not get network name: %v", err)
		return nil
	}
	h.network = tools.ParseRawNetworkName(string(network))
	return h
}

func (h *Helper) GetActorsCache() cache.IActorsCache {
	return h.actorCache
}

func (h *Helper) GetFilecoinLib() *rosettaFilecoinLib.RosettaConstructionFilecoin {
	return h.lib
}

func (h *Helper) GetFilecoinNodeClient() api.FullNode {
	return h.node
}

func (h *Helper) GetActorAddressInfo(add address.Address, key filTypes.TipSetKey, height abi.ChainEpoch) *types.AddressInfo {
	var err error
	addInfo := &types.AddressInfo{}

	if add == address.Undef {
		return addInfo
	}

	actorCid, actorName, err := h.GetActorNameFromAddress(add, int64(height), key)
	if err != nil {
		h.logger.Errorf("could not get actor cid and name from address. Err: %s", err)
	} else {
		addInfo.ActorCid = actorCid.String()
		addInfo.ActorType = actorName
	}

	addInfo.Short, err = h.actorCache.GetShortAddress(add)
	if err != nil {
		h.logger.Errorf("could not get short address for %s. Err: %v", add.String(), err)
	}

	addInfo.Robust, err = h.actorCache.GetRobustAddress(add)
	if err != nil {
		h.logger.Errorf("could not get robust address for %s. Err: %v", add.String(), err)
	}

	addInfo.IsSystemActor = h.IsSystemActor(add) || h.IsGenesisActor(add)

	return addInfo
}

func (h *Helper) GetActorNameFromAddress(add address.Address, height int64, key filTypes.TipSetKey) (cid.Cid, string, error) {
	if add == address.Undef {
		return cid.Undef, "", errors.New("address is undefined")
	}

	if ok, cid, actorName := h.isSpecialAccountActor(add, height); ok {
		return cid, actorName, nil
	}

	onChainOnly := false
	for {
		// Search for actor in cache
		actorCode, err := h.actorCache.GetActorCode(add, key, onChainOnly)
		if err != nil {
			return cid.Undef, actors.UnknownStr, err
		}

		c, err := cid.Parse(actorCode)
		if err != nil {
			h.logger.Errorf("Could not parse params. Cannot cid.parse actor code: %v", err)
			return cid.Undef, actors.UnknownStr, err
		}

		actorName, err := h.GetActorNameFromCid(c, height)
		if err != nil {
			return cid.Undef, actors.UnknownStr, err
		}

		if actorName == manifest.PlaceholderKey && !onChainOnly {
			onChainOnly = true
		} else {
			return c, actorName, nil
		}
	}
}

// isSpecialAccountActor handles actor addresses that will fail to resolve from the node for reasons documented in each case.
func (h *Helper) isSpecialAccountActor(add address.Address, height int64) (bool, cid.Cid, string) {
	if ok, cid, actorName := h.isZeroAddressAccountActor(add, height); ok {
		return true, cid, actorName
	}
	if ok, cid, actorName := h.isKeylessAccountActor(add, height); ok {
		return true, cid, actorName
	}
	return false, cid.Undef, ""
}

// The f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a is a zero address actor that existed until V10.
// Created: https://github.com/filecoin-project/lotus/blob/5750f49834deee9dfce752ff840630ae402a8b51/build/buildconstants/params_shared_vals.go#L56
// Terminated: https://github.com/filecoin-project/lotus/blob/5750f49834deee9dfce752ff840630ae402a8b51/chain/consensus/filcns/upgrades.go#L1054
func (h *Helper) isZeroAddressAccountActor(add address.Address, height int64) (bool, cid.Cid, string) {
	if h.network != tools.MainnetNetwork || add.String() != ZeroAddressAccountActor {
		return false, cid.Undef, ""
	}
	version := tools.VersionFromHeight(h.network, int64(height))
	if version.NodeVersion() < tools.V23.NodeVersion() {
		return true, accountCid, manifest.AccountKey
	}
	return false, cid.Undef, ""
}

// The f090 address was a multisig actor until V23 where it was converted to an account actor
// https://github.com/filecoin-project/lotus/releases/tag/v1.28.1
// https://github.com/filecoin-project/FIPs/blob/master/FIPS/fip-0085.md
func (h *Helper) isKeylessAccountActor(add address.Address, height int64) (bool, cid.Cid, string) {
	if h.network != tools.MainnetNetwork || add.String() != keylessAccountActor {
		return false, cid.Undef, ""
	}
	version := tools.VersionFromHeight(h.network, int64(height))
	if version.NodeVersion() < tools.V23.NodeVersion() {
		return true, msigCid, manifest.MultisigKey
	}
	return true, accountCid, manifest.AccountKey
}

// GetActorNameFromCid returns the actor name for the given cid and height from rosetta and fallsback to specialLegacyActors.
func (h *Helper) GetActorNameFromCid(cid cid.Cid, height int64) (string, error) {
	version := tools.VersionFromHeight(h.network, height)
	actorName, err := h.lib.BuiltinActors.GetActorNameFromCidByVersion(cid, version.FilNetworkVersion())
	if err != nil {
		// fallback to specialLegacyActors
		if name, ok := specialLegacyActors[cid.String()]; ok && h.network == tools.MainnetNetwork {
			return name, nil
		}
		// fallback to calibrationBuggyActors
		if name, ok := calibrationBuggyActors[cid.String()]; ok && h.network == tools.CalibrationNetwork {
			return name, nil
		}
		return "", err
	}

	return actorName, nil
}

// Deprecated: Use v2/tools.GetMethodName instead
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

	_, actorName, err := h.GetActorNameFromAddress(msg.To, height, key)
	if err != nil {
		_ = h.metrics.UpdateActorNameErrorMetric(fmt.Sprint(uint64(msg.Method)))
	}

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

// Deprecated: Use v2/tools.GetMethodName instead
// CheckCommonMethods returns the method name for the given message if Send Or Constructor, otherwise returns an empty string
func (h *Helper) CheckCommonMethods(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (string, error) {
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

	return "", nil
}

func (h *Helper) GetEVMSelectorSig(ctx context.Context, selectorID string) (string, error) {
	s, err := h.actorCache.GetEVMSelectorSig(ctx, selectorID)
	if err != nil {
		_ = h.metrics.UpdateGetEvmSelectorSigMetric()
	}
	return s, err
}

func (h *Helper) FilterTxsByActorType(ctx context.Context, txs []*types.Transaction, actorType string, tipsetKey filTypes.TipSetKey) ([]*types.Transaction, error) {
	var result []*types.Transaction
	for _, tx := range txs {
		addrTo, err := address.NewFromString(tx.TxTo)
		if err != nil {
			_ = h.metrics.UpdateParseAddressErrorMetric("to")
			h.logger.Errorf("could not parse address. Err: %s", err)
			continue
		}
		addrFrom, err := address.NewFromString(tx.TxFrom)
		if err != nil {
			_ = h.metrics.UpdateParseAddressErrorMetric("from")
			h.logger.Errorf("could not parse address. Err: %s", err)
			continue
		}

		// #nosec G115
		isType, err := h.isAnyAddressOfType(ctx, []address.Address{addrTo, addrFrom}, int64(tx.Height), tipsetKey, actorType)
		if err != nil {
			h.logger.Errorf("could not get actor type from address. Err: %s", err)
			continue
		}
		if !isType {
			continue
		}

		result = append(result, tx)
	}

	return result, nil
}

func (h *Helper) IsSystemActor(addr address.Address) bool {
	return h.actorCache.IsSystemActor(addr.String())
}
func (h *Helper) IsGenesisActor(addr address.Address) bool {
	return h.actorCache.IsGenesisActor(addr.String())
}

func (h *Helper) IsCronActor(height int64, addr address.Address, tipsetKey filTypes.TipSetKey) bool {
	_, actorName, err := h.GetActorNameFromAddress(addr, height, tipsetKey)
	if err != nil {
		return false
	}
	return strings.Contains(actorName, manifest.CronKey)
}

func (h *Helper) isAnyAddressOfType(_ context.Context, addresses []address.Address, height int64, key filTypes.TipSetKey, actorType string) (bool, error) {
	for _, addr := range addresses {
		if addr == address.Undef {
			continue
		}
		_, actorName, err := h.GetActorNameFromAddress(addr, height, key)
		if err != nil {
			return false, err
		}
		if strings.Contains(actorName, actorType) {
			return true, nil
		}
	}
	return false, nil
}
