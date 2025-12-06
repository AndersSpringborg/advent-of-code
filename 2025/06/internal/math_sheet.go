package internal

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type MathSheet struct {
	// [y][x]
	equations [][]string
}

func NewMathSheet() *MathSheet {
	return &MathSheet{
		equations: make([][]string, 0),
	}
}

func (m *MathSheet) ParseEquationInput(line string) {
	equation := ParseEquation(line)
	m.equations = append(m.equations, equation)
}

func ParseEquation(line string) []string {
	row := strings.Fields(line)
	return row
}

func (ms *MathSheet) printColumn(x int) {
	fmt.Printf("Column %d: \n", x)
	numEquations := len(ms.equations)
	for y := 0; y < numEquations; y++ {
		termStr := ms.equations[y][x]
		fmt.Printf("%s\n", termStr)
	}
	fmt.Println()

}

func (ms *MathSheet) SumEquations() int {
	sum := 0

	numEquations := len(ms.equations)
	numTerms := len(ms.equations[0])

	for c := 0; c < numTerms; c++ {
		equationResult := 0
		operator := ms.equations[numEquations-1][c]

		for r := 0; r < numEquations; r++ {
			if r == numEquations-1 {
				continue
			}
			termStr := ms.equations[r][c]
			term, err := strconv.Atoi(termStr)
			if err != nil {
				log.Fatal("Error converting term to integer:", err)
			}
			if r == 0 {
				equationResult = term
				continue
			}
			switch operator {
			case "+":
				equationResult += term
			case "*":
				equationResult *= term
			}
		}

		sum += equationResult
	}

	return sum
}

func (m *MathSheet) ToString() string {
	result := ""
	for _, equation := range m.equations {
		result += strings.Join(equation, " ") + "\n"
	}
	return result
}
