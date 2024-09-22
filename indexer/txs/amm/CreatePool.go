package amm

import (
	"encoding/json"
	"fmt"

	"github.com/elys-network/elys/indexer/types"
)

type CreatePool struct {
	Address    string      `json:"sender_address"`
	PoolId     uint64      `json:"pool_id"`
	PoolParams PoolParams  `json:"pool_params"`
	PoolAssets []PoolAsset `json:"pool_asset"`
}

type PoolParams struct {
	SwapFee                     string `json:"swap_fee"`
	ExitFee                     string `json:"exit_fee"`
	UseOracle                   bool   `json:"use_oracle"`
	WeightBreakingFeeMultiplier string `json:"weight_breaking_fee_multiplier"`
	WeightBreakingFeeExponent   string `json:"weight_breaking_fee_exponent"`
	ExternalLiquidityRatio      string `json:"external_liquidity_ratio"`
	WeightRecoveryFeePortion    string `json:"weight_recovery_fee_portion"`
	ThresholdWeightDifference   string `json:"threshold_weight_difference"`
	FeeDenom                    string `json:"fee_denom"`
}

type PoolAsset struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
	Weight string `json:"weight"`
}

func (m CreatePool) Process(database types.DatabaseManager, transaction types.BaseTransaction) (types.Response, error) {
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
