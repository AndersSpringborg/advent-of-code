package internal

import "testing"

func TestInventorySystem(t *testing.T) {
	inventory := NewInventorySystem()

	inventory.AddAvaliableIngredient("3-5")
	inventory.AddAvaliableIngredient("10-14")
	inventory.AddAvaliableIngredient("16-20")
	inventory.AddAvaliableIngredient("12-18")

	ingredients := []int{1, 5, 8, 11, 17, 32}
	fresh := 0
	for _, ingredient := range ingredients {
		if inventory.IsIngredientFresh(ingredient) {
			fresh++
		}
	}

	expected := 3
	if fresh != expected {
		t.Errorf("Expected %d, got %d", expected, fresh)
	}

	inventory.PrintAvailableRanges()
}
