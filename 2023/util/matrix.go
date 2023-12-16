package util

func Transpose[T any](grid [][]T) [][]T {
	ans := make([][]T, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		ans[i] = make([]T, len(grid))
		for j := 0; j < len(grid); j++ {
			ans[i][j] = grid[j][i]
		}
	}
	return ans
}
