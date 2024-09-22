package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elys-network/elys/x/commitment/types"
	ptypes "github.com/elys-network/elys/x/parameter/types"

	indexer "github.com/elys-network/elys/indexer"
	indexerCommitmentsTypes "github.com/elys-network/elys/indexer/txs/commitments"
)

// ClaimVesting claims already vested amount
func (k msgServer) ClaimVesting(goCtx context.Context, msg *types.MsgClaimVesting) (*types.MsgClaimVestingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the Commitments for the sender
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	commitments := k.GetCommitments(ctx, sender)

	newClaims := sdk.Coins{}
	for i, vesting := range commitments.VestingTokens {
		vestedSoFar := vesting.VestedSoFar(ctx)
		newClaim := vestedSoFar.Sub(vesting.ClaimedAmount)
		newClaims = newClaims.Add(sdk.NewCoin(vesting.Denom, newClaim))
		commitments.VestingTokens[i].ClaimedAmount = vestedSoFar
	}

	if newClaims.IsAllPositive() {
		// mint coins if vesting token is ELYS
		if newClaims.AmountOf(ptypes.Elys).IsPositive() {
			elysCoins := sdk.Coins{sdk.NewCoin(ptypes.Elys, newClaims.AmountOf(ptypes.Elys))}
			err := k.bankKeeper.MintCoins(ctx, types.ModuleName, elysCoins)
			if err != nil {
				return nil, err
			}
		}

		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, newClaims)
		if err != nil {
			return nil, err
		}
	}

	k.SetCommitments(ctx, commitments)

	// Emit blockchain event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeClaimVesting,
			sdk.NewAttribute(types.AttributeCreator, msg.Sender),
			sdk.NewAttribute(types.AttributeAmount, newClaims.String()),
		),
	)

	var indexerClaims []indexerCommitmentsTypes.Claim

	for _, claim := range newClaims {
		indexerClaims = append(indexerClaims, indexerCommitmentsTypes.Claim{
			Denom:  claim.Denom,
			Amount: claim.Amount.String(),
		})
	}

	indexer.QueueTransaction(ctx, indexerCommitmentsTypes.ClaimVesting{
		Address: sender.String(),
		Claims:  indexerClaims,
	}, []string{})

	return &types.MsgClaimVestingResponse{}, nil
}
