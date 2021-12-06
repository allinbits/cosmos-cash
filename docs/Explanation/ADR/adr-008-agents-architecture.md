# ADR 008: SSI Agents Architecture

## Changelog

- 2021-10-11: Initial draft

## Status

PROPOSED

## Abstract

Cosmos Cash relies on [self-sovereign identity (SSI)](../../Reference/GLOSSARY.md#self-sovereign-identity-ssi) for establishing secure channels of communication between 
actors. In SSI context actors establish secure channels and communicate using software components called agents.  

This document illustrates how SSI agents interacts with each other using Cosmos Cash as a [verifiable data registry (VDR)](../../Reference/GLOSSARY.md#verifiable-data-registry-vdr).

## Context

TBD

## Decision

The SSI agents general architecture is represented in the following diagram:

![ssi_agents_architecture](../../assets/diagrams/out/ssi_agents_architecture.svg)

Connections between edge agents are routed through cloud agents (see [DIDComm Mediator](https://wiki.hyperledger.org/display/ARIES/DIDComm+MediatorRouter))

### Interactions

The general architecture enables the support for different scenarios, in the Cosmos Cash research the following interacitons 
have been developed:
![ssi_agents_interactions](../../assets/diagrams/out/ssi_agents_interactions.svg)

##### Legend
- Blue Lines: on-chain commits 
- Red Lines: DIDComm exchanges
- Grey dotted lines: Out-of-band exchanges

## Consequences

N/A

### Backward Compatibility

N/A

### Positive

N/A

### Negative

N/A

### Neutral

N/A

## Further Discussions

N/A

## Test Cases [optional]

N/A

## References

N/A





