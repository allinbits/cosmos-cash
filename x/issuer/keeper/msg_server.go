package keeper

import (
	"context"
	"fmt"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateIssuer creates a new e-money token issuer
func (k msgServer) CreateIssuer(
	goCtx context.Context,
	msg *types.MsgCreateIssuer,
) (*types.MsgCreateIssuerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: pass in the did URI as an arg {msg.Id}
	// TODO: ensure this keeper can only read from store
	did, found := k.Keeper.ik.GetIdentifier(ctx, []byte("did:cash:"+msg.Owner))
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"identifer does not exists",
		)
	}

	// TODO: optimise here
	foundKey := false
	for _, auth := range did.Authentication {
		if auth.Controller == msg.Owner {
			fmt.Println("found key")
			foundKey = true
		}
	}
	if !foundKey {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"msg sender not in auth array in did document",
		)
	}

	// TODO: optimise here
	// check if the did document has the issuer credential
	hasIssuerCredential := false
	for _, service := range did.Services {
		// TODO use enum here
		if service.Type == "KYCCredential" {
			// TODO: ensure this keeper can only read from store
			vc, found := k.Keeper.vcsk.GetVerifiableCredential(ctx, []byte(service.Id))
			if !found {
				return nil, sdkerrors.Wrapf(
					types.ErrIssuerFound,
					"credential not found",
				)
			}
			hasIssuerCredential = vc.CredentialSubject.HasKyc
			// TODO: validate credential here
		}
	}
	if !hasIssuerCredential {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"did document doesnt have a credential to create issuers",
		)
	}

	_, found = k.Keeper.GetIssuer(ctx, []byte(msg.Owner))
	if found {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"issuer already exists",
		)
	}

	// TODO: should the did URI be the issuer address
	identifer := types.Issuer{
		Token:   msg.Token,
		Fee:     msg.Fee,
		Address: msg.Owner,
	}

	k.Keeper.SetIssuer(ctx, []byte(identifer.Address), identifer)

	circulatingSupply := 1000000000000

	// TODO: this needs to be refactored
	issuerToken := sdk.NewCoins(sdk.NewInt64Coin(msg.Token, int64(circulatingSupply)))

	// mint new tokens for the issuer
	if err := k.bk.MintCoins(
		ctx, types.ModuleName, issuerToken,
	); err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"cannot mint coins",
		)
	}

	receipent, _ := sdk.AccAddressFromBech32(msg.Owner)

	// send tokens from module to issuer
	if err := k.bk.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, receipent, issuerToken,
	); err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrIssuerFound,
			"cannot send tokens from module to issuer account",
		)
	}

	return &types.MsgCreateIssuerResponse{}, nil
}
