#!/bin/bash

KYC_CRED=new-kyc-cred-$RANDOM
IDENTITY_CRED=new-identity-cred-$RANDOM
ISSUER_KYC_CRED=new-issuer-kyc-cred-01

echo "Creating key for user user1"
echo 'y' | cosmos-cashd keys add user1

echo
echo "Sending tokens to user1 from validator"
cosmos-cashd tx bank send $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show user1 -a) 100000stake --from validator --chain-id cash -y

echo
echo "Creating decentralized did for user"
cosmos-cashd tx did create-did --from user1 --chain-id cash -y

echo
echo "Creating identity verifiable credential for user :validator"
cosmos-cashd tx verifiablecredentialservice create-verifiable-credential did:cash:$(cosmos-cashd keys show user1 -a) $IDENTITY_CRED secret alice dublin 1-1-1970 1234 3531234  --from validator --chain-id cash -y

echo
echo "Adding identity service to decentralized did for users"
cosmos-cashd tx did add-service did:cash:$(cosmos-cashd keys show user1 -a) $IDENTITY_CRED IdentityCredential cash:$IDENTITY_CRED --from user1 --chain-id cash -y

echo
echo "Creating verifiable credential for user :user1 with issuer :validator"
cosmos-cashd tx verifiablecredentialservice create-kyc-verifiable-credential \
	did:cash:$(cosmos-cashd keys show user1 -a) $KYC_CRED secret 1000 1000 1000  \
	--from validator --chain-id cash -y

echo
echo "Adding service to decentralized did for users"
cosmos-cashd tx did add-service did:cash:$(cosmos-cashd keys show user1 -a) $KYC_CRED KYCCredential cash:$KYC_CRED --from user1 --chain-id cash -y

echo
echo "Creating verifiable credential for user :validator with issuer :validator"
cosmos-cashd tx verifiablecredentialservice create-kyc-verifiable-credential \
	did:cash:$(cosmos-cashd keys show validator -a) $ISSUER_KYC_CRED secret 1000 1000 1000  \
	--from validator --chain-id cash -y

echo
echo "Adding service to decentralized did for validator"
cosmos-cashd tx did add-service did:cash:$(cosmos-cashd keys show validator -a) $ISSUER_KYC_CRED KYCCredential cash:$ISSUER_KYC_CRED --from user1 --chain-id cash -y

echo
echo "Querying verifiable credentials"
cosmos-cashd query verifiablecredentialservice verifiable-credential $KYC_CRED --output json | jq

echo
echo "Sending issuer tokens to user from validator"
cosmos-cashd tx bank send $(cosmos-cashd keys show validator -a) $(cosmos-cashd keys show user1 -a) 100000seuro --from validator --chain-id cash -y

echo
echo "Querying balances for user"
gaiad query bank balances $(cosmos-cashd keys show user1 -a) --output json | jq

