# ADR 009: Fungibility

## Changelog

- 2021-12-03: Initial draft.
- 2021-12-05: Updated draft.

## Status

PROPOSED

## Abstract

Cosmos Cash is a platform enabling the issuance of MICA-compliant e-money tokens.

Two approaches to e-money token Fungibility are outlined.

Cosmos Cash supports the second approach.

## Context

###EMT:

EMT = e-money token.

###Fungibility:

Fungibility is defined as an absolute, 2 EMTs of the same numeraire are Fungible if they are absolutely interchangeable; as a result, they have the same value.

###Fungibilty Approach 1:

Under this approach, the chain will have multiple EMTs - per numeraire.

Suppose n issuers issue EMTs.

The issuers do not support Fungibility, for example by agreeing to exchange and redeem the EMTs of other issuers.

EMTs are not Fungible under Approach 1.

It is believed that Fungibility Approach 1 gives a poor user experience, akin to a high street where each shop accepts different EUR notes and where
shoppers' wallets each contain a subset of the different available EUR notes.

###Fungibilty Approach 2:

Under this approach, the chain will support one EMT - per numeraire.

Suppose n issuers issue EMTs.

The issuers agree to support Fungible (at par) exchange and redemption of EMTs from other issuers.

Because they can always readily exchange EMTs from different issuers at par, users will treat EMTs from different Issuers as homogeneous.

EMTs are Fungible under Approach 2.

It is believed that Fungibility Approach 2 provides 2 benefits:
* An improved user experience.
* A network effect, increasing EMT utility (i.e. if more users accept/use a given EMT, the value from its use and, also its liquidity are increased).

Both benefits ultimately lead to greater EMT use, EMT issuance and, chain value.

Under Fungibility Approach 2, if and when an issuer holds more of another issuer's EMTs than it holds of theirs, it does face Counterparty Risk. It can mitigate this by redeeming competitor EMTs ASAP.

## Decision

Cosmos Cash will support the second approach.

* The chain will support one EMT - per numeraire.

* The chain's EMT will have a chosen name, such as EUR-E.

* Issuers support Fungibility on chain:
   - A transaction type is available as follows: for User U, licensed by Issuer I1, U can 'send' to I1 the EMT an Issuer I2, 'receiving' in returning the EMT of I1.
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

At a later date Cosmos Cash may also support Fungibilty Approach 1.

It is anticipated that some issuers may wish to issue EMTs without regard for other issuers.

It is also anticipated that these issuers may then at a later date decide to support the redemption and exchange of other issuer's tokens,
so as to encourage the issuance of their own EMT and, to encourage overall EMT issuance.

## References

N/A
