package types

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
}

type BlockMetadata struct {
	NodeInfo
}
