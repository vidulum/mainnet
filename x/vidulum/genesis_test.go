package vidulum_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vidulum/vidulum/testutil/keeper"
	"github.com/vidulum/vidulum/x/vidulum"
	"github.com/vidulum/vidulum/x/vidulum/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VidulumKeeper(t)
	vidulum.InitGenesis(ctx, *k, genesisState)
	got := vidulum.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
