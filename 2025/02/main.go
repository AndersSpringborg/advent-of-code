package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	analyser := SequenceAnalyser{}
	scanner, closer := Scan("input.txt")
	defer closer.Close()

	scanner.Scan()
	line := scanner.Text()

	ranges := strings.Split(line, ",")
	sequences := ParseSequences(ranges)

	res := analyser.Analyse(sequences)

	// task 2
	scanner2, closer2 := Scan("input.txt")
	defer closer2.Close()

	scanner2.Scan()
	line2 := scanner2.Text()
	ranges2 := strings.Split(line2, ",")
	sequences2 := ParseSequences(ranges2)

	res2 := AnalyseComplex(sequences2)

	fmt.Printf("Task 1 - Sum of invalid IDs: %d\n", res.sum)
	fmt.Printf("Task 2 - Sum of invalid IDs (complex analysis): %d\n", res2.sum)
}

func ParseSequences(ranges []string) []Sequence {
	sequences := make([]Sequence, 0)
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, errStart := strconv.Atoi(parts[0])
		end, errEnd := strconv.Atoi(parts[1])
		if errStart != nil || errEnd != nil {
			log.Fatal("Failed to parse range:", r)
		}
		sequence := Sequence{start: start, end: end}
		sequences = append(sequences, sequence)
	}
	return sequences
}
