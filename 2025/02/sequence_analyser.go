package main

type SequenceAnalyser struct {
}

func (sa *SequenceAnalyser) AnalyseSingle(sequence *Sequence) []int {
	return analyseInvalidIds(sequence)

}

func (sa *SequenceAnalyser) Analyse(sequence []Sequence) SequenceAnalyzerResult {
	allInvalidIds := make([]int, 0)
	for _, seq := range sequence {
		ids := analyseInvalidIds(&seq)
		allInvalidIds = append(allInvalidIds, ids...)
	}
	total := SumInvalidIds(allInvalidIds)
	return SequenceAnalyzerResult{sum: total, invalidIds: allInvalidIds}
}

func analyseInvalidIds(sequence *Sequence) []int {
	numbers := EnummerateSequence(sequence)
	invalidIds := make([]int, 0)
	for _, number := range numbers {
		numberAsString := NumberToString(number)
		firstHalf, secondHalf := SplitInHalves(numberAsString)
		if firstHalf == secondHalf {
			invalidIds = append(invalidIds, number)
		}
	}
	return invalidIds
}

func SplitInHalves(numberAsString string) (string, string) {
	mid := len(numberAsString) / 2
	return numberAsString[:mid], numberAsString[mid:]
}
