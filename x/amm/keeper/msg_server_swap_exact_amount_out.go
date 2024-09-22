package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/amm/types"

	indexer "github.com/elys-network/elys/indexer"
	indexerAmmTypes "github.com/elys-network/elys/indexer/txs/amm"
)

func (k msgServer) SwapExactAmountOut(goCtx context.Context, msg *types.MsgSwapExactAmountOut) (*types.MsgSwapExactAmountOutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		recipient = sender
	}
	// Try executing the tx on cached context environment, to filter invalid transactions out
	cacheCtx, _ := ctx.CacheContext()
	tokenInAmount, swapFee, discount, err := k.RouteExactAmountOut(cacheCtx, sender, recipient, msg.Routes, msg.TokenInMaxAmount, msg.TokenOut, msg.Discount)
	if err != nil {
		return nil, err
	}

	lastSwapIndex := k.GetLastSwapRequestIndex(ctx)
	k.SetSwapExactAmountOutRequests(ctx, msg, lastSwapIndex+1)
	k.SetLastSwapRequestIndex(ctx, lastSwapIndex+1)

	// Swap event is handled elsewhere
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	})

	var IndexerOutRoutes []indexerAmmTypes.SwapAmountOutRoute
	for _, OutRoute := range msg.Routes {
		IndexerOutRoutes = append(IndexerOutRoutes, indexerAmmTypes.SwapAmountOutRoute{
			PoolId:       OutRoute.PoolId,
			TokenInDenom: OutRoute.TokenInDenom,
		})
	}

	indexer.QueueTransaction(ctx, indexerAmmTypes.SwapExactAmountOut{
		Address:   msg.Sender,
		Recipient: msg.Recipient,
		TokenOut: indexerAmmTypes.Token{
			Amount: msg.TokenOut.Amount.String(),
			Denom:  msg.TokenOut.Denom,
		},
		TokenInMaxAmount: msg.TokenInMaxAmount.String(),
		Discount:         discount.String(),
		SwapFee:          swapFee.String(),
		OutRoute:         IndexerOutRoutes,
		TokenIn: indexerAmmTypes.Token{
			Denom:  msg.Routes[0].TokenInDenom,
			Amount: tokenInAmount.String(),
		},
	}, []string{msg.Recipient})

	return &types.MsgSwapExactAmountOutResponse{
		TokenInAmount: tokenInAmount,
		SwapFee:       swapFee,
		Discount:      discount,
		Recipient:     recipient.String(),
	}, nil
}
