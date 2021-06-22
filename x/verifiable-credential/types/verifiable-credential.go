package types

// Defines the accepted credential types
const (
	IdentityCredential  = "IdentityCredential"
	KYCCredential       = "KYCCredential"
	IssuerCredential    = "IssuerCredential"
	RegulatorCredential = "RegulatorCredential"
)

// IsValidCredentialType tells if a credential type is valid (accepted)
func IsValidCredentialType(credential string) bool {
	switch credential {
	case IdentityCredential,
		KYCCredential,
		IssuerCredential,
		RegulatorCredential:
		return true
	default:
		return false
	}
}

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

// NewUserCredentialSubject create a new credential subject
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

// NewProof create a new proof for a verifiable credential
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
