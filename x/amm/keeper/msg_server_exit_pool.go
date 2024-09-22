package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/amm/types"

	indexer "github.com/elys-network/elys/indexer"
	indexerAmmTypes "github.com/elys-network/elys/indexer/txs/amm"
)

func (k msgServer) ExitPool(goCtx context.Context, msg *types.MsgExitPool) (*types.MsgExitPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	exitCoins, _, err := k.Keeper.ExitPool(ctx, sender, msg.PoolId, msg.ShareAmountIn, msg.MinAmountsOut, msg.TokenOutDenom)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	})

	var IndexerMinAmountsOut []indexerAmmTypes.MinAmountOut
	for _, amount := range msg.MinAmountsOut {
		IndexerMinAmountsOut = append(IndexerMinAmountsOut, indexerAmmTypes.MinAmountOut{
			Denom:  amount.Denom,
			Amount: amount.Amount.String(),
		})
	}

	var IndexerCoinsOut []indexerAmmTypes.ExitCoin
	for _, Coin := range exitCoins {
		IndexerCoinsOut = append(IndexerCoinsOut, indexerAmmTypes.ExitCoin{
			Denom:  Coin.Denom,
			Amount: Coin.Amount.String(),
		})
	}

	indexer.QueueTransaction(ctx, indexerAmmTypes.ExitPool{
		Address:       sender.String(),
		PoolId:        msg.PoolId,
		MinAmountsOut: IndexerMinAmountsOut,
		ShareAmountIn: msg.ShareAmountIn.String(),
		ExitCoins:     IndexerCoinsOut,
	}, []string{})

	return &types.MsgExitPoolResponse{
		TokenOut: exitCoins,
	}, nil
}
