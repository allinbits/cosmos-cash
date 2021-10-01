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
	k.Logger(ctx).Info("regulator activation request", "did", msg.Credential.Issuer, "address", msg.Owner)
	// fetch the regulator address lists
	// err: not found
	addrs := k.Keeper.GetRegulatorsAddresses(ctx)
	k.Logger(ctx).Info("regulators are ", "regulator addresses", addrs)

	// if the address is not one of the regulators
	if !contains(msg.Owner, addrs) {
		k.Logger(ctx).Error("regulator activation failed", "msg", types.ErrNotARegulator)
		return nil, sdkerrors.Wrapf(
			types.ErrNotARegulator,
			msg.Owner,
		)
	}

	// verify that is a valid did
	if !didtypes.IsValidDID(msg.Credential.Issuer) {
		k.Logger(ctx).Error("regulator activation failed", "issuer did", didtypes.ErrInvalidDIDFormat)
		return nil, sdkerrors.Wrapf(
			didtypes.ErrInvalidDIDFormat,
			msg.Credential.Issuer,
		)
	}

	// store the regulator vc
	if err := k.SetVerifiableCredential(ctx, *msg.Credential); err != nil {
		k.Logger(ctx).Error("regulator activation failed", "signature verification error", err)
		return nil, sdkerrors.Wrapf(err, "credential proof could not be verified")
	}

	ctx.EventManager().EmitEvent(
		vctypes.NewCredentialCreatedEvent(msg.Owner, msg.Credential.Id),
	)
	// reply
	k.Logger(ctx).Info("regulator activation success", "did", msg.Credential.Issuer, "address", msg.Owner)

	return &types.MsgActivateResponse{}, nil
}

// IssueRegistrationCredential activates a regulator
func (k msgServer) IssueRegistrationCredential(goCtx context.Context, msg *types.MsgIssueRegistrationCredential) (*types.MsgIssueRegistrationCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("issue registration request", "address", msg.Owner, "credential", msg.Credential)

	// check that the issuer has a regulator license
	vcs := k.vcKeeper.GetVerifiableCredentialWithType(ctx, msg.Credential.Issuer, vctypes.RegulatorCredential)
	if len(vcs) != 1 { // there must be exactly one
		err := sdkerrors.Wrapf(types.ErrNotARegulator, "issuer is not a recgulator")
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// store the credentials
	if vcErr := k.SetVerifiableCredential(ctx, *msg.Credential); vcErr != nil {
		err := sdkerrors.Wrapf(vcErr, "credential proof cannot be verified")
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	k.Logger(ctx).Info("issue registration request successful", "did", msg.Credential.Issuer, "address", msg.Owner)

	ctx.EventManager().EmitEvent(
		vctypes.NewCredentialCreatedEvent(msg.Owner, msg.Credential.Id),
	)

	return &types.MsgIssueRegistrationCredentialResponse{}, nil
}

// IssueLicenseCredential activates a regulator
func (k msgServer) IssueLicenseCredential(goCtx context.Context, msg *types.MsgIssueLicenseCredential) (*types.MsgIssueLicenseCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("issue license request", "credential", msg.Credential, "address", msg.Owner)

	// check that the issuer has a regulator license
	if vcs := k.vcKeeper.GetVerifiableCredentialWithType(ctx, msg.Credential.Issuer, vctypes.RegulatorCredential); len(vcs) != 1 { // there must be exactly one
		err := sdkerrors.Wrapf(types.ErrNotARegulator, "issuer is not a recgulator")
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// check that the subject has a regulator license
	if vcs := k.vcKeeper.GetVerifiableCredentialWithType(ctx, msg.Credential.GetSubjectDID().String(), vctypes.RegistrationCredential); len(vcs) != 1 { // there must be exactly one
		err := sdkerrors.Wrapf(types.ErrNotARegulator, "subject is not registered: a registration credential is required to obtain a license")
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// store the credentials
	if vcErr := k.SetVerifiableCredential(ctx, *msg.Credential); vcErr != nil {
		err := sdkerrors.Wrapf(vcErr, "credential proof cannot be verified")
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	k.Logger(ctx).Info("issue license request successful", "did", msg.Credential.Issuer, "address", msg.Owner)

	ctx.EventManager().EmitEvent(
		vctypes.NewCredentialCreatedEvent(msg.Owner, msg.Credential.Id),
	)

	return &types.MsgIssueLicenseCredentialResponse{}, nil
}

// RevokeCredential revoke a credential
func (k msgServer) RevokeCredential(goCtx context.Context, msg *types.MsgRevokeCredential) (*types.MsgRevokeCredentialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Info("revoke license request", "credential", msg.CredentialId, "address", msg.Owner)

	if vcErr := k.Keeper.DeleteVerifiableCredential(ctx, msg.CredentialId, msg.Owner); vcErr != nil {
		err := sdkerrors.Wrapf(vcErr, "credential proof cannot be verified")
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	k.Logger(ctx).Info("revoke license request successful", "credential", msg.CredentialId, "address", msg.Owner)

	ctx.EventManager().EmitEvent(
		vctypes.NewCredentialDeletedEvent(msg.Owner, msg.CredentialId),
	)

	return &types.MsgRevokeCredentialResponse{}, nil
}
