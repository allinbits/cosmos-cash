package keeper

import (
	"context"
	"encoding/base64"

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

// VerifiableCredential queries verifiable credentials info for given verifiable credentials id
func (q Keeper) VerifiableCredential(
	c context.Context,
	req *types.QueryVerifiableCredentialRequest,
) (*types.QueryVerifiableCredentialResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.VerifiableCredentialId == "" {
		return nil, status.Error(codes.InvalidArgument, "verifiable credential id cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	vc, found := q.GetVerifiableCredential(ctx, []byte(req.VerifiableCredentialId))
	if !found {
		return nil, status.Errorf(codes.NotFound, "vc %s not found", req.VerifiableCredentialId)
	}

	return &types.QueryVerifiableCredentialResponse{VerifiableCredential: vc}, nil
}

// ValidateVerifiableCredential queries verifiable credentials info with a public key a check validity
func (q Keeper) ValidateVerifiableCredential(
	c context.Context,
	req *types.QueryValidateVerifiableCredentialRequest,
) (*types.QueryValidateVerifiableCredentialResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	vc, found := q.GetVerifiableCredential(ctx, []byte(req.VerifiableCredentialId))
	if !found {
		return nil, status.Errorf(codes.NotFound, "vc %s not found", req.VerifiableCredentialId)
	}
	pubkey, _ := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, req.IssuerPubkey)
	signature := vc.Proof.Signature
	emptyProof := types.NewProof("", "", "", "", "")
	vc.Proof = &emptyProof
	s, _ := base64.StdEncoding.DecodeString(signature)

	// TODO: this is an expesive operation, could lead to DDOS
	// TODO: we can hash this and make this less expensive
	isCorrectPubKey := pubkey.VerifySignature(
		vc.GetBytes(),
		s,
	)

	return &types.QueryValidateVerifiableCredentialResponse{
		IsValid: isCorrectPubKey,
	}, nil
}
