package types

const (
	// ModuleName defines the module name
	ModuleName = "rollapp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rollapp"
)

var (
	ParamsKey = []byte("p_rollapp")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
