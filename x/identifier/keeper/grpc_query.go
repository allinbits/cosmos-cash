package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
)

var _ types.QueryServer = Keeper{}

// DenomTrace implements the Query/DenomTrace gRPC method
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
