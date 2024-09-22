package amm

import (
	"encoding/json"
	"fmt"

	"github.com/elys-network/elys/indexer/types"
)

type SwapExactAmountIn struct {
	Address     string              `json:"address"`
	Recipient   string              `json:"recipient"`
	TokenIn     Token               `json:"token_in"`
	TokenOutMin string              `json:"token_out_min"`
	Discount    string              `json:"discount"`
	SwapFee     string              `json:"swap_fee"`
	InRoute     []SwapAmountInRoute `json:"in_route"`
	TokenOut    Token               `json:"token_out"`
}

func (m SwapExactAmountIn) Process(database types.DatabaseManager, transaction types.BaseTransaction) (types.Response, error) {
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
