package indexer

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/elys-network/elys/indexer/txs/amm"
	"github.com/elys-network/elys/indexer/txs/commitments"
	"github.com/elys-network/elys/indexer/types"
)

var txRegistry = make(map[string]reflect.Type)

func init() {
	// Commitments
	RegisterTxType("/elys.commitment.MsgStake", reflect.TypeOf(commitments.MsgStake{}))
	RegisterTxType("/elys.commitment.MsgUnstake", reflect.TypeOf(commitments.MsgUnstake{}))
	RegisterTxType("/elys.commitment.MsgUncommitTokens", reflect.TypeOf(commitments.UncommitTokens{}))
	RegisterTxType("/elys.commitment.MsgCancelVest", reflect.TypeOf(commitments.CancelVest{}))
	RegisterTxType("/elys.commitment.MsgClaimVesting", reflect.TypeOf(commitments.ClaimVesting{}))
	RegisterTxType("/elys.commitment.MsgCommitClaimedRewards", reflect.TypeOf(commitments.CommitClaimedRewards{}))
	RegisterTxType("/elys.commitment.MsgVestLiquid", reflect.TypeOf(commitments.VestLiquid{}))

	// AMM
	RegisterTxType("/elys.amm.MsgCreatePool", reflect.TypeOf(amm.CreatePool{}))
	RegisterTxType("/elys.amm.MsgExitPool", reflect.TypeOf(amm.ExitPool{}))
	RegisterTxType("/elys.amm.MsgFeedMultipleExternalLiquidity", reflect.TypeOf(amm.FeedMultipleExternalLiquidity{}))
	RegisterTxType("/elys.amm.MsgJoinPool", reflect.TypeOf(amm.JoinPool{}))
	RegisterTxType("/elys.amm.MsgSwapByDenom", reflect.TypeOf(amm.SwapByDenom{}))
	RegisterTxType("/elys.amm.MsgSwapExactAmountIn", reflect.TypeOf(amm.SwapExactAmountIn{}))
	RegisterTxType("/elys.amm.MsgSwapExactAmountOut", reflect.TypeOf(amm.SwapExactAmountOut{}))

}

func RegisterTxType(txType string, dataType reflect.Type) {
	txRegistry[txType] = dataType
}

func ParseTransaction(tx types.GenericTransaction) (string, types.Processor, error) {
	txType := tx.BaseTransaction.TxType
	dataType, ok := txRegistry[txType]
	if !ok {
		return "", nil, fmt.Errorf("unknown transaction type: %s", txType)
	}

	dataValue := reflect.New(dataType).Interface()
	dataBytes, err := json.Marshal(tx.Data)
	if err != nil {
		return "", nil, fmt.Errorf("error marshaling data: %w", err)
	}

	err = json.Unmarshal(dataBytes, dataValue)
	if err != nil {
		return "", nil, fmt.Errorf("error unmarshaling to %s: %w", dataType.Name(), err)
	}

	processor, ok := reflect.ValueOf(dataValue).Elem().Interface().(types.Processor)
	if !ok {
		return "", nil, fmt.Errorf("type %s does not implement Processor", dataType.Name())
	}

	return txType, processor, nil
}
