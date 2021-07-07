#!/bin/bash

echo "Creating decentralized identifier for user: validator"
cosmos-cashd tx identifier create-identifier validator --from validator --chain-id cash -y

echo "Creating key for user auth2"
echo 'y' | cosmos-cashd keys add auth2

echo "Adding authentication to decentralized identifier for user: validator"
cosmos-cashd tx identifier add-verification-method validator $(cosmos-cashd keys show auth2 -p) --from validator --chain-id cash -y

echo "Querying identifiers"
cosmos-cashd query identifier identifiers --output json | jq


# XXX: something odd is happening here?
# lets pair on this tomorrow
echo "Adding service to decentralized identifier for user: validator"
cosmos-cashd tx identifier add-service validator new-verifiable-cred-3 KYCCredential cosmos-cash:new-verifiable-cred-3 --from validator --chain-id cash -y

#echo "Deleting authentication from decentralized identifier for user: validator"
#cosmos-cashd tx identifier revoke-verification validator did:cash:validator#0a6bcfd0-96f7-4c30-95b1-3b578058190f --from validator --chain-id cash -y

echo "Querying identifiers"
cosmos-cashd query identifier identifiers --output json | jq

#echo "Deleting service from decentralized identifier for user: validator"
#cosmos-cashd tx identifier delete-service validator new-verifiable-cred-3 --from validator --chain-id cash -y
#
#echo "Querying identifiers"
#cosmos-cashd query identifier identifiers --output json | jq
