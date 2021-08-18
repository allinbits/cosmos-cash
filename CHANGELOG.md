<a name="unreleased"></a>
## [Unreleased]


<a name="1.0.0"></a>
## [1.0.0] - 2021-08-18
### Chore
- rename verifiable-cerdential module ([#167](https://github.com/allinbits/cosmos-cash/issues/167))
- comment formatting in cli transaction ([#161](https://github.com/allinbits/cosmos-cash/issues/161))
- **did:** renaming identfifer module to did
- **docs:** Add conventionalcommits checks ([#144](https://github.com/allinbits/cosmos-cash/issues/144)) ([#159](https://github.com/allinbits/cosmos-cash/issues/159))
- **proto:** remove ibc identifier proto files ([#105](https://github.com/allinbits/cosmos-cash/issues/105))
- **release:** add make targets and config to simplify releases ([#136](https://github.com/allinbits/cosmos-cash/issues/136))

### Docs
- expand contributing documentation ([#176](https://github.com/allinbits/cosmos-cash/issues/176))
- **adr:** adding the adr process documentation ([#120](https://github.com/allinbits/cosmos-cash/issues/120))

### Feat
- improve support for verification material ([#184](https://github.com/allinbits/cosmos-cash/issues/184))
- add support for DID Metadata ([#179](https://github.com/allinbits/cosmos-cash/issues/179))
- **ADR:** Documentation Strategy ([#129](https://github.com/allinbits/cosmos-cash/issues/129)) ([#152](https://github.com/allinbits/cosmos-cash/issues/152))
- **did:** update identifier module to version 1.0 of W3C specs
- **did:** explicitly list the supported verification relationships ([#163](https://github.com/allinbits/cosmos-cash/issues/163))
- **did:** adding commands for relationships and controllers
- **did:** add logging to the msg_server
- **did:** add blockchain account ID to verification method struct ([#135](https://github.com/allinbits/cosmos-cash/issues/135))
- **docs:** implement docs structure ([#140](https://github.com/allinbits/cosmos-cash/issues/140)) ([#157](https://github.com/allinbits/cosmos-cash/issues/157))
- **docs:** add CODEOWNERS file ([#158](https://github.com/allinbits/cosmos-cash/issues/158))
- **errors:** update issuer errors ([#110](https://github.com/allinbits/cosmos-cash/issues/110))
- **events:** add events for identifier, issuer and credenitals   ([#104](https://github.com/allinbits/cosmos-cash/issues/104))
- **swagger:** add swagger docs and ui to repo ([#162](https://github.com/allinbits/cosmos-cash/issues/162))

### Fix
- add logger to test suite
- **creds:** use enums for credential types ([#106](https://github.com/allinbits/cosmos-cash/issues/106))
- **issuer:** adding check so issuers cannot create the same token


<a name="v0.2.0"></a>
## [v0.2.0] - 2021-06-23
### Chore
- remove ibc-identifier module


<a name="v0.2.0-pre"></a>
## [v0.2.0-pre] - 2021-06-18
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


[Unreleased]: https://github.com/allinbits/cosmos-cash/compare/1.0.0...HEAD
[1.0.0]: https://github.com/allinbits/cosmos-cash/compare/v0.2.0...1.0.0
[v0.2.0]: https://github.com/allinbits/cosmos-cash/compare/v0.2.0-pre...v0.2.0
[v0.2.0-pre]: https://github.com/allinbits/cosmos-cash/compare/v21.06.03...v0.2.0-pre
