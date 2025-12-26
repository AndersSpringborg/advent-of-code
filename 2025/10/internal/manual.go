package internal

import (
	"fmt"
	"math"
	"strings"
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
	// laver light diagram om til en byte
	index := 0
	for i, light := range lights {
		if light == '#' {
			index += 1 << i
		}
	}
	// .##. = 0110 = 6
	return index
}

func (m *Machine) FewestPresses() int {
	amountOfLights := len(m.lightDiagram)
	lightMutations := int(math.Pow(2, float64(amountOfLights)))
	buttons := len(m.buttonWiring)

	// Target state: what we're trying to reach
	targetState := lightIndex(m.lightDiagram)

	// Initial state: all lights off [....]

	dpTable := make([][]int, buttons+1)
	for i := range dpTable {
		dpTable[i] = make([]int, lightMutations)
		for j := range dpTable[i] {
			dpTable[i][j] = math.MaxInt32 // infinity
		}
	}
	initialState := 0
	dpTable[0][initialState] = 0 // Base case: 0 presses to stay at initial state

	for i := 1; i <= buttons; i++ {
		for currentState := 0; currentState < lightMutations; currentState++ {
			if dpTable[i-1][currentState] == math.MaxInt32 {
				continue // Skip unreachable states
			}

			// Option 1: Don't press button i
			dpTable[i][currentState] = min(dpTable[i][currentState], dpTable[i-1][currentState])

			// Option 2: Press button i one or more times
			// Pressing the same button twice returns to the original state
			// So we only need to consider pressing it 0 or 1 times
			newState := currentState
			for _, lightIdx := range m.buttonWiring[i-1] {
				newState ^= (1 << lightIdx) // Toggle the bit
			}
			dpTable[i][newState] = min(dpTable[i][newState], dpTable[i-1][currentState]+1)
		}
	}

	result := dpTable[buttons][targetState]
	if result == math.MaxInt32 {
		return -1 // Target state unreachable
	}
	return result
}
func (m *Machine) String() string {
	return fmt.Sprintf("LightDiagram: %s, ButtonWiring: %v, JoltageRequirement: %v",
		string(m.lightDiagram), m.buttonWiring, m.joltageRequirement)
}

func (m *Machine) FewestPressesVoltage() int {
	return 0
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

func (m *Manual) ParseLine(line string) {
	parts := strings.Split(line, " ")
	machine := NewMachine()
	for _, parts := range parts {
		machine.ParsePart(parts)
	}
	m.machines = append(m.machines, machine)

}
