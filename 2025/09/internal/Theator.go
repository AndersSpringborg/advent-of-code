package internal

import (
	"fmt"
	"log"
	"math"
	"sort"
)

type Theater struct {
	redTiles    [][2]int
	insideCache map[[2]int]bool // Stores the result of isEligibleTile
}

type Square struct {
	tileAX int
	tileAY int
	tileBX int
	tileBY int
	area   int
}

type Tile struct {
	x int
	y int
}

type Edge struct {
	start Tile
	end   Tile
}

func (t *Theater) isInsideLoop(point [2]int) bool {
	x := float64(point[0]) + 0.5
	y := float64(point[1]) + 0.5
	inside := false

	for i := 0; i < len(t.redTiles); i++ {
		j := (i + 1) % len(t.redTiles)
		xi := float64(t.redTiles[i][0])
		yi := float64(t.redTiles[i][1])
		xj := float64(t.redTiles[j][0])
		yj := float64(t.redTiles[j][1])

		// Check for intersection with horizontal ray from (x, y)
		if ((yi > y) != (yj > y)) &&
			(x < (xj-xi)*(y-yi)/(yj-yi)+xi) {
			inside = !inside
		}
	}

	return inside
}

func (t *Theater) isRedTile(tile [2]int) bool {
	for _, redTile := range t.redTiles {
		if redTile[0] == tile[0] && redTile[1] == tile[1] {
			return true
		}
	}
	return false
}

