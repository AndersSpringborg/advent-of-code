package internal

import (
	"strings"
	"testing"
)

func TestMachine_FewestPresses(t *testing.T) {
	// setup multiple test cases with multiple expect
	tests := []struct {
		input         string
		fewestPresses int
	}{
		{"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}", 2},
		{"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}", 3},
		{"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}", 2},
	}

	for _, test := range tests {
		machine := NewMachine()
		for _, part := range strings.Split(test.input, " ") {
			machine.ParsePart(part)
		}
		presses := machine.FewestPresses()
		if presses != test.fewestPresses {
			t.Errorf("For input %s, expected fewest presses %d, got %d", test.input, test.fewestPresses, presses)
		}
	}
}

func TestMachine_FewestPressesVoltage(t *testing.T) {
	// setup multiple test cases with multiple expect
	tests := []struct {
		input         string
		fewestPresses int
	}{
		{"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}", 10},
		{"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}", 12},
		{"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}", 11},
	}

	for _, test := range tests {
		machine := NewMachine()
		for _, part := range strings.Split(test.input, " ") {
			machine.ParsePart(part)
		}
		presses := machine.FewestPressesVoltage()
		if presses != test.fewestPresses {
			t.Errorf("For input %s, expected fewest presses %d, got %d", test.input, test.fewestPresses, presses)
		}
	}
}
