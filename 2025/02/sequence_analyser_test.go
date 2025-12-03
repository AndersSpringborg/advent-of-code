package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 11-22 has two invalid IDs: 11 and 22.
// 95-115 has one invalid ID: 99.
// 998-1012 has one invalid ID: 1010.
// 1188511880-1188511890 has one invalid ID: 1188511885.
// 222220-222224 has one invalid ID: 222222.
// 1698522-1698528 contains no invalid IDs.
// 446443-446449 has one invalid ID: 446446.
// 38593856-38593862 has one invalid ID: 38593859.

func TestAnalyseSimple(t *testing.T) {
	testCases := []struct {
		sequence    Sequence
		expectedSum int
		description string
	}{
		{Sequence{start: 11, end: 22}, 33, "11-22 should find 11 and 22"},
		{Sequence{start: 95, end: 115}, 99, "95-115 should find 99"},
		{Sequence{start: 998, end: 1012}, 1010, "998-1012 should find 1010"},
		{Sequence{start: 1188511880, end: 1188511890}, 1188511885, "should find 1188511885"},
		{Sequence{start: 222220, end: 222224}, 222222, "should find 222222"},
		{Sequence{start: 1698522, end: 1698528}, 0, "should find nothing"},
		{Sequence{start: 446443, end: 446449}, 446446, "should find 446446"},
		{Sequence{start: 38593856, end: 38593862}, 38593859, "should find 38593859"},
		{Sequence{start: 565653, end: 565659}, 0, "should find nothing"},
		{Sequence{start: 824824821, end: 824824827}, 0, "should find nothing"},
		{Sequence{start: 2121212118, end: 2121212124}, 0, "should find nothing"},
	}

	total := 0
	for _, tc := range testCases {
		result := AnalyseSimple([]Sequence{tc.sequence})
		total += result
		assert.Equal(t, tc.expectedSum, result, tc.description)
	}
	assert.Equal(t, 1227775554, total)
}
