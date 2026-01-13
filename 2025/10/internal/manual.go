package internal

import (
	"fmt"
	"math"
	"strings"

	"github.com/lanl/clp"
)

type Machine struct {
	lightDiagram       []byte
	buttonWiring       [][]int
	joltageRequirement []int
}

func NewMachine() Machine {
	return Machine{
		lightDiagram:       make([]byte, 0),
		buttonWiring:       make([][]int, 0),
		joltageRequirement: make([]int, 0),
	}
}

func (m *Machine) ParsePart(part string) {
	if strings.HasPrefix(part, "[") {
		partAsArray := []byte(part)
		for _, char := range partAsArray {
			if char != '[' && char != ']' {
				m.lightDiagram = append(m.lightDiagram, char)
			}
		}
	} else if strings.HasPrefix(part, "(") {
		nums := strings.Split(part[1:len(part)-1], ",")
		wiring := make([]int, 0)
		for _, numStr := range nums {
			var num int
			fmt.Sscanf(numStr, "%d", &num)
			wiring = append(wiring, num)
		}
		m.buttonWiring = append(m.buttonWiring, wiring)
	} else if strings.HasPrefix(part, "{") {
		nums := strings.Split(part[1:len(part)-1], ",")
		for _, numStr := range nums {
			var num int
			fmt.Sscanf(numStr, "%d", &num)
			m.joltageRequirement = append(m.joltageRequirement, num)
		}
	}
}

func lightIndex(lights []byte) int {
	index := 0
	for i, light := range lights {
		if light == '#' {
			index += 1 << i
		}
	}
	return index
}

