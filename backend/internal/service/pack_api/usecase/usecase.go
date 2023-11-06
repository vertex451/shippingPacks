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
	var currentPackSize int

	for i := 0; i < len(uc.packSizes); i++ {
		currentPackSize = uc.packSizes[i]
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

	// If we have two packs, we should replace it with one bigger pack.
	return uc.replaceTwoSmallerPacksWithOneBigger(packs)
}

// If we have two packs, we should replace it with one bigger pack.
// In case of the input map[int]int{5000: 1, 2000: 1, 1000: 1, 500: 1, 250: 2}
// We should replace it with map[int]int{5000: 1, 2000: 1, 1000: 1, 500: 2}
// Then with map[int]int{5000: 1, 2000: 1, 1000: 2} and so on.
// The final result should be map[int]int{5000: 2}, if 5000 is the max pack size
func (uc *UseCase) replaceTwoSmallerPacksWithOneBigger(packs map[int]int) map[int]int {
	var ok bool
	var currentPackSize, prevPackSize int

	// since we have descending sorting in default pack size, we should iterate the opposite order.
	// also, to avoid out of range runtime error, lets start not from the smallest but from the next after the smallest.
	for i := len(uc.packSizes) - 2; i >= 0; i-- {
		currentPackSize = uc.packSizes[i]
		prevPackSize = uc.packSizes[i+1]
		// if we have a number of two of specific pack, lets replace them with the one bigger pack
		if packs[prevPackSize] == 2 {
			delete(packs, prevPackSize)
			packs[currentPackSize]++
		}

		// no need to iterate further if we don't have the next pack size in the resul.
		if _, ok = packs[currentPackSize]; !ok {
			break
		}
	}

	return packs
}
