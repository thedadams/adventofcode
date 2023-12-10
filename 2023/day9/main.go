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

	var ans int64
	for s.Scan() {
		input := strings.Split(s.Text(), " ")
		nums := make([]int64, len(input))
		for i, num := range input {
			n, _ := strconv.ParseInt(num, 10, 64)
			nums[i] = n
		}

		var val int64
		count := int64(len(nums))
		var neg bool
		for i := count - 1; i >= 0; i-- {
			this := choose(count, i) * nums[i]
			if neg {
				this = -this
			}
			val += this
			neg = !neg
		}

		ans += val
	}

	fmt.Printf("Answer Day Nine, Part One: %v\n", ans)
}

func choose(a, b int64) int64 {
	if a-b > b {
		b = a - b
	}

	var ans int64 = 1
	for i := b + 1; i <= a; i++ {
		ans *= i
	}
	for i := int64(2); i <= a-b; i++ {
		ans /= i
	}

	return ans
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	var ans int64
	for s.Scan() {
		input := strings.Split(s.Text(), " ")
		nums := make([]int64, len(input))
		for i, num := range input {
			n, _ := strconv.ParseInt(num, 10, 64)
			nums[i] = n
		}

		var val int64
		count := int64(len(nums)) - 1
		var neg bool
		for i := int64(0); i < count; i++ {
			this := choose(count, i+1) * nums[i]
			if neg {
				this = -this
			}
			val += this
			neg = !neg
		}

		fmt.Println(val)
		ans += val
	}

	fmt.Printf("Answer Day Nine, Part Two: %v\n", ans)
}
