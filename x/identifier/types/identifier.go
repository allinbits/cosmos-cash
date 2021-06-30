package types

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/**
Regexp generated using this ABNF specs and using https://abnf.msweet.org/index.php

did-url = did path-abempty [ "?" query ] [ "#" fragment ]
did                = "did:" method-name ":" method-specific-id
method-name        = 1*method-char
method-char        = %x61-7A / DIGIT
method-specific-id = *( *idchar ":" ) 1*idchar
idchar             = ALPHA / DIGIT / "." / "-" / "_" / pct-encoded
pct-encoded        = "%" HEXDIG HEXDIG
query       = *( pchar / "/" / "?" )
fragment    = *( pchar / "/" / "?" )
path-abempty  = *( "/" segment )
segment       = *pchar
unreserved    = ALPHA / DIGIT / "-" / "." / "_" / "~"
pchar         = unreserved / pct-encoded / sub-delims / ":" / "@"
sub-delims    = "!" / "$" / "&" / "'" / "(" / ")"
                 / "*" / "+" / "," / ";" / "="
*/

const (
	contextDIDBase            = "https://www.w3.org/ns/did/v1"
	didValidationRegexpStr    = `^did\:[a-z0-9]+\:(([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])*\:)*([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])+$`
	didURLValidationRegexpStr = `^did\:[a-z0-9]+\:(([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])*\:)*([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])+(/(([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)*)*(\?(((([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)|/|\?)*))?(#(((([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)|/|\?)*))?$`
)

var (
	didValidationRegexp    = regexp.MustCompile(didValidationRegexpStr)
	didURLValidationRegexp = regexp.MustCompile(didURLValidationRegexpStr)
)

// IsValidDID validate the input string according to the
// did specification (cfr. https://www.w3.org/TR/did-core/#did-syntax ).
func IsValidDID(input string) bool {
	return didValidationRegexp.MatchString(input)
}

// IsValidDIDURL validate the input string according to the
// did url specification (cfr. https://www.w3.org/TR/did-core/#did-url-syntax  ).
func IsValidDIDURL(input string) bool {
	return didURLValidationRegexp.MatchString(input)
}

// IsValidRFC3986Uri checks if the input string is a valid RFC3986 URI
func IsValidRFC3986Uri(input string) bool {
	if _, err := url.Parse(input); err != nil {
		return false
	}
	return true
}

// IsEmpty tells if the trimmed input is empty
func IsEmpty(input string) bool {
	if strings.TrimSpace(input) == "" {
		return true
	}
	return false
}

// NewIdentifier constructs a new Identifier
func NewIdentifier(id string, services []*Service, verifications []*Verification) (did DidDocument, err error) {

	if !IsValidDID(id) {
		err = fmt.Errorf("the document did %s is not compliant with the specification: cfr https://www.w3.org/TR/did-core/#did-syntax", id)
		return
	}

	did = DidDocument{
		Context: []string{contextDIDBase},
		Id:      id,
	}

	// used to check duplicates
	index := make(map[string]bool)

	// services must be unique
	for _, s := range services {
		// verify that the id and endpoint are valid url (according to RFC3986)
		if !IsValidRFC3986Uri(s.Id) {
			err = fmt.Errorf("service id is not a valid RFC3986 uri: %s", s.Id)
			return
		}

		// verify that the id and endpoint are valid url (according to RFC3986)
		if !IsValidRFC3986Uri(s.ServiceEndpoint) {
			err = fmt.Errorf("service endpoint is not a valid RFC3986 uri: %s", s.ServiceEndpoint)
			return
		}

		// verify that there are no duplicates in method ids
		if _, found := index[s.Id]; found {
			err = fmt.Errorf("duplicated verification method id %s", s.Id)
			return
		}
		index[s.Id] = true

		did.Services = append(did.Services, s)
	}

	err = did.AddVerifications(verifications...)

	return
}

