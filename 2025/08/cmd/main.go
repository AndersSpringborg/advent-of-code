package main

import (
	"advent_08/internal"
)

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()
	junctionBoxMap := internal.NewJunctionBoxMap()

	for scanner.Scan() {
		line := scanner.Text()
		junctionBoxMap.ParseJunctionBox(line)

	}

	// part 1
	junctionBoxMap.ConnectionAllJunctionBoxesToShortestDistance(1000)
	biggestCircuts := junctionBoxMap.GetBiggestCircuitsSum(3)

	// part 2
	biggestCircutsMultiply := junctionBoxMap.GetLastConnectingPairXMultiplyX()
	println("Sum of biggest circuts:", biggestCircuts)
	println("X multiply X of last connecting pair:", biggestCircutsMultiply)

}
