package internal

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type MathSheetReal struct {
	// [r][c]
	equations [][]string
}

func NewMathSheetReal() *MathSheetReal {
	return &MathSheetReal{
		equations: make([][]string, 0),
	}
}

func (m *MathSheetReal) ParseEquationInput(line string) {
	equation := m.parseEquation(line)
	m.equations = append(m.equations, equation)
}

func (ms *MathSheetReal) parseEquation(line string) []string {
	row := strings.Split(line, "")
	return row
}

func (ms *MathSheetReal) printColumn(x int) {
	fmt.Printf("Column %d: \n", x)
	numEquations := len(ms.equations)
	for y := 0; y < numEquations; y++ {
		termStr := ms.equations[y][x]
		fmt.Printf("%s\n", termStr)
	}
	fmt.Println()

}

func (ms *MathSheetReal) SumEquations() int {

	equationsSum := 0
	transposeEquations := ms.transposeEquations()
	for i, eq := range transposeEquations {
		fmt.Printf("Transposed Equation %d: %s\n", i, eq)
	}

	for _, equation := range transposeEquations {
		println()
		sum := 0
		numbers := strings.Fields(equation)
		operator := numbers[len(numbers)-1]
		for i, numStr := range numbers {
			// last number is operator
			if i == len(numbers)-1 {
				// skip operator char
				continue
			}

			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}

			print(num)
			print(" ", operator, " ")

			if i == 0 {
				sum = num
				continue
			}

			switch operator {
			case "+":
				sum += num
			case "*":
				sum *= num
			default:
				log.Fatalf("Unknown operator: %s", operator)
			}
		}

		equationsSum += sum
	}

	return equationsSum

}

func (m *MathSheetReal) ToString() string {
	result := ""
	for _, equation := range m.equations {
		result += strings.Join(equation, " ") + "\n"
	}
	return result
}

//  64
//  23
//  314
//   +
// bliver til
// 4 431 623 +

func (m *MathSheetReal) transposeEquations() []string {
	transposedEquations := make([]string, 0)

	equation := ""
	rows := len(m.equations)
	columns := len(m.equations[0])
	for c := columns - 1; c >= 0; c-- {
		transposedRow := ""
		for r := 0; r < rows; r++ {
			// is last row, and parsing operator
			if r == rows-1 {
				transposedRow += " "
			}
			transposedRow += m.equations[r][c]

		}

		isFirstColumn := c == 0
		if !isSpaceColumn(transposedRow) || isFirstColumn {
			equation = equation + transposedRow

		} else {
			transposedEquations = append(transposedEquations, equation)
			equation = ""
		}
	}

	return transposedEquations
}

func isSpaceColumn(row string) bool {
	return strings.TrimSpace(row) == ""

}
