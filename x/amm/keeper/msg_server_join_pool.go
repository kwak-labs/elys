package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/amm/types"

	indexer "github.com/elys-network/elys/indexer"
	indexerAmmTypes "github.com/elys-network/elys/indexer/txs/amm"
)

// JoinPool routes `JoinPoolNoSwap` where we do an abstract calculation on needed lp liquidity coins to get the designated
// amount of shares for the pool. (This is done by taking the number of shares we want and then using the total number of shares
// to get the ratio of the pool it accounts for. Using this ratio, we iterate over all pool assets to get the number of tokens we need
// to get the specified number of shares).
// Using the number of tokens needed to actually join the pool, we do a basic sanity check on whether the token does not exceed
// `TokenInMaxs`. Then we hit the actual implementation of `JoinPool` defined by each pool model.
// `JoinPool` takes in the tokensIn calculated above as the parameter rather than using the number of shares provided in the msg.
// This can result in negotiable difference between the number of shares provided within the msg
// and the actual number of share amount resulted from joining pool.
// Internal logic flow for each pool model is as follows:
// Balancer: TokensInMaxs provided as the argument must either contain no tokens or containing all assets in the pool.
// * For the case of a not containing tokens, we simply perform calculation of sharesOut and needed amount of tokens for joining the pool
func (k msgServer) JoinPool(goCtx context.Context, msg *types.MsgJoinPool) (*types.MsgJoinPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	neededLp, sharesOut, err := k.Keeper.JoinPoolNoSwap(ctx, sender, msg.PoolId, msg.ShareAmountOut, msg.MaxAmountsIn)
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

	var IndexerMaxAmountsIn []indexerAmmTypes.MaxAmountIn
	for _, AmountIn := range msg.MaxAmountsIn {
		IndexerMaxAmountsIn = append(IndexerMaxAmountsIn, indexerAmmTypes.MaxAmountIn{
			Denom:  AmountIn.Denom,
			Amount: AmountIn.Amount.String(),
		})
	}

	var IndexerTokensIn []indexerAmmTypes.TokenIn
	for _, TokenIn := range neededLp {
		IndexerTokensIn = append(IndexerTokensIn, indexerAmmTypes.TokenIn{
			Amount: TokenIn.Amount.String(),
			Denom:  TokenIn.Denom,
		})
	}

	indexer.QueueTransaction(ctx, indexerAmmTypes.JoinPool{
		Address:        sender.String(),
		PoolId:         msg.PoolId,
		MaxAmountsIn:   IndexerMaxAmountsIn,
		ShareAmountOut: msg.ShareAmountOut.String(),
		TokensIn:       IndexerTokensIn,
		SharesOut:      sharesOut.String(),
	}, []string{""})

	return &types.MsgJoinPoolResponse{
		ShareAmountOut: sharesOut,
		TokenIn:        neededLp,
	}, nil
}
