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

// NewIdentifier constructs a new Identifier
func NewIssuerVerifiableCredential(
	id string,
	vctype []string,
	issuer string,
	issuanceDate string,
	credentialSubject VerifiableCredential_IssuerCred,
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
	name string,
	address string,
	dob string,
	nationalId string,
	phoneNumber string,
	hasKyc bool,
) VerifiableCredential_UserCred {
	return VerifiableCredential_UserCred{
		&UserCredentialSubject{
			Id:          id,
			Name:        name,
			Address:     address,
			DateOfBirth: dob,
			NationalId:  nationalId,
			PhoneNumber: phoneNumber,
			HasKyc:      hasKyc,
		},
	}
}

func NewIssuerCredentialSubject(
	id string,
	isVerified bool,
) VerifiableCredential_IssuerCred {
	return VerifiableCredential_IssuerCred{
		&IssuerCredentialSubject{
			Id:         id,
			IsVerified: isVerified,
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
