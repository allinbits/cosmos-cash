# Messages

In this section we describe the processing of the staking messages and the corresponding updates to the state. All
created/modified state objects specified by each message are defined within the [state](./03_state_transitions.md)
section.

### MsgCreateIssuer

The message has the following fields

- `token` a string indicating the token denom 
- `fee` a int32 
- `issuer_did` a string 
- `license_cred_id` a string 
- `owner` a string 

#### Source

https://github.com/allinbits/cosmos-cash/blob/v2.0.0/proto/issuer/tx.proto#L24

### MsgBurnToken

TODO

#### Source

https://github.com/allinbits/cosmos-cash/blob/v2.0.0/proto/issuer/tx.proto#L38

### MsgMintToken

TODO

#### Source

https://github.com/allinbits/cosmos-cash/blob/v2.0.0/proto/issuer/tx.proto#L52

### MsgPauseToken

TODO

#### Source

https://github.com/allinbits/cosmos-cash/blob/v2.0.0/proto/issuer/tx.proto#L65

### MsgIssueUserCredential

A `MsgIssueUserCredential` is used to issue a new UserCredential, it has the following fields

- `credential` - the verifiable credential
- `owner` - a string containing the cosmos address of the private key signing the transaction

#### Source

https://github.com/allinbits/cosmos-cash/blob/v2.0.0/proto/issuer/tx.proto#L78
