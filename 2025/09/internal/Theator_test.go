package internal

import "testing"

func TestTheater_LargestAreaSquare(t *testing.T) {
	theater := NewTheater()
	theater.AddRow("7,1")
	theater.AddRow("11,1")
	theater.AddRow("11,7")
	theater.AddRow("9,7")
	theater.AddRow("9,5")
	theater.AddRow("2,5")
	theater.AddRow("2,3")
	theater.AddRow("7,3")

	largestArea := theater.LargestAreaSquare()
	expectedLargestArea := 50
	if largestArea != expectedLargestArea {
		t.Errorf("Expected largest area %d, got %d", expectedLargestArea, largestArea)
	}
}

func TestTheater_LargestAreaSquareInGreenTiles(t *testing.T) {
	theater := NewTheater()
	theater.AddRow("7,1")
	theater.AddRow("11,1")
	theater.AddRow("11,7")
	theater.AddRow("9,7")
	theater.AddRow("9,5")
	theater.AddRow("2,5")
	theater.AddRow("2,3")
	theater.AddRow("7,3")

	theater.ThaterMap()
	largestArea := theater.LargestAreaSquareInGreenTiles()
	expectedLargestArea := 24
	if largestArea != expectedLargestArea {
		t.Errorf("Expected largest area %d, got %d", expectedLargestArea, largestArea)
	}
}
