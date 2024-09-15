package indexer

import (
	"encoding/json"
	"fmt"

	indexer "github.com/elys-network/elys/indexer"
)

type CommitmentMsgStake struct {
	Sender_address    string `json:"sender_address"`
	Amount            string `json:"amount"`
	Denom             string `json:"denom"`
	Validator_address string `json:"validator_address"`
}

type mergedTransaction struct {
	BaseTransaction indexer.BaseTransaction
	Data            CommitmentMsgStake
}

func (m CommitmentMsgStake) Process(database *indexer.LMDBManager, transaction indexer.BaseTransaction) (indexer.Response, error) {
	mergedData := mergedTransaction{
		BaseTransaction: transaction,
		Data:            m,
	}

	json, err := json.Marshal(mergedData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json))

	err = database.ProcessNewTx(mergedData, transaction.Author)
	if err != nil {
		fmt.Println(err)
		return indexer.Response{}, err
	}

	fmt.Println("Successfully Stored")

	return indexer.Response{}, nil
}
