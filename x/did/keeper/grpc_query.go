package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/allinbits/cosmos-cash/x/did/resolver"
	"github.com/allinbits/cosmos-cash/x/did/types"
)

var _ types.QueryServer = Keeper{}

// DidDocuments implements the DidDocuments gRPC method
func (k Keeper) DidDocuments(
	c context.Context,
	req *types.QueryDidDocumentsRequest,
) (*types.QueryDidDocumentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	dids := k.GetAllDidDocuments(ctx)

	return &types.QueryDidDocumentsResponse{
		DidDocuments: dids,
	}, nil
}

// DidDocument implements the DidDocument gRPC method
func (k Keeper) DidDocument(
	c context.Context,
	req *types.QueryDidDocumentRequest,
) (*types.QueryDidDocumentResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "verifiable credential id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	did, found := k.GetDidDocument(ctx, []byte(req.Id))
	if !found {
		// it it isn't a key did, return error
		if !strings.HasPrefix(req.Id, types.DidKeyPrefix) {
			return nil, status.Error(codes.NotFound, "did not found: QueryDidDocument")
		}
		// autoresolve the address
		doc, meta, err := resolver.ResolveAccountDID(req.Id)
		if err != nil {
			return nil, status.Error(codes.NotFound, "did not found: QueryDidDocument")
		}
		return &types.QueryDidDocumentResponse{
			DidDocument: doc,
			DidMetadata: meta,
		}, nil
	}

	didM, _ := k.GetDidMetadata(ctx, []byte(req.Id))

	return &types.QueryDidDocumentResponse{
		DidDocument: did,
		DidMetadata: didM,
	}, nil
}