func (m *Machine) FewestPresses() int {
	amountOfLights := len(m.lightDiagram)
	lightMutations := int(math.Pow(2, float64(amountOfLights)))
	buttons := len(m.buttonWiring)

	targetState := lightIndex(m.lightDiagram)

	dpTable := make([][]int, buttons+1)
	for i := range dpTable {
		dpTable[i] = make([]int, lightMutations)
		for j := range dpTable[i] {
			dpTable[i][j] = math.MaxInt32
		}
	}
	initialState := 0
	dpTable[0][initialState] = 0

	for i := 1; i <= buttons; i++ {
		for currentState := 0; currentState < lightMutations; currentState++ {
			if dpTable[i-1][currentState] == math.MaxInt32 {
				continue
			}

			dpTable[i][currentState] = min(dpTable[i][currentState], dpTable[i-1][currentState])

			newState := currentState
			for _, lightIdx := range m.buttonWiring[i-1] {
				newState ^= (1 << lightIdx)
			}
			dpTable[i][newState] = min(dpTable[i][newState], dpTable[i-1][currentState]+1)
		}
	}

	result := dpTable[buttons][targetState]
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func (m *Machine) String() string {
	return fmt.Sprintf("LightDiagram: %s, ButtonWiring: %v, JoltageRequirement: %v",
		string(m.lightDiagram), m.buttonWiring, m.joltageRequirement)
}

// FewestPressesVoltage solves Part 2 using Integer Linear Programming.
//
// Problem: minimize sum(x) subject to Ax = b, x >= 0, x integer
// where A[i][j] = 1 if button j affects counter i, and b = joltage requirements
//
// IMPORTANT: The LP relaxation can give fractional solutions that are not achievable
// with integers, so we need proper ILP or branch-and-bound.
func (m *Machine) FewestPressesVoltage() int {
	numButtons := len(m.buttonWiring)
	numCounters := len(m.joltageRequirement)

	if numButtons == 0 || numCounters == 0 {
		return 0
	}

	// Try LP first, then if non-integer, use branch and bound
	lpResult, lpSolution := m.solveLP()
	if lpResult < 0 {
		return -1 // Infeasible
	}

	// Check if LP solution is integer
	isInteger := true
	for _, val := range lpSolution {
		if math.Abs(val-math.Round(val)) > 1e-6 {
			isInteger = false
			break
		}
	}

	if isInteger {
		return int(math.Round(lpResult))
	}

	// LP solution is fractional - need branch and bound
	return m.branchAndBound(lpResult, lpSolution)
}

func (m *Machine) solveLP() (float64, []float64) {
	numButtons := len(m.buttonWiring)
	numCounters := len(m.joltageRequirement)

	simp := clp.NewSimplex()

	// EasyLoadDenseProblem signature:
	// func (s *Simplex) EasyLoadDenseProblem(obj []float64, bounds [][2]float64, ineqs [][]float64)
	// - obj: objective coefficients (one per column/variable)
	// - bounds: [lower, upper] for each variable
	// - ineqs: each row is [lower, coef1, coef2, ..., coefN, upper] for constraint

	// Objective: minimize sum of all variables
	obj := make([]float64, numButtons)
	for i := range obj {
		obj[i] = 1.0
	}

	// Bounds for each variable: [0, +inf]
	bounds := make([][2]float64, numButtons)
	for i := range bounds {
		bounds[i] = [2]float64{0, math.Inf(1)}
	}

	// Inequality constraints: each row is [lower, coef1, ..., coefN, upper]
	// For equality Ax = b, we use lower = upper = b
	ineqs := make([][]float64, numCounters)
	for i := 0; i < numCounters; i++ {
		// Row format: [lower_bound, coef_0, coef_1, ..., coef_n-1, upper_bound]
		row := make([]float64, numButtons+2)
		row[0] = float64(m.joltageRequirement[i])            // lower bound
		row[numButtons+1] = float64(m.joltageRequirement[i]) // upper bound

		// Fill coefficients
		for btnIdx, wiring := range m.buttonWiring {
			for _, counterIdx := range wiring {
				if counterIdx == i {
					row[btnIdx+1] = 1.0
				}
			}
		}
		ineqs[i] = row
	}

	simp.EasyLoadDenseProblem(obj, bounds, ineqs)
	simp.SetOptimizationDirection(clp.Minimize)

	status := simp.Primal(clp.NoValuesPass, clp.NoStartFinishOptions)
	if status != clp.Optimal {
		return -1, nil
	}

	return simp.ObjectiveValue(), simp.PrimalColumnSolution()
}

func (m *Machine) branchAndBound(lpLowerBound float64, lpSolution []float64) int {
	numButtons := len(m.buttonWiring)
	numCounters := len(m.joltageRequirement)

	// Build constraint matrix for checking solutions
	matrix := make([][]int, numCounters)
	for i := 0; i < numCounters; i++ {
		matrix[i] = make([]int, numButtons)
	}
	for btnIdx, wiring := range m.buttonWiring {
		for _, counterIdx := range wiring {
			if counterIdx < numCounters {
				matrix[counterIdx][btnIdx] = 1
			}
		}
	}

	// Find upper bounds for each variable
	upperBounds := make([]int, numButtons)
	for j := 0; j < numButtons; j++ {
		upperBounds[j] = 0
		for i := 0; i < numCounters; i++ {
			if matrix[i][j] > 0 {
				if m.joltageRequirement[i] > upperBounds[j] {
					upperBounds[j] = m.joltageRequirement[i]
				}
			}
		}
		// Use LP solution as guide - limit to 2x LP value
		if lpSolution != nil && lpSolution[j] > 0 {
			lpBound := int(math.Ceil(lpSolution[j] * 2))
			if lpBound < upperBounds[j] {
				upperBounds[j] = lpBound
			}
		}
	}

	// Check if search space is reasonable
	searchSize := int64(1)
	for _, ub := range upperBounds {
		searchSize *= int64(ub + 1)
		if searchSize > 100000000 {
			// Too large - use rounding heuristics
			return m.roundingHeuristic(lpSolution, matrix)
		}
	}

	bestResult := math.MaxInt32
	current := make([]int, numButtons)

	var search func(varIdx int, currentSum int)
	search = func(varIdx int, currentSum int) {
		if currentSum >= bestResult {
			return // Prune
		}

		if varIdx == numButtons {
			// Check if solution is valid
			valid := true
			for i := 0; i < numCounters; i++ {
				sum := 0
				for j := 0; j < numButtons; j++ {
					sum += matrix[i][j] * current[j]
				}
				if sum != m.joltageRequirement[i] {
					valid = false
					break
				}
			}
			if valid && currentSum < bestResult {
				bestResult = currentSum
			}
			return
		}

		// Try values for this variable
		for val := 0; val <= upperBounds[varIdx]; val++ {
			current[varIdx] = val
			search(varIdx+1, currentSum+val)
		}
	}

	search(0, 0)

	if bestResult == math.MaxInt32 {
		return -1
	}
	return bestResult
}

func (m *Machine) roundingHeuristic(lpSolution []float64, matrix [][]int) int {
	numButtons := len(m.buttonWiring)
	numCounters := len(m.joltageRequirement)

	bestResult := math.MaxInt32

	// Find fractional variables
	fractionalVars := make([]int, 0)
	for i, val := range lpSolution {
		if math.Abs(val-math.Round(val)) > 1e-6 {
			fractionalVars = append(fractionalVars, i)
		}
	}

	if len(fractionalVars) > 20 {
		// Too many fractional vars - just round up and return
		total := 0
		for _, val := range lpSolution {
			total += int(math.Ceil(val))
		}
		return total
	}

	// Try all combinations of floor/ceil for fractional vars
	numCombos := 1 << len(fractionalVars)
	for combo := 0; combo < numCombos; combo++ {
		solution := make([]int, numButtons)
		sum := 0

		// Start with floor for all
		for i, val := range lpSolution {
			solution[i] = int(math.Floor(val))
		}

		// Apply ceil based on combo bits
		for bit, varIdx := range fractionalVars {
			if (combo>>bit)&1 == 1 {
				solution[varIdx] = int(math.Ceil(lpSolution[varIdx]))
			}
		}

		for _, val := range solution {
			sum += val
		}

		if sum >= bestResult {
			continue
		}

		// Check validity
		valid := true
		for i := 0; i < numCounters; i++ {
			rowSum := 0
			for j := 0; j < numButtons; j++ {
				rowSum += matrix[i][j] * solution[j]
			}
			if rowSum != m.joltageRequirement[i] {
				valid = false
				break
			}
		}

		if valid && sum < bestResult {
			bestResult = sum
		}
	}

	if bestResult == math.MaxInt32 {
		// Fallback: round up everything
		total := 0
		for _, val := range lpSolution {
			total += int(math.Ceil(val))
		}
		return total
	}
	return bestResult
}

type Manual struct {
	machines []Machine
}

func (m *Manual) String() string {
	result := ""
	for i, machine := range m.machines {
		result += fmt.Sprintf("Machine %d: %s\n", i+1, machine.String())
	}
	return result
}

func NewManual() *Manual {
	return &Manual{
		machines: make([]Machine, 0),
	}
}

func (m *Manual) FewestPresses() int {
	results := 0
	for _, machine := range m.machines {
		results += machine.FewestPresses()
	}
	return results
}

func (m *Manual) FewestPressesVoltage() int {
	results := 0
	for _, machine := range m.machines {
		results += machine.FewestPressesVoltage()
	}
	return results
}

func (m *Manual) ParseLine(line string) {
	parts := strings.Split(line, " ")
	machine := NewMachine()
	for _, part := range parts {
		machine.ParsePart(part)
	}
	m.machines = append(m.machines, machine)
}
