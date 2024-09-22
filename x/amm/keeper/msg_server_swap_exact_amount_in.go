package keeper

import (
	"context"

	// "cosmossdk.io/math"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/amm/types"

	indexer "github.com/elys-network/elys/indexer"
	indexerAmmTypes "github.com/elys-network/elys/indexer/txs/amm"
)

func (k msgServer) SwapExactAmountIn(goCtx context.Context, msg *types.MsgSwapExactAmountIn) (*types.MsgSwapExactAmountInResponse, error) {
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
	tokenOutAmount, swapFee, discount, err := k.RouteExactAmountIn(cacheCtx, sender, recipient, msg.Routes, msg.TokenIn, math.Int(msg.TokenOutMinAmount), msg.Discount)
	if err != nil {
		return nil, err
	}

	lastSwapIndex := k.GetLastSwapRequestIndex(ctx)
	k.SetSwapExactAmountInRequests(ctx, msg, lastSwapIndex+1)
	k.SetLastSwapRequestIndex(ctx, lastSwapIndex+1)

	// Swap event is handled elsewhere
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	})

	var IndexerInRoutes []indexerAmmTypes.SwapAmountInRoute
	for _, OutRoute := range msg.Routes {
		IndexerInRoutes = append(IndexerInRoutes, indexerAmmTypes.SwapAmountInRoute{
			PoolId:        OutRoute.PoolId,
			TokenOutDenom: OutRoute.TokenOutDenom,
		})
	}

	lastRoute := msg.Routes[len(msg.Routes)-1]

	indexer.QueueTransaction(ctx, indexerAmmTypes.SwapExactAmountIn{
		Address:   msg.Sender,
		Recipient: msg.Recipient,
		TokenIn: indexerAmmTypes.Token{
			Amount: msg.TokenIn.Amount.String(),
			Denom:  msg.TokenIn.Denom,
		},
		TokenOutMin: msg.TokenOutMinAmount.String(),
		Discount:    discount.String(),
		SwapFee:     swapFee.String(),
		InRoute:     IndexerInRoutes,
		TokenOut: indexerAmmTypes.Token{
			Denom:  lastRoute.TokenOutDenom,
			Amount: tokenOutAmount.String(),
		},
	}, []string{msg.Recipient})

	return &types.MsgSwapExactAmountInResponse{
		TokenOutAmount: tokenOutAmount,
		SwapFee:        swapFee,
		Discount:       discount,
		Recipient:      recipient.String(),
	}, nil
}
