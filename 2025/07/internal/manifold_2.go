package internal

import "fmt"

type Manifold2 struct {
	linesProcessed []byte
	lineLength     int
}

func NewManifold2() *Manifold2 {
	return &Manifold2{
		linesProcessed: []byte{},
		lineLength:     0,
	}
}

func (m *Manifold2) ParseLine(line string) {
	if m.lineLength == 0 {
		m.lineLength = len(line)
	}
	m.linesProcessed = append(m.linesProcessed, []byte(line)...)
}

func (m *Manifold2) PrintInput() {
	if m.lineLength == 0 {
		fmt.Println("No lines processed.")
		return
	}

	totalLines := len(m.linesProcessed) / m.lineLength
	for row := 0; row < totalLines; row++ {
		startIndex := row * m.lineLength
		endIndex := startIndex + m.lineLength
		line := m.linesProcessed[startIndex:endIndex]
		fmt.Println(string(line))
	}
}

func (m *Manifold2) GetTimelinesDP() int {
	if m.lineLength == 0 {
		return 1
	}

	width := m.lineLength
	height := len(m.linesProcessed) / width

	// Create DP table
	dp := make([]int, len(m.linesProcessed))

	// Base case: bottom row has 1 timeline
	lastRow := height - 1
	for x := 0; x < width; x++ {
		dp[lastRow*width+x] = 1
	}

	// Work backwards from second-to-last row
	for row := height - 2; row >= 0; row-- {
		for x := 0; x < width; x++ {
			idx := row*width + x
			char := m.linesProcessed[idx]

			if char == '^' {
				if x > 0 {
					timelinesOnBelowLeft := dp[(row+1)*width+(x-1)]
					dp[idx] += timelinesOnBelowLeft
				}
				if x < width-1 {
					timelinesOnBelowRight := dp[(row+1)*width+(x+1)]
					dp[idx] += timelinesOnBelowRight
				}
			} else if char == '.' || char == 'S' {
				// Continue amount of timelines from below
				timelinesOnBelow := dp[(row+1)*width+x]
				dp[idx] = timelinesOnBelow
			}
		}
	}

	// Print timelines (debug output)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := y*width + x
			fmt.Printf("%2d ", dp[idx])
		}
		fmt.Println()
	}

	// Find S and return its value
	for x := 0; x < width; x++ {
		if m.linesProcessed[x] == 'S' {
			return dp[x]
		}
	}
	return 0
}
