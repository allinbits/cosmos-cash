package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMsgCreateIdentifier(t *testing.T) {
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
					[]string{RelationshipAuthentication},
					&VerificationMethod{
						"did:auth:whatever#1",
						"type",
						"did:auth:whatever",
						"3214141231",
					},
					[]string{},
				},
			},
			Services{},
			"owner",
			true,
		},
		{
			"invalid did",
			Verifications{
				&Verification{
					[]string{RelationshipAuthentication},
					&VerificationMethod{
						"did:auth:whatever#1",
						"type",
						"cont",
						"value",
					},
					[]string{},
				},
			},
			Services{},
			"owner",
			false, // invalid did
		},
	}

	for i, tc := range tests {
		t.Logf("TestMsgCreateIdentifier#%d", i)
		msg := NewMsgCreateIdentifier(
			tc.id,
			tc.verifications,
			tc.services,
			tc.owner,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgCreateIdentifier#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgCreateIdentifier#%v", i)
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
			"did:cash:whatever",
			Verification{
				[]string{RelationshipAuthentication},
				&VerificationMethod{
					"auth",
					"type",
					"cont",
					"value",
				},
				[]string{},
			},
			"owner",
			true,
		},
	}

	for i, tc := range tests {
		t.Logf("TestMsgAddVerification#%d", i)
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
		owner      string
		expectPass bool
	}{
		{
			"did:cash:whatever",
			"did:cash:whatever#key-method-1",
			"signerAddress",
			true,
		},
	}

	for i, tc := range tests {
		t.Logf("TestMsgRevokeVerification#%d", i)
		msg := NewMsgRevokeVerification(
			tc.id,
			tc.key,
			tc.owner,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: TestMsgRevokeVerification#%v", i)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: TestMsgRevokeVerification#%v", i)
		}
	}
}
