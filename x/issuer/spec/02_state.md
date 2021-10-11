# State

This document describes the state pertaining to the issuer module. The issuer module relies partially on the
verification credential module for type definition, specifically for `MsgIssueUserCredential`

### Issuer

An issuer has the following fields:

- `token`: a string containing the denom of the token being issued
- `fee`: a int32 describing the fee amount for issuing a token
- `issuer_did`: the did of the issuer 
- `paused`: whenever the token is paused 

#### Source

https://github.com/allinbits/cosmos-cash/blob/v2.0.0/proto/issuer/issuer.proto#L7


