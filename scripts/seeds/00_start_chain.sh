#!/bin/bash

GENESIS_FILE=~/.cosmoscash/config/genesis.json
if [ -f $GENESIS_FILE ]
then
    echo "Genesis file exist, would you like to delete it? (y/n)"
    read delete_config
fi

if [[
	$delete_config == "Y" ||
	$delete_config == "y" ||
	! -f $GENESIS_FILE
   ]];
then
    rm -r ~/.cosmoscash

    echo "Initialising chain"
    cosmos-cashd init --chain-id=cash cash
    echo "y" | cosmos-cashd keys add validator
    echo "y" | cosmos-cashd keys add regulator
    echo "y" | cosmos-cashd keys add emti # e-money token issuer 
    echo "y" | cosmos-cashd keys add arti # asset-referenced token issuer 
    echo "y" | cosmos-cashd keys add bob
    echo "y" | cosmos-cashd keys add alice
 
    echo "Adding genesis account"
    cosmos-cashd add-genesis-account $(cosmos-cashd keys show validator -a) 1000000000stake --
    # this is to have the accounts on chain 
    cosmos-cashd add-genesis-account $(cosmos-cashd keys show emti -a) 1000stake
    cosmos-cashd add-genesis-account $(cosmos-cashd keys show arti -a) 1000stake
    cosmos-cashd add-genesis-account $(cosmos-cashd keys show bob -a) 1000stake
    cosmos-cashd add-genesis-account $(cosmos-cashd keys show alice -a) 1000stake
    ## add the regulator
    cosmos-cashd add-genesis-account $(cosmos-cashd keys show regulator -a) 1000stake --regulator $(cosmos-cashd keys show regulator -a) --
    cosmos-cashd gentx validator 700000000stake --chain-id cash
    cosmos-cashd collect-gentxs
fi


echo "Starting Cosmos Cash chain"
cosmos-cashd start