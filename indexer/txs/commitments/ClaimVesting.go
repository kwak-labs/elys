package commitments

import (
	"encoding/json"
	"fmt"

	"github.com/elys-network/elys/indexer/types"
)

type ClaimVesting struct {
	Address string  `json:"sender_address"`
	Claims  []Claim `json:"claims"`
}

type Claim struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

func (m ClaimVesting) Process(database types.DatabaseManager, transaction types.BaseTransaction) (types.Response, error) {
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
