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

	trees, x, y := 0, 0, 0
	for y < len(grid) {
		if grid[y][x] == "#" {
			trees++
		}
		x = (x + 3) % len(grid[y])
		y++
	}

	fmt.Printf("Answer Day Three, Part One: %v\n", trees)
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

	ans := 1
	for _, slope := range [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		trees, x, y := 0, 0, 0
		for y < len(grid) {
			if grid[y][x] == "#" {
				trees++
			}
			x = (x + slope[0]) % len(grid[y])
			y += slope[1]
		}

		ans *= trees
	}

	fmt.Printf("Answer Day Three, Part Two: %v\n", ans)
}
