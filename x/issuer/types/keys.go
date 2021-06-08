package types

// Key prefixes
var (
	IssuerKey = []byte{0x61}
	TokenKey  = []byte{0x62}
)

const (
	// ModuleName defines the module name
	ModuleName = "issuer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
