#!/bin/bash

echo "Creating issuer for user: validator"
cosmos-cashd tx issuer create-issuer seuro 100 --from validator --chain-id cash -y

echo "Querying all issuers"
cosmos-cashd query issuer issuers --output json | jq

echo "Mint tokens for issuer: validator"
cosmos-cashd tx issuer mint-token 9999seuro --from validator --chain-id cash -y

echo "Check that the tokens have been minted"
cosmos-cashd query bank total --output json | jq

echo "Burn tokens for issuer: validator"
cosmos-cashd tx issuer burn-token 9999seuro --from validator --chain-id cash -y

echo "Check that the tokens have been burned"
cosmos-cashd query bank total --output json | jq
