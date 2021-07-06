# State

This document describes the state pertaining to:

1. [DidDocument](./02_state.md#identifier)
2. [Verification](./02_state.md#authentication)
3. [Service](./02_state.md#service)


Three data structues represent a DidDocument

- DidDocument
- Verification
- Service

## DidDocument
DidDocuments are stored in the state under the `0x61` key and are stored using their ids

- DidDocument: `0x61 | DidDocument.Id -> ProtocolBuffer(DidDocument)`

### Structure


## Verification
Authentication is stored as a slice under in the DidDocument, the authentication data structure contains 4 fields and is used to store public key information. [[more_info]](https://w3c.github.io/did-core/#authentication)

### Structure
+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/identifier.proto#L38


## Service
A Service is stored as a slice under in the DidDocument data structure, the service data structure has four fields. Services are used to get data about the did subject. [[more_info]](https://w3c.github.io/did-core/#services)

### Structure
+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/identifier.proto#L41


