package dto

type TransactionParameters struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Value    string `json:"value"`
	Data     string `json:"data,omitempty"`
}
