package types

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDID(t *testing.T) {

	tests := []struct {
		did   string
		chain string
		want  string
	}{
		{
			"subject",
			"cash",
			"did:cosmos:cash:subject",
		},
		{
			"",
			"cash",
			"did:cosmos:cash:",
		},
		{
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			"cosmoshub",
			"did:cosmos:cosmoshub:cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDID#", i), func(t *testing.T) {
			if got := DID(tt.chain, tt.did); got != tt.want {
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
	tests := []struct {
		input string
		want  bool
	}{
		{
			"[][àséf",
			true,
		},
		{
			"# \u007e // / / ///// //// // / / ??? ?? 不",
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestIsValidRFC3986Uri#", i), func(t *testing.T) {
			if got := IsValidRFC3986Uri(tt.input); got != tt.want {
				t.Errorf("IsValidRFC3986Uri() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidDIDDocument(t *testing.T) {
	tests := []struct {
		didFn func() *DidDocument
		want  bool
	}{
		{
			func() *DidDocument {
				return &DidDocument{
					Context: []string{contextDIDBase},
					Id:      "did:cosmos:cash:1",
				}
			},
			true, // all good
		},
		{
			func() *DidDocument {
				return &DidDocument{
					Context: []string{},
					Id:      "did:cosmos:cash:1",
				}
			},
			false, // missing context
		},
		{
			func() *DidDocument {
				dd, _ := NewDidDocument("did:cosmos:cash:1")
				return &dd
			},
			true, // all good
		},
		{
			func() *DidDocument {
				dd, _ := NewDidDocument("")
				return &dd
			},
			false, // empty id
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestIsValidDIDDocument#", i), func(t *testing.T) {
			if got := IsValidDIDDocument(tt.didFn()); got != tt.want {
				t.Errorf("TestIsValidDIDDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidDIDMetadata(t *testing.T) {

	tests := []struct {
		didMetaFn func() *DidMetadata
		want      bool
	}{
		{
			func() *DidMetadata {
				now := time.Now()
				return &DidMetadata{
					VersionId: "d95daac05a36f93d1494208d02d1522d758466c62ea6b64c50b78999d2021f51",
					Created:   &now,
				}
			},
			true, // all good
		},
		{
			func() *DidMetadata {
				now := time.Now()
				return &DidMetadata{
					VersionId: "",
					Created:   &now,
				}
			},
			false, // missing version
		},
		{
			func() *DidMetadata {
				now := time.Now()
				return &DidMetadata{
					VersionId: "d95daac05a36f93d1494208d02d1522d758466c62ea6b64c50b78999d2021f51",
					Updated:   &now,
				}
			},
			false, // null created
		},
		{
			func() *DidMetadata {
				var now time.Time
				return &DidMetadata{
					VersionId: "d95daac05a36f93d1494208d02d1522d758466c62ea6b64c50b78999d2021f51",
					Created:   &now,
				}
			},
			false, // zero created
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestIsValidDIDMetadata#", i), func(t *testing.T) {
			if got := IsValidDIDMetadata(tt.didMetaFn()); got != tt.want {
				t.Errorf("TestIsValidDIDMetadata() = %v, want %v", got, tt.want)
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
				[]string{string(AssertionMethod)},
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

func TestNewDidDocument(t *testing.T) {
	type params struct {
		id      string
		options []DidDocumentOption
	}
	tests := []struct {
		params  params
		wantDid DidDocument
		wantErr bool
	}{
		{
			params: params{
				"did:cash:subject",
				[]DidDocumentOption{
					WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								string(Authentication),
								string(KeyAgreement),
								string(KeyAgreement), // test duplicated relationship
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
					),
					WithVerifications( // multiple verifications in separate entity
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-2",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								string(Authentication),
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
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
				Context: []string{
					"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
					contextDIDBase,
				},
				Id:         "did:cash:subject",
				Controller: []string{"did:cash:controller-1"},
				VerificationMethods: []*VerificationMethod{
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
				Services: []*Service{
					{
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
					},
				},
				Authentication: []string{"did:cash:subject#key-1", "did:cash:subject#key-2"},
				KeyAgreement:   []string{"did:cash:subject#key-1"},
			},
			wantErr: false,
		},
		{
			params: params{
				"did:cash:subject",
				[]DidDocumentOption{
					WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1", // duplicate key
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{},
						),
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
				[]DidDocumentOption{
					WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{
								"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
							},
						),
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
				options: []DidDocumentOption{},
			},
			wantDid: DidDocument{},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestNewDidDocument#", i), func(t *testing.T) {
			gotDid, err := NewDidDocument(tt.params.id, tt.params.options...)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestNewDidDocument#%v", i)
				return
			}

			require.Nil(t, err, "test: TestNewDidDocument#%v", i)
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
				dd, _ := NewDidDocument("did:cash:subject",
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
				dd, _ := NewDidDocument("did:cash:subject",
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
		{
			func() DidDocument {
				dd, _ := NewDidDocument("did:cash:subject",
					WithControllers(
						"did:cash:controller-1",
						"did:cash:controller-2",
						"did:cash:controller-3",
						"did:cash:controller-4",
					),
				)
				return dd
			},
			nil,
			false,
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
					d, _ := NewDidDocument("did:cash:subject")
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
							Authentication,
							KeyAgreement,
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
							Authentication,
							CapabilityInvocation,
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					),
				},
			},
			wantDid: DidDocument{
				Context: []string{
					"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
					contextDIDBase,
				},
				Id:         "did:cash:subject",
				Controller: nil,
				VerificationMethods: []*VerificationMethod{
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
				Services:             nil,
				Authentication:       []string{"did:cash:subject#key-1", "did:cash:subject#key-2"},
				KeyAgreement:         []string{"did:cash:subject#key-1"},
				CapabilityInvocation: []string{"did:cash:subject#key-2"},
			},
		},
		{
			wantErr: true, // duplicated existing method id
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								Authentication,
								KeyAgreement,
								KeyAgreement, // test duplicated relationship
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
							string(CapabilityDelegation),
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
			wantErr: true, // duplicated new method id
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								Authentication,
								KeyAgreement,
								KeyAgreement, // test duplicated relationship
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
							"did:cash:subject#key-2",
							"EcdsaSecp256k1VerificationKey2019",
							"did:cash:subject",
							"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
						),
						[]string{
							KeyAgreement,
						},
						[]string{
							"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
						},
					),
					NewVerification(
						NewVerificationMethod(
							"did:cash:subject#key-2",
							"EcdsaSecp256k1VerificationKey2019",
							"did:cash:subject",
							"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
						),
						[]string{
							Authentication,
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
					d, _ := NewDidDocument("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								Authentication,
								KeyAgreement,
								KeyAgreement, // test duplicated relationship
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
							string(Authentication),
							string(KeyAgreement),
							string(KeyAgreement), // test duplicated relationship
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
		{
			wantErr: true, // verification relationship does not exists
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cash:subject")
					return d
				},
				[]*Verification{
					{
						[]string{
							Authentication,
							"UNSUPPORTED RELATIONSHIP",
							KeyAgreement,
						},
						&VerificationMethod{
							"did:cash:subject#key1",
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
					d, _ := NewDidDocument("did:cash:subject",
						WithVerifications(
							NewVerification(
								NewVerificationMethod(
									"did:cash:subject#key-1",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								),
								[]string{
									Authentication,
									KeyAgreement,
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
									Authentication,
									CapabilityInvocation,
								},
								[]string{
									"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
								},
							),
						),
					)
					return d
				},
				"did:cash:subject#key-2",
			},
			wantDid: DidDocument{
				Context: []string{
					"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
					contextDIDBase,
				},
				Id:         "did:cash:subject",
				Controller: nil,
				VerificationMethods: []*VerificationMethod{
					{
						"did:cash:subject#key-1",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
				},
				Services:       nil,
				Authentication: []string{"did:cash:subject#key-1"},
				KeyAgreement:   []string{"did:cash:subject#key-1"},
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cash:subject",
						WithVerifications(
							NewVerification(
								VerificationMethod{
									"did:cash:subject#key-1",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								},
								[]string{
									Authentication,
									KeyAgreement,
								},
								nil,
							),
						),
					)
					return d
				},
				"did:cash:subject#key-1",
			},
			wantDid: DidDocument{
				Context: []string{
					contextDIDBase,
				},
				Id: "did:cash:subject",
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cash:subject",
						WithVerifications(
							NewVerification(
								VerificationMethod{
									"did:cash:subject#key-1",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								},
								[]string{
									Authentication,
									KeyAgreement,
								},
								nil,
							),
							NewVerification(
								VerificationMethod{
									"did:cash:subject#key-2",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								},
								[]string{
									Authentication,
									CapabilityInvocation,
								},
								nil,
							),
							NewVerification(
								VerificationMethod{
									"did:cash:subject#key-3",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								},
								[]string{
									Authentication,
									KeyAgreement,
									AssertionMethod,
								},
								nil,
							),
						),
					)
					return d
				},
				"did:cash:subject#key-2",
			},
			wantDid: DidDocument{
				Context: []string{
					contextDIDBase,
				},
				Id:         "did:cash:subject",
				Controller: nil,
				VerificationMethods: []*VerificationMethod{
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
				Services:        nil,
				Authentication:  []string{"did:cash:subject#key-1", "did:cash:subject#key-3"},
				KeyAgreement:    []string{"did:cash:subject#key-1", "did:cash:subject#key-3"},
				AssertionMethod: []string{"did:cash:subject#key-3"},
			},
		},
		{
			wantErr: true, // verification method not found
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cash:subject",
						WithVerifications(
							NewVerification(
								NewVerificationMethod(
									"did:cash:subject#key-1",
									"EcdsaSecp256k1VerificationKey2019",
									"did:cash:subject",
									"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
								),
								[]string{
									Authentication,
									KeyAgreement,
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
									Authentication,
									CapabilityInvocation,
								},
								[]string{
									"https://gpg.jsld.org/contexts/lds-gpg2020-v0.0.jsonld",
								},
							),
						),
					)
					return d
				},
				"did:cash:subject#key-3",
			},
			wantDid: DidDocument{},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_RevokeVerification#", i), func(t *testing.T) {
			gotDid := tt.params.malleate()

			err := gotDid.RevokeVerification(tt.params.methodID)

			if tt.wantErr {
				require.NotNil(t, err, "test: TestDidDocument_RevokeVerification#%v", i)
				return
			}

			require.Nil(t, err, "test: TestDidDocument_RevokeVerification#%v", i)

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
					dd, _ := NewDidDocument("did:cash:subject")
					return dd
				},
				methodID:      "did:cash:subject#key-1",
				relationships: []string{},
			},
			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cash:subject",
			},
		},
		{
			wantErr: true, //invalid method id
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cash:subject")
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
					dd, _ := NewDidDocument("did:cash:subject", WithVerifications(
						NewVerification(
							VerificationMethod{
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{},
						),
					))
					return dd
				},
				methodID: "did:cash:subject#key-1",
				relationships: []string{
					string(AssertionMethod),
					string(AssertionMethod), // test duplicated relationship
					string(AssertionMethod), // test duplicated relationship
					string(AssertionMethod), // test duplicated relationship
				},
			},

			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cash:subject",
				VerificationMethods: []*VerificationMethod{
					{
						"did:cash:subject#key-1",
						"EcdsaSecp256k1VerificationKey2019",
						"did:cash:subject",
						"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
					},
				},
				AssertionMethod: []string{"did:cash:subject#key-1"},
			},
		},
		{
			wantErr: false, // different delete scenarios
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cash:subject", WithVerifications(
						NewVerification(
							VerificationMethod{
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{},
						),
						NewVerification(
							VerificationMethod{
								"did:cash:subject#key-2",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{
								Authentication,
							},
							[]string{},
						),
					))
					return dd
				},
				methodID:      "did:cash:subject#key-1",
				relationships: []string{string(AssertionMethod)},
			},
			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cash:subject",
				VerificationMethods: []*VerificationMethod{
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
				Authentication:  []string{"did:cash:subject#key-2"},
				AssertionMethod: []string{"did:cash:subject#key-1"},
			},
		},
		{
			wantErr: false, // different delete scenarios
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cash:subject", WithVerifications(
						NewVerification(
							VerificationMethod{
								"did:cash:subject#key-2",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{
								Authentication,
							},
							[]string{},
						),
						NewVerification(
							VerificationMethod{
								"did:cash:subject#key-3",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{
								Authentication,
							},
							[]string{},
						),
						NewVerification(
							VerificationMethod{
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							},
							[]string{
								Authentication,
								KeyAgreement,
							},
							[]string{},
						),
					))
					return dd
				},
				methodID:      "did:cash:subject#key-1",
				relationships: []string{string(AssertionMethod)},
			},
			wantDid: DidDocument{
				Context: []string{contextDIDBase},
				Id:      "did:cash:subject",
				VerificationMethods: []*VerificationMethod{
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

				Authentication:  []string{"did:cash:subject#key-2", "did:cash:subject#key-3"},
				AssertionMethod: []string{"did:cash:subject#key-1"},
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
		signer        string
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
					dd, _ := NewDidDocument("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"signer",
							),
							[]string{
								string(Authentication),
								string(KeyAgreement),
							},
							nil,
						),
					))
					return dd
				},
				signer: "signer",
				relationships: []string{
					string(AssertionMethod),
					string(Authentication),
				},
			},
		},
		{
			expectedHasRelationship: false,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV",
							),
							[]string{
								Authentication,
								KeyAgreement,
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
								CapabilityDelegation,
							},
							nil,
						),
					))
					return dd
				},
				signer: "did:cash:subject",
				relationships: []string{
					string(CapabilityDelegation),
				},
			},
		},
		{
			expectedHasRelationship: false,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cash:subject")
					return dd
				},
				signer: "did:cash:subject",
				relationships: []string{
					string(CapabilityDelegation),
				},
			},
		},
		{
			expectedHasRelationship: false,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"signer",
							),
							[]string{
								Authentication,
								KeyAgreement,
							},
							nil,
						),
					))
					return dd
				},
				signer:        "signer",
				relationships: nil,
			},
		},
		{
			expectedHasRelationship: true,
			params: params{
				malleate: func() DidDocument {
					dd, _ := NewDidDocument("did:cash:subject", WithVerifications(
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-1",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"signer",
							),
							[]string{
								Authentication,
							},
							nil,
						),
						NewVerification(
							NewVerificationMethod(
								"did:cash:subject#key-2",
								"EcdsaSecp256k1VerificationKey2019",
								"did:cash:subject",
								"signer",
							),
							[]string{
								KeyAgreement,
							},
							nil,
						),
					))
					return dd
				},
				signer: "signer",
				relationships: []string{
					string(KeyAgreement),
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestDidDocument_SetVerificationRelationships#", i), func(t *testing.T) {
			didDoc := tt.params.malleate()
			gotHasRelationship := didDoc.HasRelationship(tt.params.signer, tt.params.relationships...)
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
					d, _ := NewDidDocument("did:cash:subject")
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
				Context: []string{contextDIDBase},
				Id:      "did:cash:subject",
				Services: []*Service{
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
		},
		{
			wantErr: true, // duplicated existing service id
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument(
						"did:cash:subject",
						WithServices(
							NewService(
								"agent:xyz",
								"DIDCommMessaging",
								"https://agent.xyz/1234",
							),
						),
					)
					return d
				},
				[]*Service{
					{
						"agent:abc",
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
			wantErr: true, // duplicated new service id
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cash:subject")
					return d
				},
				[]*Service{
					{
						"agent:xyz",
						"DIDCommMessaging",
						"https://agent.xyz/1234",
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
					d, _ := NewDidDocument("did:cash:subject")
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
					d, _ := NewDidDocument("did:cash:subject",
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
				Context: []string{contextDIDBase},
				Id:      "did:cash:subject",
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cash:subject",
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
				Context: []string{contextDIDBase},
				Id:      "did:cash:subject",
				Services: []*Service{
					{
						"agent:zyz",
						"DIDCommMessaging",
						"https://agent.abc/1234",
					},
				},
			},
		},
		{
			wantErr: false,
			params: params{
				func() DidDocument {
					d, _ := NewDidDocument("did:cash:subject",
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
				Context: []string{contextDIDBase},
				Id:      "did:cash:subject",
				Services: []*Service{
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
