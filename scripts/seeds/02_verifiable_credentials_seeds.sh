#!/bin/bash

echo "Creating verifiable credential for user :validator"
cosmos-cashd tx verifiablecredential create-license-verifiable-credential \
	did:cosmos:cash:eurolicense-credential did:cosmos:cash:vasp did:cosmos:cash:vasp \
	MICAEMI IRL "Another Financial Services Body (AFFB)" sEUR 1000 \
	--from validator --chain-id cash -y

sleep 4

echo "Querying verifiable credentials"
cosmos-cashd query verifiablecredential verifiable-credentials --output json | jq

echo "Validating verifiable credentials"
cosmos-cashd query verifiablecredential validate-verifiable-credential did:cosmos:cash:eurolicense-credential \
	$(cosmos-cashd keys show validator -p) --output json | jq

