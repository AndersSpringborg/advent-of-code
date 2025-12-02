package main

import (
	"slices"
)

type Set struct {
	numbers []int
}

func (s *Set) addNumber(n int) {
	if !slices.Contains(s.numbers, n) {
		s.numbers = append(s.numbers, n)
	}
}

func AnalyseSingleComplex(sequence *Sequence) []int {
	// split into to, three, four, until half of length of the number
	// Check if the number can be split into 3, 4, 5 equal parts
	// if all splits are equally big

	// Hvis all parts er ens, så er det et invalid  ID

	invalidIds := Set{numbers: make([]int, 0)}

	numbers := EnummerateSequence(sequence)
	for _, number := range numbers {
		numbersAsString := NumberToString(number)
		splitSizes := SplitSizes(numbersAsString)

		stringLen := len(numbersAsString)

		for _, size := range splitSizes {
			substringSize := stringLen / size

			pattern := numbersAsString[0:substringSize]
			allSame := true
			parts := SplitStringIntoParsts(numbersAsString, size)
			for _, part := range parts {
				if part != pattern {
					allSame = false
				}
			}

			if allSame {
				// so har vi et invalid ID. Tilføj, og skip resten af patterns
				invalidIds.addNumber(number)
				continue
			}

		}

	}

	return invalidIds.numbers

}

func SplitStringIntoParsts(s string, size int) []string {
	stringLen := len(s)
	substringSize := stringLen / size
	parts := make([]string, 0)
	for i := range size {
		start := i * substringSize
		end := start + substringSize
		part := s[start:end]
		parts = append(parts, part)

	}
	return parts

}

func SplitSizes(number string) []int {
	length := len(number)
	sizes := make([]int, 0)
	for divider := 2; divider <= length; divider++ {
		if length%divider == 0 {
			sizes = append(sizes, divider)
		}
	}
	return sizes
}

func AnalyseComplex(sequence []Sequence) SequenceAnalyzerResult {
	allInvalidIds := make([]int, 0)
	for _, seq := range sequence {
		ids := AnalyseSingleComplex(&seq)
		allInvalidIds = append(allInvalidIds, ids...)
	}
	total := SumInvalidIds(allInvalidIds)
	return SequenceAnalyzerResult{sum: total, invalidIds: allInvalidIds}
}
