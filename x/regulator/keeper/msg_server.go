package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/regulator/types"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
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

// contains is an helper function to search a string in a slice of strings
func contains(what string, list []string) bool {
	for _, v := range list {
		if v == what {
			return true
		}
	}
	return false
}

// Activate activates a regulator
func (k msgServer) Activate(goCtx context.Context, msg *types.MsgActivate) (*types.MsgActivateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("regulator activation request", "did", msg.Credentials.Issuer, "address", msg.Creator)
	// fetch the regulator address lists
	// err: not found
	addrs := k.Keeper.GetRegulatorsAddresses(ctx)
	k.Logger(ctx).Info("regulators are ", "regulator addresses", addrs)

	// if the address is not one of the regulators
	if !contains(msg.Creator, addrs) {
		k.Logger(ctx).Error("regulator activation failed", "msg", types.ErrNotARegulator)
		return nil, sdkerrors.Wrapf(
			types.ErrNotARegulator,
			msg.Creator,
		)
	}

	// verify that is a valid did
	if !didtypes.IsValidDID(msg.Credentials.Issuer) {
		k.Logger(ctx).Error("regulator activation failed", "issuer did", didtypes.ErrInvalidDIDFormat)
		return nil, sdkerrors.Wrapf(
			didtypes.ErrInvalidDIDFormat,
			msg.Credentials.Issuer,
		)
	}

	// store the regulator vc
	if err := k.SetVerifiableCredential(ctx, *msg.Credentials); err != nil {
		k.Logger(ctx).Error("regulator activation failed", "signature verification error", err)
		return nil, sdkerrors.Wrapf(
			vctypes.ErrMessageSigner,
			"credential proof could not be verified",
		)
	}
	// reply
	k.Logger(ctx).Info("regulator activation success", "did", msg.Credentials.Issuer, "address", msg.Creator)

	return &types.MsgActivateResponse{}, nil
}

// Activate activates a regulator
func (k msgServer) Register(goCtx context.Context, msg *types.MsgActivate) (*types.MsgActivateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("business entity registration request", "did", msg.Credentials.Issuer, "signer", msg.Creator)
	// fetch the regulator address lists
	// err: not found
	addrs := k.Keeper.GetRegulatorsAddresses(ctx)
	k.Logger(ctx).Info("regulators are ", "reg", addrs)

	// if the address is not one of the regulators
	if !contains(msg.Creator, addrs) {
		return nil, sdkerrors.Wrapf(
			types.ErrNotARegulator,
			msg.Creator,
		)
	}

	// verify that is a valid did
	if !didtypes.IsValidDID(msg.Credentials.Issuer) {
		return nil, sdkerrors.Wrapf(
			didtypes.ErrInvalidDIDFormat,
			msg.Credentials.Issuer,
		)
	}
	// verify that a credential with the same id does not exists
	_, found := k.GetVerifiableCredential(ctx, msg.Credentials.Id)
	if found {
		return nil, sdkerrors.Wrapf(
			vctypes.ErrVerifiableCredentialFound,
			"vc already exists",
		)
	}
	// verify regulator DOESN't have regulator credential
	vcs := k.GetVerifiableCredentialWithType(ctx, msg.Credentials.Issuer, vctypes.RegulatorCredential)
	if len(vcs) > 0 {
		return nil, sdkerrors.Wrapf(
			types.ErrAlreadyActive,
			msg.Creator,
		)
	}

	// build the did document and metadata
	did := didtypes.NewKeyDID(msg.Creator)
	didDoc, _ := didtypes.NewDidDocument(
		msg.Credentials.Issuer,
		didtypes.WithControllers(did.String()),
		didtypes.WithVerifications(
			didtypes.NewAccountVerification(did, ctx.ChainID(), msg.Creator, didtypes.AssertionMethod),
		),
	)
	didMeta := didtypes.NewDidMetadata(ctx.TxBytes(), ctx.BlockTime())
	// save the didDoc and the Meta
	k.SetDidDocumentWithMeta(ctx, didDoc, didMeta)
	// store the regulator vc
	err := k.SetVerifiableCredential(ctx, *msg.Credentials)
	if err != nil {
		// TODO in this case we MUS rollback did creation, perhaps did and credentials should be
		//
	}

	// reply
	k.Logger(ctx).Info("regulator activation success", "did", msg.Credentials.Issuer, "address", msg.Creator)

	return &types.MsgActivateResponse{}, nil
}
