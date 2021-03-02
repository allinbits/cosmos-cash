# Messages

In this section we describe the processing of the staking messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state_transitions.md) section.

## MsgCreateIdentifier

A decentralized identifier (DID) is created using the `MsgCreateIdentifier` service message.

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L12

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L20-L37

This service message is expected to fail if:

- another identifier with the same id is already registered

This service message creates and stores the `Identifier` object at appropriate indexes.


## MsgAddAuthentication

An authentication method is added to a decentralized identifier (DID) using the `MsgAddAuthentication` service message.

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L13

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L37-L52

This service message is expected to fail if:

- sender is not associtated with the given identifier id
- another identifier with the same id is already registered

This service message adds an authentication method to an `Identifier` and stores the `Identifier` object at the appropriate index.

## MsgDeleteAuthentication

An authentication method is removed from a decentralized identifier (DID) using the `MsgDeleteAuthentication` service message.

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L15

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L69-L80

This service message is expected to fail if:

- cannot find the given identifier id
- sender is not associtated with the given identifier id
- there will be less than 1 authentication method left in the identifier
- a given public key cannot be decoded 

This service message deletes an authentication method from a `Identifier` and stores the `Identifier` object at the appropriate index.

## MsgAddService

A service is added to a decentralized identifier (DID) using the `MsgAddService` service message.

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L14

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L54-L67

This service message is expected to fail if:

- cannot find the given identifier id

This service message adds a service to an `Identifier` and stores the `Identifier` object at the appropriate index.

## MsgDeleteService

A service is delted from a decentralized identifier (DID) using the `MsgDeleteService` service message.

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L16

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L82-L93

This service message is expected to fail if:

- cannot find the given identifier id

This service message deletes a service from an `Identifier` and stores the `Identifier` object at the appropriate index.
