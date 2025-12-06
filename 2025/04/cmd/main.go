package main

import (
	"advent_04/internal"
	"fmt"
)

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()

	grid := internal.NewGrid()

	for scanner.Scan() {
		line := scanner.Text()
		grid.AddRowInput(line)
	}

	scanner2, closer2 := internal.Scan("input.txt")
	defer closer2.Close()
	gridIT := internal.NewGridIT()
	for scanner2.Scan() {
		line := scanner2.Text()
		gridIT.AddRowInput(line)
	}

	fmt.Printf("Accessible Paper Rolls (Recursive): %d\n", grid.GetAccessibleCount())
	fmt.Printf("Accessible Paper Rolls (Iterative): %d\n", gridIT.GetAccessibleRecursive())

}
