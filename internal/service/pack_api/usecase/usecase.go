package usecase

import (
	"sort"

	"go.uber.org/zap"
)

type UseCase struct {
	packSizes []int
	// here should be other modules like database, cache, third party providers
}

func New(packSizes []int) *UseCase {
	// sort packSizes to ensure we have descending order
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})

	zap.L().Info("useCase module has been initiated",
		zap.Ints("packSizes", packSizes))

	return &UseCase{
		packSizes: packSizes,
	}
}

func (uc *UseCase) CalculatePacksNumber(itemsOrdered int) map[int]int {
	packs := make(map[int]int)

	for i := 0; i < len(uc.packSizes); i++ {
		currentPackSize := uc.packSizes[i]
		// if itemsOrdered more than current pack size, calculate how many packs we can fit within itemsOrdered
		if itemsOrdered/currentPackSize >= 1 {
			packs[currentPackSize] = itemsOrdered / currentPackSize
			// reduce items itemsOrdered by amount that already calculated
			// itemsOrdered = itemsOrdered - currentPackSize * quantity
			itemsOrdered -= currentPackSize * packs[currentPackSize]
		}
	}

	// if there are remaining items, increment the smallest pack number to fit them
	smallestPackSize := uc.packSizes[len(uc.packSizes)-1]
	if itemsOrdered > 0 {
		packs[smallestPackSize]++
	}

	return packs
}
