#!/bin/bash

echo "Creating verifiable credential for user :validator"
cosmos-cashd tx verifiablecredential create-license-verifiable-credential \
	did:cosmos:cash:eurolicense-credential did:cosmos:cash:vasp did:cosmos:cash:vasp \
	MICAEMI IRL "Another Financial Services Body (AFFB)" sEUR 1000 \
	--from validator --chain-id cash -y

sleep 5

echo "Creating verifiable credential for user :issuer"
cosmos-cashd tx verifiablecredential create-license-verifiable-credential \
	did:cosmos:cash:dollarlicense-credential did:cosmos:cash:vasp did:cosmos:cash:issuer \
	MICAEMI IRL "Another Financial Services Body (AFFB)" sUSD 1000 \
	--from validator --chain-id cash -y

sleep 5

echo "Creating key for user user1"
echo 'y' | cosmos-cashd keys add user1

echo "Creating verifiable credential for user :kyc'd user"
cosmos-cashd tx verifiablecredential create-kyc-verifiable-credential \
	did:cosmos:cash:user1 did:cosmos:cred:kyc1 did:cosmos:cash:vasp secret 1000 1000 1000  \
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

sleep 5

echo "Validating verifiable credentials"
cosmos-cashd query verifiablecredential validate-verifiable-credential did:cosmos:cash:eurolicense-credential \
	$(cosmos-cashd keys show validator -p) --output json | jq
