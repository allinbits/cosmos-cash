package types

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/crypto/blake2b"
)

type VerificationRelationship int

// A verification relationship expresses the relationship between the DID subject and a verification method.
// This enum is used to
// cfr. https://www.w3.org/TR/did-core/#verification-relationships
const (
	Authentication       = "authentication"       // https://www.w3.org/TR/did-core/#authentication
	AssertionMethod      = "assertionMethod"      // https://www.w3.org/TR/did-core/#assertion
	KeyAgreement         = "keyAgreement"         // https://www.w3.org/TR/did-core/#key-agreement
	CapabilityInvocation = "capabilityInvocation" // https://www.w3.org/TR/did-core/#capability-invocation
	CapabilityDelegation = "capabilityDelegation" // https://www.w3.org/TR/did-core/#capability-delegation
)

const (
	authentication VerificationRelationship = iota
	assertionMethod
	keyAgreement
	capabilityInvocation
	capabilityDelegation
)

// VerificationRelationships are the supported list of verification relationships
var VerificationRelationships = map[string]VerificationRelationship{
	"authentication":       authentication,
	"assertionMethod":      assertionMethod,
	"keyAgreement":         keyAgreement,
	"capabilityInvocation": capabilityInvocation,
	"capabilityDelegation": capabilityDelegation,
}

// verificationRelationships retrieve the pointer to the verification relationship
// if it exists, otherwise returns nil
func (didDoc *DidDocument) getRelationships(rel VerificationRelationship) *[]string {
	switch rel {
	case authentication:
		return &didDoc.Authentication
	case assertionMethod:
		return &didDoc.AssertionMethod
	case keyAgreement:
		return &didDoc.KeyAgreement
	case capabilityInvocation:
		return &didDoc.CapabilityInvocation
	case capabilityDelegation:
		return &didDoc.CapabilityDelegation
	default:
		return nil
	}
}

// parseRelationshipLabels parse relationships labels to a slice of VerificationRelationship
// making sure that the relationsips are not repeated
func parseRelationshipLabels(relNames ...string) (vrs []VerificationRelationship, err error) {
	names := distinct(relNames)
	vrs = make([]VerificationRelationship, len(names))
	for i, vrn := range distinct(relNames) {
		vr, validName := VerificationRelationships[vrn]
		if !validName {
			err = sdkerrors.Wrapf(ErrInvalidInput, "unsupported verification relationship %s", vrn)
			return
		}
		vrs[i] = vr
	}
	return
}

/**
Regexp generated using this ABNF specs and using https://abnf.msweet.org/index.php

did-url            = did path-abempty [ "?" query ] [ "#" fragment ]
did                = "did:" method-name ":" method-specific-id
method-name        = 1*method-char
method-char        = %x61-7A / DIGIT
method-specific-id = *( *idchar ":" ) 1*idchar
idchar             = ALPHA / DIGIT / "." / "-" / "_" / pct-encoded
pct-encoded        = "%" HEXDIG HEXDIG
query              = *( pchar / "/" / "?" )
fragment           = *( pchar / "/" / "?" )
path-abempty       = *( "/" segment )
segment            = *pchar
unreserved         = ALPHA / DIGIT / "-" / "." / "_" / "~"
pchar              = unreserved / pct-encoded / sub-delims / ":" / "@"
sub-delims         = "!" / "$" / "&" / "'" / "(" / ")"
                 / "*" / "+" / "," / ";" / "="
*/

const (
	contextDIDBase            = "https://www.w3.org/ns/did/v1"
	didValidationRegexpStr    = `^did\:[a-z0-9]+\:(([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])*\:)*([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])+$`
	didURLValidationRegexpStr = `^did\:[a-z0-9]+\:(([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])*\:)*([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])+(/(([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)*)*(\?(((([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)|/|\?)*))?(#(((([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)|/|\?)*))?$`
	rfc3986RegexpStr          = `^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\?([^#]*))?(#(.*))?$`
)

var (
	didValidationRegexp    = regexp.MustCompile(didValidationRegexpStr)
	didURLValidationRegexp = regexp.MustCompile(didURLValidationRegexpStr)
	rfc3986Regexp          = regexp.MustCompile(rfc3986RegexpStr)
)

// DID format a DID from a method specific did
// cfr.https://www.w3.org/TR/did-core/#did
func DID(didMethodSpecificDidDocument string) string {
	return fmt.Sprint(DidPrefix, didMethodSpecificDidDocument)
}

func DIDKey(didMethodSpecificDidDocument string) string {
	return fmt.Sprint(DidKeyPrefix, didMethodSpecificDidDocument)
}

