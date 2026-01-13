package main

import (
	"advent_09/internal"
	"fmt"
)

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()
	theater := internal.NewTheater()

	for scanner.Scan() {
		line := scanner.Text()
		theater.AddRow(line)
	}

	largestArea := theater.LargestAreaSquare()
	fmt.Printf("Largest area of square: %d\n", largestArea)
	largestAreaInGreenTiles := theater.LargestAreaSquareInGreenTiles()
	fmt.Printf("Largest area of square in green tiles: %d\n", largestAreaInGreenTiles)

}
