package amm

import (
	"encoding/json"
	"fmt"

	"github.com/elys-network/elys/indexer/types"
)

type SwapExactAmountOut struct {
	Address          string               `json:"address"`
	Recipient        string               `json:"recipient"`
	TokenOut         Token                `json:"token_out"`
	TokenInMaxAmount string               `json:"token_in_max_amount"`
	Discount         string               `json:"discount"`
	SwapFee          string               `json:"swap_fee"`
	OutRoute         []SwapAmountOutRoute `json:"out_route"`
	TokenIn          Token                `json:"token_in"`
}

func (m SwapExactAmountOut) Process(database types.DatabaseManager, transaction types.BaseTransaction) (types.Response, error) {
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
