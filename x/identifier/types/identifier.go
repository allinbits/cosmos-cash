package types

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// A verification relationship expresses the relationship between the DID subject and a verification method.
// This enum is used to
// cfr. https://www.w3.org/TR/did-core/#verification-relationships
const (
	RelationshipAuthentication       = "authentication"       // https://www.w3.org/TR/did-core/#authentication
	RelationshipAssertionMethod      = "assertionMethod"      // https://www.w3.org/TR/did-core/#assertion
	RelationshipKeyAgreement         = "keyAgreement"         // https://www.w3.org/TR/did-core/#key-agreement
	RelationshipCapabilityInvocation = "capabilityInvocation" // https://www.w3.org/TR/did-core/#capability-invocation
	RelationshipCapabilityDelegation = "capabilityDelegation" // https://www.w3.org/TR/did-core/#capability-delegation
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
	return strings.TrimSpace(input) == ""
}

type IdentifierOption func(*DidDocument) error

func WithVerifications(verifications []*Verification) IdentifierOption {
	return func(did *DidDocument) error {
		return did.AddVerifications(verifications...)
	}
}

func WithServices(services []*Service) IdentifierOption {
	return func(did *DidDocument) error {
		return did.AddServices(services...)
	}
}

func WithController(controller string) IdentifierOption {
	return func(did *DidDocument) (err error) {
		// if the controller is not set return error
		if !IsValidDID(controller) {
			err = fmt.Errorf("the document did %s is not compliant with the specification: cfr https://www.w3.org/TR/did-core/#did-syntax", controller)
		}
		return
	}
}

// NewIdentifier constructs a new Identifier
func NewIdentifier(id string, options ...IdentifierOption) (did DidDocument, err error) {

	if !IsValidDID(id) {
		err = fmt.Errorf("the document did %s is not compliant with the specification: cfr https://www.w3.org/TR/did-core/#did-syntax", id)
		return
	}

	did = DidDocument{
		Context: []string{contextDIDBase},
		Id:      id,
	}

	for _, fn := range options {
		if err = fn(&did); err != nil {
			return
		}
	}

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
		if !IsValidDID(v.Method.Controller) {
			err = fmt.Errorf("controller not set for verification method %s, consider using the document did as controller: %s", v.Method.Id, did.Id)
			return
		}

		// check for empty method type
		if IsEmpty(v.Method.Type) {
			err = fmt.Errorf("type not set for verification method %s", v.Method.Id)
			return
		}

		// check for empty publickey
		if IsEmpty(v.Method.PublicKeyBase58) {
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
		did.SetRelationships(v.Method.Id, v.Relationships...)
	}
	return
}

// RevokeVerification revoke a verification method
// and all relationships associated with it
func (did *DidDocument) RevokeVerification(methodID string) {

	del := func(x int) {
		lastIdx := len(did.VerificationMethods) - 1
		switch lastIdx {
		case 0:
			did.VerificationMethods = []*VerificationMethod{}
		case x:
			did.VerificationMethods = did.VerificationMethods[:lastIdx]
		default:
			did.VerificationMethods[x] = did.VerificationMethods[lastIdx]
			did.VerificationMethods = did.VerificationMethods[:lastIdx]
		}
	}

	for i, vm := range did.VerificationMethods {
		if vm.Id == methodID {
			del(i)
			break
		}
	}
}

// SetRelationships for a did document
func (did *DidDocument) SetRelationships(methodID string, relationships ...string) {

	// XXX horrible but sometimes life is like that
	del := func(relName string, x int) {
		lastIdx := len(did.VerificationMethods) - 1
		relationships := did.VerificationRelationships[relName]
		switch lastIdx {
		case 0: // remove the relationships since there is no elements left
			delete(did.VerificationRelationships, relName)
		case x: // if it's at the last position, just drop the last position
			relationships.Labels = relationships.Labels[:lastIdx]
		default: // swap and drop last position
			relationships.Labels[x] = relationships.Labels[lastIdx]
			relationships.Labels = relationships.Labels[:lastIdx]
		}
	}

	// first remove existing
	for _, relationship := range did.VerificationRelationships {
		for i, mID := range relationship.Labels {
			if mID == methodID {
				del(mID, i)
			}
		}
	}

	// then assign them
	for _, r := range relationships {
		if mIDs, exists := did.VerificationRelationships[r]; !exists {
			mIDs = &DidDocument_VerificationRelationships{
				Labels: []string{methodID},
			}
		} else {
			mIDs.Labels = append(mIDs.Labels, methodID)
		}
	}
}

// ControllerInRelationships verifies if a controller did
// exists for at least one of the relationships in the did document
func (did DidDocument) ControllerInRelationships(
	contoller string,
	relationships ...string) bool {
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
		relationships, exists := did.VerificationRelationships[r]
		if !exists {
			return false
		}

		for _, k := range relationships.Labels {
			if _, found := keyController[k]; found {
				return true
			}
		}
	}
	return false
}

func (did *DidDocument) AddServices(services ...*Service) (err error) {
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
	return
}

// GetBytes is a helper for serializing
func (did DidDocument) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&did))
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
	relationships []string,
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
