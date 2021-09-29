package regulator

import (
	"fmt"
	vctypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/allinbits/cosmos-cash/x/regulator/keeper"
	"github.com/allinbits/cosmos-cash/x/regulator/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *vctypes.MsgIssueCredential:
			switch msg.Credential.CredentialSubject.(type) {
			case *vctypes.VerifiableCredential_RegulatorCred:
				res, err := msgServer.Activate(sdk.WrapSDKContext(ctx), msg)
				return sdk.WrapServiceResult(ctx, res, err)
			case *vctypes.VerifiableCredential_RegistrationCred:
				res, err := msgServer.IssueRegistrationCredential(sdk.WrapSDKContext(ctx), msg)
				return sdk.WrapServiceResult(ctx, res, err)
			case *vctypes.VerifiableCredential_LicenseCred:
				res, err := msgServer.IssueLicenseCredential(sdk.WrapSDKContext(ctx), msg)
				return sdk.WrapServiceResult(ctx, res, err)
			default:
				errMsg := fmt.Sprintf("unrecognized credential %s message type: %T", types.ModuleName, msg)
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
			}
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
