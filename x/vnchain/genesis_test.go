package vnchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vnchain/testutil/keeper"
	"vnchain/testutil/nullify"
	"vnchain/x/vnchain"
	"vnchain/x/vnchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VnchainKeeper(t)
	vnchain.InitGenesis(ctx, *k, genesisState)
	got := vnchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
