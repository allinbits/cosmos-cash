package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
)

var _ types.QueryServer = Keeper{}

// VerifiableCredential implements the VerifiableCredentials gRPC method
func (q Keeper) VerifiableCredentials(
	c context.Context,
	req *types.QueryVerifiableCredentialsRequest,
) (*types.QueryVerifiableCredentialsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	vcs := q.GetAllVerifiableCredentials(ctx)

	return &types.QueryVerifiableCredentialsResponse{
		Vcs: vcs,
	}, nil
}
