package tools

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/zondax/golem/pkg/logger"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v11/reward"
	"github.com/filecoin-project/go-state-types/exitcode"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	blocks "github.com/ipfs/go-block-format"
)

const UnknownParserVersion = "unknown"

type Tools struct {
	Logger *logger.Logger
}

func BuildId(input ...string) string {
	h := sha256.New()
	a := make([]byte, 0)
	for _, v := range input {
		a = append(a, []byte(v)...)
	}

	h.Write(a)
	hash := h.Sum(nil)
	id := uuid.NewSHA1(uuid.Nil, hash)
	return id.String()
}

func BuildMessageId(tipsetCid, blockCid, mainMsgCid, messageCid, parentId string) string {
	return BuildId(tipsetCid, blockCid, mainMsgCid, messageCid, parentId)
}

// UniqueSiblingId returns id unchanged the first time it is seen among the
// sub-calls of one parent, and a derived id for each repeat. Two sibling calls
// with identical trace messages (same from, to, value, method, params and gas
// limit, e.g. two equal transfers sent with a fixed gas stipend) otherwise get
// the same id, and FilterDuplicated keeps only one of them, dropping a real
// transfer. The first occurrence keeps its id, so existing ids do not change.
func UniqueSiblingId(seen map[string]int, id string) string {
	n := seen[id]
	seen[id] = n + 1
	if n == 0 {
		return id
	}
	return BuildId(id, "sibling", strconv.Itoa(n))
}

func BuildFeeId(tipsetCid, blockCid, mainMsgCid string) string {
	return BuildId(tipsetCid, blockCid, mainMsgCid, "fee")
}

func BuildTipsetId(tipsetCid string) string {
	h := sha256.New()
	h.Write([]byte(tipsetCid))
	hash := h.Sum(nil)
	id := uuid.NewSHA1(uuid.Nil, hash)
	return id.String()
}

// Deprecated: Use v2/tools.GetBlockCidFromMsgCid instead
func (t *Tools) GetBlockCidFromMsgCid(msgCid, txType string, txMetadata map[string]interface{}, tipset *types.ExtendedTipSet) (string, error) {
	// Default value
	blockCid := ""

	// Process the special cases first were this kind of txs are not explicitly included in a block
	switch txType {
	case parser.MethodAwardBlockReward:
		if txMetadata == nil {
			return blockCid, fmt.Errorf("received tx of type '%s' with nil metadata", txType)
		}
		// Get the miner that received the reward
		params, ok := txMetadata["Params"]
		if !ok {
			t.Logger.Errorf("Could no get paramater 'Params' inside tx '%s'", txType)
			return blockCid, nil
		}

		rewardsParams, ok := params.(reward.AwardBlockRewardParams)
		if !ok {
			t.Logger.Errorf("Could not parse parameters for tx '%s'", txType)
			return blockCid, nil
		}
		// Get the block that this miner mined
		c, err := tipset.GetBlockMinedByMiner(rewardsParams.Miner.String())
		if err != nil {
			return blockCid, err
		}
		return c, nil
	case parser.MethodApplyRewards, parser.MethodUpdatePledgeTotal, parser.MethodCronTick,
		parser.MethodEpochTick, parser.MethodThisEpochReward, parser.MethodConfirmSectorProofsValid,
		parser.MethodActivateDeals, parser.MethodClaimAllocations, parser.MethodBurnExported,
		parser.MethodEnrollCronEvent, parser.MethodOnDeferredCronEvent, parser.MethodUpdateNetworkKPI:
		// These txs are not included in a block
		return blockCid, nil
	}

	blockCids, ok := tipset.BlockMessages[msgCid]
	if !ok {
		return blockCid, errors.New("could not find block hash for message cid")
	}

	if len(blockCids) == 0 {
		return blockCid, errors.New("could not find block hash for message cid. Slice is empty")
	} else {
		blockCid = blockCids[0].Cid
	}

	return blockCid, nil
}

func BuildCidFromMessageTrace(msg filTypes.MessageTrace, parentMsgCid string) (string, error) {
	// Serialize
	buf := new(bytes.Buffer)
	if err := msg.MarshalCBOR(buf); err != nil {
		return "", err
	}

	data := buf.Bytes()
	data = append(data, []byte(parentMsgCid)...)

	// ToStorageBlock
	c, err := abi.CidBuilder.Sum(data)
	if err != nil {
		return "", err
	}

	b, err := blocks.NewBlockWithCid(data, c)
	if err != nil {
		return "", err
	}

	return b.Cid().String(), nil
}

func SetNodeMetadata[T types.HasNodeInfo](data []T, metadata types.BlockMetadata, parserVer string) []T {
	nodeMajorMinorVersion := metadata.NodeMajorMinorVersion
	if nodeMajorMinorVersion == "" {
		nodeMajorMinorVersion = UnknownParserVersion
	}

	nodeFullVersion := metadata.NodeFullVersion
	if nodeFullVersion == "" {
		nodeFullVersion = UnknownParserVersion
	}

	for _, item := range data {
		item.SetNodeMetadata(nodeMajorMinorVersion, nodeFullVersion, parserVer)
	}
	return data

}

func ParseTxMetadata(txMetadata string) (map[string]interface{}, error) {
	var metadata map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal TxMetadata: %v", err)
	}
	return metadata, nil
}

func ParseActorName(actor string) string {
	s := strings.Split(actor, "/")
	if len(s) < 1 {
		return actor
	}
	return s[len(s)-1]
}

func GetBlocksCidByString(str string) []string {
	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")
	return strings.Split(str, ",")
}

func GetExitCodeStatus(actorName string, exitCode exitcode.ExitCode) string {
	code := exitCode.String()
	// common exit codes will be of the format Msg(code)
	// this extracts the Msg e.g Ok(0) -> Ok
	status := strings.Split(code, "(")
	if len(status) == 2 {
		return status[0]
	}
	// checking for actor specific exit codes
	return CheckExitCodeError(actorName, code)
}

// ChildrenExitCode returns the exit code a call passes down to its sub-calls,
// which they record as their status. A call that fails reverts its whole
// subtree (the FVM runs each call in a state transaction that rolls back on a
// non-zero exit), so its own error is passed down; otherwise the code inherited
// from its ancestors (ultimately the top-level message) is passed on unchanged.
func ChildrenExitCode(inherited, own exitcode.ExitCode) exitcode.ExitCode {
	if own.IsError() {
		return own
	}
	return inherited
}

func GetExitCodeOK() string {
	return GetExitCodeStatus("", exitcode.Ok)
}

func FormatTipsetKey(tipsetKey filTypes.TipSetKey) string {
	tipsetKeyStr := tipsetKey.String()
	tipsetKeyStr = strings.ReplaceAll(tipsetKeyStr, "{", "")
	tipsetKeyStr = strings.ReplaceAll(tipsetKeyStr, "}", "")
	return tipsetKeyStr
}
