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
cosmos-cashd tx regulator issue-license-credential \
did:cosmos:net:cash:emti-dollarlicense-credential did:cosmos:net:cash:regulator did:cosmos:net:cash:emti \
MICAEMI PG "Yet Another Financial Services Body (YAFFB)" sUSD 10000 \
--from regulator --chain-id cash -y

echo "Revoke the (sUSD) license"
cosmos-cashd tx regulator revoke-credential \
did:cosmos:net:cash:emti-dollarlicense-credential \
--from regulator --chain-id cash -y

sleep 5

echo "Creating User VC for user alice"
cosmos-cashd tx issuer issue-user-credential \
did:cosmos:key:$(cosmos-cashd keys show alice -a) did:cosmos:cred:emti-user-alice did:cosmos:net:cash:emti zkp_secret 1000 1000 1000  \
--from emti --chain-id cash -y

sleep 5

echo "Creating User VC for user bob"
cosmos-cashd tx issuer issue-user-credential \
did:cosmos:key:$(cosmos-cashd keys show bob -a) did:cosmos:cred:emti-user-bob did:cosmos:net:cash:emti zkp_secret 1000 1000 1000  \
--from emti --chain-id cash -y

echo "Querying verifiable credentials"
cosmos-cashd query verifiablecredential verifiable-credentials --output json | jq

echo "Revoke Bob's user credential"
cosmos-cashd tx issuer revoke-credential did:cosmos:cred:emti-user-bob \
--from emti --chain-id cash -y

sleep 5

echo "Querying verifiable credentials"
cosmos-cashd query verifiablecredential verifiable-credentials --output json | jq


echo "Validating verifiable credentials"
cosmos-cashd query verifiablecredential validate-verifiable-credential did:cosmos:net:cash:emti-eurolicense-credential \
$(cosmos-cashd keys show regulator -p) --output json | jq
