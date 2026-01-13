package main

import (
	"advent_12/internal"
	"fmt"
)

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()

	lines := ""

	for scanner.Scan() {
		line := scanner.Text()
		lines += line + "\n"
	}

	giftPacker := internal.NewGiftPacker()
	giftPacker.ParseText(lines)

	fmt.Printf("%d presents fit into the regions\n", giftPacker.FitPresentsInRegions())
}
