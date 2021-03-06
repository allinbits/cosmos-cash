# Messages

In this section we describe the processing of the staking messages and the corresponding updates to the state. All created/modified state objects specified by each message are defined within the [state](./02_state_transitions.md) section.

## MsgCreateIdentifier

A decentralized identifier (DID) is created using the `MsgCreateIdentifier` service message.

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L12

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L20-L37

This service message is expected to fail if:

- another identifier with the same id is already registered

This service message creates and stores the `Identifier` object at appropriate indexes.


## MsgAddVerification

An verification method and one or more verification relationships are added to a decentralized identifier (DID) using the `MsgAddVerification` service message.

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L13

+++ https://github.com/allinbits/cosmos-cash/blob/main/proto/identifier/tx.proto#L37-L52

This service message is expected to fail if:

- the target did does not exists
- the verification method is invalid (according to the verification method specifications)
- the sender DID composed with the address signing the transaction is not the controller of a verification method listed in the `Authorization` verification relationships
- another verification method identifier with the same id is already registered

### Caveats :warning:

- the module does not try to resolve the verification method controller

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


--------------------

### GRPC Methods 

#### Create DID Document

**Params**
- `Id`: did uri 
- `DidDocument`: did document
- `signer`: address of the account signing the transaction

**Defaults**
- `context`: https://www.w3.org/TR/did-core/

#### Update DID Document

#### Add Verification 

#### Revoke Verification

#### Set Verification Relationships

#### Add Service 


### Command Line Client


---
hic sunt leones

--- 

## Decentralized Identifiers 

The implementation is based on the [w3c draft specifications v1.0.0](https://www.w3.org/TR/did-core/)


## Supported Messages 

### CreateIdentifier 

creates a new did with the following defaults

- context: ["https://www.w3.org/ns/did/v1"]


validation:

- id must be set (and be a valid did)
- controller must be a valid did (if set)
- upper limit of 5 Verification Methods 
- upper limit of 5 Services 

### AddVerification 

add one verification to the did

validation:

- the owner of the call must have an `Authorization` verification relationship


Design decisions:
Verification Method can exists only if they exist in a relationship 


### SetVerificationRelationships

overwrites the verification relationships of an existing verification

