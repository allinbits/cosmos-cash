package types

var (
	// DidDocumentKey prefix for each key to a DidDocument
	DidDocumentKey = []byte{0x61}
)

const (
	// ModuleName defines the module name
	ModuleName = "did"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// DidPrefix defines the did prefix for this chain
	DidPrefix = "did:cash:"

	// DidKeyPrefix defines the did key prefix
	DidKeyPrefix = "did:key:"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)
