package types

type BasicBlockData struct {
	// Height contains the block height
	Height uint64 `json:"height"`
	// TipsetHash contains the tipset hash
	TipsetCid string `json:"tipset_cid"`
}

type TxBasicBlockData struct {
	BasicBlockData
	// Block Cid
	BlockCid string `json:"block_cid"`
}

type TipsetBasicBlockData struct {
	BasicBlockData
	// Blocks Cid
	BlocksCid []string `json:"blocks_cid"`

	NodeInfo
}

type BlockMetadata struct {
	NodeInfo
}
