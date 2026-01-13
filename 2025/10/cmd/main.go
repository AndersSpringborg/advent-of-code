package main

import (
	"advent_10/internal"
	"fmt"
)

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()
	manual := internal.NewManual()

	for scanner.Scan() {
		line := scanner.Text()
		manual.ParseLine(line)
	}

	fewestPresses := manual.FewestPresses()
	fewestPressesJoltage := manual.FewestPressesVoltage()
	fmt.Printf("Fewest button presses to solve all machines: %d\n", fewestPresses)
	fmt.Printf("Fewest button presses to solve all machines ojltage: %d\n", fewestPressesJoltage)
}
