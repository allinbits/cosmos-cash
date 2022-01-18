# Cosmos Cash Modules

- [Cosmos Cash Modules](#cosmos-cash-modules)
    - [DID Module](#did-module)
    - [Verifiable Credential Module](#verifiable-credential-module)
    - [Regulator Module](#regulator-module)
    - [Issuer module](#issuer-module)

The Cosmos Cash project is composed of modules. The following diagram illustrates the module dependencies

![](../assets/diagrams/out/module_dependencies.svg)


### DID Module

The decentralized identifier (DID) module implements the Cosmos [DID](../Explanation/ADR/adr-004-did.md) method and is responsible for all the
operations around DIDs.

- [Source](https://github.com/allinbits/cosmos-cash/tree/main/x/did)
- [Module docs](https://github.com/allinbits/cosmos-cash/tree/main/x/did/spec)
- Dependencies: None

### Verifiable Credential Module

The verifiable credential module is responsible for operations around verifiable credentials. It provides limited
interaction and it is main function is to provide access to shared credential storage and security around verifiable
credential proof validation.

- [Source](https://github.com/allinbits/cosmos-cash/tree/main/x/verifiable-credential)
- [Module docs](https://github.com/allinbits/cosmos-cash/tree/main/x/verifiable-credential/spec)
- Dependencies:
    - [DID Module](#did-module)
    - Account Module

### Regulator Module

The regulator module implements the [root of trust](../Explanation/ADR/adr-007-root-of-trust.md) logic to issue
[registration](../Explanation/ADR/adr-005-registration-credential.md) and [license](..
/Explanation/ADR/adr-006-license-credential.md) credentials.

- [Source](https://github.com/allinbits/cosmos-cash/tree/main/x/regulator)
- [Module docs](https://github.com/allinbits/cosmos-cash/tree/main/x/regulator/spec)
- Dependencies:
    - [Verifiable Credential Module](#verifiable-credential-module)

### Issuer module

The issuer module implements the logic for issuing [fiat-backed stablecoins](../Explanation/ADR/adr-003-issuer.md) 

- [Source](https://github.com/allinbits/cosmos-cash/tree/main/x/issuer)
- [Module docs](https://github.com/allinbits/cosmos-cash/tree/main/x/issuer/spec)
- Dependencies:
    - [DID Module](#did-module)
    - [Verifiable Credential Module](#verifiable-credential-module)
    - Bank module



