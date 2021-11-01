# Join Testnet

A Cosmos Cash testnet is available for testing purposes, this document describes the Cosmos Cash testnet coordinates and the 
satellite projects URLs.


## Testnet Coordinates

|              |                                                                                                     |
| ------------ | --------------------------------------------------------------------------------------------------- |
| Chain ID     | `cosmsoscash-testnet`                                                                               |
| Token  Denom | `cash`                                                                                              |
| Genesis File | [cosmos-cash.app.beta.starport.cloud/genesis](https://cosmos-cash.app.beta.starport.cloud/genesis?) |
| RPC URL      | `https://rpc.cosmos-cash.app.beta.starport.cloud:443`                                               |


Faucet URL:  `https://faucet.cosmos-cash.app.beta.starport.cloud`

```sh
curl -X POST \
    -d "{\"address\": \"$(cosmos-cashd keys show YOUR_WALLET_UID -a)\"}" \
    https://faucet.cosmos-cash.app.beta.starport.cloud
```

## DID Resolver

|                       |                                                                         |
| --------------------- | ----------------------------------------------------------------------- |
| DID resolver driver   | https://resolver-driver.cosmos-cash.app.beta.starport.cloud/identifier/ |
| DID resolver frontend | https://uniresolver.cosmos-cash.app.beta.starport.cloud                 |


## SSI Agent Router

|                         |                                                                  |
| ----------------------- | ---------------------------------------------------------------- |
| Router Inbound Endpoint | https://in.agent.cosmos-cash.app.beta.starport.cloud/identifier/ |
| Router WSS              | https://ws.agent.cosmos-cash.app.beta.starport.cloud             |