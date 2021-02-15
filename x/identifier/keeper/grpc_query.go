package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
)

var _ types.QueryServer = Keeper{}

// Identifers implements the Identifers gRPC method
func (q Keeper) Identifiers(c context.Context, req *types.QueryIdentifiersRequest) (*types.QueryIdentifiersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	identifiers := q.GetAllIdentifiers(ctx)

	return &types.QueryIdentifiersResponse{
		DidDocuments: identifiers,
	}, nil
}
