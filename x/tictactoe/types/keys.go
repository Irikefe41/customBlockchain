package types

const (
	// ModuleName defines the module name
	ModuleName = "tictactoe"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tictactoe"
)

var (
	ParamsKey = []byte("p_tictactoe")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
