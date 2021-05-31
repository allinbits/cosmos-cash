package types

// NewIdentifier constructs a new Identifier
func NewUserVerifiableCredential(
	id string,
	vctype []string,
	issuer string,
	issuanceDate string,
	credentialSubject VerifiableCredential_UserCred,
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
	dAtA, _ := vc.Marshal()
	return dAtA
}

func NewUserCredentialSubject(
	id string,
	root string,
	hasKyc bool,
) VerifiableCredential_UserCred {
	return VerifiableCredential_UserCred{
		&UserCredentialSubject{
			Id:     id,
			Root:   root,
			HasKyc: hasKyc,
		},
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
