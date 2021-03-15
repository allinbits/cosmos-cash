#!/bin/bash

### create identifier on two chains A => B

cosmos-cashd keys add new --home ~/.rly-cash

cosmos-cashd tx bank send $(cosmos-cashd keys show validator --home ~/.rly-cash -a)  $(cosmos-cashd keys show new --home ~/.rly-cash -a) 10000stake --from validator --chain-id cash -y --home ~/.rly-cash --node tcp://localhost:26667

cosmos-cashd tx identifier create-identifier --from new --chain-id cash -y --home ~/.rly-cash --node tcp://localhost:26667

cosmos-cashd query identifier identifiers --home ~/.rly-cash --node tcp://localhost:26667 --output json | jq

cosmos-cashd tx ibcidentifier transfer-ibc did:cash:$(cosmos-cashd keys show new --home ~/.rly-cash -a) ibcidentifier channel-0 --from validator --chain-id cash -y --home ~/.rly-cash --node tcp://localhost:26667

cosmos-cashd query identifier identifiers --home ~/.rly-cash --node tcp://localhost:26657 --output json | jq

### create identifier on two chains B => A

cosmos-cashd keys add new --home ~/.rly-cash2

cosmos-cashd keys add auth --home ~/.rly-cash2

cosmos-cashd tx bank send $(cosmos-cashd keys show validator2 --home ~/.rly-cash2 -a) $(cosmos-cashd keys show new --home ~/.rly-cash2 -a) 10000stake --from validator2 --chain-id cash2 -y --home ~/.rly-cash2 --node tcp://localhost:26657

cosmos-cashd tx identifier create-identifier --from new --chain-id cash2 -y --home ~/.rly-cash2 --node tcp://localhost:26657

cosmos-cashd tx identifier add-authentication did:cash:$(cosmos-cashd keys show new --home ~/.rly-cash2 -a) $(cosmos-cashd keys show auth --home ~/.rly-cash2 -p) --from new --chain-id cash2 --home ~/.rly-cash2 --node tcp://localhost:26657 -y

cosmos-cashd tx ibcidentifier transfer-ibc did:cash:$(cosmos-cashd keys show new --home ~/.rly-cash2 -a) ibcidentifier channel-0 --from validator2 --chain-id cash2 -y --home ~/.rly-cash2 --node tcp://localhost:26657

cosmos-cashd query identifier identifiers --home ~/.rly-cash2 --node tcp://localhost:26657 --output json | jq

echo "cosmos-cashd query identifier identifiers --node tcp://localhost:26657 --output json | jq"
echo "cosmos-cashd query identifier identifiers --node tcp://localhost:26667 --output json | jq"
