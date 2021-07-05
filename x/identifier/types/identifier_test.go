package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDID(t *testing.T) {

	tests := []struct {
		identifier string
		want       string
	}{
		{
			"subject",
			"did:cash:subject",
		},
		{
			"",
			"did:cash:",
		},
		{
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			"did:cash:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDID#", i), func(t *testing.T) {
			if got := DID(tt.identifier); got != tt.want {
				t.Errorf("DID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidDID(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"did:cash:subject", true},
		{"did:cash:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75", true},
		{"did:cash:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75#key-1", false},
		{"did:subject", false},
		{"DID:cash:subject", false},
		{"d1d:cash:subject", false},
		{"d1d:#:subject", false},
		{"", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestIsValidDID#", i), func(t *testing.T) {
			if got := IsValidDID(tt.input); got != tt.want {
				t.Errorf("IsValidDID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidDIDURL(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{"did:cash:subject", true},
		{"did:cash:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75", true},
		{"did:cash:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75#key-1", true},
		{"did:cash:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75?key=1", true},
		{"did:subject", false},
		{"DID:cash:subject", false},
		{"d1d:cash:subject", false},
		{"d1d:#:subject", false},
		{"", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestIsValidDIDURL#", i), func(t *testing.T) {
			if got := IsValidDIDURL(tt.input); got != tt.want {
				t.Errorf("IsValidDIDURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidRFC3986Uri(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidRFC3986Uri(tt.args.input); got != tt.want {
				t.Errorf("IsValidRFC3986Uri() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateVerification(t *testing.T) {
	tests := []struct {
		v       *Verification
		wantErr bool
	}{
		{
			v: NewVerification(
				NewVerificationMethod(
					"did:cash:subject#key-1",
					"EcdsaSecp256k1VerificationKey2019",
					"did:cash:subject",
					"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
				),
				nil,
				nil,
			),
			wantErr: true, // relationships are nil
		},
		{
			v:       nil,
			wantErr: true,
		},
		{
			v: NewVerification(
				NewVerificationMethod(
					"did:cash:subject#key-1",
					"EcdsaSecp256k1VerificationKey2019",
					"did:cash:subject",
					"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
				),
				[]string{RelationshipAssertionMethod},
				nil,
			),
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestValidateVerification#", i), func(t *testing.T) {
			if err := ValidateVerification(tt.v); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateVerification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateService(t *testing.T) {

	tests := []struct {
		s       *Service
		wantErr bool
	}{
		{
			s:       NewService("agent:abc", "DIDCommMessaging", "https://agent.abc/abc"),
			wantErr: false,
		},
		{
			s:       nil,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestValidateService#", i), func(t *testing.T) {
			if err := ValidateService(tt.s); (err != nil) != tt.wantErr {
				t.Errorf("ValidateService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"    a    ", false},
		{"\t", true},
		{"\n", true},
		{"   ", true},
		{"  \t \n", true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestIsEmpty#", i), func(t *testing.T) {
			if got := IsEmpty(tt.input); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIdentifier(t *testing.T) {
	type params struct {
		id      string
		options []IdentifierOption
	}
	tests := []struct {
		params  params
		wantDid DidDocument
		wantErr bool
	}{
		{
			params: params{
				"did:cash:subject",
				[]IdentifierOption{
					WithVerifications(
						&Verification{
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
								RelationshipKeyAgreement, // test duplicated relationship
							},
							&VerificationMethod{
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						},
					),
					WithVerifications( // multiple verifications in separate entity
						&Verification{
							[]string{
								RelationshipAuthentication,
							},
							&VerificationMethod{
								"did:cash:subject#key-2",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						},
					),
					WithServices(&Service{
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					}),
					WithControllers("did:cash:controller-1"),
				},
			},
			wantDid: DidDocument{
				[]string{
					"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
					contextDIDBase,
				},
				"did:cash:subject",
				[]string{"did:cash:controller-1"},
				[]*VerificationMethod{
					{
						"did:cash:subject#key-1",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
					{
						"did:cash:subject#key-2",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
				},
				[]*Service{
					{
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					},
				},
				map[string]*DidDocument_VerificationRelationships{
					RelationshipAuthentication: {
						Labels: []string{"did:cash:subject#key-1", "did:cash:subject#key-2"},
					},
					RelationshipKeyAgreement: {
						Labels: []string{"did:cash:subject#key-1"},
					},
				},
			},
			wantErr: false,
		},
		{
			params: params{
				"did:cash:subject",
				[]IdentifierOption{
					WithVerifications(
						&Verification{
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
							},
							&VerificationMethod{
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						},
						&Verification{
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
							},
							&VerificationMethod{
								"did:cash:subject#key-1", // duplicate key
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{},
						},
					),
					WithServices(&Service{
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					}),
				},
			},
			wantDid: DidDocument{},
			wantErr: true, // the error is caused by duplicated verification method id
		},
		{
			params: params{
				"did:cash:subject",
				[]IdentifierOption{
					WithVerifications(
						&Verification{
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
							},
							&VerificationMethod{
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						},
					),
					WithServices(
						&Service{
							"agent:xyz",
							"DIDCommMessaging",
							"https://agent.xyz/1234",
						},
						&Service{
							"agent:xyz",
							"DIDCommMessaging",
							"https://agent.xyz/1234",
						},
					),
				},
			},
			wantDid: DidDocument{},
			wantErr: true, //duplicated service id
		},
		{
			wantErr: true, // invalid did
			params: params{
				id:      "something not right",
				options: []IdentifierOption{},
			},
			wantDid: DidDocument{},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestNewIdentifier#", i), func(t *testing.T) {
			gotDid, err := NewIdentifier(tt.params.id, tt.params.options...)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestNewIdentifier#%v", i)
				return
			}

			require.Nil(t, err, "test: TestNewIdentifier#%v", i)
			assert.Equal(t, tt.wantDid, gotDid)
		})
	}
}

func TestDidDocument_SetControllers(t *testing.T) {

	tests := []struct {
		malleate    func() DidDocument
		controllers []string
		wantErr     bool
	}{
		{
			func() DidDocument {
				dd, _ := NewIdentifier("did:cash:subject",
					WithControllers(
						"did:cash:controller-1",
						"did:cash:controller-2",
						"did:cash:controller-3",
						"did:cash:controller-4",
						"did:cash:controller-4", // duplicate controllers
						"did:cash:controller-4",
						"did:cash:controller-4",
					),
				)
				return dd
			},
			[]string{
				"did:cash:controller-1",
				"did:cash:controller-4",
			},
			false,
		},
		{
			func() DidDocument {
				dd, _ := NewIdentifier("did:cash:subject",
					WithControllers(
						"did:cash:controller-1",
						"did:cash:controller-2",
						"did:cash:controller-3",
						"did:cash:controller-4",
					),
				)
				return dd
			},
			[]string{
				"did:cash:controller-1",
				"not a did:cash:controller-4",
			},
			true, // invalid controller did
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_SetControllers#", i), func(t *testing.T) {
			didDoc := tt.malleate()
			err := didDoc.SetControllers(tt.controllers...)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestDidDocument_SetControllers#%v", i)
				return
			}

			require.Nil(t, err, "test: TestDidDocument_SetControllers#%v", i)
			assert.Equal(t, tt.controllers, didDoc.Controller)

		})
	}
}

func TestDidDocument_AddVerifications(t *testing.T) {
	type params struct {
		malleate      func() DidDocument // build a did document
		verifications []*Verification    // input list of verifications
	}
	tests := []struct {
		params  params
		wantDid DidDocument // expected result
		wantErr bool
	}{
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject")
					return d
				},
				[]*Verification{
					NewVerification(
						NewVerificationMethod(
							"did:cash:subject#key-1",
							"EcdsaSecp256k1VerificationKey2019",
							"did:cash:subject",
							"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
						),
						[]string{
							RelationshipAuthentication,
							RelationshipKeyAgreement,
						},
						nil,
					),
					NewVerification(
						NewVerificationMethod(
							"did:cash:subject#key-2",
							"EcdsaSecp256k1VerificationKey2019",
							"did:cash:subject",
							"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
						),
						[]string{
							RelationshipAuthentication,
							RelationshipCapabilityInvocation,
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					),
				},
			},
			wantDid: DidDocument{
				[]string{
					"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
					contextDIDBase,
				},
				"did:cash:subject",
				nil,
				[]*VerificationMethod{
					{
						"did:cash:subject#key-1",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
					{
						"did:cash:subject#key-2",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
				},
				nil,
				map[string]*DidDocument_VerificationRelationships{
					RelationshipAuthentication: {
						Labels: []string{"did:cash:subject#key-1", "did:cash:subject#key-2"},
					},
					RelationshipKeyAgreement: {
						Labels: []string{"did:cash:subject#key-1"},
					},
					RelationshipCapabilityInvocation: {
						Labels: []string{"did:cash:subject#key-2"},
					},
				},
			},
		},
		{
			wantErr: true, // duplicated method id
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
								RelationshipKeyAgreement, // test duplicated relationship
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
					))
					return d
				},
				[]*Verification{
					NewVerification(
						NewVerificationMethod(
							"did:cash:subject#key-1",
							"EcdsaSecp256k1VerificationKey2019",
							"did:cash:subject",
							"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
						),
						[]string{
							RelationshipAuthentication,
							RelationshipKeyAgreement,
							RelationshipKeyAgreement, // test duplicated relationship
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					),
				},
			},
			wantDid: DidDocument{},
		},
		{
			wantErr: true, // fail validation
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
								RelationshipKeyAgreement, // test duplicated relationship
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
					))
					return d
				},
				[]*Verification{
					{
						[]string{
							RelationshipAuthentication,
							RelationshipKeyAgreement,
							RelationshipKeyAgreement, // test duplicated relationship
						},
						&VerificationMethod{
							"invalid method url",
							"EcdsaSecp256k1VerificationKey2019",
							"did:cash:subject",
							"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					},
				},
			},
			wantDid: DidDocument{},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_AddVerifications#", i), func(t *testing.T) {
			gotDid := tt.params.malleate()

			err := gotDid.AddVerifications(tt.params.verifications...)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestDidDocument_AddVerifications#%v", i)
				return
			}

			require.Nil(t, err, "test: TestDidDocument_AddVerifications#%v", i)
			assert.Equal(t, tt.wantDid, gotDid)
		})
	}
}

func TestDidDocument_RevokeVerification(t *testing.T) {
	type params struct {
		malleate func() DidDocument // build a did document
		methodID string             // input list of verifications
	}
	tests := []struct {
		params  params
		wantDid DidDocument // expected result
		wantErr bool
	}{
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject",
						WithVerifications(
							&Verification{
								[]string{
									RelationshipAuthentication,
									RelationshipKeyAgreement,
								},
								&VerificationMethod{
									"did:cash:subject#key-1",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								},
								nil,
							}, &Verification{
								[]string{
									RelationshipAuthentication,
									RelationshipCapabilityInvocation,
								},
								&VerificationMethod{
									"did:cash:subject#key-2",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								},
								[]string{
									"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
								},
							},
						),
					)
					return d
				},
				"did:cash:subject#key-2",
			},
			wantDid: DidDocument{
				[]string{
					"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
					contextDIDBase,
				},
				"did:cash:subject",
				nil,
				[]*VerificationMethod{
					{
						"did:cash:subject#key-1",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
				},
				nil,
				map[string]*DidDocument_VerificationRelationships{
					RelationshipAuthentication: {
						Labels: []string{"did:cash:subject#key-1"},
					},
					RelationshipKeyAgreement: {
						Labels: []string{"did:cash:subject#key-1"},
					},
				},
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject",
						WithVerifications(
							&Verification{
								[]string{
									RelationshipAuthentication,
									RelationshipKeyAgreement,
								},
								&VerificationMethod{
									"did:cash:subject#key-1",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								},
								nil,
							},
						),
					)
					return d
				},
				"did:cash:subject#key-1",
			},
			wantDid: DidDocument{
				[]string{
					contextDIDBase,
				},
				"did:cash:subject",
				nil,
				nil,
				nil,
				map[string]*DidDocument_VerificationRelationships{},
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject",
						WithVerifications(
							&Verification{
								[]string{
									RelationshipAuthentication,
									RelationshipKeyAgreement,
								},
								&VerificationMethod{
									"did:cash:subject#key-1",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								},
								nil,
							}, &Verification{
								[]string{
									RelationshipAuthentication,
									RelationshipCapabilityInvocation,
								},
								&VerificationMethod{
									"did:cash:subject#key-2",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								},
								nil,
							},
							&Verification{
								[]string{
									RelationshipAuthentication,
									RelationshipKeyAgreement,
									RelationshipAssertionMethod,
								},
								&VerificationMethod{
									"did:cash:subject#key-3",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								},
								nil,
							},
						),
					)
					return d
				},
				"did:cash:subject#key-2",
			},
			wantDid: DidDocument{
				[]string{
					contextDIDBase,
				},
				"did:cash:subject",
				nil,
				[]*VerificationMethod{
					{
						"did:cash:subject#key-1",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
					{
						"did:cash:subject#key-3",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
				},
				nil,
				map[string]*DidDocument_VerificationRelationships{
					RelationshipAuthentication: {
						Labels: []string{"did:cash:subject#key-1", "did:cash:subject#key-3"},
					},
					RelationshipKeyAgreement: {
						Labels: []string{"did:cash:subject#key-1", "did:cash:subject#key-3"},
					},
					RelationshipAssertionMethod: {
						Labels: []string{"did:cash:subject#key-3"},
					},
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_RevokeVerification#", i), func(t *testing.T) {
			gotDid := tt.params.malleate()

			gotDid.RevokeVerification(tt.params.methodID)

			assert.Equal(t, tt.wantDid, gotDid)
		})
	}
}

func TestDidDocument_SetVerificationRelationships(t *testing.T) {
	type params struct {
		malleate      func() DidDocument
		methodID      string
		relationships []string
	}
	tests := []struct {
		params  params
		wantDid DidDocument // expected result
		wantErr bool
	}{
		{
			wantErr: true, // empty relationships
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewIdentifier("did:cash:subject")
					return dd
				},
				methodID:      "did:cash:subject#key-1",
				relationships: []string{},
			},
			wantDid: DidDocument{
				[]string{contextDIDBase},
				"did:cash:subject",
				nil,
				nil,
				nil,
				map[string]*DidDocument_VerificationRelationships{},
			},
		},
		{
			wantErr: true, //invalid method id
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewIdentifier("did:cash:subject")
					return dd
				},
				methodID:      "did:cash:subject#key-1 invalid ",
				relationships: []string{},
			},
			wantDid: DidDocument{},
		},
		{
			wantErr: false,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewIdentifier("did:cash:subject", WithVerifications(
						&Verification{
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
							},
							&VerificationMethod{
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{},
						},
					))
					return dd
				},
				methodID: "did:cash:subject#key-1",
				relationships: []string{
					RelationshipAssertionMethod,
					RelationshipAssertionMethod, // test duplicated relationship
					RelationshipAssertionMethod, // test duplicated relationship
					RelationshipAssertionMethod, // test duplicated relationship
				},
			},

			wantDid: DidDocument{
				[]string{contextDIDBase},
				"did:cash:subject",
				nil,
				[]*VerificationMethod{
					{
						"did:cash:subject#key-1",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
				},
				nil,
				map[string]*DidDocument_VerificationRelationships{
					RelationshipAssertionMethod: {
						Labels: []string{"did:cash:subject#key-1"},
					},
				},
			},
		},
		{
			wantErr: false, // different delete scenarios
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewIdentifier("did:cash:subject", WithVerifications(
						&Verification{
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
							},
							&VerificationMethod{
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{},
						},
						&Verification{
							[]string{
								RelationshipAuthentication,
							},
							&VerificationMethod{
								"did:cash:subject#key-2",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{},
						},
					))
					return dd
				},
				methodID:      "did:cash:subject#key-1",
				relationships: []string{RelationshipAssertionMethod},
			},
			wantDid: DidDocument{
				[]string{contextDIDBase},
				"did:cash:subject",
				nil,
				[]*VerificationMethod{
					{
						"did:cash:subject#key-1",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
					{
						"did:cash:subject#key-2",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
				},
				nil,
				map[string]*DidDocument_VerificationRelationships{
					RelationshipAuthentication: {
						Labels: []string{"did:cash:subject#key-2"},
					},
					RelationshipAssertionMethod: {
						Labels: []string{"did:cash:subject#key-1"},
					},
				},
			},
		},
		{
			wantErr: false, // different delete scenarios
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewIdentifier("did:cash:subject", WithVerifications(
						&Verification{
							[]string{
								RelationshipAuthentication,
							},
							&VerificationMethod{
								"did:cash:subject#key-2",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{},
						},
						&Verification{
							[]string{
								RelationshipAuthentication,
							},
							&VerificationMethod{
								"did:cash:subject#key-3",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{},
						},
						&Verification{
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
							},
							&VerificationMethod{
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{},
						},
					))
					return dd
				},
				methodID:      "did:cash:subject#key-1",
				relationships: []string{RelationshipAssertionMethod},
			},
			wantDid: DidDocument{
				[]string{contextDIDBase},
				"did:cash:subject",
				nil,
				[]*VerificationMethod{
					{
						"did:cash:subject#key-2",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
					{
						"did:cash:subject#key-3",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
					{
						"did:cash:subject#key-1",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
				},
				nil,
				map[string]*DidDocument_VerificationRelationships{
					RelationshipAuthentication: {
						Labels: []string{"did:cash:subject#key-2", "did:cash:subject#key-3"},
					},
					RelationshipAssertionMethod: {
						Labels: []string{"did:cash:subject#key-1"},
					},
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_SetVerificationRelationships#", i), func(t *testing.T) {
			didDoc := tt.params.malleate()
			err := didDoc.SetVerificationRelationships(tt.params.methodID, tt.params.relationships...)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestDidDocument_SetVerificationRelationships#%v", i)
				return
			}

			require.Nil(t, err, "test: TestDidDocument_SetVerificationRelationships#%v", i)
			assert.Equal(t, tt.wantDid, didDoc)

		})
	}
}

func TestDidDocument_HasRelationship(t *testing.T) {

	type params struct {
		malleate      func() DidDocument
		did           string
		relationships []string
	}
	tests := []struct {
		params                  params
		expectedHasRelationship bool
	}{
		{
			expectedHasRelationship: true,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewIdentifier("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
							},
							nil,
						),
					))
					return dd
				},
				did: "did:cash:subject",
				relationships: []string{
					RelationshipAssertionMethod,
					RelationshipAuthentication,
				},
			},
		},
		{
			expectedHasRelationship: false,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewIdentifier("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
							},
							nil,
						),
						NewVerification(
							NewVerificationMethod(
								"did:cash:controller-1#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:controller-1",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								RelationshipCapabilityDelegation,
							},
							nil,
						),
					))
					return dd
				},
				did: "did:cash:subject",
				relationships: []string{
					RelationshipCapabilityDelegation,
				},
			},
		},
		{
			expectedHasRelationship: false,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewIdentifier("did:cash:subject")
					return dd
				},
				did: "did:cash:subject",
				relationships: []string{
					RelationshipCapabilityDelegation,
				},
			},
		},
		{
			expectedHasRelationship: false,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewIdentifier("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								RelationshipAuthentication,
								RelationshipKeyAgreement,
							},
							nil,
						),
					))
					return dd
				},
				did:           "did:cash:subject",
				relationships: nil,
			},
		},
		{
			expectedHasRelationship: true,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewIdentifier("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								RelationshipAuthentication,
							},
							nil,
						),
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-2",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								RelationshipKeyAgreement,
							},
							nil,
						),
					))
					return dd
				},
				did: "did:cash:subject",
				relationships: []string{
					RelationshipKeyAgreement,
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_SetVerificationRelationships#", i), func(t *testing.T) {
			didDoc := tt.params.malleate()
			gotHasRelationship := didDoc.HasRelationship(tt.params.did, tt.params.relationships...)
			assert.Equal(t, tt.expectedHasRelationship, gotHasRelationship)
		})
	}
}

