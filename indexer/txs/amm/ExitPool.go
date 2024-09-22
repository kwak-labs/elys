package amm

import (
	"encoding/json"
	"fmt"

	"github.com/elys-network/elys/indexer/types"
)

type ExitPool struct {
	Address       string         `json:"Address"`
	PoolId        uint64         `json:"pool_id"`
	MinAmountsOut []MinAmountOut `json:"min_amounts_out"`
	ShareAmountIn string         `json:"share_amount_in"`
	TokenOutDenom string         `json:"token_out_denom"`
	ExitCoins     []ExitCoin     `json:"exit_coins"`
}

type MinAmountOut struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type ExitCoin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

func (m ExitPool) Process(database types.DatabaseManager, transaction types.BaseTransaction) (types.Response, error) {
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
