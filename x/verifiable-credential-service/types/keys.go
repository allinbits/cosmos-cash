package types

var (
	// VerifiableCredentialKey prefix for each key to a DidDocument
	VerifiableCredentialKey = []byte{0x62}
)

const (
	// ModuleName defines the module name
	ModuleName = "verifiablecredentialservice"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability_vcs"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
