package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewIdentifier constructs a new Identifier
func NewIdentifier(id string, authentication Authentications) (DidDocument, error) {
	return DidDocument{
		Context:        "https://www.w3.org/ns/did/v1",
		Id:             id,
		Authentication: authentication,
	}, nil
}

// GetBytes is a helper for serialising
func (did DidDocument) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&did))
}

type Authentications []*Authentication

func NewAuthentication(
	id string,
	pubKeyType string,
	controller string,
	encodedValue string,
) Authentication {
	return Authentication{
		Id:         id,
		Type:       pubKeyType,
		Controller: controller,
		PublicKey:  encodedValue,
	}
}

type Services []*Service

func NewService(id string, serviceType string, serviceEndpoint string) Service {
	return Service{
		Id:              id,
		Type:            serviceType,
		ServiceEndpoint: serviceEndpoint,
	}
}
