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

In the context of Cosmos Cash we can identify for actors:

- Regulator
- Identity Provider (eIDAS)
- EMT/ART (crypto asset issuer)
- Account holder

## Decision

![root of trust](../../assets/diagrams/out/root_of_trust.svg)

### Credential Verfication

![credential validation](../../assets/diagrams/out/credentials_validate_proof.svg)

### Credential deletion

![credential deletion](../../assets/diagrams/out/credentials_delete.svg)

## Consequences

### Backwards Compatibility

This is a new module so backward compatibility is not a concern.

### Positive

- The implementation of the ADR provides the foundation for interoperability with the DID standard and the SSI identity
  approach.
- Closely following the W3C standard gives the best chances of successful interoperability with third-party components.

### Negative

- The implementation rigidly follows the W3C specification which leaves little room for extensibility. This approach
  might become an issue for wider adoption.

### Neutral

N/A

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section contains a summary of issues to be solved in future
iterations. The issues summarized here can reference comments from a pull request discussion. Later, this section can
optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

- The `did:key` method specifies a key format that is different from the one used in this ADR. This ADR needs to be
  amended or follow a different approach.
- The approach proposed is somewhat locked into the current implementation and will have to be revised in successive
  iterations.

## Test Cases [optional]

N/A

## References

- [DID Core](https://www.w3.org/TR/did-core)
- [DID Specification Registries](https://w3c.github.io/did-spec-registries)