// BlockchainAccountID return the account of the user with the chain id postfixed
// https://w3c.github.io/did-spec-registries/#blockchainAccountId
func BlockchainAccountID(account string) string {
	return fmt.Sprint(account, "")
}

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
// (cfr https://datatracker.ietf.org/doc/html/rfc3986#page-50)
func IsValidRFC3986Uri(input string) bool {
	return rfc3986Regexp.MatchString(input)
}

// IsValidDIDDocument tells if a DID document is valid,
// that is if it has the default context and a valid subject
func IsValidDIDDocument(didDoc *DidDocument) bool {
	if !IsValidDID(didDoc.Id) {
		return false
	}
	for _, c := range didDoc.Context {
		if c == contextDIDBase {
			return true
		}
	}
	return false
}

// ValidateVerification perform basic validation on a verification struct
// optionally validating the validation method controller against a list
// of allowed controllers.
// in case of error returns an cosmos-sdk wrapped error
// XXX: this pattern creates a ambiguous semantic (but maybe is not too severe (use WithCredentials and array of credentials))
func ValidateVerification(v *Verification, allowedControllers ...string) (err error) {
	if v == nil {
		err = sdkerrors.Wrap(ErrInvalidInput, "verification is not defined")
		return
	}
	// verify that the method id is correct
	if !IsValidDIDURL(v.Method.Id) {
		err = sdkerrors.Wrapf(ErrInvalidDIDURLFormat, "verification method id: %v", v.Method.Id)
		return
	}

	// if the controller is not set return error
	if !IsValidDID(v.Method.Controller) {
		err = sdkerrors.Wrapf(ErrInvalidDIDFormat, "verification method controller %v", v.Method.Controller)
		return
	}

	// check for empty method type
	if IsEmpty(v.Method.Type) {
		err = sdkerrors.Wrapf(ErrInvalidInput, "verification method type not set for verification method %s", v.Method.Id)
		return
	}

	// check for empty publickey
	if IsEmpty(v.Method.BlockchainAccountID) {
		err = sdkerrors.Wrapf(ErrInvalidInput, "public key not set for verification method %s", v.Method.Id)
		return
	}

	// check that there is at least a relationship
	if len(v.Relationships) == 0 {
		err = sdkerrors.Wrap(ErrEmptyRelationships, "at least a verification relationship is required")
		return
	}
	return
}

// ValidateService performs basic on a service struct
func ValidateService(s *Service) (err error) {
	if s == nil {
		err = sdkerrors.Wrap(ErrInvalidInput, "service is not defined")
		return
	}
	// verify that the id is not empty and is a valid url (according to RFC3986)
	if IsEmpty(s.Id) {
		err = sdkerrors.Wrap(ErrInvalidInput, "service id cannot be empty")
		return
	}

	if !IsValidRFC3986Uri(s.Id) {
		err = sdkerrors.Wrapf(ErrInvalidRFC3986UriFormat, "service id %s is not a valid RFC3986 uri", s.Id)
		return
	}

	// verify that the endpoint is not empty and is a valid url (according to RFC3986)
	if IsEmpty(s.ServiceEndpoint) {
		err = sdkerrors.Wrap(ErrInvalidInput, "service endpoint cannot be empty;")
		return
	}

	if !IsValidRFC3986Uri(s.ServiceEndpoint) {
		err = sdkerrors.Wrapf(ErrInvalidRFC3986UriFormat, "service endpoint %s is not a valid RFC3986 uri", s.ServiceEndpoint)
		return
	}

	// check that the service type is not empty
	if IsEmpty(s.Type) {
		err = sdkerrors.Wrap(ErrInvalidInput, "service type cannot be empty;")
		return
	}

	return
}

// IsEmpty tells if the trimmed input is empty
func IsEmpty(input string) bool {
	return strings.TrimSpace(input) == ""
}

// DidDocumentOption implements variadic pattern for optional did document fields
type DidDocumentOption func(*DidDocument) error

// WithVerifications add optional verifications
func WithVerifications(verifications ...*Verification) DidDocumentOption {
	return func(did *DidDocument) error {
		return did.AddVerifications(verifications...)
	}
}

//WithServices add optional services
func WithServices(services ...*Service) DidDocumentOption {
	return func(did *DidDocument) error {
		return did.AddServices(services...)
	}
}

// WithControllers add optional did controller
func WithControllers(controllers ...string) DidDocumentOption {
	return func(did *DidDocument) (err error) {
		return did.SetControllers(controllers...)
	}
}

// NewDidDocument constructs a new DidDocument
func NewDidDocument(id string, options ...DidDocumentOption) (did DidDocument, err error) {

	if !IsValidDID(id) {
		err = sdkerrors.Wrapf(ErrInvalidDIDFormat, "did %s", id)
		return
	}

	did = DidDocument{
		Context: []string{contextDIDBase},
		Id:      id,
	}
	// apply all the options
	for _, fn := range options {
		if err = fn(&did); err != nil {
			return
		}
	}
	return
}

