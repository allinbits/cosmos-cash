#!/bin/bash

cosmos-cashd tx liquidity create-pool 1 100000000stake,1000ibc/3C3D7B3BE4ECC85A0E5B52A3AEC3B7DFC2AA9CA47C37821E57020D6807043BE9 --from validator --chain-id cash -y

cosmos-cashd query liquidity pools --output json | jq

#cosmos-cashd tx liquidity deposit 1 100000000stake,100000000token --from validator --chain-id cash -y

#cosmos-cashd query bank balances $(cosmos-cashd keys show validator -a) --output json | jq

#cosmos-cashd tx liquidity withdraw 1 2000000pool/E4D2617BFE03E1146F6BBA1D9893F2B3D77BA29E7ED532BB721A39FF1ECC1B07 --from validator --chain-id cash -y

#cosmos-cashd query bank balances $(cosmos-cashd keys show validator -a) --output json | jq

#cosmos-cashd tx liquidity deposit 1 100000000stake,100000000token --from validator --chain-id cash -y

### Swap

#cosmos-cashd tx liquidity swap 1 1 1000stake token 0.99 0.003 --from validator --chain-id cash -y

