#!/bin/bash

echo "Create a DID doc for the regulator (by the regulator account)"
cosmos-cashd tx did create-did regulator --from regulator --chain-id cash -y 

sleep 5
cosmos-cashd query did did did:cosmos:net:cash:regulator --output json | jq

echo "Create a DID doc for the EMTi (by the validator)"
cosmos-cashd tx did create-did emti --from validator --chain-id cash -y 

sleep 5
cosmos-cashd query did did did:cosmos:net:cash:emti --output json | jq

echo "Add the EMTi account verification method to the the EMTi DID doc (by the validator account)"
cosmos-cashd tx did add-verification-method emti $(cosmos-cashd keys show emti -p) --from validator --chain-id cash -y

sleep 5
cosmos-cashd query did did did:cosmos:net:cash:emti --output json | jq

echo "Querying dids"
cosmos-cashd query did dids --output json | jq

echo "Add a service to the EMTi DID doc (by the EMTi account)"
cosmos-cashd tx did add-service emti emti-agent DIDComm "https://agents.cosmos-cash.app.beta.starport.cloud/emti" \
--from emti --chain-id cash -y

sleep 5
cosmos-cashd query did did did:cosmos:net:cash:emti --output json | jq

echo "Adding a verification relationship from decentralized did for validator"
cosmos-cashd tx did set-verification-relationship emti $(cosmos-cashd keys show validator -a) --relationship assertionMethod --relationship capabilityInvocation \
--from emti --chain-id cash -y

sleep 5
cosmos-cashd query did did did:cosmos:net:cash:emti --output json | jq

echo "Revoking verification method from decentralized did for user: validator"
cosmos-cashd tx did revoke-verification-method emti $(cosmos-cashd keys show validator -a) \
--from emti --chain-id cash -y

sleep 5
cosmos-cashd query did did did:cosmos:net:cash:emti --output json | jq

echo "Adding Alice as controller of the EMTi did (by EMTi user)"
cosmos-cashd tx did update-did-document emti $(cosmos-cashd keys show alice -a) \
--from emti --chain-id cash -y

sleep 5

echo "Querying dids"
cosmos-cashd query did dids --output json | jq

echo "Deleting service from EMTi did document (by EMTi user)"
cosmos-cashd tx did delete-service emti emti-agent \
--from emti --chain-id cash -y

sleep 5

echo "Querying dids"
cosmos-cashd query did dids --output json | jq
