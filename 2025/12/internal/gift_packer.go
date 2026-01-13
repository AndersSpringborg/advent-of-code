package internal

import (
	"fmt"
	"strings"
)

type Gift struct {
	structure [9]bool
}

func (g *Gift) area() int {
	area := 0
	for _, filled := range g.structure {
		if filled {
			area++
		}
	}
	return area
}

type Region struct {
	width            int
	height           int
	giftDistribution [6]int
}

type GiftPacker struct {
	gifts   [6]Gift
	regions []Region
}

func NewGiftPacker() *GiftPacker {
	return &GiftPacker{
		gifts: [6]Gift{},
	}
}

func (gp *GiftPacker) ParseText(text string) {
	parts := strings.Split(text, "\n\n")

	gifts := parts[0:6]
	for i, gift := range gifts {
		gp.parseGift(i, gift)
	}

	regions := parts[6]
	for _, region := range strings.Split(regions, "\n") {
		gp.parseRegion(region)
	}
}

func (gp *GiftPacker) parseRegion(region string) {
	if region == "" {
		return
	}
	var width, height int
	var giftDistribution [6]int
	fmt.Sscanf(region, "%dx%d: %d %d %d %d %d %d",
		&width, &height,
		&giftDistribution[0],
		&giftDistribution[1],
		&giftDistribution[2],
		&giftDistribution[3],
		&giftDistribution[4],
		&giftDistribution[5])

	gp.regions = append(gp.regions, Region{
		width:            width,
		height:           height,
		giftDistribution: giftDistribution,
	})
}

func (gp *GiftPacker) parseGift(giftId int, gift string) {
	lines := strings.Split(gift, "\n")
	var structure [9]bool
	for i, line := range lines[1:4] { // skip line with ID
		for j, char := range line {
			if char == '#' {
				structure[i*3+j] = true
			} else {
				structure[i*3+j] = false
			}
		}
	}
	gp.gifts[giftId] = Gift{structure: structure}
}

func (gp *GiftPacker) FitPresentsInRegions() int {
	totalFits := 0
	fmt.Println("regions to check:", len(gp.regions))
	fmt.Printf("regions details: %+v\n", gp.regions)
	for _, region := range gp.regions {
		fits := gp.fitsInRegion(region)
		if fits {
			totalFits++
		}
	}
	return totalFits
}

func (gp *GiftPacker) fitsInRegion(region Region) bool {
	// Quick check: not enough total area
	neededArea := 0
	for i, present := range gp.gifts {
		presentAreas := region.giftDistribution[i] * present.area()
		neededArea += presentAreas
	}

	if neededArea > (region.width * region.height) {
		return false
	}

	// part 2 hvor vi tjekker rotation, flip og placering med backtracking
	allGiftsInRegion := []Gift{}
	for i, count := range region.giftDistribution {
		for j := 0; j < count; j++ {
			allGiftsInRegion = append(allGiftsInRegion, gp.gifts[i])
		}
	}

	area := NewArea(region.width, region.height)
	return area.backtracking(allGiftsInRegion, 0)
}

func (a *Area) backtracking(gifts []Gift, index int) bool {
	if index >= len(gifts) {
		return true
	}

	height := len(a.area)
	width := len(a.area[0])

	gift := gifts[index]
	transformations := generateTransformations(gift)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for _, transformed := range transformations {
				if a.CanPlaceGiftAt(transformed, x, y) {
					a.PlaceGiftAt(transformed, x, y)

					if a.backtracking(gifts, index+1) {
						return true
					}

					a.RemoveGiftAt(transformed, x, y)
				}
			}
		}
	}

	return false
}

func rotateGift(gift Gift) Gift {
	var rotated Gift
	rotated.structure[0] = gift.structure[6]
	rotated.structure[1] = gift.structure[3]
	rotated.structure[2] = gift.structure[0]
	rotated.structure[3] = gift.structure[7]
	rotated.structure[4] = gift.structure[4]
	rotated.structure[5] = gift.structure[1]
	rotated.structure[6] = gift.structure[8]
	rotated.structure[7] = gift.structure[5]
	rotated.structure[8] = gift.structure[2]
	return rotated
}

func flipGift(gift Gift) Gift {
	var flipped Gift
	flipped.structure[0] = gift.structure[2]
	flipped.structure[1] = gift.structure[1]
	flipped.structure[2] = gift.structure[0]
	flipped.structure[3] = gift.structure[5]
	flipped.structure[4] = gift.structure[4]
	flipped.structure[5] = gift.structure[3]
	flipped.structure[6] = gift.structure[8]
	flipped.structure[7] = gift.structure[7]
	flipped.structure[8] = gift.structure[6]
	return flipped
}
func generateTransformations(gift Gift) []Gift {
	transformations := []Gift{}

	// Original
	transformations = append(transformations, gift)

	// Rotations
	rotated90 := rotateGift(gift)
	transformations = append(transformations, rotated90)
	rotated180 := rotateGift(rotated90)
	transformations = append(transformations, rotated180)
	rotated270 := rotateGift(rotated180)
	transformations = append(transformations, rotated270)

	// Flips
	flipped := flipGift(gift)
	transformations = append(transformations, flipped)
	flippedRotated90 := rotateGift(flipped)
	transformations = append(transformations, flippedRotated90)
	flippedRotated180 := rotateGift(flippedRotated90)
	transformations = append(transformations, flippedRotated180)
	flippedRotated270 := rotateGift(flippedRotated180)
	transformations = append(transformations, flippedRotated270)

	return transformations
}

func (g *Gift) Filled(y, x int) bool {
	return g.structure[y*3+x]
}

type Area struct {
	area [][]bool
}

func NewArea(width, height int) *Area {
	var area [][]bool = make([][]bool, height)
	for i := range area {
		area[i] = make([]bool, width)
	}
	return &Area{area: area}
}

func (a *Area) PlaceGiftAt(gift Gift, x, y int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gift.Filled(i, j) {
				a.area[y+i][x+j] = true
			}
		}
	}
}

func (a *Area) RemoveGiftAt(gift Gift, x, y int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gift.Filled(i, j) {
				a.area[y+i][x+j] = false
			}
		}
	}
}

func (a *Area) CanPlaceGiftAt(gift Gift, x, y int) bool {
	height := len(a.area)
	width := len(a.area[0])

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gift.Filled(i, j) {
				areaX := x + j
				areaY := y + i
				if areaX < 0 || areaX >= width || areaY < 0 || areaY >= height {
					return false
				}
				if a.area[areaY][areaX] {
					return false
				}
			}
		}
	}
	return true
}
