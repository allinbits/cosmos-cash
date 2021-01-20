package types

var (
	// PortKey defines the key to store the port ID in store
	PortKey = []byte{0x99}
)

const (
	// ModuleName defines the module name
	ModuleName = "ibcidentifier"

	// PortID is the default port id that transfer module binds to
	PortID = ModuleName

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
