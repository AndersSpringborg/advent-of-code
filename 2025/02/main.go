package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	sequences := getSequences("input.txt")
	sequences2 := getSequences("input2.txt")

	res1 := AnalyseSimple(sequences)
	res2 := AnalyseComplex(sequences2)

	fmt.Printf("Task 1 - Sum of invalid IDs: %d\n", res1)
	fmt.Printf("Task 2 - Sum of invalid IDs (complex analysis): %d\n", res2)
}

func getSequences(inputFile string) []Sequence {
	scanner, closer := Scan(inputFile)
	defer closer.Close()

	scanner.Scan()
	line := scanner.Text()

	ranges := strings.Split(line, ",")
	sequences := ParseSequences(ranges)
	return sequences
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
