package types

import (
	"time"
)

type FeeDetail struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}

type BaseTransaction struct {
	BlockTime         time.Time   `json:"block_time"`
	Author            string      `json:"author"`
	IncludedAddresses []string    `json:"included_addresses"`
	BlockHeight       int64       `json:"block_height"`
	TxHash            string      `json:"tx_hash"`
	TxType            string      `json:"tx_type"`
	Fees              []FeeDetail `json:"fees"`
	GasLimit          string      `json:"gas_limit"`
	GasUsed           string      `json:"gas_used"`
	Memo              string      `json:"memo"`
	Status            string      `json:"status"`
}

type GenericTransaction struct {
	BaseTransaction BaseTransaction `json:"BaseTransaction"`
	Data            interface{}     `json:"Data"`
}

type Response struct{}

type DatabaseManager interface {
	ProcessNewTx(GenericTransaction, string) error
}

type Processor interface {
	Process(DatabaseManager, BaseTransaction) (Response, error)
}
