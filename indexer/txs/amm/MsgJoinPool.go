package amm

import (
	"encoding/json"
	"fmt"

	"github.com/elys-network/elys/indexer/types"
)

type JoinPool struct {
	Address        string        `json:"Address"`
	PoolId         uint64        `json:"pool_id"`
	MaxAmountsIn   []MaxAmountIn `json:"min_amounts_out"`
	ShareAmountOut string        `json:"share_amount_out"`
	TokensIn       []TokenIn     `json:"tokens_in"`
	SharesOut      string        `json:"shares_out"`
}

type MaxAmountIn struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type TokenIn struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

func (m JoinPool) Process(database types.DatabaseManager, transaction types.BaseTransaction) (types.Response, error) {
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
