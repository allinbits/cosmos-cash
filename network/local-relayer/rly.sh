#!/bin/bash

rly config init

## Set up ibcidentifier
rly config add-chains ./chains
rly config add-paths ./paths

echo "cosmos-cash seed:"
read cash
echo "cash2 seed:"
read cash2

rly keys restore cash cash "$cash"

rly keys restore cash2 cash2 "$cash2"

rly light init cash -f
rly light init cash2 -f

rly tx path cash-identifier-cash2
#rly tx path cash-transfer-cash2

rly start cash-identifier-cash2
#rly start cash-transfer-cash2 &

## Set up ibctransfer
