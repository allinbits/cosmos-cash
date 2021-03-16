#!/bin/bash

### Chain 1
cosmos-cashd keys add validator --home ~/.rly-cash

cosmos-cashd init --chain-id cash cash --home ~/.rly-cash

sed -i -e 's/2665/2666/g' ~/.rly-cash/config/config.toml
sed -i -e 's#localhost:6060#localhost:6061#g' ~/.rly-cash/config/config.toml
sed -i -e 's#address = "0.0.0.0:9090"#address = "0.0.0.0:9092"#g' ~/.rly-cash/config/app.toml
sed -i -e 's#address = "0.0.0.0:9091"#address = "0.0.0.0:9093"#g' ~/.rly-cash/config/app.toml

cosmos-cashd add-genesis-account $(cosmos-cashd keys show validator -a --home ~/.rly-cash) 100000000000stake --home ~/.rly-cash --output json | jq
cosmos-cashd gentx validator 7000000000stake --chain-id cash --home ~/.rly-cash

cosmos-cashd collect-gentxs --home ~/.rly-cash

### Chain 2
cosmos-cashd keys add validator2 --home ~/.rly-cash2

cosmos-cashd init --chain-id cash2 cash2 --home ~/.rly-cash2

cosmos-cashd add-genesis-account $(cosmos-cashd keys show validator2 -a --home ~/.rly-cash2) 100000000000stake --home ~/.rly-cash2 --output json | jq
cosmos-cashd gentx validator2 7000000000stake --chain-id cash2 --home ~/.rly-cash2

cosmos-cashd collect-gentxs --home ~/.rly-cash2

#echo "Running both chains"
cosmos-cashd start --pruning nothing  false --home ~/.rly-cash &
cosmos-cashd start --pruning nothing  false --home ~/.rly-cash2 &

