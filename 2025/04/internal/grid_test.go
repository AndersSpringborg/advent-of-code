package internal

import "testing"

func TestGrid_GetAccessibleCount(t *testing.T) {

	grid := NewGrid()
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

	expected := 13

	if grid.GetAccessibleCount() != expected {
		t.Errorf("Expected %d, got %d", expected, grid.GetAccessibleCount())
	}

}
