package types

var (
	IdentifierKey = []byte{0x61} // prefix for each key to a DidDocument
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

	// QuerierRoute defines the module's query routing key
	DidPrefix = "did:cash:"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
