#!/bin/bash

echo "Creating decentralized identifier for user: validator"
cosmos-cashd tx identifier create-identifier vasp --from validator --chain-id cash -y

echo "Creating key for user auth2"
echo 'y' | cosmos-cashd keys add auth2

echo "Adding authentication to decentralized identifier for user: validator"
cosmos-cashd tx identifier add-verification-method vasp $(cosmos-cashd keys show auth2 -p) --from validator --chain-id cash -y

echo "Querying identifiers"
cosmos-cashd query identifier identifiers --output json | jq

echo "Adding service to decentralized identifier for user: validator"
cosmos-cashd tx identifier add-service vasp new-verifiable-cred-3 KYCCredential cosmos-cash:new-verifiable-cred-3 --from validator --chain-id cash -y

vmID=$(cosmos-cashd query identifier identifiers --output json | jq '.didDocuments[0].verificationMethods[1].id')
echo $vmID

echo "Revoking verification method from decentralized identifier for user: validator"
#cosmos-cashd tx identifier revoke-verification-method vasp $(echo $vmID)--from validator --chain-id cash -y

echo "Querying identifiers"
cosmos-cashd query identifier identifiers --output json | jq

echo "Deleting service from decentralized identifier for user: validator"
cosmos-cashd tx identifier delete-service vasp new-verifiable-cred-3 --from validator --chain-id cash -y

echo "Querying identifiers"
cosmos-cashd query identifier identifiers --output json | jq
