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
func (q Keeper) Identifiers(
	c context.Context,
	req *types.QueryIdentifiersRequest,
) (*types.QueryIdentifiersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	identifiers := q.GetAllIdentifiers(ctx)

	return &types.QueryIdentifiersResponse{
		DidDocuments: identifiers,
	}, nil
}

// Identifers implements the Identifers gRPC method
func (q Keeper) Identifier(
	c context.Context,
	req *types.QueryIdentifierRequest,
) (*types.QueryIdentifierResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "verifiable credential id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	identifier, found := q.GetIdentifier(ctx, []byte(req.Id))
	if !found {
		return nil, status.Error(codes.NotFound, "identifier not found: QueryIdentifier")
	}

	return &types.QueryIdentifierResponse{
		identifier,
	}, nil
}
