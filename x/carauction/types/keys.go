package types

const (
	// ModuleName defines the module name
	ModuleName = "carauction"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_carauction"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	LotsKey      = "Lots-value-"
	LotsCountKey = "Lots-count-"
)

const (
	AssetsKey      = "Assets-value-"
	AssetsCountKey = "Assets-count-"
)

const (
	AuctionInfoKey = "AuctionInfo-value-"
)

const (
	LotsQueueKey      = "LotsQueue-value-"
	LotsQueueCountKey = "LotsQueue-count-"
)
