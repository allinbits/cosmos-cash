#!/bin/bash

echo "Creating verifiable credential for user :validator"
cosmos-cashd tx verifiablecredential create-verifiable-credential \
	did:cash:$(cosmos-cashd keys show validator -a) new-verifiable-cred-3 secret name address dob nationalId phoneNumber \
	--from validator --chain-id cash -y

echo "Querying verifiable credentials"
cosmos-cashd query verifiablecredential verifiable-credentials --output json | jq

echo "Creating issuer verifiable credential for user :validator"
cosmos-cashd tx verifiablecredential create-issuer-verifiable-credential \
	did:cash:$(cosmos-cashd keys show validator -a) issuer-new-verifiable-cred-3 secret businessName businessRegistrationNumber businessType businessAddress \
	--from validator --chain-id cash -y

echo "Querying verifiable credentials"
cosmos-cashd query verifiablecredential verifiable-credentials --output json | jq

echo "Validating verifiable credentials"
cosmos-cashd query verifiablecredential validate-verifiable-credential new-verifiable-cred-3 $(cosmos-cashd keys show validator -p) --output json | jq

echo "Adding service to decentralized did for user: validator"
cosmos-cashd tx did add-service did:cash:$(cosmos-cashd keys show validator -a) issuer-new-verifiable-cred-3 IssuerCredential cash:new-verifiable-cred-3 --from validator --chain-id cash -y
