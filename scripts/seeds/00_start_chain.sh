#!/bin/bash

echo "Initialising chain"
cosmos-cashd init --chain-id=cash cash 
echo "y" | cosmos-cashd keys add validator

echo "Adding genesis account"
cosmos-cashd add-genesis-account $(cosmos-cashd keys show validator -a) 1000000000stake
cosmos-cashd gentx validator 700000000stake --chain-id cash
cosmos-cashd collect-gentxs 

echo "Starting cosmos cash chain"
cosmos-cashd start 
