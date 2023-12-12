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

	fmt.Printf("Answer Day Eleven, Part One: %v\n", sumOfShortestDistancesBetweenGalaxies(s, 2))
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	fmt.Printf("Answer Day Eleven, Part Two: %v\n", sumOfShortestDistancesBetweenGalaxies(s, 1000000))
}

func sumOfShortestDistancesBetweenGalaxies(s *util.Scanner, multiplier int) int {
	grid := make([][]string, 0)
	galaxies := make([][2]int, 0)
	emptyRows := make([]int, 0)
	emptyCols := make([]int, 0)
	for s.Scan() {
		grid = append(grid, strings.Split(s.Text(), ""))
		if !strings.Contains(s.Text(), "#") {
			emptyRows = append(emptyRows, len(grid)-1)
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		var hasGalaxy bool
		for j := range grid {
			if grid[j][i] == "#" {
				hasGalaxy = true
				galaxies = append(galaxies, [2]int{j, i})
			}
		}

		if !hasGalaxy {
			emptyCols = append(emptyCols, i)
		}
	}

	var ans int
	for i, galaxy := range galaxies {
		for _, other := range galaxies[i+1:] {
			start, end := galaxy[0], other[0]
			if start > end {
				start, end = end, start
			}
			for _, row := range emptyRows {
				if row > start && row < end {
					// One row turns into multiplier number of rows, so multiplier-1 extra rows.
					ans += multiplier - 1
				}
			}
			start, end = galaxy[1], other[1]
			if start > end {
				start, end = end, start
			}
			for _, col := range emptyCols {
				if col > start && col < end {
					// One col turns into multiplier number of cols, so multiplier-1 extra cols.
					ans += multiplier - 1
				}
			}
			ans += util.AbsoluteValue(galaxy[0], other[0]) + util.AbsoluteValue(galaxy[1], other[1])
		}
	}

	return ans
}
