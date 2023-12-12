package main

import (
	"embed"
	"fmt"
	"strconv"
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
	visited := make([][]bool, 0)
	for s.Scan() {
		grid = append(grid, strings.Split(s.Text(), ""))
		visited = append(visited, make([]bool, len(grid[len(grid)-1])))
	}

	var ans int
	for i, row := range grid {
		for j, item := range row {
			if _, err := strconv.Atoi(item); item != "." && err != nil {
				for _, pair := range util.ValidNeighborsWithDiagonals(grid, i, j) {
					if _, err := strconv.Atoi(grid[pair[0]][pair[1]]); visited[pair[0]][pair[1]] || err != nil {
						continue
					}
					ans += expandNumber(grid, visited, pair[0], pair[1])
				}
			}
		}
	}

	fmt.Printf("Answer Day Three, Part One: %v\n", ans)
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
	visited := make([][]bool, 0)
	for s.Scan() {
		grid = append(grid, strings.Split(s.Text(), ""))
		visited = append(visited, make([]bool, len(grid[len(grid)-1])))
	}

	var ans int
	for i, row := range grid {
		for j, item := range row {
			if item == "*" {
				for _, row := range visited {
					clear(row)
				}
				adjacent, ratio := 0, 1
				for _, pair := range util.ValidNeighborsWithDiagonals(grid, i, j) {
					if _, err := strconv.Atoi(grid[pair[0]][pair[1]]); visited[pair[0]][pair[1]] || err != nil {
						continue
					}
					adjacent++
					ratio *= expandNumber(grid, visited, pair[0], pair[1])
				}
				if adjacent == 2 {
					ans += ratio
				}
			}
		}
	}

	fmt.Printf("Answer Day Three, Part Two: %v\n", ans)
}

func expandNumber(grid [][]string, visited [][]bool, i, j int) int {
	var (
		err      error
		num, ans int
	)
	for j > 0 && err == nil {
		j--
		_, err = strconv.Atoi(grid[i][j])
	}

	if err != nil {
		j++
		err = nil
	}
	for j < len(grid[i]) && err == nil {
		num, err = strconv.Atoi(grid[i][j])
		ans = ans*10 + num
		visited[i][j] = true
		j++
	}

	if err != nil {
		return ans / 10

	}
	return ans
}
