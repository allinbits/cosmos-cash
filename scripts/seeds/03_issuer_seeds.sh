#!/bin/bash

echo "Creating issuer for user: validator"
cosmos-cashd tx issuer create-issuer did:cosmos:net:cash:emti did:cosmos:net:cash:emti-eurolicense-credential seuro 100 --from validator --chain-id cash -y

sleep 5

echo "Querying all issuers"
cosmos-cashd query issuer issuers --output json | jq

echo "Mint tokens for issuer: validator"
cosmos-cashd tx issuer mint-token did:cosmos:net:cash:emti did:cosmos:net:cash:emti-eurolicense-credential 9999seuro --from validator --chain-id cash -y

sleep 5

echo "Check that the tokens have been minted"
cosmos-cashd query bank total --output json | jq

echo "Burn tokens for issuer: validator"
cosmos-cashd tx issuer burn-token did:cosmos:net:cash:emti did:cosmos:net:cash:emti-eurolicense-credential 10000seuro --from validator --chain-id cash -y

sleep 5

echo "Check that the tokens have been burned"
cosmos-cashd query bank total --output json | jq

echo "Pause tokens for issuer: validator"
cosmos-cashd tx issuer pause-token did:cosmos:net:cash:emti did:cosmos:net:cash:emti-eurolicense-credential --from validator --chain-id cash -y

sleep 5

echo "Querying all issuers"
cosmos-cashd query issuer issuers --output json | jq

echo "Unpause tokens for issuer: validator"
cosmos-cashd tx issuer pause-token did:cosmos:net:cash:emti did:cosmos:net:cash:emti-eurolicense-credential --from validator --chain-id cash -y

sleep 5

echo "Querying all issuers"
cosmos-cashd query issuer issuers --output json | jq
