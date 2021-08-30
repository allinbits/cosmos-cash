package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMsgCreateDidDocument(t *testing.T) {
	tests := []struct {
		id            string
		verifications Verifications
		services      Services
		owner         string
		expectPass    bool
	}{
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
						"did:auth:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			true,
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						DIDVMethodTypeCosmosAccountAddress.String(),
						"did:auth:whatever",
						&VerificationMethod_BlockchainAccountID{""},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // empty pub key
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						"",
						"did:auth:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // emtpy verification method type
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{},
					&VerificationMethod{
						"did:auth:whatever#1",
						DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
						"did:auth:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // empty relationships
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#/asd 123",
						DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
						"did:auth:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // invalid method id
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
						"did:auth:whatever",
						&VerificationMethod_PublicKeyHex{""},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // empty verification key
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						"",
						"did:auth:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // empty verification method type
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						DIDVMethodTypeCosmosAccountAddress.String(),
						"",
						&VerificationMethod_BlockchainAccountID{"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // invalid verification method controller
		},
		{
			"did:auth:whatever",
			Verifications{},
			Services{},
			"owner",
			false, // empty verifications
		},

		{
			"invalid did",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
						"cont",
						&VerificationMethod_BlockchainAccountID{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // invalid did
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
						"did:auth:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{
				&Service{
					"the:agent:service",
					"DIDCommMessaging",
					"https://agent.xyz/agent/123",
				},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			true,
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
						"did:auth:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{
				&Service{
					"the:agent:service",
					"",
					"https://agent.xyz/agent/123",
				},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // empty service type
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
						"did:auth:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{
				&Service{
					"",
					"DIDCommMessaging",
					"https://agent.xyz/agent/123",
				},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // service id is not valid
		},
		{
			"did:auth:whatever",
			Verifications{
				&Verification{
					[]string{string(Authentication)},
					&VerificationMethod{
						"did:auth:whatever#1",
						DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
						"did:auth:whatever",
						&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
					},
					[]string{},
				},
			},
			Services{
				&Service{
					"this:is:fine",
					"DIDCommMessaging",
					"",
				},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // service id is not valid
		},
	}

	for i, tc := range tests {
		msg := NewMsgCreateDidDocument(
			tc.id,
			tc.verifications,
			tc.services,
			tc.owner,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgCreateDidDocument#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgCreateDidDocument#%v", i)
		}
	}
}

func TestMsgUpdateDidDocument(t *testing.T) {
	tests := []struct {
		id          string
		controllers []string
		signer      string
		expectPass  bool
	}{
		{
			"did:cash:subject",
			[]string{"did:cash:controller-1"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"did:cash:subject",
			[]string{},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			// FIXME: duplicated controller
			"did:cash:subject",
			[]string{"did:cash:controller-1", "did:cash:controller-1"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"invalid did",
			[]string{"did:cash:controller-1"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid did
		},
		{
			"did:cash:subject",
			[]string{"invalid:controller"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid controller
		},
		{
			"did:cash:subject",
			[]string{"did:cash:controller-1", "did:cash:controller-2", ""},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid controller
		},
	}

	for i, tc := range tests {
		msg := NewMsgUpdateDidDocument(
			tc.id,
			tc.controllers,
			tc.signer,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgUpdateDidDocument#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgUpdateDidDocument#%v", i)
		}
	}
}

func TestMsgAddVerification(t *testing.T) {
	tests := []struct {
		id         string
		auth       Verification
		owner      string
		expectPass bool
	}{
		{
			"did:cash:subject",
			Verification{
				[]string{string(Authentication)},
				&VerificationMethod{
					"did:cash:subject#1",
					DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
					"did:cash:subject",
					&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
				},
				[]string{},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			true,
		},
		{
			"something not right",
			Verification{
				[]string{string(Authentication)},
				&VerificationMethod{
					"did:cash:subject#1",
					DIDVMethodTypeEcdsaSecp256k1VerificationKey2019.String(),
					"did:cash:subject",
					&VerificationMethod_PublicKeyHex{"03dfd0a469806d66a23c7c948f55c129467d6d0974a222ef6e24a538fa6882f3d7"},
				},
				[]string{},
			},
			"cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8",
			false, // invalid did
		},
	}

	for i, tc := range tests {
		msg := NewMsgAddVerification(
			tc.id,
			&tc.auth,
			tc.owner,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgAddVerification#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgAddVerification#%v", i)
		}
	}
}

func TestMsgRevokeVerification(t *testing.T) {
	tests := []struct {
		id         string
		key        string
		signer     string
		expectPass bool
	}{
		{
			"did:cash:subject",
			"did:cash:subject#key-method-1",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"invalid did",
			"did:cash:subject#key-method-1",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid did
		},
		{
			"did:cash:subject",
			"did:cash:subject  #   key-method-1",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid method id
		},
		{
			"did:cash:subject",
			"did:cash:subject#key-method-1",
			"",
			true, // empty signer
		},
	}

	for i, tc := range tests {
		msg := NewMsgRevokeVerification(
			tc.id,
			tc.key,
			tc.signer,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgRevokeVerification#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgRevokeVerification#%v", i)
		}
	}
}

func TestMsgSetVerificationRelationships(t *testing.T) {
	tests := []struct {
		id            string
		key           string
		relationships []string
		signer        string
		expectPass    bool
	}{
		{
			"did:cash:subject",
			"did:cash:subject#key-method-1",
			[]string{"authorization", "keyExchange"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"did:cash:subject",
			"did:cash:subject#key-method-1",
			[]string{"authorization"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"did:cash:subject",
			"did:cash:subject  #   key-method-1",
			[]string{"authorization", "keyExchange"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid method id
		},
		{
			"invalid did",
			"did:cash:subject#key-method-1",
			[]string{"authorization", "keyExchange"},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid did
		},
		{
			"did:cash:subject",
			"did:cash:subject#key-method-1",
			[]string{},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // empty relationship
		},
		{
			"did:cash:subject",
			"did:cash:subject#key-method-1",
			[]string{"authorization", "keyExchange"},
			"",
			true, // empty signer
		},
	}

	for i, tc := range tests {
		t.Logf("TestMsgRevokeVerification#%d", i)
		msg := NewMsgSetVerificationRelationships(
			tc.id,
			tc.key,
			tc.relationships,
			tc.signer,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgSetVerificationRelationships#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgSetVerificationRelationships#%v", i)
		}
	}
}

func TestMsgAddService(t *testing.T) {
	tests := []struct {
		id         string
		service    *Service
		signer     string
		expectPass bool
	}{
		{
			"did:cash:subject",
			&Service{
				Id:              "a:valid:url",
				Type:            "DIDCommMessaging",
				ServiceEndpoint: "https://agent.xyz/validate",
			},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"invalid did",
			&Service{
				Id:              "my:agent",
				Type:            "DIDCommMessaging",
				ServiceEndpoint: "https://agent.xyz/validate",
			},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid did
		},
		{
			"did:cash:subject",
			&Service{
				Id:              "",
				Type:            "DIDCommMessaging",
				ServiceEndpoint: "https://agent.xyz/validate",
			},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // invalid agent id
		},
		{
			"did:cash:subject",
			&Service{
				Id:              "my:agent",
				Type:            "",
				ServiceEndpoint: "https://agent.xyz/validate",
			},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // empty type
		},
		{
			"did:cash:subject",
			&Service{
				Id:              "my:agent",
				Type:            "DIDCommMessaging",
				ServiceEndpoint: "",
			},
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // empty service endpoint
		},
	}

	for i, tc := range tests {
		t.Logf("TestMsgRevokeVerification#%d", i)
		msg := NewMsgAddService(
			tc.id,
			tc.service,
			tc.signer,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgAddService#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgAddService#%v", i)
		}
	}
}

func TestMsgDeleteService(t *testing.T) {
	tests := []struct {
		id         string
		serviceID  string
		signer     string
		expectPass bool
	}{
		{
			"did:cash:subject",
			"my:service:uri",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			true,
		},
		{
			"invalid did",
			"my:service:uri",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, //invalid did
		},
		{
			"did:cash:subject",
			"",
			"cosmos1uam3kpjdx3wksx46lzq6y628wwyzv0guuren75",
			false, // empty service id
		},
	}

	for i, tc := range tests {
		t.Logf("TestMsgRevokeVerification#%d", i)
		msg := NewMsgDeleteService(
			tc.id,
			tc.serviceID,
			tc.signer,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgDeleteService#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgDeleteService#%v", i)
		}
	}
}
