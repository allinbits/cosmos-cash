package resolver

import "github.com/allinbits/cosmos-cash/x/did/types"

// Error types for resolution as described in the
// resolution metadata https://www.w3.org/TR/did-core/#did-resolution-metadata
const (
	// ResolutionInvalidDID - the DID supplied to the DID resolution function does not conform to valid syntax.
	ResolutionInvalidDID = "invalidDid"
	// ResolutionNotFound - the DID resolver was unable to find the DID document resulting from this resolution request.
	ResolutionNotFound = "notFound"
	// ResolutionRepresentationNotSupported - the representation requested via the accept input metadata property is not supported
	ResolutionRepresentationNotSupported = "representationNotSupported"
)

// Accepted content types
const (
	ResolutionJSONType   = "application/json"
	ResolutionJSONLDType = "application/ld+json"
)

// resolutionContentTypes for lookups
var resolutionContentTypes = map[string]struct{}{
	ResolutionJSONType:   {},
	ResolutionJSONLDType: {},
}

// ResolutionOption are parametes for the resolver
// https://www.w3.org/TR/did-core/#did-resolution-options
type ResolutionOption struct {
	Accept string `json:"accept,omitempty"`
}

// ResolutionMetadata are info about the resolution
// https://www.w3.org/TR/did-core/#did-resolution-metadata
type ResolutionMetadata struct {
	ContentType     string            `json:"contentType,omitempty"`
	ResolutionError string            `json:"error,omitempty"`
	DidProperties   map[string]string `json:"did,omitempty"`
}

// ResolutionOk helper to get a resolution metadata for a successful resolution
func ResolutionOk(contentType string) ResolutionMetadata {
	return ResolutionMetadata{
		ContentType: contentType,
	}
}

// ResolutionErr helper to get a resolution metadata for an error
func ResolutionErr(err string) ResolutionMetadata {
	return ResolutionMetadata{
		ResolutionError: err,
	}
}

// DidResolutionReply contains the answer to a identifier endpoint
type DidResolutionReply struct {
	Document           types.DidDocument  `json:"didDocument,omitempty"`
	Metadata           types.DidMetadata  `json:"didDocumentMetadata,omitempty"`
	ResolutionMetadata ResolutionMetadata `json:"didResolutionMetadata,omitempty"`
}
