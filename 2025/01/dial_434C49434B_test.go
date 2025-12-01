package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDialMInit(t *testing.T) {
	dial := NewDial0x434C49434B(5)

	assert.Equal(t, 5, dial.Index)
}

func TestDialMRotateRight(t *testing.T) {
	dial := NewDial0x434C49434B(0)
	dial.Rotate(Right, 3)

	assert.Equal(t, 3, dial.Index)
}

func TestDialMRotateLeft(t *testing.T) {
	dial := NewDial0x434C49434B(5)
	dial.Rotate(Left, 2)

	assert.Equal(t, 3, dial.Index)
}

func TestDialMRotateRightWrapAround(t *testing.T) {
	dial := NewDial0x434C49434B(98)
	dial.Rotate(Right, 5)
	// 99, 0, 1, 2, 3
	assert.Equal(t, 3, dial.Index)
}

func TestDial0x434C49434B_WrapAroundTwiceRight(t *testing.T) {
	dial := NewDial0x434C49434B(1)
	dial.Rotate(Right, 200)

	assert.Equal(t, 1, dial.Index)
	assert.Equal(t, 2, dial.ZeroIndex)

}

func TestDial0x434C49434B_WrapAroundTwiceLeft(t *testing.T) {
	dial := NewDial0x434C49434B(1)
	dial.Rotate(Left, 200)

	assert.Equal(t, 1, dial.Index)
	assert.Equal(t, 2, dial.ZeroIndex)
}

func TestFromExample(t *testing.T) {
	dial := NewDial0x434C49434B(50)
	dial.Rotate(Left, 68) // 82
	assert.Equal(t, 1, dial.ZeroIndex)
	dial.Rotate(Left, 30)  // 52
	dial.Rotate(Right, 48) // 0
	log.Default().Print("expect 2: ", dial.ZeroIndex)
	assert.Equal(t, 2, dial.ZeroIndex)
	dial.Rotate(Left, 5)   // 95
	dial.Rotate(Right, 60) // 55
	log.Default().Print("expect 3: ", dial.ZeroIndex)
	assert.Equal(t, 3, dial.ZeroIndex)
	dial.Rotate(Left, 55) // 0
	log.Default().Print("expect 4: ", dial.ZeroIndex)
	assert.Equal(t, 4, dial.ZeroIndex)
	dial.Rotate(Left, 1)  // 99
	dial.Rotate(Left, 99) // 0
	log.Default().Print("expect 5: ", dial.ZeroIndex)
	assert.Equal(t, 5, dial.ZeroIndex)
	dial.Rotate(Right, 14) // 14
	dial.Rotate(Left, 82)  // 32
	log.Default().Print("expect 6: ", dial.ZeroIndex)
	assert.Equal(t, 6, dial.ZeroIndex)

	assert.Equal(t, 6, dial.ZeroIndex)
}

func Test55to0Left(t *testing.T) {
	dial := NewDial0x434C49434B(55)
	dial.Rotate(Left, 55)

	assert.Equal(t, 1, dial.ZeroIndex)
}

func Test45to0Right(t *testing.T) {
	dial := NewDial0x434C49434B(45)
	dial.Rotate(Right, 55)

	assert.Equal(t, 1, dial.ZeroIndex)
}

func TestWrappesRightAndHitsZero(t *testing.T) {
	dial := NewDial0x434C49434B(50)
	dial.Rotate(Right, 150)

	assert.Equal(t, 2, dial.ZeroIndex)
}

func TestWrappesLeftAndHitsZero(t *testing.T) {
	dial := NewDial0x434C49434B(50)
	dial.Rotate(Left, 150)

	assert.Equal(t, 2, dial.ZeroIndex)

	dial.Rotate(Right, 10)
	assert.Equal(t, 2, dial.ZeroIndex)
	assert.Equal(t, 10, dial.Index)
}

// 1. Starting at 0
func TestStartingAtZero(t *testing.T) {
	dial := NewDial0x434C49434B(0)
	dial.Rotate(Right, 5)

	// Should NOT count the starting position as a "pass through"
	assert.Equal(t, 0, dial.ZeroIndex)
	assert.Equal(t, 5, dial.Index)
}

// 2. Ending at 0 (but not passing through during rotation)
func TestEndingAtZeroNoPassThrough(t *testing.T) {
	dial := NewDial0x434C49434B(5)
	dial.Rotate(Left, 5)

	// Should count once (ending at 0)
	assert.Equal(t, 1, dial.ZeroIndex)
	assert.Equal(t, 0, dial.Index)
}

// 3. Starting at 0 and ending at 0
func TestStartAndEndAtZero(t *testing.T) {
	dial := NewDial0x434C49434B(0)
	dial.Rotate(Right, 100)

	// Should count once (the ending position)
	assert.Equal(t, 1, dial.ZeroIndex)
	assert.Equal(t, 0, dial.Index)
}

// 4. Rotation of 0 clicks
func TestZeroRotation(t *testing.T) {
	dial := NewDial0x434C49434B(50)
	dial.Rotate(Right, 0)

	assert.Equal(t, 0, dial.ZeroIndex)
	assert.Equal(t, 50, dial.Index)
}

// 5. Starting at 0 with zero rotation
func TestStartAtZeroNoRotation(t *testing.T) {
	dial := NewDial0x434C49434B(0)
	dial.Rotate(Right, 0)

	assert.Equal(t, 0, dial.ZeroIndex)
	assert.Equal(t, 0, dial.Index)
}

