package types

import "time"

type BasicBlockData struct {
	// Height contains the block height
	Height uint64 `json:"height" gorm:"index:idx_height"`
	// TipsetHash contains the tipset hash
	TipsetCid string `json:"tipset_cid" gorm:"index:idx_tipset_cid"`
}

type TxBasicBlockData struct {
	BasicBlockData
	// Block Cid
	BlockCid string `json:"block_cid" gorm:"index:idx_blocks_cid"`
}

type TipsetBasicBlockData struct {
	BasicBlockData
	// Blocks Cid
	BlocksCid []string `json:"blocks_cid" gorm:"type:Array(String);index:idx_blocks_cid"`

	NodeInfo
}

type BlockMetadata struct {
	NodeInfo
}

type BlockInfo struct {
	BlockCid string
	Miner    string
}

type BlocksTimestamp struct {
	TipsetBasicBlockData
	// Id is the unique identifier for this tipset
	Id string `json:"id"`
	// ParentTipsetCid
	ParentTipsetCid string `json:"parent_tipset_cid"`
	// Timestamp is the timestamp of the block
	Timestamp time.Time `json:"tipset_timestamp"`
	// BaseFee is the base fee set for the tipset measured in attoFIL/gas unit
	BaseFee uint64 `json:"base_fee"`
	// BlocksInfo contains basic info of all blocks inside this tipset
	BlocksInfo string `json:"blocks_info"`
}
