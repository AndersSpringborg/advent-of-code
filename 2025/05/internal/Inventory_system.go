package internal

import "fmt"

type InventorySystem struct {
	availableRanges [][2]int
}

func NewInventorySystem() *InventorySystem {
	return &InventorySystem{
		availableRanges: make([][2]int, 0),
	}
}

func (is *InventorySystem) AddAvaliableIngredient(ingredientsRange string) {
	start, end := parseRange(ingredientsRange)
	is.availableRanges = append(is.availableRanges, [2]int{start, end})
}

func parseRange(ingredientsRange string) (int, int) {
	var start, end int
	fmt.Sscanf(ingredientsRange, "%d-%d", &start, &end)
	return start, end
}

func ParseIngredient(line string) int {
	var ingredient int
	fmt.Sscanf(line, "%d", &ingredient)
	return ingredient
}

func (is *InventorySystem) IsIngredientFresh(ingredient int) bool {
	for _, r := range is.availableRanges {
		start, end := r[0], r[1]
		if ingredient >= start && ingredient <= end {
			return true
		}
	}
	return false
}

func (is *InventorySystem) PrintAvailableRanges() {
	for _, r := range is.availableRanges {
		fmt.Printf("%d-%d\n", r[0], r[1])
	}
}

func (is *InventorySystem) UniqueIngredientsInRanges() int {
	ranges := is.mergeOverlappingRanges()
	uniqueIngredients := 0
	for _, r := range ranges {
		uniqueIngredients += r[1] - r[0] + 1
	}

	return uniqueIngredients
}

func (is *InventorySystem) mergeOverlappingRanges() [][2]int {
	if len(is.availableRanges) == 0 {
		return [][2]int{}
	}

	sortedRanges := make([][2]int, len(is.availableRanges))
	copy(sortedRanges, is.availableRanges)
	for i := 0; i < len(sortedRanges)-1; i++ {
		for j := i + 1; j < len(sortedRanges); j++ {
			if sortedRanges[i][0] > sortedRanges[j][0] {
				sortedRanges[i], sortedRanges[j] = sortedRanges[j], sortedRanges[i]
			}
		}
	}

	mergedRanges := make([][2]int, 0)
	currentRange := sortedRanges[0]

	for i := 1; i < len(sortedRanges); i++ {
		if sortedRanges[i][0] <= currentRange[1]+1 {
			if sortedRanges[i][1] > currentRange[1] {
				currentRange[1] = sortedRanges[i][1]
			}
		} else {
			mergedRanges = append(mergedRanges, currentRange)
			currentRange = sortedRanges[i]
		}
	}
	mergedRanges = append(mergedRanges, currentRange)

	return mergedRanges

}