// SetControllers replace the controllers in the did document
func (didDoc *DidDocument) SetControllers(controllers ...string) error {
	if controllers == nil {
		didDoc.Controller = controllers
		return nil
	}
	dc := distinct(controllers)
	for _, c := range dc {
		if !IsValidDID(c) {
			return sdkerrors.Wrapf(ErrInvalidDIDFormat, "did document controller validation error '%s'", c)
		}
	}
	didDoc.Controller = dc
	return nil
}

// AddVerifications add one or more verification method and relations to a did document
func (didDoc *DidDocument) AddVerifications(verifications ...*Verification) (err error) {
	// verify that there are no duplicates in method ids
	index := make(map[string]struct{}, len(didDoc.VerificationMethods))
	// load existing verifications if any
	for _, v := range didDoc.VerificationMethods {
		index[v.Id] = struct{}{}
	}

	// loop through the verifications and look for problems
	for _, v := range verifications {
		// perform base validation checks
		if err = ValidateVerification(v); err != nil {
			return
		}

		// verify that there are no duplicates in method ids
		if _, found := index[v.Method.Id]; found {
			err = sdkerrors.Wrapf(ErrInvalidInput, "duplicated verification method id %s", v.Method.Id)
			return
		}
		index[v.Method.Id] = struct{}{}

		// first add the method to the list of methods
		didDoc.VerificationMethods = append(didDoc.VerificationMethods, v.GetMethod())

		// now add the relationships
		vrs, err := parseRelationshipLabels(v.Relationships...)
		if err != nil {
			return err
		}
		didDoc.setRelationships(v.Method.Id, vrs...)

		// update context
		didDoc.Context = union(didDoc.Context, v.Context)

	}
	return
}

// RevokeVerification revoke a verification method
// and all relationships associated with it
func (didDoc *DidDocument) RevokeVerification(methodID string) error {

	del := func(x int) {
		lastIdx := len(didDoc.VerificationMethods) - 1
		switch lastIdx {
		case 0:
			didDoc.VerificationMethods = nil
		case x:
			didDoc.VerificationMethods = didDoc.VerificationMethods[:lastIdx]
		default:
			didDoc.VerificationMethods[x] = didDoc.VerificationMethods[lastIdx]
			didDoc.VerificationMethods = didDoc.VerificationMethods[:lastIdx]
		}
	}

	// remove relationships
	didDoc.setRelationships(methodID)

	// now remove the method
	for i, vm := range didDoc.VerificationMethods {
		if vm.Id == methodID {
			del(i)
			return nil
		}
	}
	return sdkerrors.Wrapf(ErrVerificationMethodNotFound, "verification method id: %v", methodID)
}

// SetVerificationRelationships for a did document
func (didDoc *DidDocument) SetVerificationRelationships(methodID string, relationships ...string) error {
	// verify that the method id is correct
	if !IsValidDIDURL(methodID) {
		return sdkerrors.Wrapf(ErrInvalidDIDURLFormat, "verification method id: %v", methodID)
	}
	// check that there is at least a relationship
	if len(relationships) == 0 {
		return sdkerrors.Wrap(ErrEmptyRelationships, "at least a verification relationship is required")
	}
	// check that the provided relationships are valid
	vrs, err := parseRelationshipLabels(relationships...)
	if err != nil {
		return err
	}
	// update the relationships
	didDoc.setRelationships(methodID, vrs...)
	return nil
}

// setRelationships overwrite relationships for a did document
func (didDoc *DidDocument) setRelationships(methodID string, relationships ...VerificationRelationship) {

	// first remove existing relationships
	for _, vr := range VerificationRelationships {
		vrs := didDoc.getRelationships(vr)
		for i, vmID := range *vrs {
			if vmID == methodID {
				lastIdx := len(*vrs) - 1 // get the last index of the current relationship list
				switch lastIdx {
				case 0: // remove the relationships since there is no elements left
					*vrs = nil
				case i: // if it's at the last position, just drop the last position
					*vrs = (*vrs)[:lastIdx]
				default: // swap and drop last position
					(*vrs)[i] = (*vrs)[lastIdx]
					(*vrs) = (*vrs)[:lastIdx]
				}
			}
		}
	}

	// then assign the new ones
	for _, vr := range relationships {
		vrs := didDoc.getRelationships(vr)
		*vrs = append(*vrs, methodID)
	}
}

