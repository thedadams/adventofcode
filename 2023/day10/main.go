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
	counts := make([][]int, 0)
	var startX, startY int
	for s.Scan() {
		grid = append(grid, strings.Split(s.Text(), ""))
		counts = append(counts, make([]int, len(grid[len(grid)-1])))
		for i := range counts[len(counts)-1] {
			counts[len(counts)-1][i] = -1
		}
		if idx := strings.Index(s.Text(), "S"); idx > -1 {
			startX = idx
			startY = len(grid) - 1
			counts[startY][startX] = 0
		}
	}

	nextSpots := make([][4]int, 0)
	// Find the starting directions for the search
	for _, nb := range util.ValidStandardNeighbors(grid, startX, startY) {
		if next := nextNeighbor(grid, startX, startY, nb[0], nb[1]); next[0] != -1 && next[1] != -1 {
			nextSpots = append(nextSpots, [4]int{startX, startY, nb[0], nb[1]})
		}
	}

	for len(nextSpots) > 0 {
		this := nextSpots[0]
		nextSpots = nextSpots[1:]
		count := counts[this[1]][this[0]] + 1
		if counts[this[3]][this[2]] != -1 {
			fmt.Printf("Answer Day Ten, Part One: %v\n", count)
			return
		}
		counts[this[3]][this[2]] = count

		if next := nextNeighbor(grid, this[0], this[1], this[2], this[3]); counts[next[1]][next[0]] == -1 {
			nextSpots = append(nextSpots, [4]int{this[2], this[3], next[0], next[1]})
		}
	}
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
	inOut := make([][]string, 0)
	var startX, startY int
	for s.Scan() {
		grid = append(grid, strings.Split(s.Text(), ""))
		inOut = append(inOut, make([]string, len(grid[len(grid)-1])))
		if idx := strings.Index(s.Text(), "S"); idx > -1 {
			startX = idx
			startY = len(grid) - 1
			inOut[startY][startX] = "P"
		}
	}

	nextSpots := make([][4]int, 0)
	for _, nb := range util.ValidStandardNeighbors(grid, startX, startY) {
		if next := nextNeighbor(grid, startX, startY, nb[0], nb[1]); next[0] != -1 && next[1] != -1 {
			nextSpots = append(nextSpots, [4]int{startX, startY, nb[0], nb[1]})
			break
		}
	}

	for len(nextSpots) > 0 {
		this := nextSpots[0]
		nextSpots = nextSpots[1:]
		inOut[this[3]][this[2]] = "P"

		if next := nextNeighbor(grid, this[0], this[1], this[2], this[3]); inOut[next[1]][next[0]] == "" {
			nextSpots = append(nextSpots, [4]int{this[2], this[3], next[0], next[1]})
		}
	}

	var (
		count int
	)
	for i := range inOut {
		var (
			inside bool
			prev   string
		)
		for j := range inOut[i] {
			if inOut[i][j] == "P" {
				if grid[i][j] == "|" {
					inside = !inside
				} else if grid[i][j] == "-" {
					continue
				} else if prev == "L" {
					if grid[i][j] == "7" {
						inside = !inside
						prev = ""
					} else if grid[i][j] == "J" {
						prev = ""
					}
				} else if prev == "F" {
					if grid[i][j] == "J" {
						inside = !inside
						prev = ""
					} else if grid[i][j] == "7" {
						prev = ""
					}
				}
			} else if inOut[i][j] == "" && inside {
				inOut[i][j] = "I"
				count++
			} else if inOut[i][j] == "" {
				inOut[i][j] = "O"
			}
			prev = grid[i][j]
		}
	}

	for _, row := range inOut {
		for _, spot := range row {
			fmt.Printf("%2s", spot)
		}
		fmt.Println()
	}

	fmt.Printf("Answer Day Ten, Part Two: %v\n", count)
}

func nextNeighbor(grid [][]string, prevI, prevJ, i, j int) [2]int {
	nextI, nextJ := i, j
	switch grid[j][i] {
	case "|":
		if prevJ < j {
			nextJ++
		} else if prevJ > j {
			nextJ--
		} else {
			nextI, nextJ = -1, -1
		}
	case "-":
		if prevI < i {
			nextI++
		} else if prevI > i {
			nextI--
		} else {
			nextI, nextJ = -1, -1
		}
	case "L":
		if prevJ < j {
			nextI++
		} else if prevI > i {
			nextJ--
		} else {
			nextI, nextJ = -1, -1
		}
	case "J":
		if prevJ < j {
			nextI--
		} else if prevI < i {
			nextJ--
		} else {
			nextI, nextJ = -1, -1
		}
	case "7":
		if prevJ > j {
			nextI--
		} else if prevI < i {
			nextJ++
		} else {
			nextI, nextJ = -1, -1
		}
	case "F":
		if prevJ > j {
			nextI++
		} else if prevI > i {
			nextJ++
		} else {
			nextI, nextJ = -1, -1
		}
	default:
		nextI, nextJ = -1, -1
	}

	return [2]int{nextI, nextJ}
}
