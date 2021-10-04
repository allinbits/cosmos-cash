GENESIS_FILE=~/.cosmoscash/config/genesis.json

#remember:
#from root:
#go install ./cmd/cosmos-cashd

# "Adding genesis account"
cosmos-cashd init --chain-id=cash cash
cosmos-cashd keys add validator
cosmos-cashd keys add regulator
cosmos-cashd keys add emti # e-money token issuer 
cosmos-cashd keys add arti # asset-referenced token issuer 
cosmos-cashd keys add bob
cosmos-cashd keys add alice

# "Adding genesis account"
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


# "Starting cosmos cash chain"
cosmos-cashd start