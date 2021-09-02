#!/bin/bash

echo "Creating verifiable credential for user :validator"
cosmos-cashd tx verifiablecredential create-license-verifiable-credential \
	did:cosmos:cash:eurolicense-credential did:cosmos:cash:vasp did:cosmos:cash:vasp \
	MICAEMI IRL "Another Financial Services Body (AFFB)" sEUR 1000 \
	--from validator --chain-id cash -y

sleep 5

echo "Creating verifiable credential for user :issuer"
cosmos-cashd tx verifiablecredential create-license-verifiable-credential \
	did:cosmos:cash:dollarlicense-credential did:cosmos:cash:issuer did:cosmos:cash:vasp \
	MICAEMI IRL "Another Financial Services Body (AFFB)" sUSD 1000 \
	--from validator --chain-id cash -y

sleep 5

echo "Querying verifiable credentials"
cosmos-cashd query verifiablecredential verifiable-credentials --output json | jq

echo "Deleting verifiable credential for dollarlicense issuer"
cosmos-cashd tx verifiablecredential delete-verifiable-credential did:cosmos:cash:dollarlicense-credential did:cosmos:cash:vasp \
	--from validator --chain-id cash -y

sleep 5

echo "Querying verifiable credentials"
cosmos-cashd query verifiablecredential verifiable-credentials --output json | jq

### FIXME: update to new key format in 0.43 SDK
#echo "Validating verifiable credentials"
#cosmos-cashd query verifiablecredential validate-verifiable-credential did:cosmos:cash:eurolicense-credential \
#	$(cosmos-cashd keys show validator -p) --output json | jq

