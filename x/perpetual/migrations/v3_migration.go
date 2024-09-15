package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/perpetual/types"
)

func (m Migrator) V3Migration(ctx sdk.Context) error {
	params := types.NewParams()
	m.keeper.SetParams(ctx, &params)
	return nil
}
