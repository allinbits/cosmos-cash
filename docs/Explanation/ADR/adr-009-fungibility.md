# ADR 009: Fungibility

## Changelog

- 2021-12-03: Initial draft

## Status

PROPOSED

> For details on ADR workflow, see the [PROCESS](./PROCESS.md#adr-status) page.
> Use DRAFT if the ADR is in a draft stage (draft PR) or PROPOSED if it's in review.

## Abstract

It is anticipated that Elesto will have more than one Issuer issuing e-money tokens denominated in any given currency (numereraire).

Here we summarise the anticipated behaviour and driving incentives of Issuers and Users and, what this means for our choice of functionality.

## Context

For simplicity of description, assume throughout we are discussing USD-denominated tokens.

EMT = e-money token.

Issuer behaviour:

Issuers agree to support Fungible (at par) exchange and redemption of EMTs from other currently licensed issuers.

Issuer incentives for the given behaviour:

(a) Benefit: network effect leads to increased EMT use, with increased issuance and circulation.
    (This is driven by the resulting User behaviour and incentives - below. )
(b) Benefit: any given issuer would lose market share if they didn't agree to support fungility as stated.
    (Exchange and redemption take their competitors' EMTs out of circulation, put their EMT into circulation.)
(c) Cost: Issuers face Counterparty risk only while they hold competitors' EMTs - they can mitigate this by redeeming competitor tokens.

User behaviour:

(i) Users will treat EMTs from different Issuers are homogeneous.
(ii) Users will accept wallets and related use cases that equally 'handle' EMTs of different Issuers as homogenous.
     This will lead to enhanced 'ease of use' and, further adoption (see: Issuer incentives).

## Decision

We will update our functionality as follows:

(1) Our chain will support only one EMT (per currency), issued by 1 or more Issuers.

(2) Our chain's EMTs will all have the same given name? (Such as USD-E for USD.)

(3) Added Issuer functionality.
    We add a transaction type as follows:
        For User U, licensed by Issuer I1, U can 'send' to I1 the EMT an Issuer I2, 'receiving' in returning the EMT of I.
    The transaction will fail if:
        I1 is up to its issuance limit or
        The Issuer license of I2 has been revoked.

## Consequences

See Context.

### Backwards Compatibility

Expected to limit Backwards Compatibility but, to be confirmed.

### Positive

See Context.

### Negative

See Context.

### Neutral

See Context.

## Further Discussions

It is expected users will more readily adopt Fungible EMTs.

The decision given is to only support Fungible EMTs.

Alternative Decision options are:
(a) Only support non-Fungibile EMTs,
(b) Support Fungible and non-Fungible EMTs
    (For example Issuers I1 and I2 issue Fungible EMT T1 and, Issuer I3 issues non-Fungible EMT T2.)

## Test Cases [optional]

N/A

## References

https://www.notion.so/allinbits/FungibilityV2-1-e6b1deed13cc46649e9f915d4a23944b