func (t *Theater) computeAreaAllSquares() []Square {
	squares := make([]Square, 0)
	for i := 0; i < len(t.redTiles); i++ {
		tileA := t.redTiles[i]
		for j := i; j < len(t.redTiles); j++ {
			tileB := t.redTiles[j]

			if tileB[0] == tileA[0] && tileB[1] == tileA[1] {
				continue
			}

			// Calculate width and height, including the corner tiles
			width := abs(tileB[0]-tileA[0]) + 1
			height := abs(tileB[1]-tileA[1]) + 1
			area := width * height

			squares = append(squares, Square{
				tileAX: tileA[0],
				tileAY: tileA[1],
				tileBX: tileB[0],
				tileBY: tileB[1],
				area:   area,
			})
		}
	}

	return squares
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (t *Theater) ThaterMap() {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for _, tile := range t.redTiles {
		minX = min(minX, tile[0])
		minY = min(minY, tile[1])
		maxX = max(maxX, tile[0])
		maxY = max(maxY, tile[1])
	}

	// Add a border of 1 or 2 units for better visualization
	border := 2
	for y := minY - border; y <= maxY+border; y++ {
		for x := minX - border; x <= maxX+border; x++ {
			tile := [2]int{x, y}
			if t.isRedTile(tile) {
				fmt.Print("#")
			} else if t.isGreenTile(tile) {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (t *Theater) LargestAreaSquare() int {
	squares := t.computeAreaAllSquares()
	if len(squares) == 0 {
		return 0
	}
	sort.Slice(squares, func(i, j int) bool {
		return squares[i].area > squares[j].area
	})

	return squares[0].area
}

func (t *Theater) isOnLoopBoundary(tile [2]int) bool {
	if len(t.redTiles) < 2 {
		return false
	}
	// Check if the tile is on an axis-aligned segment between two consecutive red tiles
	for i := 0; i < len(t.redTiles); i++ {
		j := (i + 1) % len(t.redTiles)
		start := t.redTiles[i]
		end := t.redTiles[j]

		tileX, tileY := tile[0], tile[1]

		// Only check for axis-aligned segments as per the puzzle description.
		// Check if tile is on the segment between start and end
		if start[0] == end[0] { // Vertical segment
			minY := min(start[1], end[1])
			maxY := max(start[1], end[1])

			if tileX == start[0] && tileY >= minY && tileY <= maxY {
				return true
			}
		} else if start[1] == end[1] { // Horizontal segment
			minX := min(start[0], end[0])
			maxX := max(start[0], end[0])

			if tileY == start[1] && tileX >= minX && tileX <= maxX {
				return true
			}
		}
	}
	return false
}

func (t *Theater) isEligibleTile(tile [2]int) bool {
	key := tile
	if isEligible, found := t.insideCache[key]; found {
		return isEligible
	}

	isRed := t.isRedTile(tile)
	isOnBoundary := t.isOnLoopBoundary(tile)
	isInside := t.isInsideLoop(tile)

	isEligible := isRed || isOnBoundary || isInside
	t.insideCache[key] = isEligible
	return isEligible
}

func (t *Theater) isGreenTile(tile [2]int) bool {
	return t.isEligibleTile(tile) && !t.isRedTile(tile)
}

func NewTheater() *Theater {
	return &Theater{
		redTiles:    make([][2]int, 0),
		insideCache: make(map[[2]int]bool),
	}
}

func (t *Theater) AddRow(row string) {
	var x, y int
	_, err := fmt.Sscanf(row, "%d,%d", &x, &y)

	if err != nil {
		log.Fatalf("Failed to parse row: %s, error: %v", row, err)
	}
	newTile := [2]int{x, y}
	t.redTiles = append(t.redTiles, newTile)
}

// It returns the 2D Prefix Sum Array S, the compressed X coordinates, and compressed Y coordinates.
//
//	{1,    2, 3,  4},	 	{1,   2,  3,  4},
//	{5,   6,  7,  8},  -->  {5,   6,  7,  8},
//	{9,  10, 11, 12},		{9,  10, 11, 12},
//	{13, 14, 15, 16}		{13, 14, 15, 16}

// Her laver vi en matrix med alle mulige koordinater, og fylder den med arealerne af de gyldige felter. Så vi ved at hvis vi har et felt med værdi 4, så er der 4 gyldige felter i det område.
func (t *Theater) buildPrefixSumGrid() ([][]int, []int, []int) {
	// 1. Collect and Compress Coordinates
	xCoords, yCoords := make(map[int]bool), make(map[int]bool)

	// Collect all relevant boundary coordinates: corners and the points just outside
	for _, r := range t.redTiles {
		xCoords[r[0]] = true
		xCoords[r[0]+1] = true
		yCoords[r[1]] = true
		yCoords[r[1]+1] = true
	}

	// Convert map keys to sorted slices (the compressed axes)
	var uniqueX, uniqueY []int
	for x := range xCoords {
		uniqueX = append(uniqueX, x)
	}
	for y := range yCoords {
		uniqueY = append(uniqueY, y)
	}
	sort.Ints(uniqueX)
	sort.Ints(uniqueY)

	// Map: Original Coordinate -> Compressed Index
	xMap, yMap := make(map[int]int), make(map[int]int)
	for i, x := range uniqueX {
		xMap[x] = i
	}
	for i, y := range uniqueY {
		yMap[y] = i
	}

	width := len(uniqueX)
	height := len(uniqueY)

	// Initialize the Eligibility Grid
	grid := make([][]int, width)
	for i := range grid {
		grid[i] = make([]int, height)
	}

	// 2. Fill the Eligibility Grid (Area by Area)
	for i := 0; i < width-1; i++ {
		for j := 0; j < height-1; j++ {
			tileX := uniqueX[i]
			tileY := uniqueY[j]

			// Check eligibility for the entire cell
			if t.isEligibleTile([2]int{tileX, tileY}) {
				cellWidth := uniqueX[i+1] - tileX
				cellHeight := uniqueY[j+1] - tileY

				// Set the value in the grid to the area of the eligible region within this cell
				// For a fully eligible cell, this is the area of the cell itself.
				grid[i][j] = cellWidth * cellHeight
			}
		}
	}

	sum := make([][]int, width+1)
	for i := range sum {
		sum[i] = make([]int, height+1)
	}

	for i := 1; i <= width-1; i++ {
		for j := 1; j <= height-1; j++ {
			// sum[i][j] is the sum of areas in the rectangle from (0,0) to (i-1, j-1) of the grid
			sum[i][j] = grid[i-1][j-1] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}

	return sum, uniqueX, uniqueY
}

func (t *Theater) LargestAreaSquareInGreenTiles() int {
	sumMatrix, uniqueX, uniqueY := t.buildPrefixSumGrid()

	xMap, yMap := make(map[int]int), make(map[int]int)
	for i, x := range uniqueX {
		xMap[x] = i
	}
	for i, y := range uniqueY {
		yMap[y] = i
	}

	squares := t.computeAreaAllSquares()
	sort.Slice(squares, func(i, j int) bool {
		return squares[i].area > squares[j].area
	})

	best := 0

	for _, sq := range squares {
		if sq.area <= best {
			break
		}

		x1 := min(sq.tileAX, sq.tileBX)
		x2 := max(sq.tileAX, sq.tileBX)
		y1 := min(sq.tileAY, sq.tileBY)
		y2 := max(sq.tileAY, sq.tileBY)

		idx_max_x := xMap[x2+1]
		idx_max_y := yMap[y2+1]

		idx_min_x_minus_1 := xMap[x1]
		idx_min_y_minus_1 := yMap[y1]

		// Calculate the sum of eligible areas in the rectangle using the prefix sum array
		// https://upload.wikimedia.org/wikipedia/commons/thumb/5/58/Summed_area_table.png/250px-Summed_area_table.png
		sum := sumMatrix[idx_max_x][idx_max_y] -
			sumMatrix[idx_min_x_minus_1][idx_max_y] -
			sumMatrix[idx_max_x][idx_min_y_minus_1] +
			sumMatrix[idx_min_x_minus_1][idx_min_y_minus_1]

		totalTiles := (x2 - x1 + 1) * (y2 - y1 + 1)

		if sum == totalTiles {
			// The total eligible area in the rectangle is equal to the rectangle's area,
			// meaning it is fully covered by eligible tiles.
			best = sq.area
			break
		}
	}

	return best
}
