package main

import (
	"embed"
	"fmt"
	"strings"

	"github.com/thedadams/adventofcode/2023/util"
)

//go:embed input.txt
var f embed.FS

func main() {
	partOne()
	partTwo()
}

func partOne() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	grids := make([][][]int, 0)
	grid := make([][]int, 0)
	for s.Scan() {
		if s.Text() == "" {
			grids = append(grids, grid)
			grid = make([][]int, 0)
		} else {
			row := strings.Split(s.Text(), "")
			grid = append(grid, make([]int, len(row)))
			for i, spot := range row {
				if spot == "." {
					grid[len(grid)-1][i] = 0
				} else {
					grid[len(grid)-1][i] = 1
				}
			}
		}
	}

	if len(grid) > 0 {
		grids = append(grids, grid)
	}

	var ans int
	for _, g := range grids {
		if score := findReflectionScore(g, 0); score > 0 {
			ans += 100 * score
		} else {
			ans += findReflectionScore(util.Transpose(g), 0)
		}
	}

	fmt.Printf("Answer Day Thirteen, Part One: %v\n", ans)
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	grids := make([][][]int, 0)
	grid := make([][]int, 0)
	for s.Scan() {
		if s.Text() == "" {
			grids = append(grids, grid)
			grid = make([][]int, 0)
		} else {
			row := strings.Split(s.Text(), "")
			grid = append(grid, make([]int, len(row)))
			for i, spot := range row {
				if spot == "." {
					grid[len(grid)-1][i] = 0
				} else {
					grid[len(grid)-1][i] = 1
				}
			}
		}
	}

	if len(grid) > 0 {
		grids = append(grids, grid)
	}

	var ans int
	for _, g := range grids {
		if score := findReflectionScore(g, 1); score > 0 {
			ans += 100 * score
		} else {
			ans += findReflectionScore(util.Transpose(g), 1)
		}
	}

	fmt.Printf("Answer Day Thirteen, Part Two: %v\n", ans)
}

func findReflectionScore(grid [][]int, totalDifference int) int {
	possibleReflectionPoints := make([]int, 0)
	for i := 1; i < len(grid); i++ {
		var xorSum int
		for j := 0; j < len(grid[i]); j++ {
			xorSum += grid[i][j] ^ grid[i-1][j]
		}
		if xorSum <= totalDifference {
			possibleReflectionPoints = append(possibleReflectionPoints, i-1)
		}
	}

	for _, i := range possibleReflectionPoints {
		j := 1
		reflection := true
		var xorSum int
		for i+j < len(grid) && i-j+1 >= 0 {
			for k := 0; k < len(grid[i+j]); k++ {
				xorSum += grid[i+j][k] ^ grid[i-j+1][k]
			}
			if xorSum > totalDifference {
				reflection = false
				break
			}
			j++
		}
		if reflection && xorSum == totalDifference {
			return i + 1
		}
	}

	return 0
}
