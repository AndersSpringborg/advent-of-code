package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSequenceAnalyser1122(t *testing.T) {
	s := Sequence{start: 11, end: 22}
	analyser := SequenceAnalyser{}

	result := analyser.AnalyseSingle(&s)

	expected := []int{11, 22}
	assert.Equal(t, expected, result)
}

// 95-115 has one invalid ID, 99.
//998-1012 has one invalid ID, 1010.
//1188511880-1188511890 has one invalid ID, 1188511885.
//222220-222224 has one invalid ID, 222222.
//1698522-1698528 contains no invalid IDs.
//446443-446449 has one invalid ID, 446446.
//38593856-38593862 has one invalid ID, 38593859.

func TestSequenceAnalyserVariousRanges(t *testing.T) {
	testCases := []struct {
		sequence Sequence
		expected []int
	}{
		{Sequence{start: 11, end: 22}, []int{11, 22}},
		{Sequence{start: 95, end: 115}, []int{99}},
		{Sequence{start: 998, end: 1012}, []int{1010}},
		{Sequence{start: 1188511880, end: 1188511890}, []int{1188511885}},
		{Sequence{start: 222220, end: 222224}, []int{222222}},
		{Sequence{start: 1698522, end: 1698528}, []int{}},
		{Sequence{start: 446443, end: 446449}, []int{446446}},
		{Sequence{start: 38593856, end: 38593862}, []int{38593859}},
		{Sequence{start: 565653, end: 565659}, []int{}},
		{Sequence{start: 824824821, end: 824824827}, []int{}},
		{Sequence{start: 2121212118, end: 2121212124}, []int{}},
	}

	analyser := SequenceAnalyser{}
	total := 0

	for _, tc := range testCases {
		result := analyser.AnalyseSingle(&tc.sequence)
		for _, n := range result {
			total += n
		}
		assert.Equal(t, tc.expected, result, "Failed for sequence %+v", tc.sequence)
	}
	assert.Equal(t, 1227775554, total)
}
