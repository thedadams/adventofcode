package util

import "fmt"

func ValidStandardNeighbors[T comparable](grid [][]T, i, j int) [][2]int {
	nbrs := make([][2]int, 0)
	for _, pair := range [][2]int{
		{i, j + 1},
		{i, j - 1},
		{i + 1, j},
		{i - 1, j},
	} {
		if pair[0] >= 0 && pair[1] >= 0 && pair[0] < len(grid) && pair[1] < len(grid[pair[0]]) {
			nbrs = append(nbrs, pair)
		}
	}

	return nbrs
}

func ValidNeighborsWithDiagonals[T comparable](grid [][]T, x, y int) [][2]int {
	nbrs := make([][2]int, 0)
	for _, pair := range [][2]int{
		{x - 1, y - 1},
		{x + 1, y - 1},
		{x + 1, y + 1},
		{x - 1, y + 1},
		{x, y + 1},
		{x, y - 1},
		{x + 1, y},
		{x - 1, y},
	} {
		if pair[0] >= 0 && pair[1] >= 0 && pair[0] < len(grid) && pair[1] < len(grid[pair[0]]) {
			nbrs = append(nbrs, pair)
		}
	}

	return nbrs
}

func PrintGrid[T any](grid [][]T) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}
