package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewIdentifier constructs a new Identifier
func NewVerifiableCredential(
	id string,
	vctype []string,
	issuer string,
	issuanceDate string,
	credentialSubject CredentialSubject,
	proof Proof,
) VerifiableCredential {
	return VerifiableCredential{
		Context:           []string{"https://www.w3.org/TR/vc-data-model/"},
		Id:                id,
		Type:              vctype,
		Issuer:            issuer,
		IssuanceDate:      issuanceDate,
		CredentialSubject: &credentialSubject,
		Proof:             &proof,
	}
}

// GetBytes is a helper for serialising
func (vc VerifiableCredential) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&vc))
}

func NewCredentialSubject(
	id string,
	hasKyc bool,
) CredentialSubject {
	return CredentialSubject{
		Id:     id,
		HasKyc: hasKyc,
	}
}

func NewProof(
	proofType string,
	created string,
	proofPurpose string,
	verificationMethod string,
	signature string,
) Proof {
	return Proof{
		Type:               proofType,
		Created:            created,
		ProofPurpose:       proofPurpose,
		VerificationMethod: verificationMethod,
		Signature:          signature,
	}
}
