#!/bin/bash

echo "Creating issuer for user: emti"
cosmos-cashd tx issuer create-issuer \
 did:cosmos:net:cash:emti did:cosmos:net:cash:emti-eurolicense-credential sEUR 100 \
 --from emti \
 --chain-id cash -y --broadcast-mode block


echo "Querying all issuers"
cosmos-cashd query issuer issuers --output json | jq

echo "Mint tokens for issuer: emti"
cosmos-cashd tx issuer mint-token \
 did:cosmos:net:cash:emti did:cosmos:net:cash:emti-eurolicense-credential 9999sEUR \
 --from emti \
 --chain-id cash -y --broadcast-mode block


echo "Check that the tokens have been minted"
cosmos-cashd query bank total --output json | jq

echo "Burn tokens for issuer: emti"
cosmos-cashd tx issuer burn-token \
 did:cosmos:net:cash:emti did:cosmos:net:cash:emti-eurolicense-credential 1000sEUR \
 --from emti \
 --chain-id cash -y --broadcast-mode block


echo "Check that the tokens have been burned"
cosmos-cashd query bank total --output json | jq

echo "Pause tokens for issuer: emti"
cosmos-cashd tx issuer pause-token \
 did:cosmos:net:cash:emti did:cosmos:net:cash:emti-eurolicense-credential \
 --from emti \
 --chain-id cash -y --broadcast-mode block


echo "Querying all issuers"
cosmos-cashd query issuer issuers --output json | jq

echo "Unpause tokens for issuer: emti"
cosmos-cashd tx issuer pause-token \
 did:cosmos:net:cash:emti did:cosmos:net:cash:emti-eurolicense-credential \
 --from emti \
 --chain-id cash -y --broadcast-mode block


echo "Querying all issuers"
cosmos-cashd query issuer issuers --output json | jq
