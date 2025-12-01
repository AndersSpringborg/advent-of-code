package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDialInit(t *testing.T) {
	dial := NewDial(5)

	assert.Equal(t, uint(5), dial.Index)
}

func TestDialRotateRight(t *testing.T) {
	dial := NewDial(0)
	dial.Rotate(Right, 3)

	assert.Equal(t, uint(3), dial.Index)
}

func TestDialRotateLeft(t *testing.T) {
	dial := NewDial(5)
	dial.Rotate(Left, 2)

	assert.Equal(t, uint(3), dial.Index)
}

func TestDialRotateRightWrapAround(t *testing.T) {
	dial := NewDial(98)
	dial.Rotate(Right, 5)
	// 99, 0, 1, 2, 3
	assert.Equal(t, uint(3), dial.Index)
}

func TestDialRotateLeftWrapAround(t *testing.T) {
	dial := NewDial(2)
	dial.Rotate(Left, 5)
	// 1, 0, 99, 98, 97
	assert.Equal(t, uint(97), dial.Index)
}

func TestDial_Counts_ZeroIndex(t *testing.T) {
	dial := NewDial(1)
	dial.Rotate(Left, 1)

	assert.Equal(t, uint(1), dial.ZeroIndex)
}
