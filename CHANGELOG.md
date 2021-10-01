<a name="unreleased"></a>
## [Unreleased]


<a name="v2.0.0"></a>
## [v2.0.0] - 2021-10-01
### Chore
- upgrade cosmos-sdk to version v0.44 ([#237](https://github.com/allinbits/cosmos-cash/issues/237))
- use EcdsaSecp256k1VerificationKey2019 as vm type ([#221](https://github.com/allinbits/cosmos-cash/issues/221))

### Docs
- add links to how-to readme page ([#217](https://github.com/allinbits/cosmos-cash/issues/217))
- ADR for the DID module ([#175](https://github.com/allinbits/cosmos-cash/issues/175))
- link documentation in the root README ([#203](https://github.com/allinbits/cosmos-cash/issues/203))

### Feat
- verfiable credential proof verification ([#258](https://github.com/allinbits/cosmos-cash/issues/258))
- add regulator module  ([#247](https://github.com/allinbits/cosmos-cash/issues/247))
- validate controller did ([#252](https://github.com/allinbits/cosmos-cash/issues/252))
- ADR License and registration credential design ([#177](https://github.com/allinbits/cosmos-cash/issues/177))
- add license credential to verifiable credential module ([#200](https://github.com/allinbits/cosmos-cash/issues/200))
- improve support for the DID module (according to adr) ([#230](https://github.com/allinbits/cosmos-cash/issues/230))
- support cosmos-sdk-v0.43 key format ([#220](https://github.com/allinbits/cosmos-cash/issues/220))
- **creds:** restructure vc module to make refactoring easier ([#235](https://github.com/allinbits/cosmos-cash/issues/235))
- **creds:** adding functionality to delete verifiable credentials fr… ([#224](https://github.com/allinbits/cosmos-cash/issues/224))
- **docs:** READMEs for docs, Explanation and Reference Documentation ([#148](https://github.com/allinbits/cosmos-cash/issues/148)) ([#168](https://github.com/allinbits/cosmos-cash/issues/168))
- **issuer:** issuer can pause their emoney token ([#251](https://github.com/allinbits/cosmos-cash/issues/251))
- **issuer:** credential checks for  sending/minting/burning tokens ([#246](https://github.com/allinbits/cosmos-cash/issues/246))

### Fix
- ensure safe and consistent updates for verification relationships ([#219](https://github.com/allinbits/cosmos-cash/issues/219))

### Test
- fix compiler error ([#226](https://github.com/allinbits/cosmos-cash/issues/226))
- **ante:** adding ante handler tests ([#254](https://github.com/allinbits/cosmos-cash/issues/254))

### BREAKING CHANGE

the did method schema has changed according to the ADR

replace verification method type EcdsaSecp256k1RecoveryMethod2020 with EcdsaSecp256k1VerificationKey2019

replace the cli command add-verification-relationship with set-verification-relationships

remove support for legacy key format


<a name="v1.0.0"></a>
## [v1.0.0] - 2021-08-18
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


[Unreleased]: https://github.com/allinbits/cosmos-cash/compare/v2.0.0...HEAD
[v2.0.0]: https://github.com/allinbits/cosmos-cash/compare/v1.0.0...v2.0.0
[v1.0.0]: https://github.com/allinbits/cosmos-cash/compare/v0.2.0...v1.0.0
[v0.2.0]: https://github.com/allinbits/cosmos-cash/compare/v0.2.0-pre...v0.2.0
[v0.2.0-pre]: https://github.com/allinbits/cosmos-cash/compare/v21.06.03...v0.2.0-pre
