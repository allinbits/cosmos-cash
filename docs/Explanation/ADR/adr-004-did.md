# ADR 4: DID

## Changelog

- 2021-08-02: Initial draft

## Status

DRAFT
## Abstract

[Decentralized identifiers](https://www.w3.org/TR/did-core) (DIDs) are a type of identifier that enables verifiable, decentralized digital identity. A DID refers to any subject (for example, a person, organization, thing, data model, abstract entity, and so on) as determined by the controller of the DID.

This document specifies the DID method for a Cosmos SDK-based implementation of the W3C recommendation, its properties, operations, and an explanation of the process to resolve DIDs to the resources that they represent. 

## Context

The aim of the Cosmos Cash project is to provide a state-of-the-art platform for the hosting of collateralized stable coins that is compliant with:

 - EU regulations such as General Data Protection Regulation (GDPR) and Markets in Crypto-Assets (MiCA)
 - International recommendations such as the Financial Action Task Force (FATF) "Travel Rule"
 - Local anti-money laundering (AML) regulations

The Cosmos Cash platform is based on the following principles:

- Open financial infrastructure is a public good
- Money laundering prevention also benefits society
- Users should benefit from a strict privacy-respecting approach (GDPR)

The self-sovereign identity (SSI) approach to tackling the identity and privacy challenge has been gaining momentum in recent years. Coupled with distributed ledger technology (DLT) technology, the SSI approach has been capturing the attention of both the private and public sectors. 

The SSI approach relies on two building blocks: decentralized identifiers (DID) and verifiable credentials (VC). This architecture decision record (ADR) describes the DID implementation in a Cosmos SDK-based blockchain.

The goal of this ADR is to define a foundation for the necessary components to realize the Cosmos Cash objectives while ensuring the implementation of the DID is fully compliant with the W3C specifications. **Successive iterations will address API ergonomics and standard compatibility issues.** 

## Decision


The Cosmos Cash implementation for DIDs will follow the [DID W3C core recommendations](https://github.com/w3c/did-core) with the goal of maximizing compatibility with 3rd party tools and projects.


### DID Method Name

The namestring that shall identify the Cosmos Cash DID method is: `cosmos`.

A DID that uses the Cosmos Cash method MUST begin with the following prefix: `did:cosmos`. Per the [W3C DID specification](https://www.w3.org/TR/did-core), this prefix string MUST be in lowercase. The remainder of the DID, after the prefix, is as follows:

#### Method Specific Identifier


The namespace specific identifier is defined by the following ABNF:

```ABNF
cosmos-did                = "did:cosmos:" identifier-type
identifier-type           = cosmos-key / cosmos-network
cosmos-key                = "key:" 1*255id-char "1" 20*255HEXDIG 
cosmos-network            = "net:" 1*255id-char ":" unique-identifier
unique-identifier         = 38*255id-char
id-char                   = ALPHA / DIGIT / (ALPHA "-") / (DIGIT "-")
```

For the `unique-identifier` it is RECOMMENDED to use a UUID.

The `identifier-type` distinguish between two did types: 

1. the `key` type (inspired from the [`did:key`](https://w3c-ccg.github.io/did-method-key/) method)
2. the `net` type that identifies did and the originating network of the did

DIDs of `key` type are ephemeral and immutable: they are are always generated from the cosmos blockchain address they refer to. The following are examples of DIDs of `key` type 

- `did:cosmos:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj` 
- `did:cosmos:key:cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2`

DIDs of `net` type are persistent and mutable: they are stored in the node database and can be created and updated according to rules described in the [DID Operations](#did-operations) section. The following are the examples of DIDs with `net` type:

- `did:cosmos:net:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb`
- `did:cosmos:net:cosmoshub:1cc7813c-bb31-4999-a768-19424e6c10fa`


##### DID Operations

DID and associated DID documents are managed by a Cosmos SDK module that uses the gRPC communication protocol. This section defines the [CRUD operations](https://www.w3.org/TR/did-core/#method-operations) for a Cosmos DID. 

###### Create

To create and publish a DID document use the message 

```golang
MsgCreateDidDocument(id string, signerPubKey string)
```

The message parameters are the DID to be created and the `signerPubKey` MUST be the public key of the account signing the transaction. The public key MUST be used to attach a verification method of type `EcdsaSecp256k1RecoveryMethod2020` with the value of `publicKeyHex` containing the public key encoded in hexadecimal.

The verification method controller MUST be one of the following

- the DID of the document
- the key DID of the address signing the transaction

The verification method id SHOULD be generated as:

```
{verificationMethodController}#{CosmosAddressFromPubKey}
```

The verification method id MUST be listed in the `authentication` relationships. 

If the input DID is not a valid DID for the Cosmos method, or if the DID already exists on-chain, the message returns an error. 

Contextually with the creation of a DID document, a [DID document metadata](#did-document-metadata) MUST be created with the following values:

- the hash of the transaction as `versionId`
- the block time for the `created` and `updated` fields
- `false` for the `deactivated` field


To address privacy concerns:

- Do not use an id that is the same as the blockchain account address
- Isolate the verification methods to the DID subject (for example, during key rotation)


> **Note:** A more fine-grained DID creation method can be implemented with the goal of saving in gas by executing a single transaction in a complex DID scenario.


###### Resolve and Verify

The integrity of the DID documents stored on the ledger is guaranteed by the underlying blockchain protocol. 

A DID can be resolved using the gRPC message:

```golang
QueryDidDocumentRequest(did string)
```

Example of a did document resolved via gRPC interface:

```javascript
{
  "didDocument": {
    "context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:net:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
    "controller": [
      "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
    ],
    "verificationMethod": [
      {
        "controller": "did:cosmos:net:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
        "id": "did:cosmos:net:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1RecoveryMethod2020"
      }
    ],
    "authentication": [
      "did:cosmos:net:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ]
  },
  "didMetadata": {
    "versionId": "9f7c547dc852af60c9da1fd514e1497d407b6a3d8ae3e52b626d536519dc8f4c",
    "created": "2021-08-23T08:24:26.972761898Z",
    "updated": "2021-08-24T15:54:40.902858856Z",
    "deactivated": false
  }
}
```



Or via a REST endpoint that MUST be compatible with the [W3C DID recommendations](https://www.w3.org/TR/did-core) and pass the [did test suite](https://w3c.github.io/did-test-suite/):

```
{NODE_URL}:{NODE_REST_PORT}/identifier/{did}
```

Example of a did document resolved via a REST endpoint:

```javascript
{
   "didDocument":{
      "@context":[
         "https://www.w3.org/ns/did/v1"
      ],
      "id":"did:cosmos:net:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
      "controller":[
         "did:cosmos:key:cosmos1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"
      ],
      "verificationMethod":[
         {
            "controller":"did:cosmos:net:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e",
            "id":"did:cosmos:net:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
            "publicKeyHex":"0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
            "type":"EcdsaSecp256k1RecoveryMethod2020"
         }
      ],
      "authentication":[
         "did:cosmos:net:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e#cosmos1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
      ]
   },
   "didDocumentMetadata":{
      "versionId":"4f0f8857ab36bdeee0ddb541ea7e7b9cb509d29e1103cc7def44d3d1e8220c22",
      "created":"2021-08-23T08:24:26.972761898Z",
      "updated":"2021-08-24T15:54:40.902858856Z"
   },
   "didResolutionMetadata":{
      "contentType":"application/ld+json",
      "did":{
         "method":"cosmos",
         "methodSpecificId":"net:cosmoscash-testnet:900d82bc-2bfe-45a7-ab22-a8d11773568e"
      }
   }
}
```


###### Update

There are two ways of updating a DID document:

- Manage DID controllers
- Manipulate verification methods and relationships 

In both cases, the target DID must exist on-chain and the `signerAccount` must exist as a verification method in a verification relationship of type `authentication` or it's `key` DID be listed as a DID controller. Additionally the DID document MUST NOT be in a deactivated status.

Each update operation MUST update the `versionId` and the `updated` field of the associated [DID document metadata](#did-document-metadata) with the transaction hash and the block time respectively.

**Manage DID Controllers** 

Set the DID controllers using the gRPC message:

```golang
MsgUpdateDidDocument(did string, controllers []string, signerAccount string)
``` 

The parameters are as follows:

 - `did` identifies the did document
 - `controllers` are a list of cosmos addresses that will replace the DID document controllers list 
 - `signerAccount` is the account address that is signing the transaction

Controllers will be added using the `did:cosmos:key:` method type. A controller of a DID document MUST be a did of type `key`.

**Manipulate Verification Methods and Relationships**

Add a new verification method using the gRPC message:

```golang
MsgAddVerification(did string, accountId string, relationships []string, signerAccount string)
```

The parameters are as follows:

 - `did` identifies the did document
 - `accountId` is the account to be added to the verification method
 - `relationships` is the list of relationships that the `accountId` will be registered into
 - `signerAccount` is the account address that is signing the transaction

The list of relationships must contain only valid [relationships names](#verification-relationships) as defined in the DID document:


Set the relationships of a verification method using the gRPC message:

```
MsgSetVerificationRelationships(did string, accountId string, relationships []string, signerAccount string)
```

The list of relationships MUST contain only valid [relationships names](#verification-relationships) and MUST be non empty.


**Services**

A service MUST be an entity with the following properties:
  
- `id`: a valid RFC3986 URI string. 
- `type`: a non empty string.
- `serviceEndpoint`: a valid RFC3986 URI string.  

A service can be added using the gRPC method:

```golang
MsgAddService(did string, service_data Service, signerAccount string)
```

The `id` of a service MUST be unique within the DID document.


A service can be deleted using the gRPC message:

```golang
MsgDeleteService(did string, service_id string, signerAccount string)
```


###### Deactivate

A DID can be deactivated using the gRPC message:

```golang
MsgDeactivateDid(did string, signerAccount string)
```

The operation MUST update the DID document metadata and set the `deactivated` value to true. The operation is not reversible.


### Method-specific Properties 

#### DID-core Verification Material

The [Verification Material](https://www.w3.org/TR/did-core/#verification-material) type MUST support

- type `EcdsaSecp256k1RecoveryMethod2020` with `pubKeyHex` to encode a cosmos account public key in hexadecimal format
- type `CosmosAccountAddress` with `blockchainAccountID` to represent a cosmos account 


Support for other verification materials can be added. 

#### Verification Relationships

The DID document MUST support the following [verification relationships](https://www.w3.org/TR/did-core/#verification-relationships):

- [`authentication`](https://www.w3.org/TR/did-core/#authentication) - authorizes amends to the DID document
- [`assertionMethod`](https://www.w3.org/TR/did-core/#assertion)
- [`keyAgreement`](https://www.w3.org/TR/did-core/#key-agreement)
- [`capabilityInvocation`](https://www.w3.org/TR/did-core/#capability-invocation)
- [`capabilityDelegation`](https://www.w3.org/TR/did-core/#capability-delegation)


#### DID Document Metadata

The implementation for [DID document metadata](https://www.w3.org/TR/did-core/#did-document-metadata) MUST report the following properties for a DID document:

- `created`: a [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) string of the creation date that is the UTC date associated with the block height when the DID document was submitted the first time
- `updated`: a [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) string of the last update date that is the UTC date associated with the block height when the DID document was submitted the last time
- `deactivated`: a boolean field that indicates if the DID document is [deactivated](#Deactivate) 
- `versionId`: a hex-encoded BLAKE2b hash of the transaction that created or updated the DID

#### DID Resolution Metadata

The [DID Resolution Metadata](https://www.w3.org/TR/did-core/#did-resolution-metadata) is outside the scope of the gRPC interface and is not covered in this ADR.

#### DID URL Syntax

The [DID URL Syntax](https://www.w3.org/TR/did-core/#did-url-syntax) is outside the scope of the gRPC interface and is not covered in this ADR.

#### DID Query Parameters

The [DID Query parameters](https://www.w3.org/TR/did-core/#did-parameters) URL is outside the scope of the gRPC interface and is not covered in this ADR.

<!-- 

The implementation MUST support the following query parameters:

- `versionId` - to retrieve a DID document with a specific version 
- `versionTime` - to retrieve the version of the DID document valid at a specific time, the parameter MUST be a valid [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime).   

The format for the queries is:
-->
## Consequences

The Cosmos ecosystem WILL HAVE a DID module that is compatible with the W3C standard and offers a high chance of compatibility with third-party components such as cloud and edge agents, resolvers, and so on.

### Backwards Compatibility

This is a new module so backward compatibility is not a concern.

### Positive

- The implementation of the ADR provides the foundation for interoperability with the DID standard and the SSI identity approach.
- Closely following the W3C standard gives the best chances of successful interoperability with third-party components.

### Negative

- The implementation rigidly follows the W3C specification which leaves little room for extensibility. This approach might become an issue for wider adoption.

### Neutral

N/A

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section contains a summary of issues to be solved in future iterations. The issues summarized here can reference comments from a pull request discussion.
Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

- The `did:key` method specifies a key format that is different from the one used in this ADR. This ADR needs to be amended or follow a different approach.
- The approach proposed is somewhat locked into the current implementation and will have to be revised in successive iterations. 
## Test Cases [optional]

N/A

## References

- [DID Core](https://www.w3.org/TR/did-core)
- [DID Specification Registries](https://w3c.github.io/did-spec-registries)





