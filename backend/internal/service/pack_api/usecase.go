package pack_api

type UseCase interface {
	CalculatePacksNumber(itemsOrdered int) (packs map[int]int)
}
