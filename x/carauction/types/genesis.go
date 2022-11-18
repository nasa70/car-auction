package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LotsList:   []Lots{},
		AssetsList: []Assets{},
		AuctionInfo: &AuctionInfo{
			FirstInQueueLotId: 0,
			LastInQueueLotId:  0,
			FirstFinishTime:   0,
		},
		LotsQueueList: []LotsQueue{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in lots
	lotsIdMap := make(map[uint64]bool)
	lotsCount := gs.GetLotsCount()
	for _, elem := range gs.LotsList {
		if _, ok := lotsIdMap[elem.LotId]; ok {
			return fmt.Errorf("duplicated id for lots")
		}
		if elem.LotId >= lotsCount {
			return fmt.Errorf("lots id should be lower or equal than the last id")
		}
		lotsIdMap[elem.LotId] = true
	}
	// Check for duplicated ID in assets
	assetsIdMap := make(map[uint64]bool)
	assetsCount := gs.GetAssetsCount()
	for _, elem := range gs.AssetsList {
		if _, ok := assetsIdMap[elem.AssetId]; ok {
			return fmt.Errorf("duplicated id for assets")
		}
		if elem.AssetId >= assetsCount {
			return fmt.Errorf("assets id should be lower or equal than the last id")
		}
		assetsIdMap[elem.AssetId] = true
	}
	// Check for duplicated ID in lotsQueue
	lotsQueueIdMap := make(map[uint64]bool)
	lotsQueueCount := gs.GetLotsQueueCount()
	for _, elem := range gs.LotsQueueList {
		if _, ok := lotsQueueIdMap[elem.LotId]; ok {
			return fmt.Errorf("duplicated id for lotsQueue")
		}
		if elem.LotId >= lotsQueueCount {
			return fmt.Errorf("lotsQueue id should be lower or equal than the last id")
		}
		lotsQueueIdMap[elem.LotId] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
