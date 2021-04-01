#!/bin/bash

echo "Creating key for user user1"
echo 'y' | cosmos-cashd keys add user1
echo 'y' | cosmos-cashd keys add user2
echo 'y' | cosmos-cashd keys add user3

echo "Sending tokens to user 1 from validator"
cosmos-cashd tx bank send $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show user1 -a) 100000stake --from validator --chain-id cash -y
cosmos-cashd tx bank send $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show user2 -a) 100000stake --from validator --chain-id cash -y
cosmos-cashd tx bank send $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show user3 -a) 100000stake --from validator --chain-id cash -y

echo "Creating decentralized identifier for users"
cosmos-cashd tx identifier create-identifier --from user1 --chain-id cash -y
cosmos-cashd tx identifier create-identifier --from user2 --chain-id cash -y
cosmos-cashd tx identifier create-identifier --from user3 --chain-id cash -y

echo "Creating verifiable credential for user :validator"
cosmos-cashd tx verifiablecredentialservice create-verifiable-credential did:cash:$(cosmos-cashd keys show user1 -a) kyc-cred-1 --from validator --chain-id cash -y
cosmos-cashd tx verifiablecredentialservice create-verifiable-credential did:cash:$(cosmos-cashd keys show user2 -a) kyc-cred-2 --from validator --chain-id cash -y
cosmos-cashd tx verifiablecredentialservice create-verifiable-credential did:cash:$(cosmos-cashd keys show user3 -a) kyc-cred-3 --from validator --chain-id cash -y

echo "Adding service to decentralized identifier for users"
cosmos-cashd tx identifier add-service did:cash:$(cosmos-cashd keys show user1 -a) kyc-cred-1 KYCCredential cash:kyc-cred-1 --from user1 --chain-id cash -y
cosmos-cashd tx identifier add-service did:cash:$(cosmos-cashd keys show user2 -a) kyc-cred-2 KYCCredential cash:kyc-cred-2 --from user2 --chain-id cash -y
cosmos-cashd tx identifier add-service did:cash:$(cosmos-cashd keys show user3 -a) kyc-cred-3 KYCCredential cash:kyc-cred-3 --from user3 --chain-id cash -y

echo "Querying all data"
cosmos-cashd query identifier identifiers --output json | jq
cosmos-cashd query verifiablecredentialservice verifiable-credentials --output json | jq

echo "Sending issuer tokens to users from validator"
cosmos-cashd tx bank send $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show user1 -a) 100000seuro --from validator --chain-id cash -y
cosmos-cashd tx bank send $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show user2 -a) 100000seuro --from validator --chain-id cash -y
cosmos-cashd tx bank send $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show user3 -a) 100000seuro --from validator --chain-id cash -y

echo "Querying balances for users"
cosmos-cashd query bank balances $(cosmos-cashd keys show user1 -a) --output json | jq
cosmos-cashd query bank balances $(cosmos-cashd keys show user2 -a) --output json | jq
cosmos-cashd query bank balances $(cosmos-cashd keys show user3 -a) --output json | jq

