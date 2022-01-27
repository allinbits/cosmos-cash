package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/v3/x/issuer/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Issuers(
	c context.Context,
	req *types.QueryIssuersRequest,
) (*types.QueryIssuersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	issuers := k.GetAllIssuers(ctx)

	return &types.QueryIssuersResponse{
		Issuers: issuers,
	}, nil
}
