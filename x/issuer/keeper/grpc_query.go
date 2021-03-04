package keeper

import (
	"context"

	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"

	"github.com/allinbits/cosmos-cash/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (q Keeper) Issuers(
	c context.Context,
	req *types.QueryIssuersRequest,
) (*types.QueryIssuersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	issuers := q.GetAllIssuers(ctx)

	return &types.QueryIssuersResponse{
		Issuers: issuers,
	}, nil
}
