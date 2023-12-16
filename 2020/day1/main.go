package main

import (
	"embed"
	"fmt"

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

	numbersThatMatch := make(map[int]struct{})
	for s.Scan() {
		num := util.MustAtoi(s.Text())
		if _, ok := numbersThatMatch[2020-num]; ok {
			fmt.Printf("Answer Day One, Part One: %v\n", num*(2020-num))
			return
		}
		numbersThatMatch[num] = struct{}{}
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

	numbersThatMatch := make(map[int]struct{})
	for s.Scan() {
		num := util.MustAtoi(s.Text())
		for n := range numbersThatMatch {
			if _, ok := numbersThatMatch[2020-num-n]; ok {
				fmt.Printf("Answer Day One, Part Two: %v\n", num*n*(2020-num-n))
				return
			}
		}
		numbersThatMatch[num] = struct{}{}
	}
}
