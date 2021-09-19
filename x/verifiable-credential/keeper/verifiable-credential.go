package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

// SetVerifiableCredential commit a verifiable credential to the storage
func (q Keeper) SetVerifiableCredential(ctx sdk.Context, key []byte, vc types.VerifiableCredential) (err error) {
	err = ValidateProof(ctx, q, vc)
	if err != nil {
		return
	}
	q.Set(ctx, key, types.VerifiableCredentialKey, vc, q.MarshalVerifiableCredential)
	return
}

func (q Keeper) GetVerifiableCredential(ctx sdk.Context, key []byte) (types.VerifiableCredential, bool) {
	val, found := q.Get(ctx, key, types.VerifiableCredentialKey, q.UnmarshalVerifiableCredential)
	return val.(types.VerifiableCredential), found
}

func (q Keeper) DeleteVerifiableCredentialFromStore(ctx sdk.Context, key []byte) {
	q.Delete(ctx, key, types.VerifiableCredentialKey)
}

func (q Keeper) UnmarshalVerifiableCredential(value []byte) (interface{}, bool) {
	vc := types.VerifiableCredential{}

	err := q.cdc.Unmarshal(value, &vc)
	if err != nil {
		return types.VerifiableCredential{}, false
	}

	if vc.Id == "" {
		return types.VerifiableCredential{}, false
	}

	return vc, true
}

func (q Keeper) MarshalVerifiableCredential(value interface{}) []byte {
	did := value.(types.VerifiableCredential)

	bytes, _ := q.cdc.Marshal(&did)

	return bytes
}

func (q Keeper) GetAllVerifiableCredentialsWithCondition(
	ctx sdk.Context,
	key []byte,
	vcSelector func(votes types.VerifiableCredential) bool,
) (vcs []types.VerifiableCredential) {
	iterator := q.GetAll(ctx, key)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		vc, _ := q.UnmarshalVerifiableCredential(iterator.Value())
		vcTyped := vc.(types.VerifiableCredential)
		if vcSelector(vcTyped) {
			vcs = append(vcs, vcTyped)
		}
	}

	return vcs
}

func (q Keeper) GetAllVerifiableCredentials(ctx sdk.Context) []types.VerifiableCredential {
	return q.GetAllVerifiableCredentialsWithCondition(
		ctx,
		types.VerifiableCredentialKey,
		func(vc types.VerifiableCredential) bool { return true },
	)
}

// ValidateProof validate the proof of a verifiable credential
func ValidateProof(ctx sdk.Context, k Keeper, vc types.VerifiableCredential) error {
	did, err := func() (did didtypes.DidDocument, err error) {
		if strings.HasPrefix(vc.Issuer, didtypes.DidKeyPrefix) {
			did, _, err = didtypes.ResolveAccountDID(vc.Issuer, ctx.ChainID())
			return
		}
		did, found := k.didKeeper.GetDidDocument(ctx, []byte(vc.Issuer))
		if !found {
			err = didtypes.ErrDidDocumentNotFound
		}
		return
	}()

	if err != nil {
		return sdkerrors.Wrapf(
			err, "issuer DID is not resolvable",
		)
	}

	// verify the signature
	if vc.Proof == nil {
		return sdkerrors.Wrapf(
			types.ErrMessageSigner,
			"proof is nil %v",
			err,
		)
	}
	// get the address in the verification method
	issuerAddress, err := did.GetVerificationMethodAddress(vc.Proof.VerificationMethod)
	if err != nil {
		return sdkerrors.Wrapf(
			types.ErrMessageSigner,
			"the issuer address cannot be retrieved due to %v",
			err,
		)
	}
	// verify the relationships
	issuerBID := didtypes.NewBlockchainAccountID(ctx.ChainID(), issuerAddress)
	if !did.HasRelationship(issuerBID, didtypes.Authentication, didtypes.AssertionMethod) {
		return sdkerrors.Wrapf(
			types.ErrMessageSigner,
			"signer is not in issuer did",
		)
	}
	// verify that is the same of the vc
	issuerAccount, err := sdk.AccAddressFromBech32(issuerAddress)
	if err != nil {
		return sdkerrors.Wrapf(
			types.ErrMessageSigner,
			"failed to convert the issuer address to account %v",
			err,
		)
	}
	// get the public key from the account
	pk, err := k.accountKeeper.GetPubKey(ctx, issuerAccount)
	if err != nil {
		return sdkerrors.Wrapf(
			types.ErrMessageSigner,
			"issuer pubkey not found %v",
			err,
		)
	}
	//
	if isValid := vc.Validate(pk); !isValid {
		return sdkerrors.Wrapf(
			types.ErrMessageSigner,
			"verification error %v",
			err,
		)
	}
	return nil
}
