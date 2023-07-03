package d8

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

var field = make([][]int, 0)

func Day8() int {
	reader := bufio.NewScanner(os.Stdin)
	// Parse input
	for reader.Scan() {
		line := reader.Text()
		row := strings.Split(line, "")
		iRow := make([]int, 0)
		for i := 0; i < len(row); i++ {
			val, _ := strconv.Atoi(row[i])
			iRow = append(iRow, val)
		}
		field = append(field, iRow)
	}
	//visible := findParameter()
	//interior := 0
	// For part 2 store the scenic scores in a list for the trees
	treeViews := make([]int, 0)
	for i := 1; i < len(field)-1; i++ {
		for j := 1; j < len(field[0])-1; j++ {
			treeViews = append(treeViews, checkUp(i, j)*checkDown(i, j)*checkLeft(i, j)*checkRight(i, j))
		}
	}
	maxTreeScenic := treeViews[0]
	for _, treeScenic := range treeViews {
		maxTreeScenic = int(math.Max(float64(maxTreeScenic), float64(treeScenic)))
	}
	return maxTreeScenic
}

func findParameter() int {
	parameter := 0

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if i == 0 || i == len(field)-1 || j == 0 || j == len(field[0])-1 {
				parameter++
			}
		}
	}
	return parameter
}

func checkUp(y int, x int) int {
	visible := 0
	tree := field[y][x]
	for i := y - 1; i >= 0; i-- {
		if field[i][x] < tree {
			visible++
		} else {
			return visible + 1
		}
	}
	return visible
}

// Slightly change conditionals for Part 2
func checkDown(y int, x int) int {
	visible := 0
	tree := field[y][x]
	for i := y + 1; i < len(field); i++ {
		if field[i][x] < tree {
			visible++
		} else {
			return visible + 1
		}
	}
	return visible
}

func checkLeft(y int, x int) int {
	visible := 0
	tree := field[y][x]

	for j := x - 1; j >= 0; j-- {
		if field[y][j] < tree {
			visible++
		} else {
			return visible + 1
		}
	}
	return visible
}

func checkRight(y int, x int) int {
	visible := 0
	tree := field[y][x]

	for j := x + 1; j < len(field[y]); j++ {
		if field[y][j] < tree {
			visible++
		} else {
			return visible + 1
		}
	}
	return visible
}
