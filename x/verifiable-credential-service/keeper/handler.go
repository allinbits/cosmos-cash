package keeper

import (
	"fmt"

	"github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k Keeper) sdk.Handler {
	msgServer := NewMsgServerImpl(k)
	// this line is used by starport scaffolding # handler/msgServer

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateVerifiableCredential:
			res, err := msgServer.CreateVerifiableCredential(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateIssuerVerifiableCredential:
			res, err := msgServer.CreateIssuerVerifiableCredential(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
