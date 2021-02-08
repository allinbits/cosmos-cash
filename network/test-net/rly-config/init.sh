#!/bin/bash

rly config init

rly config add-chains config/chains
rly config add-paths config/paths

echo "cosmos-cash seed: "
read cash
echo "gaia seed:"
read gaia

rly keys restore cash cash $cash
rly keys restore gaia gaia $gaia

rly light init cash -f
rly light init gaia -f

rly tx path cash-transfer-gaia

rly start cash-transfer-gaia

