<a name="unreleased"></a>
## [Unreleased]


<a name="v0.2.0"></a>
## [v0.2.0] - 2021-06-17
### Chore
- fix lint error
- minor update seed for selective disclosure

### Feat
- add mint to the issuer module ([#92](https://github.com/allinbits/cosmos-cash/issues/92))
- **issuer:** add issuer burning token functionality
- **kyc-cred:** add signature verification to eKYC check
- **kyc-cred:** ante handler for kyc credentials
- **seed:** add params to switch chain for selective disclosure


<a name="v21.06.03"></a>
## v21.06.03 - 2021-06-03
### Chore
- added starport config.yml
- update README.md
- Revert "cleanup proto"
- Revert "revert go mod"
- update deps
- remove unused scaffolded code
- **proto:** remove outdated proto files

### Docs
- add testnet coordinates
- add testnet coordinates

### Feat
- add testing framework
- add identifier module
- add handler and msg_server tests
- update unit tests to test suite
- hash credential attributes using hmac before creating the credentials
- add service unit and integration tests
- add ibc identifier transfer module
- initial commit for issuer module
- add integration test suite
- add authentication data structure to did document
- add service to did document
- **aml:** adding money mule scenario
- **auth:** add authentication to did document
- **creds:** remove issuer creds in favor of merkle root generic creds
- **creds:** using merkle tree to compute root for credentials
- **creds:** update issuer credential to contain relevant fields
- **did:** query for individual did document
- **did:** remove auth from did document
- **docs:** adding verifiable credential spec
- **docs:** adding identifier spec
- **issuer:** distribute tokens to issuer on creation
- **issuer:** adding issuer module scaffolding
- **issuer:** adding create issuer functionality
- **issuer:** use ante handler for issuer authorization
- **rly:** adding basic relayer set up
- **seeds:** add query and validate commands to vsc seeds
- **seeds:** add selective disclosure seed
- **service:** remove a service from a did document
- **test:** add tests for msgs
- **test:** increase test coverage on keeper and client packages
- **tests:** add test framework for issuer keeper
- **tests:** add integration test framework for issuer client package
- **tests:** integration tests for verifiable credential module
- **tests:** increase vcs keeper test coverage
- **vcs:** create verifiable credential data structure
- **vcs:** create verifiable credential proof
- **vcs:** init verifiable credential module
- **vcs:** setup unit and integration test suites
- **vcs:** adding issuer credentials to verifiable credentials module
- **vcs:** checking user credentials an ante handler
- **vcs:** restructure code and updating seeds
- **verify-cred:** validate creds in a query

### Fix
- verifiable credential tests call the incorrect function
- base64 encode public key when adding auth
- use correct app in integration test setup
- relayer round trip broken on acknowledge packet from unregistered interface
- readme
- **integration-test:** add wait times to tranaction commands
- **integration-test:** restructure app.go and fix RegisterTxService


[Unreleased]: https://github.com/allinbits/cosmos-cash/compare/v0.2.0...HEAD
[v0.2.0]: https://github.com/allinbits/cosmos-cash/compare/v21.06.03...v0.2.0
