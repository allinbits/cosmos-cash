#!/bin/bash

cosmos-cashd tx identifier create-identifier --from validator --chain-id cash -y

echo 'y' | cosmos-cashd keys add auth2

cosmos-cashd tx identifier add-authentication did:cash:$(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show auth2 -p) --from validator --chain-id cash -y

cosmos-cashd tx identifier add-service did:cash:$(cosmos-cashd keys show validator -a) new-verifiable-cred-3 KYCCredential cosmos-cash:new-verifiable-cred-3 --from validator --chain-id cash -y