// AddVerifications add one or more verification method and relations to a did document
func (did *DidDocument) AddVerifications(verifications ...*Verification) (err error) {

	// verify that there are no duplicates in method ids
	index := make(map[string]bool)
	// load existing verifications if any
	for _, v := range did.VerificationMethods {
		index[v.Id] = true
	}
	// loop through the verifications and look for problems
	for _, v := range verifications {
		// verify that the method id is correct
		if !IsValidDIDURL(v.Method.Id) {
			err = fmt.Errorf("the verification method id %s is not compliant with the specification: cfr https://www.w3.org/TR/did-core/#did-url-syntax", v.Method.Id)
			return
		}

		// if the controller is not set return error
		if v.Method.Controller == "" {
			err = fmt.Errorf("controller not set for verification method %s, consider using the document did as controller: %s", v.Method.Id, did.Id)
			return
		}

		// check for empty method type
		if v.Method.Type == "" {
			err = fmt.Errorf("type not set for verification method %s", v.Method.Id)
			return
		}

		// check for empty publickey
		if v.Method.PublicKeyBase58 == "" {
			err = fmt.Errorf("public key not set for verification method %s", v.Method.Id)
			return
		}

		// verify that there are no duplicates in method ids
		if _, found := index[v.Method.Id]; found {
			err = fmt.Errorf("duplicated verification method id %s", v.Method.Id)
			return
		}
		index[v.Method.Id] = true

		// first add the method to the list of methods
		did.VerificationMethods = append(did.VerificationMethods, v.GetMethod())

		// now add the relationships
		for _, r := range v.Relationships {
			mIDs, errR := did.getRelationships(r)
			if errR != nil {
				err = errR
				return
			}
			*mIDs = append(*mIDs, v.Method.Id)
		}
	}
	return
}

// RevokeVerification revoke a verification method
func (did *DidDocument) RevokeVerification(methodID string) {
	//TODO Implement
}

// ControllerInRelationships verifies if a controller did
// exists for at least one of the relationships in the did document
func (did DidDocument) ControllerInRelationships(
	contoller string,
	relationships ...VerificationRelationship) bool {
	keyController := make(map[string]string)
	// first check if the controller exists
	for _, vm := range did.VerificationMethods {
		if vm.Controller != contoller {
			continue
		}
		keyController[vm.Id] = vm.Controller
	}
	// if controller was not found then return
	if len(keyController) == 0 {
		return false
	}
	// now see if the controller key is in the relationship
	for _, r := range relationships {
		methodIDs, _ := did.getRelationships(r)
		for _, k := range *methodIDs {
			if _, found := keyController[k]; found {
				return true
			}
		}
	}
	return false
}

// GetBytes is a helper for serializing
func (did DidDocument) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&did))
}

// getRelationships returns the pointer to the relationship slice identified
// by r. This is done to improve ergonomics to access relationships data
func (did *DidDocument) getRelationships(r VerificationRelationship) (methodIDs *[]string, err error) {

	lazyGet := func(data *[]string) *[]string {
		if data == nil {
			newData := make([]string, 0)
			return &newData
		}
		return data
	}

	switch r {
	case VerificationRelationship_authentication:
		methodIDs = lazyGet(&did.Authentication)
	case VerificationRelationship_assertionMethod:
		methodIDs = lazyGet(&did.AssertionMethod)
	case VerificationRelationship_keyAgreement:
		methodIDs = lazyGet(&did.KeyAgreement)
	case VerificationRelationship_capabilityInvocation:
		methodIDs = lazyGet(&did.CapabilityInvocation)
	case VerificationRelationship_capabilityDelegation:
		methodIDs = lazyGet(&did.CapabilityDelegation)
	default:
		err = fmt.Errorf("invalid verification relationship: %v", r)
	}
	return
}

// Verifications is a list of verification
type Verifications []*Verification

// NewVerification build a new verification to be
// attached to a identifier document
func NewVerification(
	id string,
	pubKeyType string,
	controller string,
	pubKey []byte,
	relationships []VerificationRelationship,
	contexts []string,
) Verification {
	return Verification{
		Context: contexts,
		Method: &VerificationMethod{
			Id:              id,
			Type:            pubKeyType,
			Controller:      controller,
			PublicKeyBase58: base58.Encode(pubKey),
		},
		Relationships: relationships,
	}
}

// GetBytes is a helper for serializing
func (did Verification) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&did))
}

// Services are a list of services
type Services []*Service

// NewService creates a new service
func NewService(id string, serviceType string, serviceEndpoint string) Service {
	return Service{
		Id:              id,
		Type:            serviceType,
		ServiceEndpoint: serviceEndpoint,
	}
}
