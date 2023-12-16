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

	grid := make([][]string, 0)
	for s.Scan() {
		grid = append(grid, strings.Split(s.Text(), ""))
	}

	tiltNorth(grid)
	fmt.Printf("Answer Day Fourteen, Part One: %v\n", supportLoadOnNorth(grid))
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	grid := make([][]string, 0)
	for s.Scan() {
		grid = append(grid, strings.Split(s.Text(), ""))
	}

	var (
		nodes = make(map[string]int)
		limit int
	)
	var loops int
	for limit == 0 || loops < limit {
		str := concat(grid)
		if loop, ok := nodes[str]; ok {
			// The tilting has started to repeat. Therfore, we can calculate where we would stop if repeated 1000000000 times.
			limit = loops + ((1000000000 - loop) % (loops - loop))
		} else {
			nodes[str] = loops
		}
		tiltNorth(grid)
		tiltWest(grid)
		tiltSouth(grid)
		tiltEast(grid)
		loops++
	}

	fmt.Printf("Answer Day Fourteen, Part Two: %v\n", supportLoadOnNorth(grid))
}

func supportLoadOnNorth(grid [][]string) int {
	var ans int
	numRows := len(grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "O" {
				ans += numRows - i
			}
		}
	}
	return ans
}

func concat(grid [][]string) string {
	var ans string
	for _, row := range grid {
		ans += strings.Join(row, "")
	}
	return ans
}

func tiltNorth(grid [][]string) {
	blanks := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "O" {
				grid[i][j] = "."
				grid[i-blanks[j]][j] = "O"
			} else if grid[i][j] == "." {
				blanks[j]++
			} else {
				blanks[j] = 0
			}
		}
	}
}

func tiltWest(grid [][]string) {
	blanks := make([]int, len(grid))
	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == "O" {
				grid[i][j] = "."
				grid[i][j-blanks[i]] = "O"
			} else if grid[i][j] == "." {
				blanks[i]++
			} else {
				blanks[i] = 0
			}
		}
	}
}

func tiltSouth(grid [][]string) {
	blanks := make([]int, len(grid[0]))
	for i := len(grid) - 1; i >= 0; i-- {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "O" {
				grid[i][j] = "."
				grid[i+blanks[j]][j] = "O"
			} else if grid[i][j] == "." {
				blanks[j]++
			} else {
				blanks[j] = 0
			}
		}
	}
}

func tiltEast(grid [][]string) {
	blanks := make([]int, len(grid))
	for j := len(grid[0]) - 1; j >= 0; j-- {
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == "O" {
				grid[i][j] = "."
				grid[i][j+blanks[i]] = "O"
			} else if grid[i][j] == "." {
				blanks[i]++
			} else {
				blanks[i] = 0
			}
		}
	}
}
