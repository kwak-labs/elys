package commitments

import (
	"encoding/json"
	"fmt"

	"github.com/elys-network/elys/indexer/types"
)

type MsgUnstake struct {
	Address          string `json:"sender_address"`
	Amount           string `json:"amount"`
	Denom            string `json:"denom"`
	ValidatorAddress string `json:"validator_address"`
}

func (m MsgUnstake) Process(database types.DatabaseManager, transaction types.BaseTransaction) (types.Response, error) {
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
