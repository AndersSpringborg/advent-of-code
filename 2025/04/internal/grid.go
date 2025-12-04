package internal

type Grid struct {
	// 2D grid of integers
	// [row][column]
	grid [][]int
	row  int
}

func NewGrid() *Grid {
	return &Grid{
		grid: make([][]int, 0),
		row:  0,
	}
}

func (g *Grid) AddRowInput(line string) {
	row := ParseRow(line)
	g.AddRow(row)
}

func (g *Grid) GetAccessibleCount() int {
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
			}
		}
	}

	return accessibleCount
}

func ParseRow(line string) []int {
	row := make([]int, len(line))
	for i, ch := range line {
		roll := ch == '@'
		if roll {
			row[i] = 1
		} else {
			row[i] = 0
		}
	}
	return row
}

func (g *Grid) getValuue(row int, col int) int {
	if row < 0 || row >= g.row {
		return 0
	}

	if col < 0 || col >= len(g.grid[row]) {
		return 0
	}

	return g.grid[row][col]
}

func (g *Grid) AddRow(row []int) {
	g.grid = append(g.grid, row)
	g.row++

}

func (g *Grid) Print() {
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
