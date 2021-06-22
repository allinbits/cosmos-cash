package helpers

import (
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/crypto/blake2b"

	credential "github.com/allinbits/cosmos-cash/x/verifiable-credential-service/types"
)

// SignCredential adds proof (signature) to the verifiable credentials
// returns error when the signature cannot be computed
func SignCredential(k keyring.Keyring, account sdk.AccAddress, vc *credential.VerifiableCredential) (err error) {
	// TODO: signature verification for big messages can be expensive and slow down the node?
	// reset proof if existing
	proof := credential.EmptyProof()
	vc.Proof = &proof
	signature, pubKey, err := k.SignByAddress(account, vc.GetBytes())
	if err != nil {
		return
	}
	// TODO: what is the keys-1 in the proof verification method?
	p := credential.NewProof(
		pubKey.Type(),
		vc.IssuanceDate,
		"assertionMethod",
		account.String()+"#keys-1",
		base64.StdEncoding.EncodeToString(signature),
	)
	vc.Proof = &p
	return
}

// VerifyCredentialSignature verifies the signature within the verifiable credentials
// return a nil error if the signature is verified successfully
func VerifyCredentialSignature(pubkey types.PubKey, vc credential.VerifiableCredential) (err error) {
	// pubkey := account.GetPubKey()
	if pubkey == nil {
		err = fmt.Errorf("account pubkey is nil")
		return
	}
	// extract the signature
	signature, err := base64.StdEncoding.DecodeString(vc.Proof.Signature)
	if err != nil {
		return
	}
	// remove the proof from the message
	emptyProof := credential.EmptyProof()
	vc.Proof = &emptyProof
	// verify the signature
	verified := pubkey.VerifySignature(
		vc.GetBytes(),
		signature,
	)
	if !verified {
		err = fmt.Errorf("signature mismatch")
	}
	return
}

// SignCredentialHash adds proof (signature) to the hash of the verifiable credentials
// returns error when the signature cannot be computed
func SignCredentialHash(k keyring.Keyring, account sdk.AccAddress, vc *credential.VerifiableCredential) (err error) {
	// reset the proof
	proof := credential.EmptyProof()
	vc.Proof = &proof
	// hash the data
	sigHash := blake2b.Sum256(vc.GetBytes())
	signature, pubKey, err := k.SignByAddress(account, sigHash[:])
	if err != nil {
		return
	}
	// TODO: what is the keys-1 in the proof verification method?
	p := credential.NewProof(
		pubKey.Type(),
		vc.IssuanceDate,
		"assertionMethod",
		account.String()+"#keys-1",
		base64.StdEncoding.EncodeToString(signature),
	)
	vc.Proof = &p
	return
}

// VerifyCredentialHashSignature verifies the signature within the verifiable credentials
// return a nil error if the signature is verified successfully
func VerifyCredentialHashSignature(pubkey types.PubKey, vc credential.VerifiableCredential) (err error) {
	// get the pubkey (verify if empty)
	if pubkey == nil {
		err = fmt.Errorf("account pubkey is nil")
		return
	}
	// extract the signature
	signature, err := base64.StdEncoding.DecodeString(vc.Proof.Signature)
	if err != nil {
		return
	}
	// remove the proof from the message
	emptyProof := credential.NewProof("", "", "", "", "")
	vc.Proof = &emptyProof
	// hash the verifiable credentials
	sigHash := blake2b.Sum256(vc.GetBytes())
	// verify
	verified := pubkey.VerifySignature(
		sigHash[:],
		signature,
	)
	if !verified {
		err = fmt.Errorf("signature mismatch")
	}
	return
}
