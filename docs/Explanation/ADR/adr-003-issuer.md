# ADR 003: Cosmos Cash Issuance model

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

All major stablecoins, Tether and USDC are based on Ethereum and are composed from a number smart contract standards. Those of interest here include:

* [ERC-20](https://www.notion.so/allinbits/USDC-vs-USDT-vs-CASH-ISSUER-5e1e6530247c433292caabf4e96038bf#b035a266c0954b3abdde3d6577d74908) - transfer, transferFrom and approve
* [Mintable](https://www.notion.so/allinbits/USDC-vs-USDT-vs-CASH-ISSUER-5e1e6530247c433292caabf4e96038bf#22bd3b30f86b4b4fbfe7097c04707f68)
* [Burnable](https://www.notion.so/allinbits/USDC-vs-USDT-vs-CASH-ISSUER-5e1e6530247c433292caabf4e96038bf#0d815a8305824674ad7b61ce3393f790)
* [Pausable](https://www.notion.so/allinbits/USDC-vs-USDT-vs-CASH-ISSUER-5e1e6530247c433292caabf4e96038bf#22bd3b30f86b4b4fbfe7097c04707f68)
* [Ownable](https://www.notion.so/allinbits/USDC-vs-USDT-vs-CASH-ISSUER-5e1e6530247c433292caabf4e96038bf#b78214caa3c748968bdb4773cbae5d02)
* [RBAC](https://www.notion.so/allinbits/USDC-vs-USDT-vs-CASH-ISSUER-5e1e6530247c433292caabf4e96038bf#e5a99af185a24f25bb847dc63604c7cf) - manage contract/role access

In addition these implementations also have a range of bespoke functions:

* Blocklist, which implies adding/removing addresses from a block list and also rescuing or burning a blocklisted users' funds (USDT and USDC)
* Set minting limits for a minter (USDC)
* Set spending limits by address (USDC)

These contracts then have role based access to perform certain actions such as mint/burn tokens. Both tokens are deemed to be regulatory compliant in terms of functionality.

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

The functionality SHALL USE Verifiable Credentials to establish Role Based Access to functions. See [ADR-OO5 License Credential](https://github.com/allinbits/cosmos-cash/blob/main/docs/Explanation/ADR/adr-005-license-credential.md) for further details regarding issuance, revocation etc.


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
