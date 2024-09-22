package commitments

import (
	"encoding/json"
	"fmt"

	"github.com/elys-network/elys/indexer/types"
)

type VestLiquid struct {
	Address              string `json:"sender_address"`
	Amount               string `json:"amount"`
	Denom                string `json:"denom"`
	StartBlock           int64  `json:"start_block"`
	NumBlocks            int64  `json:"numb_blocks"`
	VestStartedTimestamp int64  `json:"vested_started_timestamp"`
}

func (m VestLiquid) Process(database types.DatabaseManager, transaction types.BaseTransaction) (types.Response, error) {
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
