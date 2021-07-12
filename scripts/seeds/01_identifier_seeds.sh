#!/bin/bash

echo "Creating decentralized did for user: validator"
cosmos-cashd tx did create-did vasp --from validator --chain-id cash -y

echo "Creating key for user auth2"
echo 'y' | cosmos-cashd keys add auth2

echo "Adding authentication to decentralized did for user: validator"
cosmos-cashd tx did add-verification-method vasp $(cosmos-cashd keys show auth2 -p) --from validator --chain-id cash -y

echo "Querying dids"
cosmos-cashd query did dids --output json | jq

echo "Adding service to decentralized did for user: validator"
cosmos-cashd tx did add-service vasp new-verifiable-cred-3 KYCCredential cosmos-cash:new-verifiable-cred-3 --from validator --chain-id cash -y

vmID=$(cosmos-cashd query did dids --output json | jq '.didDocuments[0].verificationMethods[1].id')
echo $vmID

echo "Revoking verification method from decentralized did for user: validator"
#cosmos-cashd tx did revoke-verification-method vasp $(echo $vmID)--from validator --chain-id cash -y

echo "Querying dids"
cosmos-cashd query did dids --output json | jq

echo "Deleting service from decentralized did for user: validator"
cosmos-cashd tx did delete-service vasp new-verifiable-cred-3 --from validator --chain-id cash -y

echo "Querying dids"
cosmos-cashd query did dids --output json | jq
