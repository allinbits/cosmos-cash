# State Transitions

This document describes the state transitions pertaining to:

1. [Identifier || DidDocument](./02_state.md#identifier)
2. [Authentication](./02_state.md#authentication)
3. [Service](./02_state.md#service)

## Identifier || DidDocument

### Setting a decentralized identifier (DID) in the store
+++ https://github.com/allinbits/cosmos-cash/blob/main/x/identifier/keeper/identifier.go#L8-L10

### Adding an identifier
+++ https://github.com/allinbits/cosmos-cash/blob/main/x/identifier/keeper/msg_server.go#L24-L43


## Authentication
Authentication methods are added and removed from DidDocuments, we edit the structure then set the identifier in the store

### Adding an Authentication method

+++ https://github.com/allinbits/cosmos-cash/blob/main/x/identifier/keeper/msg_server.go#L45-L74

### Deleting an Authentication method

+++ https://github.com/allinbits/cosmos-cash/blob/main/x/identifier/keeper/msg_server.go#L98-L144 

## Service
Services are added and removed from DidDocuments, we edit the structure then set the identifier in the store


### Adding a Service

+++ https://github.com/allinbits/cosmos-cash/blob/main/x/identifier/keeper/msg_server.go#L77-L95

### Deleting a Service

+++ https://github.com/allinbits/cosmos-cash/blob/main/x/identifier/keeper/msg_server.go#L147-L191