func TestDidDocument_AddServices(t *testing.T) {
	type params struct {
		malleate func() DidDocument // build a did document
		services []*Service         // input list of verifications
	}
	tests := []struct {
		params  params
		wantDid DidDocument // expected result
		wantErr bool
	}{
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject")
					return d
				},
				[]*Service{
					NewService(
						"agent:abc",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					),
					NewService(
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					),
				},
			},
			wantDid: DidDocument{
				[]string{contextDIDBase},
				"did:cash:subject",
				nil,
				nil,
				[]*Service{
					NewService(
						"agent:abc",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					),
					NewService(
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					),
				},
				nil,
			},
		},
		{
			wantErr: true, // duplicated service id
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject")
					return d
				},
				[]*Service{
					{
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					}, {
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					},
				},
			},
			wantDid: DidDocument{},
		},
		{
			wantErr: true, // fail validation
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject")
					return d
				},
				[]*Service{
					{
						"agent:abc",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					}, {
						"",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					},
				},
			},
			wantDid: DidDocument{},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_AddServices#", i), func(t *testing.T) {
			gotDid := tt.params.malleate()

			err := gotDid.AddServices(tt.params.services...)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestDidDocument_AddServices#%v", i)
				return
			}

			require.Nil(t, err, "test: TestDidDocument_AddServices#%v", i)
			assert.Equal(t, tt.wantDid, gotDid)
		})
	}
}

