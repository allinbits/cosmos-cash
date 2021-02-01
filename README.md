# cosmos-cash

Cosmos Cash is a protocol designed to be regulatory compliant that offers the same guarantees as traditional banking systems. Features that enable these guarantees are KYC (Know your customer), AML (Anti-Money Laundering) tracking, FAFT travel rule, and identity management. We use a novel approach to identity management by leveraging W3C specifications for decentralized identifiers and verifiable credentials.

### How to build

- `make install`
- `cosmos-cashd -h`

### How to set up a chain locally

- `cosmos-cashd init --chain-id=cash cash`
- `cosmos-cashd keys add validator`
- `cosmos-cashd keys add-genesis-account $(shell cosmos-cashd keys show validator -a) 1000000000stake`
- `cosmos-cashd keys gentx validator 700000000stake --chain-id cash`
- `cosmos-cashd collect-gentxs`
- `cosmos-cashd start`

