package internal

import "testing"

func TestGridIT_GetAccessibleRecursive(t *testing.T) {

	grid := NewGridIT()
	grid.AddRowInput("..@@.@@@@.")
	grid.AddRowInput("@@@.@.@.@@")
	grid.AddRowInput("@@@@@.@.@@")
	grid.AddRowInput("@.@@@@..@.")
	grid.AddRowInput("@@.@@@@.@@")
	grid.AddRowInput(".@@@@@@@.@")
	grid.AddRowInput(".@.@.@.@@@")
	grid.AddRowInput("@.@@@.@@@@")
	grid.AddRowInput(".@@@@@@@@.")
	grid.AddRowInput("@.@.@@@.@.")

	expected := 43

	if grid.GetAccessibleRecursive() != expected {
		t.Errorf("Expected %d, got %d", expected, grid.GetAccessibleRecursive())
	}

}
