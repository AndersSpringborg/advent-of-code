package main

import (
	"advent_05/internal"
	"fmt"
)

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()

	iventorySystem := internal.NewInventorySystem()
	storage := true
	fresh := 0
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if line == "" {
			fmt.Printf("Finished reading storage at line %d\n", lineNumber)
			storage = false
			continue
		} else if storage {
			iventorySystem.AddAvaliableIngredient(line)
		} else {

			ingredient := internal.ParseIngredient(line)
			if iventorySystem.IsIngredientFresh(ingredient) {
				fresh++
			}
		}
	}

	iventorySystem.PrintAvailableRanges()

	fmt.Printf("fresh ingredients: %d\n", fresh)
	fmt.Printf("unique available ingredients: %d\n", iventorySystem.UniqueIngredientsInRanges())

}
