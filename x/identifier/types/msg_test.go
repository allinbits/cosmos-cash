package types

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgCreateIdentifier(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		auth       Authentications
		owner      string
		expectPass bool
	}{
		{
			"Pass: ",
			"id1",
			Authentications{
				&Authentication{
					"auth",
					"type",
					"cont",
					"value",
				},
			},
			"owner",
			true,
		},
	}

	for _, tc := range tests {
		msg := NewMsgCreateIdentifier(
			tc.id,
			tc.auth,
			tc.owner,
		)

		if tc.expectPass {
			require.Nil(t, msg.ValidateBasic(), "test: %v", tc.name)
		} else {
			require.NotNil(t, msg.ValidateBasic(), "test: %v", tc.name)
		}
	}
}

func TestMsgAddAuthentication(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		auth       Authentication
		owner      string
		expectPass bool
	}{
		{
			"Pass: ",
			"id1",
			Authentication{
				"auth",
				"type",
				"cont",
				"value",
			},
			"owner",
			true,
		},
	}

	for _, tc := range tests {
		msg := NewMsgAddAuthentication(
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

func TestMsgDeleteAuthentication(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		key        string
		owner      string
		expectPass bool
	}{
		{
			"Pass: ",
			"id1",
			"key",
			"owner",
			true,
		},
	}

	for _, tc := range tests {
		msg := NewMsgDeleteAuthentication(
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
