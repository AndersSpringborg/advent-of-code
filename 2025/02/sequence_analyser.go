package main

import "strconv"

func AnalyseSimple(sequences []Sequence) int {
	sum := 0
	for _, seq := range sequences {
		for num := seq.start; num <= seq.end; num++ {
			s := strconv.Itoa(num)
			length := len(s)

			// Only even-length numbers can split into equal halves
			if length%2 == 0 {
				mid := length / 2
				if s[:mid] == s[mid:] {
					sum += num
				}
			}
		}
	}
	return sum
}
