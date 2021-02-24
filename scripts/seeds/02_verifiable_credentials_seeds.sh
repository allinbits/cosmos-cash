#!/bin/bash

echo "Creating verifiable credential for user :validator"
cosmos-cashd tx verifiablecredentialservice create-verifiable-credential did:cash:$(cosmos-cashd keys show validator -a) --from validator --chain-id cash -y
