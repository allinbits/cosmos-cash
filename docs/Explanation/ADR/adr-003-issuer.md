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
    - [Setting up an Issuer](#setting-up-an-issuer)
    - [Revocation Lists](#revocation-lists)
    - [Out-of-Scope](#out-of-scope)
  - [Consequences](#consequences)
    - [Backward Compatibility](#backward-compatibility)
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

Regulators and issuers are issued licenses through a regulator or other government body. The licenses then permit the holder to perform certain activities such as holding client money, and so on. In the UK, the [FCA register](https://register.fca.org.uk/s/) acts as the public record of licensed entities and the services the entities can provide. For example, the record for UBS AG on FCA is [UBS AG
Reference number: 186958](https://register.fca.org.uk/s/firm?id=001b000000MfHZiAAN). 

---

## Decision

WE SHALL implement an issuer module with functionality that is similar to Tether and USDC. This module will have the following functionality:

* Create an issuer and a payment token
* Mint and burn tokens
* Pause and unpause token circulation
* Set an upper bound on token circulation
* Set a fee rate
* Rescue funds from frozen accounts

### Actors

There SHALL BE the following actors with permissions actions:

* **Regulator** - this actor SHALL HAVE role-based permissions to: 
    * Create and revoke an issuer.
    * Set and update an upper bound on the total amount of tokens in circulation.
* **Issuer** - this actor SHALL HAVE role-based permissions to:
    * Mint and burn tokens
    * Pause and unpause tokens in circulation
    * Redeem tokens from block-listed address 

### DID

The Regulator and the Issuer Decentralized Identifier Document (DID) SHALL HAVE multiple DID controllers. For an Issuer, these controllers represent the different functions within a financial institution - Operations, Compliance, and so on. This DID SHALL support multiple signing such that `Pause and unpause` requires two parties to make the change. This requirement would mimic a case where Operations would request the pause and Compliance or a Senior Manager would approve the action.

Likewise, the Regulator DID will also have multiple controllers with multi-signatories that are required for creating and removing an issuer and setting circulation bounds.

### Verifiable Credentials


The functionality SHALL USE Verifiable Credentials to establish role-based access to functions. See [ADR-OO6 License Credential](https://github.com/allinbits/cosmos-cash/blob/main/docs/Explanation/ADR/adr-006-license-credential.md) for details regarding issuance, revocation, and so on.


### Setting up an Issuer

* A Regulator actor WILL BE defined. 
* This Regulator WILL ISSUE signed license verifiable credentials to an Issuer.
* The Regulator address and DID document SHALL BE defined in Genesis.

> Does the Regulator also do this from Verifiable Credential?

### Revocation Lists

Given that role-based access uses Verifiable Credentials then the address must be added to a revocation list to override the credential to revoke and unrevoke access.

* Revocation lists will be maintained as accumulators on the identity chain.
* There WILL BE revocation lists for issuers and users.
* A Regulator actor cannot be revoked.


### Out-of-Scope

The issuer module for USDT and USDC both handle blocklisting with a defined Blacklister role. 

As proposed by this ADR, the Issuer module WILL NOT have the functionality to maintain access. These permissions will be handled through Verifiable Credentials which will be off-chain. Revocation will be handled as described in the [Revocation Lists](#revocation-lists) section.

User transaction limits will also be handled in [Verifiable Credentials](#verifiable-credentials). A user will present a credential each time they perform a transaction. The credential will prove if the transaction is within the limit for that user.


---

## Consequences

### Backward Compatibility

This is a new module so backward compatibility is not a concern.

### Positive

* Allows Cosmos Cash to support multiple issuers and multiple payment tokens on the same network 
* The issuer module will be regulatory compliant because it offers:
    * Restrictions to Issuer creation so that Issuer creation is possible only with correct regulatory sign-off
    * Functionality to pause and unpause tokens
    * Controls to manage token circulation
* The issuer module will use Decentralized Identity and Verifiable Credentials authentication and authorization of actions.

### Negative

* This module is not compatible with USDC or USDT issuers module because it separates the permission model from the actual functionality.

---

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section contains a summary of issues to be solved in future iterations. The issues summarized here can reference comments from a pull request discussion.
Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.


## References

- [OpenVASP Core Data types](https://github.com/OpenVASP/ovips/blob/master/ovip-0013.md)
- [OpenVASP Credential](https://github.com/OpenVASP/ovips/blob/master/ovip-0015.md)
