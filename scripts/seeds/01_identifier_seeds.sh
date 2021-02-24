#!/bin/bash

echo "Creating decentralized identifier for user: validator"
cosmos-cashd tx identifier create-identifier --from validator --chain-id cash -y

echo "Creating key for user auth2"
echo 'y' | cosmos-cashd keys add auth2

echo "Adding authentication to decentralized identifier for user: validator"
cosmos-cashd tx identifier add-authentication did:cash:$(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show auth2 -p) --from validator --chain-id cash -y

echo "Querying identifiers"
cosmos-cashd query identifier identifiers --output json | jq

echo "Adding service to decentralized identifier for user: validator"
cosmos-cashd tx identifier add-service did:cash:$(cosmos-cashd keys show validator -a) new-verifiable-cred-3 KYCCredential cosmos-cash:new-verifiable-cred-3 --from validator --chain-id cash -y

echo "Deleting authentication from decentralized identifier for user: validator"
cosmos-cashd tx identifier delete-authentication did:cash:$(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show auth2 -p) --from validator --chain-id cash -y

echo "Querying identifiers"
cosmos-cashd query identifier identifiers --output json | jq

echo "Deleting service from decentralized identifier for user: validator"
cosmos-cashd tx identifier delete-service did:cash:$(cosmos-cashd keys show validator -a) new-verifiable-cred-3 --from validator --chain-id cash -y
