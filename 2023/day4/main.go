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

	var ans int
	for s.Scan() {
		_, nums, _ := strings.Cut(s.Text(), ":")
		winningNumbers, myNumbers, _ := strings.Cut(nums, "|")

		winningSet := make(map[string]struct{}, len(winningNumbers))
		for _, num := range strings.Split(winningNumbers, " ") {
			if num = strings.TrimSpace(num); num != "" {
				winningSet[strings.TrimSpace(num)] = struct{}{}
			}
		}

		var score int
		for _, num := range strings.Split(myNumbers, " ") {
			if _, ok := winningSet[strings.TrimSpace(num)]; ok {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		ans += score
	}

	fmt.Printf("Answer Day Four, Part One: %v\n", ans)
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	cards := make([]int, 0)
	var count int
	for s.Scan() {
		count++
		for len(cards) <= count {
			cards = append(cards, 1)
		}

		_, nums, _ := strings.Cut(s.Text(), ":")
		winningNumbers, myNumbers, _ := strings.Cut(nums, "|")

		winningSet := make(map[string]struct{}, len(winningNumbers))
		for _, num := range strings.Split(winningNumbers, " ") {
			if num = strings.TrimSpace(num); num != "" {
				winningSet[strings.TrimSpace(num)] = struct{}{}
			}
		}

		score := 1
		for _, num := range strings.Split(myNumbers, " ") {
			if _, ok := winningSet[strings.TrimSpace(num)]; ok {
				for len(cards) <= count+score {
					cards = append(cards, 1)
				}

				cards[count+score] += cards[count]
				score++
			}
		}
	}

	var ans int
	for _, count := range cards[1:] {
		ans += count
	}
	fmt.Printf("Answer Day Four, Part Two: %v\n", ans)
}
