
# Get Testnet Tokens (Faucet)

The Cosmos Cash native token denom is `cash`.  

A keypair is created with the following command:

```sh
cosmos-cashd keys add alice
```

To obtain obtain `cash` tokens for the alice account from the testnet faucet, use the following command:

```sh 
curl -X POST \
 -d "{\"address\": \"$(cosmos-cashd keys show alice -a)\"}" \
 https://faucet.cosmos-cash.app.beta.starport.cloud
```
