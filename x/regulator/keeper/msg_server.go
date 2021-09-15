package keeper

import (
	"context"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/regulator/types"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
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

func (k msgServer) Activate(goCtx context.Context, msg *types.MsgActivate) (*types.MsgActivateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// fetch the regulator address lists
	// err: not found
	addrs := k.Keeper.GetRegulatorsAddresses(ctx)
	k.Logger(ctx).Info("regulators are ", "reg", addrs)
	contains := func(what string, list []string) bool {
		for _, v := range list {
			if v == what {
				return true
			}
		}
		return false
	}

	// if the address is not one of the regulators
	if !contains(msg.Creator, addrs) {
		return nil, sdkerrors.Wrapf(
			types.ErrNotARegulator,
			msg.Creator,
		)
	}
	
	// generate a new did
	did := didtypes.NewChainDID(
		ctx.ChainID(),
		msg.DidId,
	).String()

	// verify that is a valid did
	if !didtypes.IsValidDID(did) {
		return nil, sdkerrors.Wrapf(
			didtypes.ErrInvalidDIDFormat,
			did,
		)
	}

	// build the did document and metadata
	didDoc, _ := didtypes.NewDidDocument(
		did,
		didtypes.WithControllers(didtypes.NewKeyDID(msg.Creator).String()),
	)
	didMeta := didtypes.NewDidMetadata(ctx.TxBytes(), ctx.BlockTime())
	// save the didDoc and the Meta
	k.SetDidDocumentWithMeta(ctx, didDoc, didMeta)

	k.Logger(ctx).Info("regulator activation request", "did", did, "address", msg.Creator)
	// verify regulator DOESN't have regulator credential
	vcs := k.GetVerifiableCredentialWithType(ctx, did, vctypes.RegulatorCredential)
	if len(vcs) > 0 {
		return nil, sdkerrors.Wrapf(
			types.ErrAlreadyActive,
			msg.Creator,
		)
	}
	// generate a regulator vc
	vc := vctypes.NewRegulatorVerifiableCredential(
		"id",
		did,
		ctx.BlockTime(),
		vctypes.NewRegulatorCredentialSubject(
			"id2",
			msg.Name,
			msg.Country,
		),
	)
	// store the regulator vc
	k.SetVerifiableCredential(ctx, vc)
	// reply
	k.Logger(ctx).Info("regulator activation success", "did", did, "address", msg.Creator)

	return &types.MsgActivateResponse{
		Did:  did,
		VcId: vc.Id,
	}, nil
}
