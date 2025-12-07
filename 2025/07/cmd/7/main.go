package main

import (
	"advent_07/internal"
)

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()

	manifold := internal.NewManifold(141)
	manifold2 := internal.NewManifold2()

	for scanner.Scan() {
		line := scanner.Text()
		manifold.ParseLine(line)
		manifold2.ParseLine(line)
	}

	splits := manifold.GetSplits()
	timelines := manifold2.GetTimelinesDP()
	println("Total splits:", splits)
	println("Total timelines:", timelines)
}
