package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/allinbits/cosmos-cash/x/issuer/types"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

// NewHandler ...
func NewHandler(k Keeper) sdk.Handler {
	msgServer := NewMsgServerImpl(k)
	// this line is used by starport scaffolding # handler/msgServer

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateIssuer:
			res, err := msgServer.CreateIssuer(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBurnToken:
			res, err := msgServer.BurnToken(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgMintToken:
			res, err := msgServer.MintToken(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgPauseToken:
			res, err := msgServer.PauseToken(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *vctypes.MsgIssueCredential:
			if _, ok := msg.Credential.CredentialSubject.(*vctypes.VerifiableCredential_UserCred); !ok {
				errMsg := fmt.Sprintf("unsupported credential type %s message type: %T", types.ModuleName, msg)
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
			}
			res, err := msgServer.IssueUserCredential(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
