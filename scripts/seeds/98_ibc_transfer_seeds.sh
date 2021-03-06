#!/bin/bash

cosmos-cashd keys add ibctransfer --home ~/.rly-cash2

cosmos-cashd tx ibc-transfer transfer transfer channel-1 $(cosmos-cashd keys show ibctransfer --home ~/.rly-cash -a) 100000stake --home ~/.rly-cash --from validator --chain-id cash -y --node tcp://localhost:26667

echo "Balance on cash chain"
cosmos-cashd query bank balances $(cosmos-cashd keys show ibctransfer --home ~/.rly-cash -a) --node tcp://localhost:26667

echo "Balance on cash2 chain"
cosmos-cashd query bank balances $(cosmos-cashd keys show ibctransfer --home ~/.rly-cash -a) --node tcp://localhost:26657

cosmos-cashd tx ibc-transfer transfer transfer channel-1 $(cosmos-cashd keys show ibctransfer --home ~/.rly-cash -a) 500000stake --home ~/.rly-cash2 --from validator2 --chain-id cash2 -y --node tcp://localhost:26657

echo "Balance on cash chain"
cosmos-cashd query bank balances $(cosmos-cashd keys show ibctransfer --home ~/.rly-cash -a) --node tcp://localhost:26667

cosmos-cashd tx bank send $(cosmos-cashd keys show ibctransfer --home ~/.rly-cash -a) $(cosmos-cashd keys show validator --home ~/.rly-cash -a) 1000ibc/3C3D7B3BE4ECC85A0E5B52A3AEC3B7DFC2AA9CA47C37821E57020D6807043BE9 --from ibctransfer --chain-id cash -y --home ~/.rly-cash --node tcp://localhost:26667

cosmos-cashd tx ibc-transfer transfer transfer channel-1 $(cosmos-cashd keys show ibctransfer --home ~/.rly-cash -a) 1000ibc/3C3D7B3BE4ECC85A0E5B52A3AEC3B7DFC2AA9CA47C37821E57020D6807043BE9 --home ~/.rly-cash --from validator --chain-id cash -y --node tcp://localhost:26667

echo "cosmos-cashd query bank balances $(cosmos-cashd keys show ibctransfer --home ~/.rly-cash -a) --node tcp://localhost:26657"
echo "cosmos-cashd query bank balances $(cosmos-cashd keys show ibctransfer --home ~/.rly-cash -a) --node tcp://localhost:26667"
