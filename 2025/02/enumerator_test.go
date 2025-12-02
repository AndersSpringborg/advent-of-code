package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnummerateSequence(t *testing.T) {
	s := Sequence{start: 1, end: 5}

	result := EnummerateSequence(&s)

	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, result)
}

func TestEnummerateSequenceBigNumbers(t *testing.T) {
	s := Sequence{start: 1188511880, end: 1188511890}

	result := EnummerateSequence(&s)

	expected := []int{
		1188511880,
		1188511881,
		1188511882,
		1188511883,
		1188511884,
		1188511885,
		1188511886,
		1188511887,
		1188511888,
		1188511889,
		1188511890}
	assert.Equal(t, expected, result)
}

func TestEnummerateSequence_Empty(t *testing.T) {
	s := Sequence{start: 5, end: 5}

	result := EnummerateSequence(&s)

	expected := []int{5}
	assert.Equal(t, expected, result)
}

func TestNumberToString(t *testing.T) {
	result := NumberToString(1234)

	expected := "1234"
	assert.Equal(t, expected, result)
}
