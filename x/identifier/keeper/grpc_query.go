package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/allinbits/cosmos-cash/x/identifier/types"
)

var _ types.QueryServer = Keeper{}

// Identifiers implements the Identifiers gRPC method
func (k Keeper) Identifiers(
	c context.Context,
	req *types.QueryIdentifiersRequest,
) (*types.QueryIdentifiersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	identifiers := k.GetAllIdentifiers(ctx)

	return &types.QueryIdentifiersResponse{
		DidDocuments: identifiers,
	}, nil
}

// Identifier implements the Identifier gRPC method
func (k Keeper) Identifier(
	c context.Context,
	req *types.QueryIdentifierRequest,
) (*types.QueryIdentifierResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "verifiable credential id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	identifier, found := k.GetIdentifier(ctx, []byte(req.Id))
	if !found {
		return nil, status.Error(codes.NotFound, "identifier not found: QueryIdentifier")
	}

	return &types.QueryIdentifierResponse{
		DidDocument: identifier,
	}, nil
}
