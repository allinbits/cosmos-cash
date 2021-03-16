#!/bin/bash

cosmos-cashd tx ibc-transfer transfer transfer channel-1 $(cosmos-cashd keys show validator --home ~/.rly-cash -a) 15000000stake --home ~/.rly-cash2 --from validator2 --chain-id cash2 -y --node tcp://localhost:26657

cosmos-cashd tx liquidity create-pool 1 10000000stake,10000000ibc/3C3D7B3BE4ECC85A0E5B52A3AEC3B7DFC2AA9CA47C37821E57020D6807043BE9 --from validator --chain-id cash --node tcp://localhost:26667 --home ~/.rly-cash

cosmos-cashd query liquidity pools --output json --node tcp://localhost:26667 | jq

#cosmos-cashd tx liquidity withdraw 1 2000000pool/E4D2617BFE03E1146F6BBA1D9893F2B3D77BA29E7ED532BB721A39FF1ECC1B07 --from validator --chain-id cash -y

cosmos-cashd query bank balances $(cosmos-cashd keys show validator --home ~/.rly-cash -a) --node tcp://localhost:26667 --output json | jq
cosmos-cashd query bank balances $(cosmos-cashd keys show ibctransfer --home ~/.rly-cash -a) --node tcp://localhost:26667 --output json | jq

#cosmos-cashd tx liquidity deposit 1 100000000stake,100000000token --from validator --chain-id cash -y

### Swap

cosmos-cashd tx liquidity swap 1 1 999stake ibc/3C3D7B3BE4ECC85A0E5B52A3AEC3B7DFC2AA9CA47C37821E57020D6807043BE9 0.99 0.003 --from validator --chain-id cash --node tcp://localhost:26667 --home ~/.rly-cash -y

cosmos-cashd tx liquidity swap 1 1 12345ibc/3C3D7B3BE4ECC85A0E5B52A3AEC3B7DFC2AA9CA47C37821E57020D6807043BE9 stake 0.99 0.003 --from validator --chain-id cash --node tcp://localhost:26667 --home ~/.rly-cash -y
