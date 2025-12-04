package main

import (
	"advent_04/internal"
)

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()

	grid := internal.NewGrid()

	for scanner.Scan() {
		line := scanner.Text()
		grid.AddRowInput(line)
	}

	println(grid.GetAccessibleCount())
}
