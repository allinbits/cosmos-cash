# Trust over IP

- [Trust over IP](#trust-over-ip)
  - [Summary](#summary)
  - [ToIP](#toip)
  - [Why is it important?](#why-is-it-important)
  - [Features](#features)
    - [Technology stack](#technology-stack)
      - [DID Utilities](#did-utilities)
      - [DIDComm Protocol](#didcomm-protocol)
      - [Data Exchange protocols](#data-exchange-protocols)
      - [Application Ecosystem](#application-ecosystem)
    - [Governance Stack](#governance-stack)
      - [Public Utility Governance](#public-utility-governance)
      - [DIDComm Protocol Governance](#didcomm-protocol-governance)
      - [Credential Governance](#credential-governance)
      - [Ecosystem Governance](#ecosystem-governance)
  - [Conclusions](#conclusions)
  - [References](#references)

## Summary 

This article is a summary of the [Trust Over IP](https://trustoverip.org/). ToIP is important because it defines the governance framework for a Cosmos Cash identity chain. 

## ToIP

ToIP is a meta model for a governance framework based on Self-Sovereign Identity. It describes a set of desired governance features that can be implemented in a number of ways based on technology, identity domain and business case. Examples of specific governance frameworks include

- Sovrin
- etc

## Why is it important?

Many examples of governance frameworks already exist. For example, Mastercard etc have governance frameworks to protect stakeholders, define behaviour protocols, build confidence in the network, resolve disputes etc.

Blockchain Technology such as Tendermint may be a trusted network built out of trustless components, but doing the same for humans involves some kind of the governance.

## Features

ToIP is a model based on four layers and two pillars: Technology Stack and Governance Stack.

[![](https://www.evernym.com/wp-content/uploads/2020/05/ToIP-stack-1024x631.jpg)]()

Each layer of the technology stack is supported by an equivalent governace process. Often feature of the governance could be implemented in the technology.


### Technology stack

The layers of the technology stack serve two distinct purposes:

- **Cryptographic trust** through
  - **DID utilities**, this corresponds to the DLT or other Verifiable Data Registry technology that support DID's and 
  - **DIDComm protocol**
- **Human trust** through
  - **Data Exchange Protocols**
  - **Application Ecosystems**

The following section will describe each in turn:

#### DID Utilities

This is the lowest level of the technology stack. This level covers the implementation of the W3C Decentralized Identity. This includes

- [Decentralized Identifiers (DIDs)](https://www.w3.org/TR/did-core/) fromats
- [DID Methods](https://www.w3.org/TR/did-core/#methods) which define how DID's is resolved, updated and deactivated.
- DID resolvers which take a DID and resolve it to a DID document
 
These are often implemented on blockchain, Decentralized Ledger Technologies or decentralized file storage.

#### DIDComm Protocol

This is the technology that underpins off-chain peer-to-peer identity interactions. In particular it refers to

- `did:peer` is a specific implementation of the W3c DID method standard that doesn't require the use of DID's hosted on a public utility.
- The [DIDComm protocol](https://identity.foundation/didcomm-messaging/spec/) is the secure communications layer.

This layer includes the hardware and software that supports these protocols. This includes

- Agent and controller software
- DID resolvers where needed to connect to public DIDs
- Key Management Systems (also known as wallets)
- Secure data storage, where data is secured through encryption 

Key observations include:

- DID's are generated from keys in the KMS
- `did:peer` conenctions can be broken if one of peers no longer needs the connection
- Each agent is paired with **ONE** wallet/KMS
- Can have more than one secure store and the stores can be synchronized

#### Data Exchange protocols

This is the layer that establishes human trust through the use of Verifiable Credentials. It comprises three parts:

- [W3C Verifiable Credential Data model](https://www.w3.org/TR/vc-data-model/) - the "Trust Triangle" of Holder, Issuer and Verifier
- The Credential proof types - this includes the following options:
    - JSON Web Tokens
    - Linked Data signatures
    - Zero Knowledge Proofs
- [Credential Exchange protocols](https://www.w3.org/TR/vc-imp-guide/) - this uses DIDComm as layer for the issuance and verification of credentials. The implementation depends on the proof type because each proof type has a different request/response model.

[![](https://freecontent.manning.com/wp-content/uploads/the-building-blocks-of-ssi_03.png)]()

#### Application Ecosystem

This is rather hazy as a definition, but is based on the analogy that ToIP applications are similar to any application that uses TCP/IP to communicate.

The more practical example would be Hyperledger Aries architecture, whereby a "Controller" is the application layer and uses Aries as the ToIP interactions.

### Governance Stack

Each level of the governance pillar comprises rules and standards that support each corresponding level of the technology stack. 

One view is to consider this as the rules that

#### Public Utility Governance

Layer One is the governance that support the technology stack at the public utility level. For a system such as Cosmos Cash this will be based on a public, permissionless blockchain using delegated proof-of-stake model to incentivize validators to secure the chain.

However, there do need to be some consideration about who has a DID's on a public utility. At the a minimum this will be Credential issuers and governance authorities. However, Sovrin makes the point that if the blockchain is public and permissionless then any private individual can write a public DID to chain, but it can't be deleted, which means there is an infringement of GDPR's right to be forgotten rules. Sovrin address this through the use of a TAA, which has to be signed and stored and chain to ensure that users understand their rights (or foregoing in this instance).

#### DIDComm Protocol Governance

DIDComm protocol is an open standard maintained by the [Decentralized Identity Foundation](https://identity.foundation/) (DIF). However there is much in terms of best practice and governance that is required, specifically

- **Hardware developers** who build secure enclaves, trusted execution environments, [Hardware Security Modules](https://en.wikipedia.org/wiki/Hardware_security_module) etc
- **Software developers** who develop ToIP compliant agents, controllers, Key Management Systems and secure data stores

It also applies to those parties that offer services for other actors in the idenity system, specifically:

- **Agencies** who host ToIP compliant cloud agents,  
- **Guardians** who provide agent/wallet services for a another party. This role comes with specific issues since the hosting of keys and other cryptographic materials involves risks (operational, legal & regulatory) since there is a vital link between keys to credentials and decentralized identity. 

#### Credential Governance

Outside the W3C specification, there are the additional rules that define policies that support the issuance and verfification of Verifiable Credentials. These policies will also define the interoperability of such credentials. Topics include

- Who maintains and manages Verifiable Credential schema definitions
- Rules & metadata regarding who can act as an **Authoritative Issuers** i.e. who can issue particular type of credential
- Policies for the issuance and revocation of credentials
- Authoritative Issuers, such that bad actors can be have issuance rights revoked.  
- Business models, liability frameworks and insurance models

Depending on the identity domain, this framework also covers topics such as **credential registries**. These are verifiable credentials stored on a public utility. 

#### Ecosystem Governance

This layer of governance adds the Governance Authority to the triangle of trust. 

[![](https://freecontent.manning.com/wp-content/uploads/the-building-blocks-of-ssi_13.png)]()

This is the authority that defines legal and techncial rules under which members agree to operate so that everyone has trust in the network. Artefacts may include


Threat include

- Issuers issuing untrustworthy credentials
- Verfiers coercing credential holders to reveal credentials when they are not needed and reveal more information that is required

Credential governance


## Conclusions

This topic document outlines at a high level the main principles of Trust Over IP as a meta model. This [topic paper](cosmos_cash_toip.md) will define how these could be defined to Cosmos Cash Identities.



## References

- [Trust Over IP](https://trustoverip.org/) 
- [DIDComm protocol](https://identity.foundation/didcomm-messaging/spec/)
- [W3C Verifiable Credential Data model](https://www.w3.org/TR/vc-data-model/)
- [Verifiable Credential Implementation](https://www.w3.org/TR/vc-imp-guide/)
- [Aries RFCs](https://github.com/hyperledger/aries-rfcs/tree/main/features)
- [DID Peer Method specification](https://identity.foundation/peer-did-method-spec/)
- [Sovrin TAA Agreement](https://sovrin.org/wp-content/uploads/Transaction-Author-Agreement-V1.1.pdf)
- [Preparing for TAA](https://sovrin.org/preparing-for-the-sovrin-transaction-author-agreement/)



