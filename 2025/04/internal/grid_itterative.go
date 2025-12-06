package internal

import "fmt"

type GridIT struct {
	// 2D grid of integers
	// [row][column]
	grid                 [][]int
	row                  int
	paperRollsForRemoval []Index
}

func NewGridIT() *GridIT {
	return &GridIT{
		grid:                 make([][]int, 0),
		row:                  0,
		paperRollsForRemoval: make([]Index, 0),
	}
}

func (g *GridIT) AddRowInput(line string) {
	row := ParseRow(line)
	g.AddRow(row)
}

type Index struct {
	Row int
	Col int
}

func (g *GridIT) GetAccessibleRecursive() int {
	totalAccessible := 0
	accessibleCount := -1
	for accessibleCount != 0 {
		accessibleCount = g.GetAccessibleCount()
		fmt.Printf("accessibleCount: %d\n", accessibleCount)
		fmt.Printf("paperRollsForRemoval: %v\n", g.paperRollsForRemoval)
		totalAccessible += accessibleCount
		g.removePaperRolls()
	}

	return totalAccessible

}

func (g *GridIT) removePaperRolls() {
	for _, index := range g.paperRollsForRemoval {
		g.grid[index.Row][index.Col] = 0
	}
	g.paperRollsForRemoval = make([]Index, 0)
}

func (g *GridIT) GetAccessibleCount() int {
	accessibleCount := 0
	for i := 0; i < g.row; i++ {
		for j := 0; j < len(g.grid[i]); j++ {
			if g.grid[i][j] == 0 {
				continue
			}
			topLeft := g.getValuue(i-1, j-1)
			top := g.getValuue(i-1, j)
			topRight := g.getValuue(i-1, j+1)
			left := g.getValuue(i, j-1)
			right := g.getValuue(i, j+1)
			bottomLeft := g.getValuue(i+1, j-1)
			bottom := g.getValuue(i+1, j)
			bottomRight := g.getValuue(i+1, j+1)

			sumOfNeighbors := topLeft + top + topRight + left + right + bottomLeft + bottom + bottomRight
			if sumOfNeighbors < 4 {
				accessibleCount++
				g.paperRollsForRemoval = append(g.paperRollsForRemoval, Index{Row: i, Col: j})
			}
		}
	}

	return accessibleCount
}

func (g *GridIT) getValuue(row int, col int) int {
	if row < 0 || row >= g.row {
		return 0
	}

	if col < 0 || col >= len(g.grid[row]) {
		return 0
	}

	return g.grid[row][col]
}

func (g *GridIT) AddRow(row []int) {
	g.grid = append(g.grid, row)
	g.row++

}

func (g *GridIT) Print() {
	for _, row := range g.grid {
		for _, val := range row {
			if val == 1 {
				print("@")
			} else {
				print(".")
			}
		}
		println()
	}
}
