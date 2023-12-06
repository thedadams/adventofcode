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

	times := make([]int, 0)
	distances := make([]int, 0)

	s.Scan()
	_, nums, _ := strings.Cut(s.Text(), ":")
	for _, num := range strings.Split(nums, " ") {
		if num := strings.TrimSpace(num); num != "" {
			n, _ := strconv.Atoi(num)
			times = append(times, n)
		}
	}

	s.Scan()
	_, nums, _ = strings.Cut(s.Text(), ":")
	for _, num := range strings.Split(nums, " ") {
		if num := strings.TrimSpace(num); num != "" {
			n, _ := strconv.Atoi(num)
			distances = append(distances, n)
		}
	}

	product := 1
	for i := range times {
		var ways int
		for j := 0; j < times[i]; j++ {
			if j*j-times[i]*j+distances[i] < 0 {
				ways++
			}
		}

		product *= ways
	}

	fmt.Printf("Answer Day Six, Part One: %v\n", product)
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	s.Scan()
	_, nums, _ := strings.Cut(s.Text(), ":")
	totalTime, _ := strconv.Atoi(strings.ReplaceAll(nums, " ", ""))

	s.Scan()
	_, nums, _ = strings.Cut(s.Text(), ":")
	totalDistance, _ := strconv.Atoi(strings.ReplaceAll(nums, " ", ""))

	var ways int
	for j := 0; j < totalTime; j++ {
		if j*j-totalTime*j+totalDistance < 0 {
			ways++
		}
	}

	fmt.Printf("Answer Day Six, Part Two: %v\n", ways)
}
