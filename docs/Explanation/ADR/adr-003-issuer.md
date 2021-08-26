# ADR 003: Cosmos Cash Issuance model

## Table of Contents

- [ADR 003: Cosmos Cash Issuance model](#adr-003-cosmos-cash-issuance-model)
  - [Table of Contents](#table-of-contents)
  - [Changelog](#changelog)
  - [Status](#status)
  - [Abstract](#abstract)
  - [Context](#context)
    - [Ethereum based Stablecoin issuers](#ethereum-based-stablecoin-issuers)
    - [Licenses](#licenses)
  - [Decision](#decision)
    - [Actors](#actors)
    - [DID](#did)
    - [Verifiable Credentials](#verifiable-credentials)
    - [Setting up an issuer](#setting-up-an-issuer)
    - [Revocation Lists](#revocation-lists)
    - [Out-of-Scope](#out-of-scope)
  - [Consequences](#consequences)
    - [Backwards Compatibility](#backwards-compatibility)
    - [Positive](#positive)
    - [Negative](#negative)
  - [Further Discussions](#further-discussions)
  - [References](#references)

## Changelog

* 29-Jul-2021: Initial Draft

## Status

DRAFT - Not Implemented

---

## Abstract

In order to issue a payment token a regulatory compliant issuer module is required. This must be able to 

* Initialize an issuer and a payment token
* Mint/Burn tokens
* Freeze/Unfreeze accounts
* Rescue funds from frozen accounts
* Pause/Unpause token circulation

This functionality will use Role Based Access based on Verifiable Credentials with blocklists maintained through the use of revocation lists.

---

## Context

### Ethereum based Stablecoin issuers

All major stablecoins, USDT and USDC are based on Ethereum and are composed from a [number smart contract standards with Role Based Access Control](../topics/Compare USDC vs USDT vs CASH-ISSUER.md). Both tokens are deemed to be regulatory compliant in terms of functionality.

[Cosmos Cash Proof of Concept issuer](https://github.com/allinbits/cosmos-cash-poc/tree/master/x/issuer) has some similar functionality:

* `createIssuer(token, fee, owner)`
* `mintToken(amount, owner)`
* `burnToken(amount, owner)`

### Licenses

Regulators and issuers are issued licenses through a regulator or other government body. The licenses then permit the holder to perform certain activities such as hold client money etc. In the UK the [FCA register](https://register.fca.org.uk/s/) acts as public record of licensed entities and the services they can provide. For example, this the record for [UBS AG](https://register.fca.org.uk/s/firm?id=001b000000MfHZiAAN)

---

## Decision

WE SHALL implement an Issuer module with functionality similar to those shown by Tether and USDC. This module will have the following functionality:

* Create an issuer and a payment token
* Mint/Burn tokens
* Pause/Unpause token circulation
* Set an upper bound on token circulation
* Set a fee rate
* Rescue funds from frozen accounts

### Actors

There SHALL BE the following actors with permissions actions:

* **Regulator** - this actor SHALL HAVE permissions to 
    * Create/Revoke an issuer.
    * Set/update an upper bound on total amount of tokens in circulation.
* **Issuer** - this actor SHALL HAVE permissions to
    * Mint/Burn tokens
    * Pause/Unpause tokens in circulation
    * Redeem tokens from block listed address 

### DID

The Regulator and the Issuer Decentralized Identifier Document (DID) SHALL HAVE multiple DID controllers. For an Issuer these represent the different functions within a financial insitutions - Operations, Compliance etc. This SHALL support multiple signing such that `Pause/Unpause` requires two parties to make the change. This would mimic a case where Operations would request the pause and Compliance or a Senior Manager would approve the action.

Likewise, the Regulator DID will also have multiple controllers with multi-signatories required for creating/removing an issuer and setting circulation bounds

### Verifiable Credentials

The functionality SHALL USE Verifiable Credentials to establish Role Based Access to functions. See [ADR-OO5 License Credential](https://github.com/allinbits/cosmos-cash/blob/main/docs/Explanation/ADR/adr-006-license-credential.md) for further details regarding issuance, revocation etc.


### Setting up an issuer

* A Regulator actor WILL BE defined. 
* This Regulator WILL ISSUE signed license verifiable credentials to an Issuer.
* The Regulator address and DID document SHALL BE defined in Genesis

> Does the Regulator also do this from Verifiable Credential?

### Revocation Lists

Given that Role Based Access uses Verifiable Credentials then in order to revoke and unrevoke access then address needs to be added to a revocation list in order to override the credential.

* Revocation lists wil be maintain as accumulators on the identity chain.
* There WILL BE revocation lists for issuers and users.
* A regulator actor can not be revoked


### Out-of-Scope

The issuer module for USDT and USDC both handle blocklisting with a defined Blacklister role. 

As propsoed by this ADR, the Issuer module WILL NOT have functionality to maintain access. These permissions will be handled through Verifiable Credentials which will be off-chain. Revocation will be handled as described above.

User transaction limits will also be handled in verifiable credentials. When a user performs a transaction they will present a credential. The credential will prove if the transaction is within the limit for that user.


---

## Consequences

### Backwards Compatibility

This is a new module so backward compatibility is not a concern.

### Positive

* Allows Cosmos Cash to support multiple issuers and multiple payment tokens on the same network 
* It will be regulatory compliant because it offers
    * Issuers can only be created with correct regulatory sign-off
    * Can pause and unpause tokens
    * Controls to manage token circulation
* It will be utilise Decentralized Identity and Verifiable Credentials authentication and authorization of actions.

### Negative

* This won't be compatible with USDC or USDT issuers module, but it separates permission model from the actual functionality.

---

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section contains a summary of issues to be solved in future iterations. The issues summarized here can reference comments from a pull request discussion.
Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.


## References

- [OpenVASP Core Data types](https://github.com/OpenVASP/ovips/blob/master/ovip-0013.md)
- [OpenVASP Credential](https://github.com/OpenVASP/ovips/blob/master/ovip-0015.md)
