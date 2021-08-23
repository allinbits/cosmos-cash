package types

import (
	"time"

	"github.com/cosmos/cosmos-sdk/types"
)

// Defines the accepted credential types
const (
	IdentityCredential  = "IdentityCredential"
	KYCCredential       = "KYCCredential"
	IssuerCredential    = "IssuerCredential"
	RegulatorCredential = "RegulatorCredential"
	LicenseCredential   = "LicenseCredential"
)

// IsValidCredentialType tells if a credential type is valid (accepted)
func IsValidCredentialType(credential string) bool {
	switch credential {
	case IdentityCredential,
		KYCCredential,
		IssuerCredential,
		RegulatorCredential,
		LicenseCredential:
		return true
	default:
		return false
	}
}

// NewUserVerifiableCredential constructs a new VerifiableCredential instance
func NewUserVerifiableCredential(
	id string,
	issuer string,
	issuanceDate time.Time,
	credentialSubject VerifiableCredential_UserCred,
) VerifiableCredential {
	return VerifiableCredential{
		Context:           []string{"https://www.w3.org/TR/vc-data-model/"},
		Id:                id,
		Type:              []string{"VerifiableCredential", KYCCredential},
		Issuer:            issuer,
		IssuanceDate:      &issuanceDate,
		CredentialSubject: &credentialSubject,
		Proof:             nil,
	}
}

// NewLicenseVerifiableCredential constructs a new VerifiableCredential instance
func NewLicenseVerifiableCredential(
	id string,
	issuer string,
	issuanceDate time.Time,
	credentialSubject VerifiableCredential_LicenseCred,
) VerifiableCredential {
	return VerifiableCredential{
		Context:           []string{"https://www.w3.org/TR/vc-data-model/"},
		Id:                id,
		Type:              []string{"VerifiableCredential", LicenseCredential},
		Issuer:            issuer,
		IssuanceDate:      &issuanceDate,
		CredentialSubject: &credentialSubject,
		Proof:             nil,
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

// NewLicenseCredentialSubject create a new license credential subject
func NewLicenseCredentialSubject(
	id string,
	licenseType string,
	country string,
	authority string,
	circulationLimit types.Coin,
) VerifiableCredential_LicenseCred {
	return VerifiableCredential_LicenseCred{
		&LicenseCredentialSubject{
			Id:               id,
			LicenseType:      licenseType,
			Country:          country,
			Authority:        authority,
			CirculationLimit: circulationLimit,
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
