#!/bin/bash

echo "Creating issuer for user: validator"
cosmos-cashd tx issuer create-issuer seuro 100 --from validator --chain-id cash -y

echo "Querying all issuers"
cosmos-cashd query issuer issuers --output json | jq
