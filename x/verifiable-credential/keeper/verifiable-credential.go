package keeper

import (
	"encoding/base64"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"github.com/allinbits/cosmos-cash/x/verifiable-credential/types"
)

func (q Keeper) SetVerifiableCredential(ctx sdk.Context, key []byte, vc types.VerifiableCredential) (err error) {

	// XXX: where to do this check
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

	// XXX: Paddy you are going to love this one
	did, err := func() (did didtypes.DidDocument, err error) {
		if strings.HasPrefix(vc.Issuer, didtypes.DidKeyPrefix) {
			did, _, err = didtypes.ResolveAccountDID(vc.Issuer, ctx.ChainID())
		} else {
			var found bool
			did, found = k.didKeeper.GetDidDocument(ctx, []byte(vc.Issuer))
			if !found {
				err = didtypes.ErrDidDocumentNotFound
			}
		}
		return
	}()

	if err != nil {
		return sdkerrors.Wrapf(
			err, "issuer DID is not resolvable",
		)
	}

	// verify the signature
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
	signerID := didtypes.NewBlockchainAccountID(ctx.ChainID(), issuerAddress)
	if !did.HasRelationship(signerID, didtypes.Authentication, didtypes.AssertionMethod) {
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
	// verify the signature
	signature, err := base64.StdEncoding.DecodeString(vc.Proof.Signature)
	if err != nil {
		return sdkerrors.Wrapf(
			types.ErrMessageSigner,
			"signature decoding error %v",
			err,
		)
	}
	if isValid := pk.VerifySignature(vc.GetPayload(), signature); !isValid {
		return sdkerrors.Wrapf(
			types.ErrMessageSigner,
			"verification error %v",
			err,
		)
	}
	return nil
}
