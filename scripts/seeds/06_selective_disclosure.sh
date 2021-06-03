#!/bin/bash

## Context
echo
echo "Bob wants to obtain some emoney a.k.a collaterized stablecoin for his vacation"
echo "since bob lives in the EU, he needs to follow regulation and adhear to AML policies to do this he needs to prove his identity"
echo "since he is an EU citizen the process has to be privacy respecting so the identity of bob is not disclosed"

## Technical Details: The steps. 
echo
echo "DID: Bob needs to publish a his own decentralized identifier on-chain, this allows him own credentials"
echo "VERIFIABLE_CREDENTIALS: An eIDAS identity provider needs to publish Bobs privacy respecting credentials on-chain"
echo "VERIFIABLE_CREDENTIALS: The emoney issuer needs to be able to verify Bobs credentials"
echo "VERIFIABLE_CREDENTIALS: Bob needs to be able to prove he owns the credential"

echo 
echo "STEP 1: Bob generates his keys"
read cont
cosmos-cashd keys add bob

echo
echo "STEP 2: Bob gets some native tokens for the chain"
read cont
cosmos-cashd tx bank send $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show bob -a) 100000stake --from validator --chain-id cash -y

echo
echo "STEP 3: Bob publishes his own decentralized identifier"
read cont
cosmos-cashd tx identifier create-identifier --from bob --chain-id cash  -y

echo
echo "STEP 3-a: Now we can query the identifier"
read cont
cosmos-cashd query identifier identifiers --output json  | jq


echo
echo "STEP 4: An eIDAS identity provider publishes Bobs privacy respecting credentials"
echo "Bob shares his secret and credential with the eIDAS compliant identity provider using a secure channel"
echo "secret:  secret"
echo "name:    bob"
echo "DOB:     1-1-1970"
echo "address: berlin/germany"
echo "id:      1234567"
echo "number:  3531234567"
read cont
cosmos-cashd tx verifiablecredentialservice create-verifiable-credential \
	did:cash:$(cosmos-cashd keys show bob -a) what-a-demo-1 secret bob 1-1-1970 berlin/germany 1234567 3531234567 \
	--from validator --chain-id cash  -y 

echo
echo "STEP 4-a: Now we can query the verifiable credential"
read cont
cosmos-cashd query verifiablecredentialservice verifiable-credential what-a-demo-1 --output json | jq 

echo
echo "STEP 4-b: What does this credential look like?"
read cont 
vlc ../../temp/sample.png
read cont 

echo
echo "STEP 5: Bob associates the credential with his DID"
read cont
cosmos-cashd tx identifier add-service did:cash:$(cosmos-cashd keys show bob -a) what-a-demo-1 KYCCredential cash:what-a-demo-1 --from bob --chain-id cash -y

echo
echo "STEP 5-a: Now the credential is associated with Bobs DID"
read cont
cosmos-cashd query identifier identifier did:cash:$(cosmos-cashd keys show bob -a) --output json  | jq

echo
echo "How can Bob prove that the data in the credential is accurate?"

echo
echo "How it all works: What Bob can do with his credential"
read cont

echo
echo "Bobs options are:"
echo "1. Bob could reveal all his attributes in his credential"
echo "2. Bob could reveal one of the attributes of his creds using a proof"
read cont

echo
echo "Bob decides to reveal his age by revealing his DOB: 1-1-1970"
read cont

echo "Bob needs to generate a proof for his credentials"
echo "Bob uses a tools to input his secret and his credential to generate this proof"
echo "Bob then distributes this proof and his DOB to the emoney token issuer via a secure channel"
read cont

echo "dlv debug --continue ~/git/PaddyMc/selective-disclosire/go"

echo
echo "Now using a webui we can as the emoney token provider verify Bobs credential is correct"
echo "Issuer webui is outlined below:"
echo "https://credentials.tendermint.prototyp.xyz/"

# --node https://cosmos-cash.app.beta.starport.cloud:443
# curl -X POST -d "{\"address\": \"$(cosmos-cashd keys show bob -a)\"}" https://faucet.cosmos-cash.app.beta.starport.cloud
