#!/bin/bash

### Initialisation
cosmos-cashd init cash --chain-id cash
echo "y" | cosmos-cashd keys add validator
cosmos-cashd add-genesis-account $(cosmos-cashd keys show validator -a) 1000000000stake
cosmos-cashd gentx validator 700000000stake --chain-id cash
cosmos-cashd collect-gentxs

echo "Run: cosmos-cashd start"

### Update config.toml to expose rpc
sed -i -r 's#laddr = "tcp://127.0.0.1:26657"#laddr = "tcp://0.0.0.0:26657"#' ~/.cosmoscash/config/config.toml

cosmos-cashd start --pruning=nothing
