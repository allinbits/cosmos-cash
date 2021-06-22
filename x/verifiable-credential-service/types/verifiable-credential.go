package types

// NewUserVerifiableCredential constructs a new VerifiableCredential instance
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

// GetBytes is a helper for serializing
func (m VerifiableCredential) GetBytes() []byte {
	dAtA, _ := m.Marshal()
	return dAtA
}

// NewUserCredentialSubject build a new user credential subject
func NewUserCredentialSubject(
	id string,
	root string,
	isVerified bool,
) VerifiableCredential_UserCred {
	return VerifiableCredential_UserCred{
		&UserCredentialSubject{
			Id:         id,
			Root:       root,
			IsVerified: isVerified,
		},
	}
}

// NewProof returns a new proof
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

// EmptyProof - helper function to build an empty (not initialized) proof
func EmptyProof() Proof {
	return Proof{}
}
