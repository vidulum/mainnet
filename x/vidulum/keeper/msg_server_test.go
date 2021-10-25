package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/vidulum/vidulum/testutil/keeper"
	"github.com/vidulum/vidulum/x/vidulum/keeper"
	"github.com/vidulum/vidulum/x/vidulum/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VidulumKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
