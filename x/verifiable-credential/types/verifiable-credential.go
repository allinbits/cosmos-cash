package types

import (
	"encoding/base64"
	didtypes "github.com/allinbits/cosmos-cash/x/did/types"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Defines the accepted credential types
const (
	IdentityCredential     = "IdentityCredential"
	UserCredential         = "UserCredential"
	IssuerCredential       = "IssuerCredential"
	RegulatorCredential    = "RegulatorCredential"
	RegistrationCredential = "RegistrationCredential"
	LicenseCredential      = "LicenseCredential"
)

// IsValidCredentialType tells if a credential type is valid (accepted)
func IsValidCredentialType(credential string) bool {
	switch credential {
	case IdentityCredential,
		UserCredential,
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
		Type:              []string{"VerifiableCredential", UserCredential},
		Issuer:            issuer,
		IssuanceDate:      &issuanceDate,
		CredentialSubject: &credentialSubject,
		Proof:             nil,
	}
}

// NewRegistrationVerifiableCredential constructs a new VerifiableCredential instance
func NewRegistrationVerifiableCredential(
	id string,
	issuer string,
	issuanceDate time.Time,
	credentialSubject VerifiableCredential_RegistrationCred,
) VerifiableCredential {
	return VerifiableCredential{
		Context:           []string{"https://www.w3.org/TR/vc-data-model/"},
		Id:                id,
		Type:              []string{"VerifiableCredential", RegistrationCredential},
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

// NewRegulatorVerifiableCredential constructs a new VerifiableCredential instance
func NewRegulatorVerifiableCredential(
	id string,
	issuer string,
	issuanceDate time.Time,
	credentialSubject VerifiableCredential_RegulatorCred,
) VerifiableCredential {
	return VerifiableCredential{
		Context:           []string{"https://www.w3.org/TR/vc-data-model/"},
		Id:                id,
		Type:              []string{"VerifiableCredential", RegulatorCredential},
		Issuer:            issuer,
		IssuanceDate:      &issuanceDate,
		CredentialSubject: &credentialSubject,
		Proof:             nil,
	}
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
	circulationLimit sdk.Coin,
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

// NewRegistrationCredentialSubject create a new license credential subject
// TODO: placeholder implementation, refactor it
func NewRegistrationCredentialSubject(
	id string,
	country string,
	shortName string,
	longName string,
) VerifiableCredential_RegistrationCred {
	return VerifiableCredential_RegistrationCred{
		&RegistrationCredentialSubject{
			Id: id,
			Address: &Address{
				Country: country,
			},
			LegalPersons: []*LegalPerson{
				{
					Names: []*Name{
						{
							Type: "SN",
							Name: shortName,
						},
						{
							Type: "LN",
							Name: longName,
						},
					},
				},
			},
			Ids: []*Id{
				{
					Id:   "529900W6B9NEA233DS71",
					Type: "LEIX",
				},
			},
		},
	}
}

// NewRegulatorCredentialSubject create a new regulator credential subject
func NewRegulatorCredentialSubject(
	subjectID string,
	name string,
	country string,
) VerifiableCredential_RegulatorCred {
	return VerifiableCredential_RegulatorCred{
		&RegulatorCredentialSubject{
			Id:      subjectID,
			Name:    name,
			Country: country,
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

// Validate validates a verifiable credential against a provided public key
func (vc VerifiableCredential) Validate(
	pk cryptotypes.PubKey,
) bool {
	s, err := base64.StdEncoding.DecodeString(vc.Proof.Signature)
	if err != nil {
		panic(err)
	}

	// reset the proof
	vc.Proof = nil

	// TODO: this is an expensive operation, could lead to DDOS
	// TODO: we can hash this and make this less expensive
	isCorrectPubKey := pk.VerifySignature(
		vc.GetBytes(),
		s,
	)

	return isCorrectPubKey
}

// Sign signs a credential with a provided private key
func (vc VerifiableCredential) Sign(
	keyring keyring.Keyring,
	address sdk.Address,
	verificationMethodID string,
) (VerifiableCredential, error) {
	tm := time.Now()
	// reset the proof
	vc.Proof = nil
	// TODO: this could be expensive review this signing method
	// TODO: we can hash this an make this less expensive
	signature, pubKey, err := keyring.SignByAddress(address, vc.GetBytes())
	if err != nil {
		return vc, err
	}

	p := NewProof(
		pubKey.Type(),
		tm.Format(time.RFC3339),
		// TODO: define proof purposes
		"assertionMethod",
		verificationMethodID,
		base64.StdEncoding.EncodeToString(signature),
	)
	vc.Proof = &p
	return vc, nil
}

func (vc VerifiableCredential) Hash() string {
	// TODO: implement the hashing of creds for signing
	return "TODO"
}

// HasType tells whenever a credential has a specific type
func (vc VerifiableCredential) HasType(vcType string) bool {
	for _, vct := range vc.Type {
		if vct == vcType {
			return true
		}
	}
	return false
}

// GetSubjectDID return the credential DID subject, that is the holder
// of the credentials
func (vc VerifiableCredential) GetSubjectDID() didtypes.DID {
	switch subj := vc.CredentialSubject.(type) {
	case *VerifiableCredential_RegistrationCred:
		return didtypes.DID(subj.RegistrationCred.Id)
	case *VerifiableCredential_LicenseCred:
		return didtypes.DID(subj.LicenseCred.Id)
	case *VerifiableCredential_UserCred:
		return didtypes.DID(subj.UserCred.Id)
	case *VerifiableCredential_RegulatorCred:
		return didtypes.DID(subj.RegulatorCred.Id)
	default:
		// TODO, not great
		return didtypes.DID("")
	}
}

// GetIssuerDID returns the did of the issuer
func (vc VerifiableCredential) GetIssuerDID() didtypes.DID {
	return didtypes.DID(vc.Issuer)
}

// GetBytes is a helper for serializing
func (vc VerifiableCredential) GetBytes() []byte {
	dAtA, _ := vc.Marshal()
	return dAtA
}
