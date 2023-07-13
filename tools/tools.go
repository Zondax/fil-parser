package tools

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v11/reward"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	blocks "github.com/ipfs/go-block-format"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

func BuildMessageId(tipsetCid, blockCid, messageCid string) string {
	h := sha256.New()
	h.Write([]byte(tipsetCid + blockCid + messageCid))
	hash := h.Sum(nil)
	id := uuid.NewSHA1(uuid.Nil, hash)
	return id.String()
}

func BuildTipsetId(tipsetCid string) string {
	h := sha256.New()
	h.Write([]byte(tipsetCid))
	hash := h.Sum(nil)
	id := uuid.NewSHA1(uuid.Nil, hash)
	return id.String()
}

func BuildFeeId(tipsetCid, blockCid, messageCid string) string {
	h := sha256.New()
	h.Write([]byte(tipsetCid + blockCid + messageCid + "fee"))
	hash := h.Sum(nil)
	id := uuid.NewSHA1(uuid.Nil, hash)
	return id.String()
}

func GetBlockCidFromMsgCid(msgCid, txType string, txMetadata map[string]interface{}, tipset *types.ExtendedTipSet) (string, error) {
	// Default value
	blockCid := tipset.GetCidString()

	// Process the special cases first were this kind of txs are not explicitly included in a block
	switch txType {
	case parser.MethodAwardBlockReward:
		if txMetadata == nil {
			return blockCid, fmt.Errorf("received tx of type '%s' with nil metadata", txType)
		}
		// Get the miner that received the reward
		params, ok := txMetadata["Params"]
		if !ok {
			zap.S().Errorf("Could no get paramater 'Params' inside tx '%s'", txType)
			return blockCid, nil
		}

		rewardsParams, ok := params.(reward.AwardBlockRewardParams)
		if !ok {
			zap.S().Errorf("Could not parse parameters for tx '%s'", txType)
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

func BuildCidFromMessageTrace(msg *filTypes.MessageTrace) (string, error) {
	// Serialize
	buf := new(bytes.Buffer)
	if err := msg.MarshalCBOR(buf); err != nil {
		return "", err
	}

	data := buf.Bytes()

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
