package cli

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash/x/did/types"
)

// request send a json http request
func request(method, url string, requestBody io.Reader, val interface{}) (err error) {
	var client = &http.Client{
		Timeout: time.Second * 1,
	}
	req, err := http.NewRequestWithContext(context.Background(), method, url, requestBody)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyBytes, &val)
	return
}

// post send a post http request
func post(url string, requestBody, val interface{}) error {
	if requestBody != nil {
		return request("POST", url, serialize(requestBody), val)
	}
	return request("POST", url, nil, val)
}

// serialize  interface to json bytes
func serialize(in interface{}) io.Reader {
	v, err := json.Marshal(in)
	if err != nil {
		// TODO: NO PANIC
		panic(err.Error())
	}
	return bytes.NewBuffer(v)
}

// keySetCreateReq struct to make a request to generate a new keySet
type keySetCreateReq struct {
	KeyType string `json:"keyType"`
}

// keySetCreateRsp response to generate a new keySet
type keySetCreateRsp struct {
	KeyID     string `json:"keyId"`
	PublicKey string `json:"publicKey"`
}

// toBytes get the public key to bytes
func (kscr keySetCreateRsp) toBytes() ([]byte, error) {

	return base64.RawURLEncoding.DecodeString(kscr.PublicKey)
}

// x25519ECDHKW struct containing a decoded aries pubkey reply for X25519ECDHKW
type x25519ECDHKW struct {
	Kid   string `json:"kid"`
	X     string `json:"x"`
	Curve string `json:"curve"`
	Type  string `json:"type"`
}

// processAriesKeySet process key in case something else has to be done
// the key generated for a keyType X25519ECDHKW is actually
// a json structure containing the key itself
func processAriesKeySet(keyType string, kscr *keySetCreateRsp) (err error) {
	if keyType != "X25519ECDHKW" {
		return
	}
	/*
		for this particular key the result from aries is something like this:
		eyJraWQiOiJKdEhpSVFaN0JENTRhalBwN0xhUU5ZR2ZiRFQ3bnN5ZkJiTDUwWHVTRjhRIiwieCI6IkNZc1c0eG1HeXNFejJRcm9vZlViOCsvUGpBb043ejQrajhveWdCcW5JREU9IiwiY3VydmUiOiJYMjU1MTkiLCJ0eXBlIjoiT0tQIn0
		that is actually a RawURLEncoding of this structure
		{"kid":"JtHiIQZ7BD54ajPp7LaQNYGfbDT7nsyfBbL50XuSF8Q","x":"CYsW4xmGysEz2QroofUb8+/PjAoN7z4+j8oygBqnIDE=","curve":"X25519","type":"OKP"}
		that contains, repeated, the kid and the actual pubkey standard base64 encoded in the X
	*/

	// decode the kscr.PubKey
	keyRawJSON, err := base64.RawURLEncoding.DecodeString(kscr.PublicKey)
	if err != nil {
		return
	}
	// now get into a struct
	var xPubKey x25519ECDHKW
	err = json.Unmarshal(keyRawJSON, &xPubKey)
	if err != nil {
		return
	}
	// now convert the base64 encoding
	rawPubKey, err := base64.StdEncoding.DecodeString(xPubKey.X)
	if err != nil {
		return
	}
	// re-encode as expected
	kscr.PublicKey = base64.RawURLEncoding.EncodeToString(rawPubKey)
	return
}

func createKeySetOnAgent(agentURL, keyType string) (keyID string, pubKey []byte, err error) {
	var ksr keySetCreateRsp
	err = post(fmt.Sprint(agentURL, "/kms/keyset"), &keySetCreateReq{KeyType: keyType}, &ksr)
	if err != nil {
		return
	}
	// this is a dodgy one
	err = processAriesKeySet(keyType, &ksr)
	if err != nil {
		return
	}
	// now get the bytes
	pubKey, err = ksr.toBytes()
	if err != nil {
		return
	}
	keyID = ksr.KeyID
	return
}

// NewLinkAriesAgentCmd link an aries
func NewLinkAriesAgentCmd() *cobra.Command {

	var keyType string

	cmd := &cobra.Command{
		Use:     "link-aries-agent [id] [aries-agent-api-url] [service-endpoint]",
		Short:   "create a service and add a keyExchange method and relationship from an aries agent",
		Example: `cosmos-cashd tx link-aries-agent alice http://localhost:7090 http://localhost:7091`,
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			signer := clientCtx.GetFromAddress()
			ariesAPIURL := args[1]
			serviceEndpoint := args[2]
			// send a http request to the aries endpoint to generate a new key
			keyID, pubKey, err := createKeySetOnAgent(ariesAPIURL, keyType)
			if err != nil {
				return err
			}
			// build the target DID
			did := types.NewChainDID(clientCtx.ChainID, args[0])
			// add a new verification method based on the key generated by aries
			vmID := did.NewVerificationMethodID(keyID)
			verification := types.NewVerification(
				types.NewVerificationMethod(
					vmID,
					did,
					types.NewPublicKeyMultibase(pubKey, types.DIDVMethodTypeX25519KeyAgreementKey2019),
				),
				[]string{types.KeyAgreement},
				nil,
			)
			// add verification message
			msgAV := types.NewMsgAddVerification(
				did.String(),
				verification,
				signer.String(),
			)

			// add service that links the aries agent that holds the private key
			service := types.NewService(
				args[0]+"-agent",
				"DIDCommMessaging",
				serviceEndpoint,
			)

			msgAS := types.NewMsgAddService(
				did.String(),
				service,
				signer.String(),
			)
			// broadcast the messages
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgAV, msgAS)

		},
	}

	cmd.Flags().StringVar(&keyType, "key-type", "X25519ECDHKW", "the key type that the aries node should generate")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
