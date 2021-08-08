# ADR 4: DID

## Changelog

- 2021-08-02: Initial draft

## Status

DRAFT
## Abstract

[Decentralized identifiers](https://www.w3.org/TR/did-core) (DIDs) are a new type of identifier that enables verifiable, decentralized digital identity. A DID refers to any subject (e.g., a person, organization, thing, data model, abstract entity, etc.) as determined by the controller of the DID.

This document specifies the DID method for a cosmos based implementation of the W3C recommendation, its properties, operations and an explanation of the process of resolving DIDs to the resources that they represent. 

## Context

The aim of the Cosmos Cash project is to provide a state of the art collateralised stable-coin implementation compliant with the EU regulations such as GDPR and MICA and international recommendations such as the FATF "Travel Rule" and local AML regulations.

A state of the art implementation includes

- A public financial infrastructure (public goods) 
- Auditing and identification of bad actors (AML regulations)
- A strictly privacy respecting approach (GDPR)

An approach to tackle the identity and privacy challenge that has been gaining momentum in the recent years is the SSI approach, that, coupled with DLT technology, has been capturing the attention of both the private and public sector. 

The Self-Sovereign Identity (SSI) relays on two building blocks: decentralized identifiers (DIDs) and verifiable credentials (VC). This ADR is about describing the DID implementation in a cosmos-sdk based blockchain.

The goal of this ADR is to define a foundation for the necessary components to realize the Cosmos Cash objectives and on the same time to keep the implementation of the DID fully compliant with the W3C specifications. **Successive iterations will address API ergonomics and standard compatibility issues.** 

## Decision


The DID W3C specification are designed using the "open world assumption" approach for data modelling. 
For the Cosmos Cash implementation we will be following the did-core specification: custom properties will require a fork and a different implementation of the cosmos-cash did module.

The following is an example of a DID Document:

```javascript
{
    "context": [
        "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cosmos:ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
    "verificationMethods" : [
        {
            "id": "did:cosmos:cosmos:ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj#key-1",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:cosmos:cosmos:ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "blockchainAccountID": "cosmos1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
        }
    ],
    "authentication": [
        "did:cosmos:cosmos:ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj#key-1"
    ],
    "services": [
        {
            "id":"agent:ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "type":"DIDCommMessaging",
            "serviceEndpoint":"https://agent.xyz/ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
        }
    ]
}
```

### DID Method name

The namestring that shall identify this DID method is: `cosmos`.

A DID that uses this method MUST begin with the following prefix: `did:cosmos`. Per this [DID specification](https://www.w3.org/TR/did-core), this string MUST be in lowercase. The remainder of the DID, after the prefix, is specified below.

#### Method Specific Identifier


The namespace specific identifier is defined by the following ABNF:

```
cosmos-did                = "did:cosmos:" cosmos-specific-id-string
cosmos-specific-id-string = cosmos-chain-name ":" unique-identifier
cosmos-chain-name         = 1*255id-char
unique-identifier         = 38*256id-char
id-char                   = ALPHA / DIGIT / (ALPHA "-") / (DIGIT "-")
```

For the `unique-identifier` it is RECOMMENDED to use a cosmos account address (without chain prefix and separator) or an UUID 

Examples using a cosmos address: 

- `did:cosmos:cash:ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj`  from address `cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj`
- `did:cosmos:cosmoshub:ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj`  from address `cosmos1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj`

Examples using an UUID:

- `did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb`
  

##### [DID Operations](https://www.w3.org/TR/did-core/#method-operations)

DID and associated DID documents are managed by a cosmos-sdk module using gRPC as communication protocol. In this section the CRUD operations for a cosmos DID are defined.
###### Create

To create and publish a DID document use the message 

```
MsgCreateDidDocument(id string, signerAccount string)
```

The parameter for the message are the DID to be created and the key that will be used as the initial verification method in the authentication relationship in the DID document. 

If the input did is not a valid did for the cosmos method or the DID already exists on chain, the message will return an error. 

Example message and resulting DID Document on a Cosmos Cash chain:
```javascript
/* gRPC message */
MsgCreateDidDocument(
    "806e557e-ecdb-4e80-ab0d-a82ad35c9ceb", 
    "cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
)

/* DID document */
{
    "context": [
        "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb",
    "verificationMethods" : [
        {
            "id": "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj#cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "blockchainAccountID": "cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
        }
    ],
    "authentication": [
        "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
    ]
}

/* DID metadata */
{
  "created": "2021-03-23T06:35:22Z",
  "updated": "2021-03-23T06:35:22Z", 
  "versionId": "ae11692325525e82337167fcfab34d45d1904ff786e2d4bf4be2d1c4878cd34c" /* hex(blake2b(tx)) */
}

```

The [`did:key`](https://w3c-ccg.github.io/did-method-key/) method is supported by the module and resolves automatically blockchain addresses. 

It is RECOMMENDED to use an id that is not equals to the blockchain account address for privacy concerns and to isolate the verification methods to the did subject (for example during key rotation)


> NOTE: a more fine grained way of creating a DID MAY be implemented with a `MsgCreateDidDocumentWitOptions` with the goal of saving in gas by executing a single transactions in a complex DID scenario.


###### Resolve/Verify

The integrity of the DID documents stored on the ledger are guaranteed by the underlying blockchain protocol. 

A DID MAY be resolved using the gRPC message:

```
QueryDidDocumentRequest(did string)
```

Example
```javascript
/* gRPC message */
QueryDidDocumentRequest("did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb")

/* DID document */
{
    "context": [
        "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb",
    "verificationMethods" : [
        {
            "id": "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj#cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "blockchainAccountID": "cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
        }
    ],
    "authentication": [
        "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
    ]
}

/* DID metadata */
{
  "created": "2021-03-23T06:35:22Z",
  "updated": "2021-03-23T06:35:22Z", 
  "versionId": "ae11692325525e82337167fcfab34d45d1904ff786e2d4bf4be2d1c4878cd34c" /* hex(blake2b(tx)) */
}
```

Note that the representation is not compatible with the JSON-LD standard due to some specificity of the Protobuf message format.

###### Update

There are two ways of updating a DID document:

- manage DID controllers
- manipulate verification methods and relationships 

In both cases the target DID MUST exists on chain and the `signerAccount` MUST exists as a verification method (property `blockchainAccountID`) in a verification relationship of type `authentication` or being listed as a DID controller.


**DID controllers** 

The DID controllers MAY be set using the gRPC message:

`MsgUpdateDidDocument(did string, controllers []string, signerAccount string)` 

The parameter `did` identifies the did document, the `controllers` are a list of DIDs that will replace the DID document controllers list and the `signerAccount` is the account address signing the transaction.

Controllers will be added using the `did:keys` method.


Example:

```javascript
/* gRPC message */
MsgUpdateDidDocument(
    "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb", 
    ["cosmos195rlq2hjnn2tmagskys4xtsnsey6gjljg8zxrn"],
    "cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
)

/* DID document */
{
    "context": [
        "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb",
    "controller": [
        "did:key:cosmos195rlq2hjnn2tmagskys4xtsnsey6gjljg8zxrn"  // <-- new controller added
    ],
    "verificationMethods" : [
        {
            "id": "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj#cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "blockchainAccountID": "cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
        }
    ],
    "authentication": [
        "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj#cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
    ]
}

/* DID metadata */
{
  "created": "2021-03-23T06:35:22Z",
  "updated": "2021-04-23T06:35:22Z",  // <--  update field modified
  "versionId": "96b3504be7e37a6aa55faff3cd41266bf4db3b0654263e1d9b779d3b30174dd1", /* hex(blake2b(tx)) */  //<-- new hash computed
  "deactivated": false
}

```


**Verification methods and relationships**

A new verification method MAY be added using the gRPC message

```
MsgAddVerification(did string, accountId string, relationships []string, signerAccount string)
```

The parameter `did` identifies the did document, the `accountId` is the account to be added to the verification method, the `relationships` are the list of relationships that the `accountId` shall be registered into.  The `signerAccount` is the account address signing the transaction.

The list of relationships MUST contains only valid [relationships names](#DID_document)

Example:

```javascript
/* gRPC message */
MsgAddVerification(
    "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb", 
    "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
    ["authentication", "keyAgreement"],
    "cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj" // <-- the signer has authorization relationship
)

/* DID document */
{
    "context": [
        "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb",
    "controller": [
        "did:key:cosmos195rlq2hjnn2tmagskys4xtsnsey6gjljg8zxrn" 
    ],
    "verificationMethods" : [
        {
            "id": "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj#cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "blockchainAccountID": "cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
        },
        {   // <--  a new verification method is added 
            "id": "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
            "blockchainAccountID": "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"
        }
    ],
    "authentication": [
        "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj#cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
        "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"  //<-- new verification method added to authentication relationship
    ],
    "keyAgreement": {
         "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2" //<-- new verification method added to keyAgreement relationship
    }
}

/* DID metadata */
{
  "created": "2021-03-23T06:35:22Z",
  "updated": "2021-05-23T06:35:22Z",  // <--  update field modified
  "versionId": "262495b1159c1cd7faf0da56deb6521bb2980d04435818906e417aae47604027", /* hex(blake2b(tx)) */ // <-- new hash computed
  "deactivated": false
}

```

The relationships of a verification method MAY be set using the gRPC message:

```
MsgSetVerificationRelationships(did string, accountId string, relationships []string, signerAccount string)
```

The list of relationships MUST contains only valid [relationships names](#DID_document)


Example:


```javascript
/* gRPC message */
MsgAddVerification(
    "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb", 
    "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
    ["authentication"],
    "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"
)

/* DID document */
{
    "context": [
        "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb",
    "controller": [
        "did:key:cosmos195rlq2hjnn2tmagskys4xtsnsey6gjljg8zxrn" 
    ],
    "verificationMethods" : [
        {
            "id": "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj#cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
            "blockchainAccountID": "cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj"
        },
        { 
            "id": "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
            "blockchainAccountID": "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"
        }
    ],
    "authentication": [
        "did:key:cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj#cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
        "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"  
    ]
    // <-- keyAgreement has been removed 
}

/* DID metadata */
{
  "created": "2021-03-23T06:35:22Z",
  "updated": "2021-06-23T06:35:22Z",  // <--  update field modified
  "versionId": "9066f968940392f3b36a580737f8e1c0bb4bc5ea6757f4d981cb7252c58710e5", /* hex(blake2b(tx)) */ //<-- new hash computed
  "deactivated": false
}
```

A verification method MAY be removed using the same gRPC message leaving the `relationships` fields empty. 

The `signerAccount` MUST exists as a verification method (property `blockchainAccountID`) in a verification relationship of type `authentication` or being listed as a DID controller.

Example:


```javascript
/* gRPC message */
MsgAddVerification(
    "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb", 
    "cash1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
    [],
    "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"
)

/* DID document */
{
    "context": [
        "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb",
    "controller": [
        "did:key:cosmos195rlq2hjnn2tmagskys4xtsnsey6gjljg8zxrn" 
    ],
    "verificationMethods" : [
        // <-- original verification method removed ==> keys have been rotated
        { 
            "id": "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
            "blockchainAccountID": "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"
        }
    ],
    "authentication": [
        "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"  
    ]
}

/* DID metadata */
{
  "created": "2021-03-23T06:35:22Z",
  "updated": "2021-06-23T06:35:22Z",  // <--  update field modified
  "versionId": "9066f968940392f3b36a580737f8e1c0bb4bc5ea6757f4d981cb7252c58710e5", /* hex(blake2b(tx)) */ // <-- new hash computed
  "deactivated": false
}
```

**Services**

A service MUST be an entity with the following properties:
  
- `id`: a valid RFC3986 URI string. 
- `type`: a non empty string.
- `serviceEndpoint`: a valid RFC3986 URI string.  

A service MAY be added using the gRPC method:

```
MsgAddService(did string, service_data Service, signerAccount string)
```

The `id` of a service MUST be unique within the DID document

Example:


```javascript
/* gRPC message */
MsgAddService(
    "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb", 
    {
        "id": "did:example:123#edv",
        "type": "EncryptedDataVault",
        "serviceEndpoint": "https://edv.example.com/"
    },
    "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"
)

/* DID document */
{
    "context": [
        "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb",
    "controller": [
        "did:key:cosmos195rlq2hjnn2tmagskys4xtsnsey6gjljg8zxrn" 
    ],
    "verificationMethods" : [
        // <-- original verification method removed ==> keys have been rotated
        { 
            "id": "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
            "blockchainAccountID": "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"
        }
    ],
    "authentication": [
        "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"  
    ],
    "services": [  // <-- the service is added to the DID document
        {
            "id": "did:example:123#edv",
            "type": "EncryptedDataVault",
            "serviceEndpoint": "https://edv.example.com/"
        }
    ]
}

/* DID metadata */
{
  "created": "2021-03-23T06:35:22Z",
  "updated": "2021-07-23T06:35:22Z",  // <--  update field modified
  "versionId": "d41260e87b4124ece80641207e44ea339ff06865fd5ce204e943e608d4b22268", /* hex(blake2b(tx)) */ // <-- new hash computed
  "deactivated": false
}
```

A service MAY be deleted using the gRPC message:

```
MsgDeleteService(did string, service_id string, signerAccount string)
```

Example:

```javascript
/* gRPC message */
MsgDeleteService(
    "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb", 
    "did:example:123#edv",
    "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"
)

/* DID document */
{
    "context": [
        "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb",
    "controller": [
        "did:key:cosmos195rlq2hjnn2tmagskys4xtsnsey6gjljg8zxrn" 
    ],
    "verificationMethods" : [
        // <-- original verification method removed ==> keys have been rotated
        { 
            "id": "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
            "type": "EcdsaSecp256k1RecoveryMethod2020",
            "controller": "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2",
            "blockchainAccountID": "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"
        }
    ],
    "authentication": [
        "did:key:cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2#cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"  
    ],
    "services": []  // <-- the service list is now empty
}

/* DID metadata */
{
  "created": "2021-03-23T06:35:22Z",
  "updated": "2021-08-23T06:35:22Z",  // <--  update field modified
  "versionId": "35ba49b94b20ac0a1805aa5283035ea71b003f56e3aeda6a4e9027779fe4aef5", /* hex(blake2b(tx)) */ // <-- new hash computed
  "deactivated": false
}
```

###### Deactivate

A DID can be deactivated using the gRPC message:

```
MsgDeactivateDid(did string, signerAccount string)
```

The DID identified by the parameter `did` MUST exists on chain and the `signerAccount` MUST exists as a verification method (property `blockchainAccountID`) in a verification relationship of type `authentication` or being listed as a DID controller.

This operation MUST remove all the verification methods and controllers and set the metadata property `deactivated` to true. This operation is not reversible.

Example:

```javascript
/* gRPC message */
MsgAddDeactivateDid(
    "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb", 
    "cash1n90e4s33ljxn00lhucvnmnfjng773efma38dc2"
)

/* DID document */
{
    "context": [
        "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:cash:806e557e-ecdb-4e80-ab0d-a82ad35c9ceb",
    "controller": [], // <-- controllers are removed 
    "verificationMethods" : [] // <-- verification methods are removed
}

/* DID metadata */
{
  "created": "2021-03-23T06:35:22Z",
  "updated": "2021-09-23T06:35:22Z",  // <--  update field modified
  "versionId": "e5ca728b93d19daa54180e20eec76d0d5614656c2ee3509df5a21a3abc1249ea", /* hex(blake2b(tx)) */ // <-- new hash computed
  "deactivated": true  // <-- field deactivated set to true
}
```

### Method specific properties 

#### [Verification material](https://www.w3.org/TR/did-core/#verification-material)

The verification material type SHOULD be `EcdsaSecp256k1RecoveryMethod2020`

The content of the verification material SHOULD be `blockchainAccountID`, but for interoperability reason the verification material should support also `publicKeyHex`. 

Support for other verification materials MAY be introduced. 

### [Verification relationships]((https://www.w3.org/TR/did-core/#verification-relationships))

The DID document MUST support the following verification relationships:

- [`authentication`](https://www.w3.org/TR/did-core/#authentication) - authorizes amends to the DID document
- [`assertionMethod`](https://www.w3.org/TR/did-core/#assertion)
- [`keyAgreement`](https://www.w3.org/TR/did-core/#key-agreement)
- [`capabilityInvocation`](https://www.w3.org/TR/did-core/#capability-invocation)
- [`capabilityDelegation`](https://www.w3.org/TR/did-core/#capability-delegation)


##### [DID Document Metadata](https://www.w3.org/TR/did-core/#did-document-metadata)

The implementation for metadata MUST report the following properties for a DID document

- `created`: a [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) string of the creation date, that is the utc date associated to the block height when the DID document was submitted the first time
- `updated`: a [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) string of the last update date, that is the utc date associated to the block height when the DID document was submitted the last time
- `deactivated`: a boolean field that indicates whenever the DID document is [deactivated](#Deactivate) 
- `versionId`: for the version id we use the hex encoded blake2b hash of the transaction that created/updated the DID

##### [DID Resolution Metadata](https://www.w3.org/TR/did-core/#did-resolution-metadata)

The resolution metadata are outside the scope of the gRPC interface and are not covered in this ADR

#### [DID URL Syntax](https://www.w3.org/TR/did-core/#did-url-syntax)

The DID URL is outside the scope of the gRPC interface and are not covered in this ADR
##### [Query parameters](https://www.w3.org/TR/did-core/#did-parameters)

The query parameters URL is outside the scope of the gRPC interface and are not covered in this ADR

<!-- 

The implementation MUST support the following query parameters:

- `versionId` - to retrieve a DID document with a specific version 
- `versionTime` - to retrieve the version of the DID document valid at a specific time, the parameter MUST be a valid [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime).   

The format for the queries is:
-->
## Consequences

The cosmos ecosystem will have at its disposition a DID module compatible with the W3C standard and with a high chance of compatibility with 3rd party components such as cloud and edge agents, resolvers and so on.

### Backwards Compatibility

This is a new module so backward compatibility is not a concern.

### Positive

- The implementation of the ADR provides the foundation for interoperability with the DID standard and more in general with SSI identity approach
- Closely following the W3C standard gives the best chances of successful interoperability with 3rd party components.

### Negative

- The implementation follows rigidly the W3C specification leaving little room for extensibility, that might became an issue for a wider adoption 

### Neutral

N/A

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section contains a summary of issues to be solved in future iterations. The issues summarized here can reference comments from a pull request discussion.
Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

- The `did:key` method specifies a key format that is not the one used in this ADR. The ADR needs to be amended or to follow a different approach
- The approach proposed is somewhat locked in to the current implementation, it will have to be revised in successive iterations. 
## Test Cases [optional]

N/A

## References

- [DID Core](https://www.w3.org/TR/did-core)
- [DID Specification Registries](https://w3c.github.io/did-spec-registries)