// GetVerificationRelationships returns the relationships associated with the
// verification method id.
func (didDoc DidDocument) GetVerificationRelationships(methodID string) []string {
	relationships := []string{}
	for vrn, vr := range VerificationRelationships {
		for _, vmID := range *didDoc.getRelationships(vr) {
			if vmID == methodID {
				relationships = append(relationships, vrn)
			}
		}
	}
	return relationships
}

// HasRelationship verifies if a controller did
// exists for at least one of the relationships in the did document
func (didDoc DidDocument) HasRelationship(
	signer string,
	relationships ...string,
) bool {
	// first check if the controller exists
	for _, vm := range didDoc.VerificationMethods {
		if vm.BlockchainAccountID != BlockchainAccountID(signer) {
			continue
		}

		vrs := didDoc.GetVerificationRelationships(vm.Id)
		if len(intersection(vrs, relationships)) > 0 {
			return true
		}
	}
	return false
}

// AddServices add services to a did document
func (didDoc *DidDocument) AddServices(services ...*Service) (err error) {
	if didDoc.Services == nil {
		didDoc.Services = []*Service{}
	}

	// used to check duplicates
	index := make(map[string]struct{}, len(didDoc.Services))

	// load existing services
	for _, s := range didDoc.Services {
		index[s.Id] = struct{}{}
	}

	// services must be unique
	for _, s := range services {
		if err = ValidateService(s); err != nil {
			return
		}

		// verify that there are no duplicates in method ids
		if _, found := index[s.Id]; found {
			err = sdkerrors.Wrapf(ErrInvalidInput, "duplicated verification method id %s", s.Id)
			return
		}
		index[s.Id] = struct{}{}

		didDoc.Services = append(didDoc.Services, s)
	}
	return
}

// DeleteService delete an existing service from a did document
func (didDoc *DidDocument) DeleteService(serviceID string) {
	del := func(x int) {
		lastIdx := len(didDoc.Services) - 1
		switch lastIdx {
		case 0: // remove the relationships since there is no elements left
			didDoc.Services = nil
		case x: // if it's at the last position, just drop the last position
			didDoc.Services = didDoc.Services[:lastIdx]
		default: // swap and drop last position
			didDoc.Services[x] = didDoc.Services[lastIdx]
			didDoc.Services = didDoc.Services[:lastIdx]
		}
	}

	for i, s := range didDoc.Services {
		if s.Id == serviceID {
			del(i)
			break
		}
	}
}

// GetBytes is a helper for serializing
func (didDoc DidDocument) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&didDoc))
}

// Verifications is a list of verification
type Verifications []*Verification

// NewVerification build a new verification to be
// attached to a did document
func NewVerification(
	method VerificationMethod,
	relationships []string,
	contexts []string,
) *Verification {
	return &Verification{
		Context:       contexts,
		Method:        &method,
		Relationships: relationships,
	}
}

// NewVerificationMethod build a new verification method
// TODO: this only uses BlockchainAccountID
func NewVerificationMethod(id, keyType, controller, key string) VerificationMethod {
	return VerificationMethod{
		Id:                  id,
		Type:                keyType,
		Controller:          controller,
		BlockchainAccountID: key,
	}
}

// GetBytes is a helper for serializing
func (did Verification) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&did))
}

// Services are a list of services
type Services []*Service

// NewService creates a new service
func NewService(id string, serviceType string, serviceEndpoint string) *Service {
	return &Service{
		Id:              id,
		Type:            serviceType,
		ServiceEndpoint: serviceEndpoint,
	}
}

// NewDidMetadata returns a DidMetadata strcut that has equals created and updated date,
// and with deactivated field set to false
func NewDidMetadata(versionData []byte, created time.Time) DidMetadata {
	// compute the hash from the version data
	txH := blake2b.Sum256(versionData)
	return DidMetadata{
		VersionId:   hex.EncodeToString(txH[:]),
		Created:     &created,
		Updated:     &created,
		Deactivated: false,
	}
}

// union perform union, distinct amd sort operation between two slices
// duplicated element in list a are
func union(a, b []string) []string {
	if len(b) == 0 {
		return a
	}
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	for _, item := range b {
		if _, ok := m[item]; !ok {
			m[item] = struct{}{}
		}
	}
	u := make([]string, 0, len(m))
	for k := range m {
		u = append(u, k)
	}
	sort.Strings(u)
	return u
}

func intersection(a, b []string) []string {
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	i := []string{}
	for _, item := range distinct(b) {
		if _, ok := m[item]; ok {
			i = append(i, item)
		}
	}
	sort.Strings(i)
	return i
}

// distinct remove duplicates and sorts from a list of strings
func distinct(a []string) []string {
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	d := make([]string, 0, len(m))
	for k := range m {
		d = append(d, k)
	}
	sort.Strings(d)
	return d
}
