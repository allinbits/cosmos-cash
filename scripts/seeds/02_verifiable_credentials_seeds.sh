#!/bin/bash


echo "Create Regulator VC to activate the Regulator did"
cosmos-cashd tx regulator activate-regulator-credential TheAuthority EU --did did:cosmos:net:cash:regulator \
--from regulator --chain-id cash -y

sleep 5

echo "Create Registration VC for EMTi did"
cosmos-cashd tx regulator issue-registration-credential \
did:cosmos:net:cash:emti-registration-credential did:cosmos:net:cash:regulator did:cosmos:net:cash:emti \
EU "First Galactic Bank" "FGB" \
--from regulator --chain-id cash -y

sleep 5

echo "Create License VC for EMTi did (sEUR)"
cosmos-cashd tx regulator issue-license-credential \
did:cosmos:net:cash:emti-eurolicense-credential did:cosmos:net:cash:regulator did:cosmos:net:cash:emti \
MICAEMI IRL "Another Financial Services Body (AFFB)" sEUR 10000 \
--from regulator --chain-id cash -y

sleep 5

echo "Create License VC for EMTi did (sUSD)"
cosmos-cashd tx verifiablecredential create-license-verifiable-credential \
did:cosmos:net:cash:emti-dollarlicense-credential did:cosmos:net:cash:vasp did:cosmos:net:cash:issuer \
MICAEMI IRL "Another Financial Services Body (AFFB)" sUSD 10000 \
--from regulator --chain-id cash -y

sleep 5

echo "Creating key for user user1"
echo 'y' | cosmos-cashd keys add user1

echo "Creating verifiable credential for user :kyc'd user"
cosmos-cashd tx verifiablecredential create-kyc-verifiable-credential \
	did:cosmos:net:cash:user1 did:cosmos:cred:kyc1 did:cosmos:net:cash:vasp secret 1000 1000 1000  \
	--from validator --chain-id cash -y


sleep 5

echo "Querying verifiable credentials"
cosmos-cashd query verifiablecredential verifiable-credentials --output json | jq

echo "Deleting verifiable credential for dollarlicense issuer"
cosmos-cashd tx verifiablecredential delete-verifiable-credential did:cosmos:net:cash:dollarlicense-credential did:cosmos:net:cash:vasp \
	--from validator --chain-id cash -y

sleep 5

echo "Querying verifiable credentials"
cosmos-cashd query verifiablecredential verifiable-credentials --output json | jq

sleep 5

echo "Validating verifiable credentials"
cosmos-cashd query verifiablecredential validate-verifiable-credential did:cosmos:net:cash:eurolicense-credential \
	$(cosmos-cashd keys show validator -p) --output json | jq
