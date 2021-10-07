# ADR 007: Root of trust

## Changelog

- 2021-10-05: Initial draft

## Status

PROPOSED

## Abstract

Blockchain technology provides a distributed and tamper-proof data layer where all the agents of the protocol interact
using asymmetric encryption where it is guaranteed that a piece of information has been signed by the controller of a
private key, but it doesn't provide any other information beside the consistency of the data.

This document illustrates a technique to define a root of trust for a blockchain system, that is to define a set of
address that are trusted on a protocol level. This document describes also how the root of trust is used as a based to
implement a public [verifiable credential](https://verifiablecredential.io/) scheme within the Cosmos Cash project.

## Context

In the context of Cosmos Cash there are three roles on chain:

- Regulator
- EMT/ART (crypto asset issuer)
- Account holder

The roles form a chain of trust illustrated in the following diagram:

![chain of trust](../../assets/diagrams/out/chain_of_trust.svg)

The chain of trust reliability is determined by the root of trust.

## Decision

To provide a reliable determination of the root of trust the chain initialization follows the procedure described in the
following diagram:

![root of trust](../../assets/diagrams/out/root_of_trust.svg)

The details of the off-chain ceremony where the root of trust addresses are collected are outside the scope of this ADR.

> Note: an improvement to this process is the introduction of root of trust governance model.

Once the root of trust is defined, we define a set of verification credentials, compliant wit the W3C verifiable
credential specification, that implements the chain of trust.

The credentials and their issuance constraints are represented in the following diagram:

![credentials](../../assets/diagrams/out/credentials.svg)

### Issuing credentials

Since the verifiable credential mechanism is independent of the blockchain transaction verification process, a 
separate verification process MUST be provided for issuing credentials. Such process is described in the following 
diagram:

![credential validation](../../assets/diagrams/out/credentials_validate_proof.svg)

The consequence of the verification process is that in line of principle the actor issuing credentials doesn't have 
to be the same actor that sign, pays and broadcasts the transaction containing the verifiable credentials.

### Deleting (Revoking) credentials

The credential revocation verification process is described in the following diagram:

![credential deletion](../../assets/diagrams/out/credentials_delete.svg)

## Consequences

### Backwards Compatibility

This is a new module so backward compatibility is not a concern.

### Positive

- The root of trust and the chain of trust models are compatible with the SSI verifiable credential process, and 
  therefore compatible with 3rd party applicaitons

### Negative

- The root of trust model is somewhat centralised

### Neutral

N/A

## Further Discussions

N/A

## Test Cases [optional]

N/A

## References

N/A





