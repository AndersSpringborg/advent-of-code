package main

import "strconv"

type Sequence struct {
	start int
	end   int
}

type SequenceAnalyzerResult struct {
	invalidIds []int
	sum        int
}

func EnummerateSequence(sequence *Sequence) []int {
	result := make([]int, 0)
	for i := sequence.start; i <= sequence.end; i++ {
		result = append(result, i)
	}

	return result
}

func NumberToString(number int) string {
	return strconv.Itoa((number))
}

func SumInvalidIds(invalidIds []int) int {
	total := 0
	for _, id := range invalidIds {
		total += id
	}
	return total
}
