#!/bin/bash

### Initialisation
gaiad init gaia --chain-id gaia
echo "y" | gaiad keys add validator --keyring-backend test
gaiad add-genesis-account $(gaiad keys show validator -a  --keyring-backend test) 1000000000stake
gaiad gentx validator 700000000stake --chain-id gaia --keyring-backend test
gaiad collect-gentxs

echo "Run: gaiad start"

### Update config.toml to expose rpc
sed -i -r 's#laddr = "tcp://127.0.0.1:26657"#laddr = "tcp://0.0.0.0:26657"#' ~/.gaia/config/config.toml

gaiad start --pruning=nothing

