package main

import (
	"embed"
	"fmt"
	"slices"
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

	grid := make([][][]string, 0)
	for s.Scan() {
		grid = append(grid, make([][]string, 0))
		for _, spot := range strings.Split(s.Text(), "") {
			grid[len(grid)-1] = append(grid[len(grid)-1], []string{spot})
		}
	}

	energized := make(map[[2]int]struct{})
	energizeGrid(grid, energized, [2]int{0, 0}, "E")

	fmt.Printf("Answer Day Fifteen, Part One: %v\n", len(energized))
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	grid := make([][][]string, 0)
	for s.Scan() {
		grid = append(grid, make([][]string, 0))
		for _, spot := range strings.Split(s.Text(), "") {
			grid[len(grid)-1] = append(grid[len(grid)-1], []string{spot})
		}
	}

	var ans int
	for j := 0; j < len(grid[0]); j++ {
		energized := make(map[[2]int]struct{})
		energizeGrid(grid, energized, [2]int{0, j}, "S")
		if len(energized) > ans {
			ans = len(energized)
		}

		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				grid[x][y] = grid[x][y][:1]
			}
		}

		energized = make(map[[2]int]struct{})
		energizeGrid(grid, energized, [2]int{len(grid) - 1, j}, "N")
		if len(energized) > ans {
			ans = len(energized)
		}

		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				grid[x][y] = grid[x][y][:1]
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		energized := make(map[[2]int]struct{})
		energizeGrid(grid, energized, [2]int{i, 0}, "E")
		if len(energized) > ans {
			ans = len(energized)
		}

		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				grid[x][y] = grid[x][y][:1]
			}
		}

		energized = make(map[[2]int]struct{})
		energizeGrid(grid, energized, [2]int{i, len(grid[i]) - 1}, "W")
		if len(energized) > ans {
			ans = len(energized)
		}

		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				grid[x][y] = grid[x][y][:1]
			}
		}
	}

	fmt.Printf("Answer Day Fifteen, Part Two: %v\n", ans)
}

func energizeGrid(grid [][][]string, energized map[[2]int]struct{}, laserPositions [2]int, direction string) {
	if laserPositions[0] < 0 || laserPositions[1] < 0 {
		return
	}

	symbols := grid[laserPositions[0]][laserPositions[1]]
	if slices.Contains(symbols, "v") && direction == "S" ||
		slices.Contains(symbols, ">") && direction == "E" ||
		slices.Contains(symbols, "<") && direction == "W" ||
		slices.Contains(symbols, "^") && direction == "N" {
		// If we've already been down this road, then we don't need to do it again.
		return
	}

	switch direction {
	case "N":
		grid[laserPositions[0]][laserPositions[1]] = append(grid[laserPositions[0]][laserPositions[1]], "^")
	case "S":
		grid[laserPositions[0]][laserPositions[1]] = append(grid[laserPositions[0]][laserPositions[1]], "v")
	case "E":
		grid[laserPositions[0]][laserPositions[1]] = append(grid[laserPositions[0]][laserPositions[1]], ">")
	case "W":
		grid[laserPositions[0]][laserPositions[1]] = append(grid[laserPositions[0]][laserPositions[1]], "<")
	}

	energized[laserPositions] = struct{}{}

	switch symbols[0] {
	case ".":
		energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, direction), direction)
	case "-":
		switch direction {
		case "E", "W":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, direction), direction)
		case "N", "S":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "E"), "E")
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "W"), "W")
		}

	case "|":
		switch direction {
		case "N", "S":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, direction), direction)
		case "E", "W":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "N"), "N")
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "S"), "S")
		}

	case "/":
		switch direction {
		case "N":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "E"), "E")
		case "S":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "W"), "W")
		case "E":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "N"), "N")
		case "W":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "S"), "S")
		}

	case "\\":
		switch direction {
		case "N":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "W"), "W")
		case "S":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "E"), "E")
		case "E":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "S"), "S")
		case "W":
			energizeGrid(grid, energized, nextNeighbor(grid, laserPositions, "N"), "N")
		}
	}
}

func nextNeighbor(grid [][][]string, pos [2]int, direction string) [2]int {
	i, j := pos[0], pos[1]
	switch direction {
	case "N":
		i--
	case "S":
		i++
	case "E":
		j++
	case "W":
		j--
	}

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		i, j = -1, -1
	}
	return [2]int{i, j}
}
