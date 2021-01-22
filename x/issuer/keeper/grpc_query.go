package keeper

import (
	"context"

	"github.com/allinbits/cosmos-cash/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// Issuers implements the Query/issuers gRPC method
func (q Keeper) Issuers(c context.Context, req *types.QueryIssuersRequest) (*types.QueryIssuersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	issuers := q.GetAllIssuers(ctx)

	return &types.QueryIssuersResponse{
		Issuers: issuers,
	}, nil
}
