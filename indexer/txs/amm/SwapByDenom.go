package amm

import (
	"encoding/json"
	"fmt"

	"github.com/elys-network/elys/indexer/types"
)

type SwapByDenom struct {
	Address     string               `json:"address"`
	Recipient   string               `json:"recipient"`
	TokenIn     Token                `json:"token_in"`
	TokenOutMin Token                `json:"token_out_min"`
	MaxAmount   Token                `json:"max_amount"`
	Discount    string               `json:"discount"`
	SwapFee     string               `json:"swap_fee"`
	SpotPrice   string               `json:"spot_price"`
	InRoute     []SwapAmountInRoute  `json:"in_route"`
	OutRoute    []SwapAmountOutRoute `json:"out_route"`
	TokenOut    Token                `json:"token_out"`
}

type SwapAmountInRoute struct {
	PoolId        uint64 `json:"pool_id"`
	TokenOutDenom string `json:"token_out_denom"`
}
type SwapAmountOutRoute struct {
	PoolId       uint64 `json:"pool_id"`
	TokenInDenom string `json:"token_in_denom"`
}

type Token struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

func (m SwapByDenom) Process(database types.DatabaseManager, transaction types.BaseTransaction) (types.Response, error) {
	mergedData := types.GenericTransaction{
		BaseTransaction: transaction,
		Data:            m,
	}

	jsonData, err := json.Marshal(mergedData)
	if err != nil {
		return types.Response{}, fmt.Errorf("error marshaling data: %w", err)
	}

	fmt.Println(string(jsonData))

	err = database.ProcessNewTx(mergedData, transaction.Author)
	if err != nil {
		return types.Response{}, fmt.Errorf("error processing transaction: %w", err)
	}

	fmt.Println("Successfully Stored")

	return types.Response{}, nil
}
