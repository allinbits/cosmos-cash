#!/bin/bash

echo "Creating key for user user1"
echo 'y' | cosmos-cashd keys add user1
echo 'y' | cosmos-cashd keys add user2

echo "Sending tokens to user 1 from validator"
cosmos-cashd tx bank send \
 $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show user1 -a) 100000stake \
 --from validator \
 --chain-id cash -y --broadcast-mode block

echo "Sending tokens to user 2 from validator"
cosmos-cashd tx bank send \
 $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show user2 -a) 100000stake \
 --from validator \
 --chain-id cash -y --broadcast-mode block


echo "Creating decentralized did for user1"
cosmos-cashd tx did create-did user1 \
 --from user1 \
 --chain-id cash -y --broadcast-mode block

echo "Creating decentralized did for user2"
cosmos-cashd tx did create-did user2 \
 --from user2 \
 --chain-id cash -y --broadcast-mode block


echo "Creating verifiable credential for user :user1"
cosmos-cashd tx issuer issue-user-credential \
	did:cosmos:net:cash:emti did:cosmos:net:cash:user1 secret 1000 1000 1000  \
	--credential-id did:cosmos:net:cash:user1-pokyc \
	--from emti \
	--chain-id cash -y --broadcast-mode block


echo "Creating verifiable credential for user :user2"
cosmos-cashd tx issuer issue-user-credential \
	did:cosmos:net:cash:emti did:cosmos:net:cash:user2 secret 1000 1000 1000  \
	--credential-id did:cosmos:net:cash:user2-pokyc \
	--from emti \
	--chain-id cash -y --broadcast-mode block


echo "Self-issuing proof of kyc to itself: emti"
cosmos-cashd tx issuer issue-user-credential \
	did:cosmos:net:cash:emti did:cosmos:net:cash:emti secret 1000 1000 1000  \
	--credential-id did:cosmos:net:cash:emti-pokyc \
	--from emti \
	--chain-id cash -y --broadcast-mode block


echo "Querying all data"
cosmos-cashd query did dids --output json | jq
cosmos-cashd query verifiablecredential verifiable-credentials --output json | jq

echo "Sending issuer tokens to users from emti"
cosmos-cashd tx bank send \
 $(cosmos-cashd keys show emti -a) $(cosmos-cashd keys show user1 -a) 10sEUR \
 --from emti \
 --chain-id cash -y --broadcast-mode block


echo "Sending issuer tokens to users from emti"
cosmos-cashd tx bank send \
 $(cosmos-cashd keys show emti -a) $(cosmos-cashd keys show user2 -a) 10sEUR \
 --from emti \
 --chain-id cash -y --broadcast-mode block


echo "Querying balances for users"
cosmos-cashd query bank balances $(cosmos-cashd keys show user1 -a) --output json | jq
cosmos-cashd query bank balances $(cosmos-cashd keys show user2 -a) --output json | jq

echo "Pause tokens for issuer: validator"
cosmos-cashd tx issuer pause-token \
 did:cosmos:net:cash:emti did:cosmos:net:cash:eurolicense-credential \
 --from regulator \
 --chain-id cash -y --broadcast-mode block


echo "Sending paused issuer tokens to user from validator: should fail"
cosmos-cashd tx bank send \
 $(cosmos-cashd keys show user1 -a) $(cosmos-cashd keys show user2 -a) 10sEUR \
 --from user1 \
 --chain-id cash -y --broadcast-mode block


echo "Querying balances for user2 should be 10sEUR"
cosmos-cashd query bank balances $(cosmos-cashd keys show user2 -a) --output json | jq
