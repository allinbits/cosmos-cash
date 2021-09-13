package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
		// if it isn't a key did, return error
		if !strings.HasPrefix(req.Id, types.DidKeyPrefix) {
			return nil, status.Error(codes.NotFound, fmt.Sprint("resolution failed for did: ", req.Id))
		}
		// auto-resolve the address
		doc, meta, err := types.ResolveAccountDID(types.DID(req.Id), ctx.ChainID())
		if err != nil {
			return nil, status.Error(codes.Unavailable, "cosmos address account resolution error")
		}
		return &types.QueryDidDocumentResponse{
			DidDocument: doc,
			DidMetadata: meta,
		}, nil
	}
	// now fetch the metadata
	didM, _ := k.GetDidMetadata(ctx, []byte(req.Id))

	return &types.QueryDidDocumentResponse{
		DidDocument: did,
		DidMetadata: didM,
	}, nil
}
