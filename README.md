# Cosmos Cash

[![Go Reference](https://pkg.go.dev/badge/github.com/allinbits/cosmos-cash.svg)](https://pkg.go.dev/github.com/allinbits/cosmos-cash)
[![build](https://github.com/allinbits/cosmos-cash/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/allinbits/cosmos-cash/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/allinbits/cosmos-cash/branch/main/graph/badge.svg?token=NLT5ZWM460)](https://codecov.io/gh/allinbits/cosmos-cash)
[![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/allinbits/cosmos-cash)](https://libraries.io/go/github.com%2Fallinbits%2Fcosmos-cash)

![](https://miro.medium.com/max/1000/1*8Wx44uvyJxpZUVS0WojMNw.png)


### Summary

Cosmos Cash is a protocol designed to be regulatory compliant that offers the same guarantees as traditional banking systems. Features that enable these guarantees are KYC (Know your customer), AML (Anti-Money Laundering) tracking, FAFT travel rule, and identity management. We use a novel approach to identity management by leveraging W3C specifications for decentralized identifiers and verifiable credentials.

### Research paper

For more information on the research behind the cosmos cash protocol, please look at the cosmos cash research paper:

[Cosmos Cash: Investigation into EU regulations affecting E-Money tokens](https://drive.google.com/file/d/1zmEyA8kA0uAIRGDKxYElOKvjtz4f_Ep5/view)


### Architecture

The Cosmos Cash approach leverage open standards to reach it's goals and to offer an open model compatible with 3rd party projects that uses the open standards. In particular we use 
- SSI 
- DID
- Verifiable Credentials
- Zero knowledge proof  

For a detailed architecture description and design choices, visit the [ADR](./docs/Explanation/ADR) section. 

### Documentation

The Cosmos Cash documentation is available under the [docs](./docs) folder. 

Technical reference is bundled within each module, visit the [module index](./docs/Reference/MODULES.md) page for links to individual modules.

Links to presentations, discussions and interviews are available in the [presentations page](./docs/Explanation/presentations.md).

### Getting started 

To get start and contribute to the project visit the [technical setup](./TECHNICAL-SETUP.md) page and the [contributing](./CONTRIBUTING.md) page.


A Cosmos Cash testnet is available at this coordinates:

- **Chain ID**: `cosmoscash-testnet`
- **Token Denom**: `cash`
- **Genesis**: [cosmos-cash.app.beta.starport.cloud/genesis](https://cosmos-cash.app.beta.starport.cloud/genesis?)
- **RPC URL**:  `https://rpc.cosmos-cash.app.beta.starport.cloud:443`

A Cosmos Cash DID resolver endpoint (compatible with the [universal resolver](https://github.com/decentralized-identity/universal-resolver) driver specifications) is available at:

```
https://api.cosmos-cash.app.beta.starport.cloud/identifier/
```

A universal resolver frontend is also available for testing and verification purposes at:
- https://uniresolver.cosmos-cash.app.beta.starport.cloud

--- 

Have questions or want to get in touch? drop an email at *cosmos-cash@tendermint.com*
