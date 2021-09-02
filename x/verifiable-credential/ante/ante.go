package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didkeeper "github.com/allinbits/cosmos-cash/x/did/keeper"
	didtypes "github.com/allinbits/cosmos-cash/x/did/types"

	vcskeeper "github.com/allinbits/cosmos-cash/x/verifiable-credential/keeper"
	vcstypes "github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

// CheckSignerHasDIDDecorator checks the signer of the message
// has a did in a preprocessing hook
type CheckSignerHasDIDDecorator struct {
	didk didkeeper.Keeper
	vcsk vcskeeper.Keeper
}

func NewCheckSignerHasDIDDecorator(
	didk didkeeper.Keeper,
	vcsk vcskeeper.Keeper,
) CheckSignerHasDIDDecorator {
	return CheckSignerHasDIDDecorator{
		didk: didk,
		vcsk: vcsk,
	}
}

// AnteHandle verifies the did provided is associated with the verifiable credential issuer
func (cshd CheckSignerHasDIDDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *vcstypes.MsgDeleteVerifiableCredential:
			did, found := cshd.didk.GetDidDocument(ctx, []byte(msg.IssuerDid))
			if !found {
				return ctx, sdkerrors.Wrapf(
					nil,
					"did does not exists",
				)
			}
			vcs, found := cshd.vcsk.GetVerifiableCredential(ctx, []byte(msg.VerifiableCredentialId))

			if vcs.Issuer != did.Id {
				return ctx, sdkerrors.Wrapf(
					nil,
					"provided vc and did issuer do not match",
				)
			}

			if !did.HasRelationship(msg.Owner, didtypes.Authentication) {
				return ctx, sdkerrors.Wrapf(
					nil,
					"signer is not in issuer did",
				)
			}
		}
	}
	return next(ctx, tx, simulate)
}
