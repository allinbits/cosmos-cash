package types

var (
	// IdentifierKey prefix for each key to a DidDocument
	IdentifierKey = []byte{0x61}
)

const (
	// ModuleName defines the module name
	ModuleName = "identifier"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// DidPrefix defines the did prefix for this chain
	DidPrefix = "did:cash:"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)
