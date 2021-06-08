# Cosmos Cash

[![build](https://github.com/allinbits/cosmos-cash/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/allinbits/cosmos-cash/actions/workflows/ci.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/allinbits/cosmos-cash.svg)](https://pkg.go.dev/github.com/allinbits/cosmos-cash)
[![codecov](https://codecov.io/gh/allinbits/cosmos-cash/branch/main/graph/badge.svg?token=NLT5ZWM460)](https://codecov.io/gh/allinbits/cosmos-cash)
[![Go Report Card](https://goreportcard.com/badge/github.com/allinbits/cosmos-cash)](https://goreportcard.com/report/github.com/allinbits/cosmos-cash)

![](https://miro.medium.com/max/1000/1*8Wx44uvyJxpZUVS0WojMNw.png)


### Summary

Cosmos Cash is a protocol designed to be regulatory compliant that offers the same guarantees as traditional banking systems. Features that enable these guarantees are KYC (Know your customer), AML (Anti-Money Laundering) tracking, FAFT travel rule, and identity management. We use a novel approach to identity management by leveraging W3C specifications for decentralized identifiers and verifiable credentials.

### Research paper

For more information on the research behind the cosmos cash protocol, please look at the cosmos cash research paper:

[Cosmos Cash: Investigation into EU regulations affecting E-Money tokens](https://drive.google.com/file/d/1zmEyA8kA0uAIRGDKxYElOKvjtz4f_Ep5/view)

### How to build

- `make install`
- `cosmos-cashd -h`

### How to run the chain

- `make start-dev`

### How to seed the chain

- `make seed`

### How to test

- `make test`

## Testnet 

The Cosmos Cash testnet coordinates are:

- **Chain ID**: `cosmoscash-testnet`
- **Token Denom**: `cash`
- **Genesis**: [cosmos-cash.app.beta.starport.cloud/genesis](https://cosmos-cash.app.beta.starport.cloud/genesis?)
- **RPC URL**:  `https://rpc.cosmos-cash.app.beta.starport.cloud:443`
- **Explorer URL**: TBD
