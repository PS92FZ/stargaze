package alloc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v12/x/alloc/keeper"
	"github.com/public-awesome/stargaze/v12/x/alloc/types"
)

// InitGenesis initializes the alloc module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	k.GetModuleAccount(ctx, types.FairburnPoolName)
	k.GetModuleAccount(ctx, types.SupplementPoolName)
	err := k.FundCommunityPool(ctx)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis returns the alloc module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}
