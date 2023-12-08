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

	var validPasswords int
	for s.Scan() {
		criteria, password, _ := strings.Cut(s.Text(), ":")
		criteriaParts := strings.Split(criteria, " ")
		minCount, _ := strconv.Atoi(strings.Split(criteriaParts[0], "-")[0])
		maxCount, _ := strconv.Atoi(strings.Split(criteriaParts[0], "-")[1])

		var count int
		for _, c := range strings.TrimSpace(password) {
			if string(c) == criteriaParts[1] {
				count++
			}
		}
		if count >= minCount && count <= maxCount {
			validPasswords++
		}
	}

	fmt.Printf("Answer Day Two, Part One: %v\n", validPasswords)
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	var validPasswords int
	for s.Scan() {
		criteria, password, _ := strings.Cut(s.Text(), ":")
		criteriaParts := strings.Split(criteria, " ")
		firstIndex, _ := strconv.Atoi(strings.Split(criteriaParts[0], "-")[0])
		secondIndex, _ := strconv.Atoi(strings.Split(criteriaParts[0], "-")[1])

		var firstMatch, secondMatch bool
		for i, c := range strings.TrimSpace(password) {
			if string(c) == criteriaParts[1] {
				if firstIndex == (i + 1) {
					firstMatch = true
				}
				if secondIndex == (i + 1) {
					secondMatch = true
				}
			}
		}
		if (firstMatch || secondMatch) && firstMatch != secondMatch {
			validPasswords++
		}
	}

	fmt.Printf("Answer Day Two, Part One: %v\n", validPasswords)
}
