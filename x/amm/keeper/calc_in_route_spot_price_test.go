package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/elys-network/elys/x/amm/types"
	ptypes "github.com/elys-network/elys/x/parameter/types"
)

func (suite *KeeperTestSuite) TestCalcInRouteSpotPrice() {
	poolInitBalance := sdk.Coins{sdk.NewInt64Coin(ptypes.Elys, 1000000), sdk.NewInt64Coin(ptypes.BaseCurrency, 1000000)}
	pool2InitBalance := sdk.Coins{sdk.NewInt64Coin("uusda", 1000000), sdk.NewInt64Coin(ptypes.BaseCurrency, 1000000)}
	senderInitBalance := sdk.Coins{sdk.NewInt64Coin(ptypes.Elys, 1000000), sdk.NewInt64Coin(ptypes.BaseCurrency, 1000000)}

	suite.SetupTest()
	suite.SetupStableCoinPrices()

	// bootstrap accounts
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	poolAddr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	pool2Addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	treasuryAddr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	treasury2Addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())

	// bootstrap balances
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, senderInitBalance)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, sender, senderInitBalance)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, poolInitBalance)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, poolAddr, poolInitBalance)
	suite.Require().NoError(err)

	// execute function
	for _, coin := range poolInitBalance {
		suite.app.AmmKeeper.SetDenomLiquidity(suite.ctx, types.DenomLiquidity{
			Denom:     coin.Denom,
			Liquidity: coin.Amount,
		})
	}
	pool := types.Pool{
		PoolId:            1,
		Address:           poolAddr.String(),
		RebalanceTreasury: treasuryAddr.String(),
		PoolParams: types.PoolParams{
			UseOracle:                   false,
			ExternalLiquidityRatio:      sdk.NewDec(2),
			WeightBreakingFeeMultiplier: sdk.ZeroDec(),
			WeightBreakingFeeExponent:   sdk.NewDecWithPrec(25, 1), // 2.5
			WeightRecoveryFeePortion:    sdk.NewDecWithPrec(10, 2), // 10%
			ThresholdWeightDifference:   sdk.ZeroDec(),
			SwapFee:                     sdk.ZeroDec(),
			FeeDenom:                    ptypes.BaseCurrency,
		},
		TotalShares: sdk.Coin{},
		PoolAssets: []types.PoolAsset{
			{
				Token:  poolInitBalance[0],
				Weight: sdk.NewInt(10),
			},
			{
				Token:  poolInitBalance[1],
				Weight: sdk.NewInt(10),
			},
		},
		TotalWeight: sdk.ZeroInt(),
	}
	pool2 := types.Pool{
		PoolId:            2,
		Address:           pool2Addr.String(),
		RebalanceTreasury: treasury2Addr.String(),
		PoolParams: types.PoolParams{
			SwapFee:  sdk.ZeroDec(),
			FeeDenom: ptypes.BaseCurrency,
		},
		TotalShares: sdk.Coin{},
		PoolAssets: []types.PoolAsset{
			{
				Token:  pool2InitBalance[0],
				Weight: sdk.NewInt(10),
			},
			{
				Token:  pool2InitBalance[1],
				Weight: sdk.NewInt(10),
			},
		},
		TotalWeight: sdk.ZeroInt(),
	}
	suite.app.AmmKeeper.SetPool(suite.ctx, pool)
	suite.app.AmmKeeper.SetPool(suite.ctx, pool2)

	tokenIn := sdk.NewCoin(ptypes.Elys, sdk.NewInt(100))
	routes := []*types.SwapAmountInRoute{{PoolId: 1, TokenOutDenom: ptypes.BaseCurrency}}
	spotPrice, _, _, _, _, _, _, _, err := suite.app.AmmKeeper.CalcInRouteSpotPrice(suite.ctx, tokenIn, routes, sdk.ZeroDec(), sdk.ZeroDec())
	suite.Require().NoError(err)
	suite.Require().Equal(spotPrice.String(), sdk.OneDec().String())

	routes = []*types.SwapAmountInRoute{
		{PoolId: 1, TokenOutDenom: ptypes.BaseCurrency},
		{PoolId: 2, TokenOutDenom: "uusda"},
	}
	spotPrice, _, _, _, _, _, _, _, err = suite.app.AmmKeeper.CalcInRouteSpotPrice(suite.ctx, tokenIn, routes, sdk.ZeroDec(), sdk.ZeroDec())
	suite.Require().NoError(err)
	suite.Require().Equal(spotPrice.String(), sdk.OneDec().String())

	// Test no routes
	_, _, _, _, _, _, _, _, err = suite.app.AmmKeeper.CalcInRouteSpotPrice(suite.ctx, tokenIn, nil, sdk.ZeroDec(), sdk.ZeroDec())
	suite.Require().Error(err)

	// Test invalid pool
	routes = []*types.SwapAmountInRoute{{PoolId: 9999, TokenOutDenom: "uusda"}}
	_, _, _, _, _, _, _, _, err = suite.app.AmmKeeper.CalcInRouteSpotPrice(suite.ctx, tokenIn, routes, sdk.ZeroDec(), sdk.ZeroDec())
	suite.Require().Error(err)
}
