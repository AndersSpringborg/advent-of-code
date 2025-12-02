package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//11-22 still has two invalid IDs, 11 and 22.
//95-115 now has two invalid IDs, 99 and 111.
//998-1012 now has two invalid IDs, 999 and 1010.
//1188511880-1188511890 still has one invalid ID, 1188511885.
//222220-222224 still has one invalid ID, 222222.
//1698522-1698528 still contains no invalid IDs.
//446443-446449 still has one invalid ID, 446446.
//38593856-38593862 still has one invalid ID, 38593859.
//565653-565659 now has one invalid ID, 565656.
//824824821-824824827 now has one invalid ID, 824824824.
//2121212118-2121212124 now has one invalid ID, 2121212121.

func TestAnalyseComplex(t *testing.T) {
	testCases := []struct {
		sequence Sequence
		expected []int
	}{
		{Sequence{start: 11, end: 22}, []int{11, 22}},
		{Sequence{start: 95, end: 115}, []int{99, 111}},
		{Sequence{start: 998, end: 1012}, []int{999, 1010}},
		{Sequence{start: 1188511880, end: 1188511890}, []int{1188511885}},
		{Sequence{start: 222220, end: 222224}, []int{222222}},
		{Sequence{start: 1698522, end: 1698528}, []int{}},
		{Sequence{start: 446443, end: 446449}, []int{446446}},
		{Sequence{start: 38593856, end: 38593862}, []int{38593859}},
		{Sequence{start: 565653, end: 565659}, []int{565656}},
		{Sequence{start: 824824821, end: 824824827}, []int{824824824}},
		{Sequence{start: 2121212118, end: 2121212124}, []int{2121212121}},
	}

	total := 0
	for _, tc := range testCases {

		result := AnalyseSingleComplex(&tc.sequence)
		for _, n := range result {
			total += n
		}
		assert.Equal(t, tc.expected, result, "Failed for sequence %+v", tc.sequence)
	}
	assert.Equal(t, 4174379265, total)
}

func TestSplitSizes(t *testing.T) {
	testCases := []struct {
		number             string
		expectedSplitSizes []int
	}{
		{number: "12", expectedSplitSizes: []int{2}},
		{number: "123", expectedSplitSizes: []int{3}},
		{number: "1234", expectedSplitSizes: []int{2, 4}},
		{number: "12345", expectedSplitSizes: []int{5}},
		{number: "123456", expectedSplitSizes: []int{2, 3, 6}},
	}

	for _, tc := range testCases {
		result := SplitSizes(tc.number)
		assert.Equal(t, tc.expectedSplitSizes, result, "Failed for sequence %+v", tc.number)
	}
}

func TestSplitIntoParts(t *testing.T) {
	testCases := []struct {
		number             string
		size               int
		expectedSplitSizes []string
	}{
		{number: "12", size: 2, expectedSplitSizes: []string{"1", "2"}},
		{number: "123", size: 3, expectedSplitSizes: []string{"1", "2", "3"}},
		{number: "112233", size: 3, expectedSplitSizes: []string{"11", "22", "33"}},
		{number: "112233", size: 3, expectedSplitSizes: []string{"11", "22", "33"}},
	}

	for _, tc := range testCases {
		result := SplitStringIntoParsts(tc.number, tc.size)
		assert.Equal(t, tc.expectedSplitSizes, result, "Failed for sequence %+v", tc.number)
	}
}
