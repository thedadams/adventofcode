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

	s.Scan()
	pattern := strings.Split(s.Text(), "")
	network := make(map[string][2]string)
	// Blank line
	s.Scan()
	for s.Scan() {
		keyVal := strings.Split(s.Text(), " = ")
		network[keyVal[0]] = *((*[2]string)(strings.Split(keyVal[1][1:len(keyVal[1])-1], ", ")))
	}

	var hops int
	start := "AAA"
	for start != "ZZZ" {
		for _, direction := range pattern {
			if direction == "L" {
				start = network[start][0]
			} else {
				start = network[start][1]
			}
			hops++
			if start == "ZZZ" {
				break
			}
		}
	}

	fmt.Printf("Answer Day Eight, Part One: %v\n", hops)
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
	pattern := strings.Split(s.Text(), "")
	network := make(map[string][2]string)
	positions := make([]string, 0)
	// Blank line
	s.Scan()
	for s.Scan() {
		keyVal := strings.Split(s.Text(), " = ")
		network[keyVal[0]] = *((*[2]string)(strings.Split(keyVal[1][1:len(keyVal[1])-1], ", ")))
		if strings.HasSuffix(keyVal[0], "A") {
			positions = append(positions, keyVal[0])
		}
	}

	firstZs := make([]int, len(positions))
	var hops int
	for {
		for _, direction := range pattern {
			hops++
			for i := range positions {
				if direction == "L" {
					positions[i] = network[positions[i]][0]
				} else {
					positions[i] = network[positions[i]][1]
				}
				if strings.HasSuffix(positions[i], "Z") {
					if firstZs[i] == 0 {
						firstZs[i] = hops
					}
				}
			}

			allZsFound := true
			for _, z := range firstZs {
				if z == 0 {
					allZsFound = false
					break
				}
			}
			if allZsFound {
				fmt.Printf("Answer Day Eight, Part Two: %v\n", lcmSlice(firstZs))
				return
			}
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmSlice(s []int) int {
	result := s[0]
	for i := 1; i < len(s); i++ {
		result = lcm(result, s[i])
	}
	return result
}
