package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMsgCreateIdentifier(t *testing.T) {
	tests := []struct {
		name          string
		id            string
		verifications Verifications
		services      Services
		owner         string
		expectPass    bool
	}{
		{
			"Pass: ",
			"id1",
			Verifications{
				&Verification{
					[]string{RelationshipAuthentication},
					&VerificationMethod{
						"auth",
						"type",
						"cont",
						"value",
					},
					[]string{},
				},
			},
			Services{},
			"owner",
			true,
		},
	}

	for _, tc := range tests {
		msg := NewMsgCreateIdentifier(
			tc.id,
			tc.verifications,
			tc.services,
			tc.owner,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}

func TestMsgAddVerification(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		auth       Verification
		owner      string
		expectPass bool
	}{
		{
			"Pass: ",
			"id1",
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

	for _, tc := range tests {
		msg := NewMsgAddVerification(
			tc.id,
			&tc.auth,
			tc.owner,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}

func TestMsgRevokeVerification(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		key        string
		owner      string
		expectPass bool
	}{
		{
			"Pass: ",
			"id1:cash:31",
			"key",
			"owner",
			true,
		},
	}

	for _, tc := range tests {
		msg := NewMsgRevokeVerification(
			tc.id,
			tc.key,
			tc.owner,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}
