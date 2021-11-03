
# Get Testnet Tokens (Faucet)

The Cosmos Cash native token denom is `cash`.  Assuming that a keypair has been created with the follwing command:

```sh
cosmos-cashd keys add alice
```

to obtain obtain `cash` tokens for the alice account from the testnet faucet use the following command:

```sh 
curl -X POST \
 -d "{\"address\": \"$(cosmos-cashd keys show alice -a)\"}" \
 https://faucet.cosmos-cash.app.beta.starport.cloud
```
