package client

import "database/sql"

type Receipts struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		BlockHash         string      `json:"blockHash"`
		BlockNumber       string      `json:"blockNumber"`
		ContractAddress   interface{} `json:"contractAddress"`
		CumulativeGasUsed string      `json:"cumulativeGasUsed"`
		EffectiveGasPrice string      `json:"effectiveGasPrice"`
		From              string      `json:"from"`
		GasUsed           string      `json:"gasUsed"`
		Logs              []struct {
			Address          string   `json:"address"`
			Topics           []string `json:"topics"`
			Data             string   `json:"data"`
			BlockNumber      string   `json:"blockNumber"`
			TransactionHash  string   `json:"transactionHash"`
			TransactionIndex string   `json:"transactionIndex"`
			BlockHash        string   `json:"blockHash"`
			LogIndex         string   `json:"logIndex"`
			Removed          bool     `json:"removed"`
		} `json:"logs"`
		LogsBloom        string `json:"logsBloom"`
		Status           string `json:"status"`
		To               string `json:"to"`
		TransactionHash  string `json:"transactionHash"`
		TransactionIndex string `json:"transactionIndex"`
		Type             string `json:"type"`
	} `json:"result"`
}

type TransactionLogs struct {
	Topic0       sql.NullString
	Topic1       sql.NullString
	Topic2       sql.NullString
	Topic3       sql.NullString
	TokenAddress string
	Data         string `json:"data"`
}

type Result struct {
	Type         string
	TransferData TransferData
}

type TransferData struct {
	Token    string `json:"token"`
	Sender   string `json:"sender,omitempty"`
	Receiver string `json:"receiver,omitempty"`
	NftId    string `json:"nft_id,omitempty"`
	Amount   string `json:"amount,omitempty"`
}
