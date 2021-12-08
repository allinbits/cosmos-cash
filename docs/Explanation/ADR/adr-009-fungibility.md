# ADR 009: Fungibility

## Changelog

- 2021-12-03: Initial draft

## Status

PROPOSED

> For details on ADR workflow, see the [PROCESS](./PROCESS.md#adr-status) page.
> Use DRAFT if the ADR is in a draft stage (draft PR) or PROPOSED if it's in review.

## Abstract

Cosmos Cash is a platform enabling the issuance of MICA-compliant e-money tokens.

Two approaches to e-money token Fungibility are outlined.

Cosmos Cash supports the second approach.

## Context

###EMT:

EMT = e-money token.

For simplicity of description, assume throughout we are discussing EUR-denominated EMTs.

###Fungibility:

Fungibility is defined as an absolute, 2 EMTs are Fungible if they are absolutely interchangeable; as a result, they have the same value.

###Fungibilty Approach 1:

Under this approach, the chain will have multiple EMTs - per numeraire.

Suppose n issuers issue EMTs.

The issuers do not support Fungibility, for example by agreeing to exchange and redeem the tokens of other issuers.

EMTs are not Fungible.

It is believed that Fungibility Approach 1 gives a poor user experience, akin to a high street where each shop accepts different EUR notes and,
shoppers' wallets each contain a subset of the different available EUR notes.

###Fungibilty Approach 2:

Under this approach, the chain will support one EMT - per numeraire.

Suppose n issuers issue tokens.

The issuers agree to support Fungible (at par) exchange and redemption of EMTs from other issuers.

Because they can always readily exchange EMTs from different issuers at par, users will treat EMTs from different Issuers are homogeneous.

It is believed that Fungibility Approach 2 provides:
(i) An improved user experience
(ii) A network effect, increasing EMT utility (i.e. if more users accept/use a given EMT, the value from its use and, also its liquidity are increased)

Both (i) and (ii) should ultimately lead to greater EMT use, EMT issuance and, chain value.

## Decision

Cosmos Cash will support the second approach.

(1) The chain will support one EMT - per numeraire.

(2) The chain's EMT will have a chosen name, such as EUR-E.

(3) Issuers support Fungibility on chain via:
A transaction type is available as follows:
For User U, licensed by Issuer I1, U can 'send' to I1 the EMT an Issuer I2, 'receiving' in returning the EMT of I.
The transaction will fail if:
I1 is up to its issuance limit or
The Issuer license of I2 has been revoked.

## Consequences

See Further Discussions.

### Backwards Compatibility

N/A.

### Positive

N/A.

### Negative

N/A.

### Neutral

N/A.

## Further Discussions

###Firstly:

At a later date the chain may support 'auto-clearing' where 2 issuers both hold n EMTs of each other's tokens.

###Secondly:

At a later date the chain may also support Fungibilty Approach 1.

It is anticipated that some issuers may wish to issue EMTs without regard for other issuers.

It is also anticipated that these issuers may then at a later date decide to support the redemption and exchange of other issuer's tokens,
so as to encourage the issuance of their own EMT and, to encourage overall EMT issuance.

## Test Cases [optional]

N/A

## References

N/A
