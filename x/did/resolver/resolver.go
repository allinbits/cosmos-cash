package resolver

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/allinbits/cosmos-cash/x/did/types"
)

// ResolveRepresentation resolve a did document with a specific representation
func ResolveRepresentation(ctx client.Context, did string, opts ResolutionOption) (drr DidResolutionReply) {
	drr.ResolutionMetadata = ResolutionOk(opts.Accept)
	// fail if the content type is not recognized
	if _, ok := resolutionContentTypes[opts.Accept]; !ok {
		drr.ResolutionMetadata = ResolutionErr(ResolutionRepresentationNotSupported)
		return
	}
	// fail if the did is not valid
	if !types.IsValidDID(did) {
		drr.ResolutionMetadata = ResolutionErr(ResolutionInvalidDID)
		return
	}
	// now check the type of did method, if it did:key then generate the did doc
	if strings.HasPrefix(did, types.DidKeyPrefix) {
		didDoc, didMeta, err := resolveAccountDID(did)
		if err != nil {
			drr.ResolutionMetadata = ResolutionErr(ResolutionInvalidDID)
			return
		}
		drr.Document = didDoc
		drr.Metadata = didMeta
		return

	}

	// fail if it is not found
	qc := types.NewQueryClient(ctx)
	qr, err := qc.DidDocument(context.Background(), &types.QueryDidDocumentRequest{Id: did})
	if err != nil {
		drr.ResolutionMetadata = ResolutionErr(ResolutionNotFound)
		return
	}
	// build the resolution
	drr.Document = qr.DidDocument
	drr.Metadata = qr.DidMetadata
	return
}

// resolveAccountDID generates a DID document from an address
func resolveAccountDID(did string) (didDoc types.DidDocument, didMeta types.DidMetadata, err error) {
	didPieaces := strings.Split(did, ":")
	// account is the last part
	account := didPieaces[len(didPieaces)-1]
	// compose the metadata
	didMeta = types.NewDidMetadata([]byte(account), time.Now())
	// compose the did document
	// TODO: hexAddress should be used here
	didDoc, err = types.NewDidDocument(did, types.WithVerifications(
		types.NewVerification(
			types.NewVerificationMethod(
				fmt.Sprint(did, "#", account),
				"CosmosAccountAddress",
				did,
				account,
			),
			[]string{
				types.Authentication,
				types.KeyAgreement,
				types.AssertionMethod,
				types.CapabilityInvocation,
				types.CapabilityDelegation,
			},
			nil,
		),
	))
	return
}
