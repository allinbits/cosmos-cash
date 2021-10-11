# Messages

In this section we describe the processing of the staking messages and the corresponding updates to the state. All
created/modified state objects specified by each message are defined within the [state](./03_state_transitions.md)
section.

### MsgIssueRegulatorCredential

A `MsgIssueRegulatorCredential` is used to issue a new RegulatorCredential, it has the following fields

- `credential` - the verifiable credential
- `owner` - a string containing the cosmos address of the private key signing the transaction

#### Source

https://github.com/allinbits/cosmos-cash/blob/v2.0.0/proto/regulator/tx.proto#L27

### MsgIssueRegistrationCredential

A `MsgIssueRegistrationCredential` is used to issue a new RegistrationCredential, it has the following fields

- `credential` - the verifiable credential
- `owner` - a string containing the cosmos address of the private key signing the transaction

#### Source

https://github.com/allinbits/cosmos-cash/blob/v2.0.0/proto/regulator/tx.proto#L40

### MsgIssueLicenseCredential

A `MsgIssueLicenseCredential` is used to issue a new LicenseCredential, it has the following fields

- `credential` - the verifiable credential
- `owner` - a string containing the cosmos address of the private key signing the transaction

#### Source

https://github.com/allinbits/cosmos-cash/blob/v2.0.0/proto/regulator/tx.proto#L51
