package helpers

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	credential "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	bobMnemonic = "imitate burden rural clarify plate devote between alone kidney praise execute angry afford enough crane tent mosquito eye human fatigue cave cook pass argue"
	password    = "testPasspqo3498htqpwiojefp!@"
)

var (
	k     keyring.Keyring
	tests []testItem
)

// testItem data structure that holds the test data
type testItem struct {
	params    testParam
	wantErr   bool
	signature string
}

// testParam parameters for executing a test item
type testParam struct {
	account keyring.Info
	vc      credential.VerifiableCredential
}

// helper function to create an account
func buildAccount(name, mnemonic string) keyring.Info {
	acc, err := k.NewAccount(name,
		mnemonic,
		password,
		sdk.FullFundraiserPath,
		hd.Secp256k1,
	)
	if err != nil {
		panic(err)
	}
	return acc
}

// helper to build a credential object
func buildCred(a sdk.AccAddress, subjectData string, proof credential.Proof) credential.VerifiableCredential {
	cred := credential.NewUserVerifiableCredential(
		fmt.Sprint("did:cash:", a.String()),
		[]string{"VerifiableCredential", "IdentityCredential"},
		a.String(),
		time.Now().Format(time.RFC3339),
		credential.NewUserCredentialSubject(
			fmt.Sprint("did:cash:", a.String()),
			hex.EncodeToString([]byte(subjectData)),
			true,
		),
		proof,
	)
	return cred
}

// fixture for tests
func init() {
	k = keyring.NewInMemory()
	acc := buildAccount("bob", bobMnemonic)
	tests = []testItem{
		{
			params: testParam{
				acc, buildCred(
					acc.GetAddress(),
					fmt.Sprintf("%0*s", 0, `9074a74de0f34ece3f046403ae88d2eea400281da0ed6ebfa76c949016fa672d`),
					credential.EmptyProof()),
			},
			wantErr:   false,
			signature: "H7GR/Q2AOgEJ3qfB+XjH/pBG4kdUv8WPHWQBMDpUvL9zoSuTs/hyOjj03TI6OMnr2cvb30JgQfAyLxlMfZpcsA==",
		},
		{
			params: testParam{
				acc, buildCred(
					acc.GetAddress(),
					fmt.Sprintf("%0*s", 1000000, `9074a74de0f34ece3f046403ae88d2eea400281da0ed6ebfa76c949016fa672d`),
					credential.EmptyProof()),
			},
			wantErr:   false,
			signature: "H7GR/Q2AOgEJ3qfB+XjH/pBG4kdUv8WPHWQBMDpUvL9zoSuTs/hyOjj03TI6OMnr2cvb30JgQfAyLxlMfZpcsA==",
		},
	}
}

// TestSignCredential test sign+verification of a verifiable credential structure
func TestSignCredential(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestSignCredential#", i), func(t *testing.T) {
			if err := SignCredential(k, tt.params.account.GetAddress(), &tt.params.vc); (err != nil) != tt.wantErr {
				t.Errorf("SignCredential() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := VerifyCredentialSignature(tt.params.account.GetPubKey(), tt.params.vc); (err != nil) != tt.wantErr {
				t.Errorf("SignCredential() error = %v, wantErr %v", err, tt.wantErr)
			}

			// TODO: is the signature not deterministic?
			// if tt.signature != tt.params.vc.GetProof().Signature {
			// 	t.Errorf("SignCredential() \nexpected  %v\ngot       %v", tt.signature, tt.params.vc.GetProof().Signature)
			// }
		})
	}
}

// TestSignCredentialHash test signing+verification of hashed verifiable credentials
func TestSignCredentialHash(t *testing.T) {
	for i, tt := range tests {
		t.Run(fmt.Sprint("TestSignCredentialHash#", i), func(t *testing.T) {
			if err := SignCredentialHash(k, tt.params.account.GetAddress(), &tt.params.vc); (err != nil) != tt.wantErr {
				t.Errorf("SignCredentialHash() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := VerifyCredentialHashSignature(tt.params.account.GetPubKey(), tt.params.vc); (err != nil) != tt.wantErr {
				t.Errorf("SignCredentialHash() error = %v, wantErr %v", err, tt.wantErr)
			}

			// TODO:  is the signature not deterministic?
			// if tt.signature != tt.params.vc.GetProof().Signature {
			// 	t.Errorf("SignCredentialHash() \nexpected  %v\ngot       %v", tt.signature, tt.params.vc.GetProof().Signature)
			// }

		})
	}
}

func BenchmarkSignatureSignCredential(b *testing.B) {
	test := tests[0].params
	for i := 0; i < b.N; i++ {
		err := SignCredential(k, test.account.GetAddress(), &test.vc)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSignatureSignCredentialHash(b *testing.B) {
	test := tests[0].params
	for i := 0; i < b.N; i++ {
		err := SignCredentialHash(k, test.account.GetAddress(), &test.vc)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSignature(b *testing.B) {
	var (
		err  error
		vc   = tests[0].params.vc
		addr = tests[0].params.account.GetAddress()
		pk   = tests[0].params.account.GetPubKey()
	)
	for i := 0; i < b.N; i++ {
		err = SignCredential(k, addr, &vc)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		err = VerifyCredentialSignature(pk, vc)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkSignatureHash(b *testing.B) {
	var (
		err  error
		vc   = tests[0].params.vc
		addr = tests[0].params.account.GetAddress()
		pk   = tests[0].params.account.GetPubKey()
	)
	for i := 0; i < b.N; i++ {
		err = SignCredentialHash(k, addr, &vc)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		err = VerifyCredentialHashSignature(pk, vc)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

// malicious bench

func BenchmarkSignatureSignCredentialMalicious(b *testing.B) {
	test := tests[1].params
	for i := 0; i < b.N; i++ {
		err := SignCredential(k, test.account.GetAddress(), &test.vc)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSignatureSignCredentialHashMalicious(b *testing.B) {
	test := tests[1].params
	for i := 0; i < b.N; i++ {
		err := SignCredentialHash(k, test.account.GetAddress(), &test.vc)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSignatureMalicious(b *testing.B) {
	var (
		err  error
		vc   = tests[1].params.vc
		addr = tests[1].params.account.GetAddress()
		pk   = tests[1].params.account.GetPubKey()
	)
	for i := 0; i < b.N; i++ {
		err = SignCredential(k, addr, &vc)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		err = VerifyCredentialSignature(pk, vc)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkSignatureHashMalicious(b *testing.B) {
	var (
		err  error
		vc   = tests[1].params.vc
		addr = tests[1].params.account.GetAddress()
		pk   = tests[1].params.account.GetPubKey()
	)
	for i := 0; i < b.N; i++ {
		err = SignCredentialHash(k, addr, &vc)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		err = VerifyCredentialHashSignature(pk, vc)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

// need to have a valid signature for this one
// func BenchmarkSignatureVerifyCredential(b *testing.B) {
// 	test := tests[0].params
// 	for i := 0; i < b.N; i++ {
// 		err := VerifyCredentialSignature(test.account.GetPubKey(), &test.vc)
// 		if err != nil {
// 			b.Error(err)
// 		}
// 	}
// }

// need to have a valid signature for this one
// func BenchmarkSignatureVerifyCredentialHash(b *testing.B) {
// 	test := tests[0].params
// 	for i := 0; i < b.N; i++ {
// 		err := VerifyCredentialHashSignature(test.account.GetPubKey(), &test.vc)
// 		if err != nil {
// 			b.Error(err)
// 		}
// 	}
// }
