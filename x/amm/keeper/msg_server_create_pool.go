package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/amm/types"
	ptypes "github.com/elys-network/elys/x/parameter/types"

	indexer "github.com/elys-network/elys/indexer"
	indexerAmmTypes "github.com/elys-network/elys/indexer/txs/amm"
)

// CreatePool attempts to create a pool returning the newly created pool ID or an error upon failure.
// The pool creation fee is used to fund the community pool.
// It will create a dedicated module account for the pool and sends the initial liquidity to the created module account.
func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Pay pool creation fee
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	params := k.GetParams(ctx)

	if !params.PoolCreationFee.IsNil() && params.PoolCreationFee.IsPositive() {
		feeCoins := sdk.Coins{sdk.NewCoin(ptypes.Elys, params.PoolCreationFee)}
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, feeCoins); err != nil {
			return nil, err
		}
	}

	poolId, err := k.Keeper.CreatePool(ctx, msg)
	if err != nil {
		return &types.MsgCreatePoolResponse{}, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeEvtPoolCreated,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(poolId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.GetSigners()[0].String()),
		),
	})

	var IndexerPoolAssets []indexerAmmTypes.PoolAsset
	for _, asset := range msg.PoolAssets {
		IndexerPoolAssets = append(IndexerPoolAssets, indexerAmmTypes.PoolAsset{
			Amount: asset.Token.Amount.String(),
			Denom:  asset.Token.Denom,
			Weight: asset.Weight.String(),
		})
	}

	indexer.QueueTransaction(ctx, indexerAmmTypes.CreatePool{
		Address: sender.String(),
		PoolId:  poolId,
		PoolParams: indexerAmmTypes.PoolParams{
			SwapFee:                     msg.PoolParams.SwapFee.String(),
			ExitFee:                     msg.PoolParams.ExitFee.String(),
			UseOracle:                   msg.PoolParams.UseOracle,
			WeightBreakingFeeMultiplier: msg.PoolParams.WeightBreakingFeeMultiplier.String(),
			WeightBreakingFeeExponent:   msg.PoolParams.WeightBreakingFeeExponent.String(),
			ThresholdWeightDifference:   msg.GetPoolParams().ThresholdWeightDifference.String(),
			FeeDenom:                    msg.PoolParams.FeeDenom,
		},
		PoolAssets: IndexerPoolAssets,
	}, []string{})

	return &types.MsgCreatePoolResponse{
		PoolID: poolId,
	}, nil
}
