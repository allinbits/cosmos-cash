#!/bin/bash

echo "Creating issuer for user: validator"
cosmos-cashd tx issuer create-issuer did:cosmos:net:cash:vasp did:cosmos:net:cash:eurolicense-credential seuro 100 --from validator --chain-id cash -y

sleep 5

echo "Querying all issuers"
cosmos-cashd query issuer issuers --output json | jq

# TODO: update mint tokens to take a license cred
#echo "Mint tokens for issuer: validator"
#cosmos-cashd tx issuer mint-token 9999seuro --from validator --chain-id cash -y

echo "Check that the tokens have been minted"
cosmos-cashd query bank total --output json | jq

# TODO: update burn tokens to take a license cred
#echo "Burn tokens for issuer: validator"
#cosmos-cashd tx issuer burn-token 9999seuro --from validator --chain-id cash -y

#echo "Check that the tokens have been burned"
#cosmos-cashd query bank total --output json | jq
