# ADR 009: Fungibility

## Changelog

- 2021-12-03: Initial draft.
- 2021-12-09: Updated draft.

## Status

PROPOSED

## Abstract

Cosmos Cash is a platform enabling the issuance of regulatory-compliant E-Money tokens.

This document defines token fungibility within the Cosmos Cash context and outlines possible approaches and describes the approach chosen for the Cosmos Cash project. 

Two approaches to E-Money token Fungibility are outlined. It is decided that Cosmos Cash will support the second approach.

## Context

Recall that Cosmos Cash is a research project into a Cosmos SDK based, regulatory compliant finance protocol that can host tokens backed by a given fiat currency (or a digital representation thereof).

Under MiCA regulations a Cosmos Cash token would qualify as an E-Money token since it holds the following properties:
* It is a crypto asset.
* It is an electronic surrogate for coins and banknotes and is used for making payments.
* It maintains a stable value by referring to the value of one fiat currency (1:1 pegging).

I.e. E-Money tokens are collateralized stablecoins pegged to a given single fiat currency - typically the Euro - and with multiple E-Money issuers a technical challenge arises in terms of token fungibility.

###Fungibility:

Fungibility herein is defined as an absolute, where two E-Money tokens of the same numeraire are fungible if they are treated as absolutely interchangeable and, as a result, they have the same value.

###Fungibility Approach 1:

Under this approach, Cosmos Cash will have multiple E-Money tokens, per numeraire. (The word numeraire translates as "money," "coinage," or "face value" and just means per currency used to measure amounts; i.e. with less specificity, 'per numeraire' means 'per currency'.)

Following Approach 1 leads us to a situation complementing existing blockchain ecosystems where typically there are different stablecoins pegged to the same fiat currency where nominally they have the same value but their exchange rates vary slightly (for example USDC, USDT, BUSD, UST, ...).

In this 1st approach, issuers do not support Fungibility natively, for example by agreeing to exchange and redeem the EMTs of other issuers, and therefore their EMTs are not fungible according to our definition.

From a user's perspective this approach gives a poor user experience, akin to a high street where each shop accepts different EUR notes and where shoppers' wallets each contain a subset of the different available EUR notes.

###Fungibility Approach 2:

Under this approach, Cosmos Cash will have multiple e-money providers that issue a single E-Money token, per numeraire. This is different but similar to the approach implemented by the [Circle](https://www.circle.com/en/) USDC project, where a consortium of organizations issue USDC tokens. 

Issuers agree to support Fungible (at par) exchange and redemption of E-Money tokens from other issuers.

Because they can always readily exchange E-Money tokens from different issuers at par, users will treat E-Money tokens from different Issuers as homogeneous.

E-Money tokens are Fungible under Approach 2.

It is expected that Fungibility Approach 2 provides the following benefits:
* An improved user experience.
* A network effect, increasing E-Money token utility (i.e. if more users accept/use a given E-Money token, the value from its use and also its liquidity are increased).

Both benefits ultimately lead to greater E-Money token use, E-Money token issuance and chain economic value.

Under Fungibility Approach 2, if and when an issuer holds more of another issuer's EMTs than it holds of theirs, it does face Counterparty Risk. It can mitigate this by redeeming competitor E-Money token ASAP.

## Decision

In consideration of the benefits and complexity of the approaches, Cosmos Cash will support the second approach to Fungibility.

Recall that Cosmos Cash's ultimate goal is to bring true liquidity to the Cosmos ecosystem through a reliable, regulatory-compliant protocol that guarantees the same guarantees as the traditional banking sector.

The second approach is conceptually more likely to provide a value proposition for e-money issuers while at the same time ensuring an optimal user experience for token holders. Nevertheless it does present significant technical implementation challenges.

* The chain will support one E-Money token - per numeraire.

* The chain's E-Money token will have a given name, such as EUR-E.

* Issuers support Fungibility on chain:
   - A transaction type is available as follows: for User U, licensed by Issuer I1, U can 'send' to I1 the E-Money token of Issuer I2, 'receiving' in returning the E-Money token of I1.
   - The transaction will fail if: I1 is up to its issuance limit or the Issuer license of I2 has been revoked.

## Consequences

See Further Discussions.

### Backwards Compatibility

This ADR is expected to limit Backward Compatibility, but this compatibility remains to be confirmed.

### Positive

N/A.

### Negative

N/A.

### Neutral

N/A.

## Further Discussions

###Firstly:

At a later date Cosmos Cash may support 'auto-clearing' where 2 issuers both hold n EMTs of each other's tokens.

###Secondly:

At a later date Cosmos Cash may also support Fungibility Approach 1 - note that it is possible to support both outlined approaches to Fungibility at the same time.

It is anticipated that some issuers may wish to issue EMTs without regard for other issuers.

It is also anticipated that these issuers may then at a later date decide to support the redemption and exchange of other issuer's tokens,
so as to encourage the issuance of their own EMT and, to encourage overall EMT issuance.

## References
 
N/A.
