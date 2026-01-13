package internal

import (
	"testing"
)

func TestGiftPacker_ParseText(t *testing.T) {
	text := `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`
	gp := NewGiftPacker()

	gp.ParseText(text)

	if len(gp.gifts) != 6 {
		t.Errorf("Expected 6 gifts, got %d", len(gp.gifts))
	}
	giftZero := gp.gifts[0]
	if !giftZero.Filled(0, 0) || giftZero.Filled(2, 2) {
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				if giftZero.Filled(x, y) {
					print("#")
				} else {
					print(".")
				}
			}
			println("\n")
		}
		t.Errorf("Gift 0 structure not parsed correctly")
	}

	if len(gp.regions) != 3 {
		t.Errorf("Expected 2 regions, got %d", len(gp.regions))
	}
}

// TestGiftPacker_RegionsFitPresents (305.41s)  -- VERY SLOW, needs cache
func TestGiftPacker_RegionsFitPresents(t *testing.T) {
	text := `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2
`
	gp := NewGiftPacker()

	gp.ParseText(text)

	fitPresents := gp.FitPresentsInRegions()
	if fitPresents != 2 {
		t.Errorf("Expected 2 regions that can fit presents, got %d", fitPresents)
	}

}
