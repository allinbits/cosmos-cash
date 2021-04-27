#!/bin/bash

GENESIS_FILE=~/.cosmoscash/config/genesis.json
if [ -f $GENESIS_FILE ]
then
    echo "Genesis file exist, would you like to delete it? (y/n)"
    read delete_config
fi

if [[
	$delete_config == "Y" ||
	! -f $GENESIS_FILE
   ]];
then
    rm -r ~/.cosmoscash

    echo "Initialising chain"
    cosmos-cashd init --chain-id=cash cash
    echo "y" | cosmos-cashd keys add validator

    echo "Adding genesis account"
    cosmos-cashd add-genesis-account $(cosmos-cashd keys show validator -a) 1000000000stake
    cosmos-cashd gentx validator 700000000stake --chain-id cash
    cosmos-cashd collect-gentxs
fi


echo "Starting cosmos cash chain"
cosmos-cashd start

