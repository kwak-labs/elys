package amm

import (
	"encoding/json"
	"fmt"

	"github.com/elys-network/elys/indexer/types"
)

type FeedMultipleExternalLiquidity struct {
	Address   string              `json:"address"`
	Liquidity []ExternalLiquidity `json:"liquidity"`
}

type ExternalLiquidity struct {
	PoolId          uint64             `json:"pool_id"`
	AmountDepthInfo []AssetAmountDepth `json:"amount_depth_info"`
}

type AssetAmountDepth struct {
	Asset  string `json:"asset"`
	Amount string `json:"amount"`
	Depth  string `json:"depth"`
}

func (m FeedMultipleExternalLiquidity) Process(database types.DatabaseManager, transaction types.BaseTransaction) (types.Response, error) {
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
