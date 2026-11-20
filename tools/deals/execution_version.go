package deals

import (
	"github.com/zondax/fil-parser/tools"
)

// executedNodeVersion is the network version whose actors executed the messages (and cron) of the
// tipset at height. Lotus runs a network upgrade's state migration while computing the first
// tipset AFTER the upgrade height, before its messages: ApplyBlocks calls HandleStateForks for
// i in [parentEpoch, epoch) (chain/consensus/compute_state.go), and GetNetworkVersion(h) returns
// the old version for h <= upgrade height (chain/stmgr/stmgr.go, atOrBelow = upgrade.Height). So
// the tipset AT an upgrade height still executes with the old actors; e.g. StateNetworkVersion is
// 21 at mainnet 3,855,360 (Dragon / NV22) and 22 from 3,855,361.
//
// tools.VersionFromHeight labels the upgrade height itself with the new version, so the version
// that executed height h is the one tools.VersionFromHeight gives for h-1. This is used where the
// deals code depends on the actors version and the decoded params cannot tell; the actor parser
// still decodes the upgrade tipset with the new version's types.
func executedNodeVersion(network string, height uint64) uint {
	if height == 0 {
		return tools.VersionFromHeight(network, 0).NodeVersion()
	}
	//#nosec G115
	return tools.VersionFromHeight(network, int64(height)-1).NodeVersion()
}
