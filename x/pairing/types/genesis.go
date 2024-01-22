package types

import (
	"fmt"

	fixationtypes "github.com/lavanet/lava/x/fixationstore/types"
	timerstoretypes "github.com/lavanet/lava/x/timerstore/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		UniquePaymentStorageClientProviderList: []UniquePaymentStorageClientProvider{},
		ProviderPaymentStorageList:             []ProviderPaymentStorage{},
		EpochPaymentsList:                      []EpochPayments{},
		BadgeUsedCuList:                        []BadgeUsedCu{},
		BadgesTS:                               *timerstoretypes.DefaultGenesis(),
		ProviderQosFS:                          *fixationtypes.DefaultGenesis(),
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in uniquePaymentStorageClientProvider
	uniquePaymentStorageClientProviderIndexMap := make(map[string]struct{})

	for _, elem := range gs.UniquePaymentStorageClientProviderList {
		index := string(UniquePaymentStorageClientProviderKey(elem.Index))
		if _, ok := uniquePaymentStorageClientProviderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for uniquePaymentStorageClientProvider")
		}
		uniquePaymentStorageClientProviderIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in providerPaymentStorage
	providerPaymentStorageIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProviderPaymentStorageList {
		index := string(ProviderPaymentStorageKey(elem.Index))
		if _, ok := providerPaymentStorageIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for providerPaymentStorage")
		}
		providerPaymentStorageIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in epochPayments
	epochPaymentsIndexMap := make(map[string]struct{})

	for _, elem := range gs.EpochPaymentsList {
		index := string(EpochPaymentsKey(elem.Index))
		if _, ok := epochPaymentsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for epochPayments")
		}
		epochPaymentsIndexMap[index] = struct{}{}
	}

	// check the badgeUsedCuIndex map is empty
	if len(gs.BadgeUsedCuList) > 0 {
		return fmt.Errorf("badgeUsedCuList is not empty")
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
