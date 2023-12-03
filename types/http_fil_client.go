package types

type Block struct {
	Miner         string `json:"Miner"`
	ParentBaseFee string `json:"ParentBaseFee"`
}

type Result struct {
	Blocks []Block `json:"Blocks"`
}

type GetTipsetByHeightResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
	ID      int    `json:"id"`
}

type ApiRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}
