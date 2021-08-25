#!/bin/bash

echo "Creating decentralized did for user: validator"
cosmos-cashd tx did create-did vasp --from validator --chain-id cash -y

sleep 3

echo "Creating key for user auth2"
echo 'y' | cosmos-cashd keys add auth2

# FIXME: the keys subcommand returns a different value for public keys now
# this needs to be updated
#echo "Adding authentication to decentralized did for user: validator"
#cosmos-cashd tx did add-verification-method vasp $(cosmos-cashd keys show auth2 -p) --from validator --chain-id cash -y

sleep 3

echo "Querying dids"
cosmos-cashd query did dids --output json | jq

echo "Adding service to decentralized did for user: validator"
cosmos-cashd tx did add-service vasp new-verifiable-cred-3 KYCCredential cosmos-cash:new-verifiable-cred-3 --from validator --chain-id cash -y

sleep 3

# FIXME: the keys subcommand returns a different value for public keys now
# this needs to be updated, as this uses the previously created verificationMethods
#vmID=$(cosmos-cashd query did dids --output json | jq '.didDocuments[0].verificationMethods[1].id')
#vmID=${vmID:17:-1}
#
#echo "Adding a verification relationship from decentralized did for user: validator"
#cosmos-cashd tx did add-verification-relationship vasp $vmID assertionMethod --from validator --chain-id cash -y

echo "Querying dids"
cosmos-cashd query did dids --output json | jq

# FIXME: the keys subcommand returns a different value for public keys now
# this needs to be updated, as this uses the previously created verificationMethods
#echo "Revoking verification method from decentralized did for user: validator"
#cosmos-cashd tx did revoke-verification-method vasp $vmID --from validator --chain-id cash -y

sleep 3

echo "Creating key for user didcontroller"
echo 'y' | cosmos-cashd keys add didcontroller

echo "Adding a controller to a did document for user: validator"
cosmos-cashd tx did update-did-document vasp $(cosmos-cashd keys show didcontroller -a) --from validator --chain-id cash -y

sleep 3

echo "Querying dids"
cosmos-cashd query did dids --output json | jq

echo "Deleting service from decentralized did for user: validator"
cosmos-cashd tx did delete-service vasp new-verifiable-cred-3 --from validator --chain-id cash -y

sleep 3

echo "Querying dids"
cosmos-cashd query did dids --output json | jq
