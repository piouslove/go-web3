package dto

type TransactionParameters struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Value    string `json:"value"`
	Data     string `json:"data,omitempty"`
}

type TransactionResponse struct {
	Hash             string `json:"hash"`
	Nonce            int    `json:"nonce"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      int64  `json:"blockNumber"`
	TransactionIndex int64  `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	GasPrice         string `json:"gasPrice,omitempty"`
	Gas              string `json:"gas,omitempty"`
	Data             string `json:"data,omitempty"`
}
