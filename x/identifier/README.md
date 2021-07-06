# DID Module 


This module implements the DID core spec. For more information please see the [DID W3C core-spec](https://w3c.github.io/did-core/).

Decentralized identifiers (DIDs) are a type of identifier that enables verifiable, decentralized digital identity. A DID identifies any subject (e.g., a person, organization, thing, data model, abstract entity, etc.) that the controller of the DID decides that it identifies.

A DID is a globally unique persistent identifier that does not require a centralized registration authority and is often generated and/or registered cryptographically.

## Overview 

The DID module provides the functionalities to create and manage the lifecycle of a DID document on chain.

> :warning: **The implementation default DID method is did:cash**

The lifecycle of the DID document is represented in the following sequence diagram ([src board](https://whimsical.com/did-module-tx-process-263fkjvvNa939A1yjzZUep)):

![DID Creation](./docs/assets/DID%20Module%20Tx%20process.png)

## Reference

1. [State](./docs/02_state.md)
2. [Transitions](./docs/03_state_transitions.md)
3. [Messages](./docs/04_messages.md)

