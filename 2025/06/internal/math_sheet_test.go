package internal

import (
	"fmt"
	"testing"
)

func TestMathSheet(t *testing.T) {
	ms := NewMathSheet()
	ms.ParseEquationInput("123 328  51 64")
	ms.ParseEquationInput(" 45 64  387 23")
	ms.ParseEquationInput("  6 98  215 314")
	ms.ParseEquationInput("*   +   *   +")

	fmt.Printf("mathsheet: \n%s", ms.ToString())

	sum := ms.SumEquations()

	expectedSum := 4277556

	if sum != expectedSum {
		t.Errorf("Expected sum %d, got %d", expectedSum, sum)
	}

}
