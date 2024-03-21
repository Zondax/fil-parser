package tools

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v11/reward"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"go.uber.org/zap"
)

const defaultCodeCid = "bafy2bzaceab3xcn7qkcuj5oyifa6dn3ihke55bdmerphef4r6aorjdhk3uriq"

type Tools struct {
	Logger *zap.Logger
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
			t.Logger.Sugar().Errorf("Could no get paramater 'Params' inside tx '%s'", txType)
			return blockCid, nil
		}

		rewardsParams, ok := params.(reward.AwardBlockRewardParams)
		if !ok {
			t.Logger.Sugar().Errorf("Could not parse parameters for tx '%s'", txType)
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
		return blockCid, fmt.Errorf("could not find block hash for message cid '%s'", msgCid)
	}

	if len(blockCids) == 0 {
		return blockCid, fmt.Errorf("could not find block hash for message cid '%s'. Slice is empty", msgCid)
	} else {
		blockCid = blockCids[0].Cid
	}

	return blockCid, nil
}

func BuildCidFromMessageTrace(actor filTypes.ActorTrace, parentMsgCid string) (string, error) {

	// v1.23 has not CodeCid field. It was introduced on v1.24.
	// In order to make the parser v2 backward compatible to v1.23 files (avoid to create a new parser)
	// we set a default value for this field as constant.
	//
	// v1.26 removed the CodeCid from the trace.Msg object and added a new InvokedActor field in the Execution Trace.
	// The InvokedActor State containes the the actor's code CID
	if !actor.State.Code.Defined() {
		defaultCodeCid, err := cid.Parse(defaultCodeCid)
		if err != nil {
			return "", err
		}

		actor.State.Code = defaultCodeCid
	}

	// Serialize
	buf := new(bytes.Buffer)

	if err := actor.State.MarshalCBOR(buf); err != nil {
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

func SetNodeMetadataOnTxs(txs []*types.Transaction, metadata types.BlockMetadata, parserVer string) []*types.Transaction {
	nodeMajorMinorVersion := metadata.NodeMajorMinorVersion
	if nodeMajorMinorVersion == "" {
		nodeMajorMinorVersion = "unknown"
	}

	nodeFullVersion := metadata.NodeFullVersion
	if nodeFullVersion == "" {
		nodeFullVersion = "unknown"
	}

	for _, tx := range txs {
		tx.NodeMajorMinorVersion = nodeMajorMinorVersion
		tx.NodeFullVersion = nodeFullVersion
		tx.ParserVersion = parserVer
	}

	return txs
}
