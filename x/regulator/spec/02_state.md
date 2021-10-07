# State

This document describes the state pertaining to the regulator module. The regulator module relies on the verification
credential module for type definition.

## Genesis

The regulator module requires custom genesis parameters

### Regulators

The regulators message contains the following attributes:

`addresses`: the list of root of trust addresses

#### Source

+++ https://github.com/allinbits/cosmos-cash/blob/v2.0.0/proto/regulator/genesis.proto#L19


