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
	config := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	s, err := util.ReadInputFile(f)
	if err != nil {
		return
	}

	var ans, id int
	for s.Scan() {
		possible := true
		gameInfo := strings.Split(s.Text(), ":")
		id, _ = strconv.Atoi(strings.TrimPrefix(gameInfo[0], "Game "))

	outer:
		for _, round := range strings.Split(gameInfo[1], ";") {
			for _, kind := range strings.Split(round, ",") {
				kind = strings.TrimSpace(kind)
				if count, _ := strconv.Atoi(strings.Split(kind, " ")[0]); count > config[strings.Split(kind, " ")[1]] {
					possible = false
					break outer
				}
			}
		}

		if possible {
			ans += id
		}
	}

	fmt.Printf("Answer Day Two, Part One: %v\n", ans)
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		return
	}

	var ans int
	for s.Scan() {
		config := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		gameInfo := strings.Split(s.Text(), ":")
		for _, round := range strings.Split(gameInfo[1], ";") {
			for _, kind := range strings.Split(round, ",") {
				kind = strings.TrimSpace(kind)
				if count, _ := strconv.Atoi(strings.Split(kind, " ")[0]); count > config[strings.Split(kind, " ")[1]] {
					config[strings.Split(kind, " ")[1]] = count
				}
			}
		}

		power := 1
		for _, count := range config {
			power *= count
		}
		ans += power
	}

	fmt.Printf("Answer Day Two, Part Two: %v\n", ans)
}
