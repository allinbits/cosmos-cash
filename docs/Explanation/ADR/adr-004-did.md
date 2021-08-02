# ADR 4: DID

## Changelog

- 2021-08-02: Initial draft

## Status

DRAFT


## Abstract

The ADR describes an implementation of the [W3C DID specification](https://www.w3.org/TR/did-core) 


## Context

The W3C DID specification are a building block to realize a SSI platform for tendermint based chains. 

In the context of cosmos-cash project the SSI approach is a viable path to provide a technical solution for issues around identity, privacy and security that is compliant with regulations such as GDPR.

## Decision

> This section describes our response to these forces. It is stated in full sentences, with active voice. "We will ..."
> {decision body}

The DID W3C specification are designed using the "open world assumption" approach for data modelling. 
For the cosmos-cash implementation we will be following the did-core specification: custom properties will require a fork and a different implementation of the cosmos-cash did module.

### Method schema

The schema for the cosmos-cash DID implementation is the following

`did`:`cash`:`{CHAIN_NAME}`:`{ACCOUNT_ADDRESS (without chain prefix)}`

i.e. `did:cash:cosmos:ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj`

where the `CHAIN_NAME` is used by a resolver as routing and the `ACCOUNT_ADDRESS` as the unique identifier for the DID

> TODO: this is a placeholder, a decision must be taken in this regards, the method name could be more generically indicate the module like i.e. `tendermint`


### DID Document

The DID document will support the following attributes:

- `context` - for json-ld specification
- `id` - indicating the DID of the document
- `controller` - a list of DID controllers for the DID document
- `verificationMethods` - a list of verification method objects
- `services` - a list of services

We will support the following verification relationships:

- `authentication`
- `assertionMethod`
- `keyExchange`
- `invocationMethod`
- `delegationMethod`


##### [DID Operations](https://www.w3.org/TR/did-core/#method-operations)


###### Create

The publication on chain of a DID document is allowed only when the submitted DID document contains at least one verification method that is listed in the verification relationship `authentication` and whose `blockchainAccountID` matches the transaction signer for the address part (ignoring the chain prefix).

###### Resolve/Verify

The integrity of the DID documents stored on the ledger are guaranteed by the underlying blockchain protocol. 

> TODO: anything to add about resolution? see [169#issuecomment-891189527](https://github.com/allinbits/cosmos-cash/issues/169#issuecomment-891189527)

###### Update

Amends to a DID document are allowed only to:

- controllers of one of the verification method listed in the `authentication` verification relationship
- controllers of one of the verification method listed in the `authentication` verification relationship of the DID document identified in the the controller of the current DID document 

```
example: 

giving did did:cash:abc:1111 described by document

{
    "@context": ...
    "id": "did:cash:abc:1111",           <--- DID 111 subject
    "controller": [
        "did:cash:abc:2222"              <--- Controller for the document describing did:cash:abc:1111 
    ]
    "verificationMethods" : [
        "id": "did:cash:abc:1111#key-1",
        "type": "EcdsaSecp256k1VerificationKey2019",
        "controller": "did:cash:abc:1111",   <--- Controller for the verification method identified by did:cash:abc:1111#key-1
        "blockchainAccountID": "cosmos1g4fdrtxjjvzxvazkmj7zgz634a4hn93nry0p9u"
    ],
    "authentication": [
        "id": "did:cash:abc:1111#key-1"      <--- Authentication relationships for key did:cash:abc:1111#key-1
    ]
}

and did:cash:abc:2222

{
    "@context": ...
    "id": "did:cash:abc:2222",         <--- DID subject 
    "verificationMethods" : [
        "id": "did:cash:abc:2222#key-1",
        "type": "EcdsaSecp256k1VerificationKey2019",
        "controller": "did:cash:abc:2222", <--- Controller for the verification method identified by did:cash:abc:1111#key-1
        "blockchainAccountID": "cosmos1ts9ejqg7k4ht2sm53hycty875362yqxqmt9grj",
    ],
    "authentication": [
        "id": "did:cash:abc:2222#key-1"    <--- Authentication relationships for key did:cash:abc:2222#key-1
    ] 
}

according to the rules the DID did:cash:abc:2222 is authorized to amend the DID document for did:cash:abc:1111  
```


The transaction signer address MUST match the verification method `blockchainAccountID` for the address part (ignoring the chain prefix).

> XXX: the consequences of that is that a controller has the same powers of the creator of the did document

###### Deactivate

The deactivation of a DID is obtained by removing all the values for the properties:

- `controller`
- `verificationMethods` 

the deactivation is allowed following the same rules af for the [update operation](#update) 

> 

#### Metadata


##### [DID Document Metadata](https://www.w3.org/TR/did-core/#did-document-metadata)

The implementation for metadata MUST report the following properties for a DID document

- `created`: a [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) string of the creation date, that is the utc date associated to the block height when the DID document was submitted the first time
- `deactivated`: whenever the DID document is [deactivated](#Deactivate) 
- `versionId`: for the version id we use the hash of the transaction that created/updated the DID


##### [DID Resolution Metadata](https://www.w3.org/TR/did-core/#did-resolution-metadata)

This part of the W3C DID specification is not covered by this ADR


#### [DID URL Syntax](https://www.w3.org/TR/did-core/#did-url-syntax)

No `paths` or `fragments` are defined for this DID method. 
##### [Query parameters](https://www.w3.org/TR/did-core/#did-parameters)

The implementation MUST support the following query parameters:

- `versionId` - to retrieve a DID document with a specific version 
- `versionTime` - to retrieve the version of the DID document valid at a specific time, the parameter MUST be a valid [datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime).   


The format for the queries is:


## Consequences

> This section describes the resulting context after applying the decision. List all consequences here, taking care not to list only the "positive" consequences. A particular decision may have positive, negative, and neutral consequences, but all of the consesquences affect the team and project in the future.



### Backwards Compatibility

N/A

### Positive

- The implementation of the ADR provides the foundation for interoperability with the DID standard and more in general with SSI identity approach
- Closely following the W3C standard gives the best chances of successful interoperability with 3rd party components  

### Negative

- The implementation follows rigidly the W3C specification leaving little room for extensibility, that might became an issue for a wider adoption 

### Neutral

N/A

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section contains a summary of issues to be solved in future iterations. The issues summarized here can reference comments from a pull request discussion.
Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

- 

## Test Cases [optional]

N/A

## References

- [DID Core](https://www.w3.org/TR/did-core)
- [DID Specification Registries](https://w3c.github.io/did-spec-registries)
