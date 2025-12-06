package main

import (
	"advent_06/internal"
)

func main() {
	scanner, closer := internal.Scan("input.txt")
	defer closer.Close()
	mathsheet := internal.NewMathSheet()
	mathSheetReal := internal.NewMathSheetReal()

	for scanner.Scan() {
		line := scanner.Text()
		mathsheet.ParseEquationInput(line)
		mathSheetReal.ParseEquationInput(line)
	}

	sum := mathsheet.SumEquations()
	sumReal := mathSheetReal.SumEquations()
	println("Sum of equations:", sum)
	println("Sum of equations (real):", sumReal)
}