func TestDidDocument_DeleteService(t *testing.T) {
	type params struct {
		malleate func() DidDocument // build a did document
		methodID string             // input list of verifications
	}
	tests := []struct {
		params  params
		wantDid DidDocument // expected result
		wantErr bool
	}{
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject",
						WithServices(
							&Service{
								"agent:abc",
								"DIDCommMessaging",
								"https://agent.abc/1234",
							},
						),
					)
					return d
				},
				"agent:abc",
			},
			wantDid: DidDocument{
				[]string{contextDIDBase},
				"did:cash:subject",
				nil,
				nil,
				nil,
				nil,
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject",
						WithServices(
							&Service{
								"agent:zyz",
								"DIDCommMessaging",
								"https://agent.abc/1234",
							},
							&Service{
								"agent:abc",
								"DIDCommMessaging",
								"https://agent.abc/1234",
							},
						),
					)
					return d
				},
				"agent:abc",
			},
			wantDid: DidDocument{
				[]string{contextDIDBase},
				"did:cash:subject",
				nil,
				nil,
				[]*Service{
					{
						"agent:zyz",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					},
				},
				nil,
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewIdentifier("did:cash:subject",
						WithServices(
							&Service{
								"agent:zyz",
								"DIDCommMessaging",
								"https://agent.abc/1234",
							},
							&Service{
								"agent:abc",
								"DIDCommMessaging",
								"https://agent.abc/1234",
							},
							&Service{
								"agent:007",
								"DIDCommMessaging",
								"https://agent.abc/007",
							},
						),
					)
					return d
				},
				"agent:abc",
			},
			wantDid: DidDocument{
				[]string{contextDIDBase},
				"did:cash:subject",
				nil,
				nil,
				[]*Service{
					{
						"agent:zyz",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					}, {
						"agent:007",
						"DIDCommMessaging",
						"https://agent.abc/007",
					},
				},
				nil,
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_DeleteService#", i), func(t *testing.T) {
			gotDid := tt.params.malleate()

			gotDid.DeleteService(tt.params.methodID)

			assert.Equal(t, tt.wantDid, gotDid)
		})
	}
}