// 6. Exact multiple of 100 (should pass through 0 multiple times)
func TestExactMultipleOf100(t *testing.T) {
	dial := NewDial0x434C49434B(0)
	dial.Rotate(Right, 300)

	// Should count 3 times (at 100, 200, 300 which wraps to 0)
	assert.Equal(t, 3, dial.ZeroIndex)
	assert.Equal(t, 0, dial.Index)
}

// 7. Almost a full rotation (99 clicks)
func TestAlmostFullRotation(t *testing.T) {
	dial := NewDial0x434C49434B(1)
	dial.Rotate(Right, 99)

	// Should pass through 0 once
	assert.Equal(t, 1, dial.ZeroIndex)
	assert.Equal(t, 0, dial.Index)
}

// 8. Single click that crosses 0
func TestSingleClickCrossZeroRight(t *testing.T) {
	dial := NewDial0x434C49434B(99)
	dial.Rotate(Right, 1)

	assert.Equal(t, 1, dial.ZeroIndex)
	assert.Equal(t, 0, dial.Index)
}

func TestSingleClickCrossZeroLeft(t *testing.T) {
	dial := NewDial0x434C49434B(0)
	dial.Rotate(Left, 1)

	// Going left from 0 wraps to 99, does NOT pass through 0
	assert.Equal(t, 0, dial.ZeroIndex)
	assert.Equal(t, 99, dial.Index)
}

// Test the exact example step by step
func TestExampleStepByStep(t *testing.T) {
	dial := NewDial0x434C49434B(50)

	dial.Rotate(Left, 68) // 50→82, passes through 0
	assert.Equal(t, 82, dial.Index)
	assert.Equal(t, 1, dial.ZeroIndex, "L68 from 50 should pass through 0 once")

	dial.Rotate(Left, 30) // 82→52, no pass
	assert.Equal(t, 52, dial.Index)
	assert.Equal(t, 1, dial.ZeroIndex, "L30 from 82 should not pass through 0")

	dial.Rotate(Right, 48) // 52→0, ends at 0
	assert.Equal(t, 0, dial.Index)
	assert.Equal(t, 2, dial.ZeroIndex, "R48 from 52 should end at 0")

	dial.Rotate(Left, 5) // 0→95, no pass (going backwards from 0)
	assert.Equal(t, 95, dial.Index)
	assert.Equal(t, 2, dial.ZeroIndex, "L5 from 0 should not pass through 0")

	dial.Rotate(Right, 60) // 95→55, passes through 0
	assert.Equal(t, 55, dial.Index)
	assert.Equal(t, 3, dial.ZeroIndex, "R60 from 95 should pass through 0 once")

	dial.Rotate(Left, 55) // 55→0, ends at 0
	assert.Equal(t, 0, dial.Index)
	assert.Equal(t, 4, dial.ZeroIndex, "L55 from 55 should end at 0")

	dial.Rotate(Left, 1) // 0→99, no pass
	assert.Equal(t, 99, dial.Index)
	assert.Equal(t, 4, dial.ZeroIndex, "L1 from 0 should not pass through 0")

	dial.Rotate(Left, 99) // 99→0, ends at 0
	assert.Equal(t, 0, dial.Index)
	assert.Equal(t, 5, dial.ZeroIndex, "L99 from 99 should end at 0")

	dial.Rotate(Right, 14) // 0→14, no pass
	assert.Equal(t, 14, dial.Index)
	assert.Equal(t, 5, dial.ZeroIndex, "R14 from 0 should not pass through 0")

	dial.Rotate(Left, 82) // 14→32, passes through 0
	assert.Equal(t, 32, dial.Index)
	assert.Equal(t, 6, dial.ZeroIndex, "L82 from 14 should pass through 0 once")
}

// Key edge case: moving away from 0 should NOT count
func TestMovingAwayFromZero(t *testing.T) {
	dial := NewDial0x434C49434B(0)
	dial.Rotate(Right, 5)
	assert.Equal(t, 0, dial.ZeroIndex, "Moving away from 0 should not count 0")

	dial = NewDial0x434C49434B(0)
	dial.Rotate(Left, 5)
	assert.Equal(t, 0, dial.ZeroIndex, "Moving away from 0 should not count 0")
}

// Key edge case: ending at 0 counts
func TestEndingAtZeroCounts(t *testing.T) {
	dial := NewDial0x434C49434B(10)
	dial.Rotate(Left, 10)
	assert.Equal(t, 1, dial.ZeroIndex, "Ending at 0 should count")
	assert.Equal(t, 0, dial.Index)
}

// Key edge case: passing through 0 during rotation
func TestPassingThroughZero(t *testing.T) {
	dial := NewDial0x434C49434B(95)
	dial.Rotate(Right, 10) // 95→96→97→98→99→0→1→2→3→4→5
	assert.Equal(t, 1, dial.ZeroIndex, "Should count 0 when passing through")
	assert.Equal(t, 5, dial.Index)
}

// Multiple complete rotations
func TestMultipleCompleteRotations(t *testing.T) {
	dial := NewDial0x434C49434B(50)
	dial.Rotate(Right, 250) // 2.5 rotations
	// Should pass 0 at: 100, 200, and end at 300%100=0
	assert.Equal(t, 3, dial.ZeroIndex)
	assert.Equal(t, 0, dial.Index)
}

// Rotation that doesn't cross 0
func TestNoZeroCrossing(t *testing.T) {
	dial := NewDial0x434C49434B(20)
	dial.Rotate(Right, 30)
	assert.Equal(t, 0, dial.ZeroIndex, "Should not count 0 if not crossed")
	assert.Equal(t, 50, dial.Index)
}
