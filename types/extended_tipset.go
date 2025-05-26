package types

import (
	"encoding/json"
	"fmt"

	lotusChainTypes "github.com/filecoin-project/lotus/chain/types"
)

type LightBlockHeader struct {
	Cid        string
	BlockMiner string
}

// FIXME LightBlockHeader shouldn't be a slice
type BlockMessages map[string][]LightBlockHeader // map[MessageCid][]LightBlockHeader

type ExtendedTipSet struct {
	lotusChainTypes.TipSet
	BlockMessages
}

func (e *ExtendedTipSet) GetCidString() string {
	cid, _ := e.Key().Cid()
	return cid.String()
}
func (e *ExtendedTipSet) GetParentCidString() string {
	cid, _ := e.Parents().Cid()
	return cid.String()
}

func (e *ExtendedTipSet) GetBlockMiner(blockCid string) (string, error) {
	blockMessages := e.Blocks()
	for _, block := range blockMessages {
		if block.Cid().String() == blockCid {
			return block.Miner.String(), nil
		}
	}
	return "", fmt.Errorf("could not find miner that mined block '%s'", blockCid)
}

func (e *ExtendedTipSet) GetBlockMinedByMiner(minerAddress string) (string, error) {
	blocks := e.Blocks()
	for _, block := range blocks {
		if block.Miner.String() == minerAddress {
			return block.Cid().String(), nil
		}
	}
	// fallback to checking messages
	for _, block := range e.BlockMessages {
		for _, message := range block {
			if message.BlockMiner == minerAddress {
				return message.Cid, nil
			}
		}
	}
	return "", fmt.Errorf("could not find block mined by miner '%s'", minerAddress)
}

func (e *ExtendedTipSet) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(&struct {
		lotusChainTypes.ExpTipSet
		BlockMessages
	}{
		ExpTipSet: lotusChainTypes.ExpTipSet{
			Cids:   e.TipSet.Cids(),
			Blocks: e.TipSet.Blocks(),
			Height: e.TipSet.Height(),
		},
		BlockMessages: e.BlockMessages,
	})

	return data, err
}

func (e *ExtendedTipSet) UnmarshalJSON(data []byte) error {
	auxTipset := &struct {
		lotusChainTypes.TipSet
	}{}

	if err := json.Unmarshal(data, &auxTipset); err != nil {
		// try other way
		auxTipset := &struct {
			Tipset lotusChainTypes.TipSet
		}{}

		if err := json.Unmarshal(data, &auxTipset); err != nil {
			return err
		}
		e.TipSet = auxTipset.Tipset
	} else {
		e.TipSet = auxTipset.TipSet
	}

	auxMessages := &struct {
		BlockMessages
	}{}
	if err := json.Unmarshal(data, &auxMessages); err != nil {
		return err
	}

	e.BlockMessages = auxMessages.BlockMessages
	return nil
}
