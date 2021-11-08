# Cosmos Cash

[![Go Reference](https://pkg.go.dev/badge/github.com/allinbits/cosmos-cash.svg)](https://pkg.go.dev/github.com/allinbits/cosmos-cash)
[![build](https://github.com/allinbits/cosmos-cash/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/allinbits/cosmos-cash/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/allinbits/cosmos-cash/branch/main/graph/badge.svg?token=NLT5ZWM460)](https://codecov.io/gh/allinbits/cosmos-cash)
[![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/allinbits/cosmos-cash)](https://libraries.io/go/github.com%2Fallinbits%2Fcosmos-cash)

![](https://miro.medium.com/max/1000/1*8Wx44uvyJxpZUVS0WojMNw.png)

### Summary

Cosmos Cash is a protocol designed to be regulatory compliant that offers the same guarantees as traditional banking
systems. Features that enable these guarantees are Know Your Customer (KYC), anti-money laundering (AML) tracking, Financial Action Task Force (FATF) travel rule, and identity management. Cosmos Cash uses a novel approach to identity management by leveraging W3C specifications
for decentralized identifiers and verifiable credentials.

### Research paper

For more information on the research behind the Cosmos Cash protocol, please look at the Cosmos Cash research paper:

[Cosmos Cash: Investigation into EU regulations affecting E-Money tokens](https://drive.google.com/file/d/1zmEyA8kA0uAIRGDKxYElOKvjtz4f_Ep5/view)

### Architecture

The Cosmos Cash approach leverages open standards to reach its goals and to offer an open model that is compatible with
third-party projects that use the open standards. In particular, the Cosmos Cash project uses:

- Self-sovereign identity ([SSI](./Reference/GLOSSARY.md#self-sovereign-identity-ssi))
- Decentralized identifier ([DID](./Reference/GLOSSARY.md#decentralized-identifier-did))
- Verifiable credentials ([VC](./Reference/GLOSSARY.md#verifiable-credential-vc))
- Zero-knowledge proofs ([ZKP](./Reference/GLOSSARY.md#zero-knowledge-proof-zkp))

For a detailed architecture description and design choices, visit the [ADR](./docs/Explanation/ADR) section.

### Documentation

The Cosmos Cash documentation is available on the [Cosmos Cash Documentation Portal](https://docs.cosmos-cash.app.beta.starport.cloud/).

Technical reference is bundled within each module, visit the [Cosmos Cash Modules](./docs/Reference/MODULES.md) page for links
to individual modules.

Links to presentations, discussions and interviews are available in
[Presentations](./docs/Explanation/presentations.md).

### Getting Started

To get started and contribute to the project, see [Technical Setup](./TECHNICAL-SETUP.md) and the
[Contributing](./CONTRIBUTING.md) page.


#### Testnet

To join or interact using the testnet, see [testnet coordinates](https://docs.cosmos-cash.app.beta.starport.cloud/Networks/testnet/)

--- 

Do you have questions or want to get in touch? Send us an email at *cosmos-cash@tendermint.com*.
