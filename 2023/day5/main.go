package main

import (
	"embed"
	"fmt"
	"math"
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

	seeds := make([]int, 0)
	s.Scan()
	_, nums, _ := strings.Cut(s.Text(), ":")
	for _, num := range strings.Split(nums, " ") {
		if num := strings.TrimSpace(num); num != "" {
			n := util.MustAtoi(num)
			seeds = append(seeds, n)
		}
	}

	// Blank line
	s.Scan()
	var maps []*node
	// Map label
	s.Scan()
	// Seed to soil
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Soil to fertilizer
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Fertilizer to water
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Water to light
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Light to temperature
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Temperature to humidity
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Humidity to location
	maps = append(maps, populateMapBST(s))

	minLocation := math.MaxInt
	for _, seed := range seeds {
		for _, m := range maps {
			seed = m.valFromBST(seed)
		}

		minLocation = min(minLocation, seed)
	}

	fmt.Printf("Answer Day Five, Part One: %v\n", minLocation)
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	seeds := make([]int, 0)
	s.Scan()
	_, nums, _ := strings.Cut(s.Text(), ":")
	for _, num := range strings.Split(nums, " ") {
		if num := strings.TrimSpace(num); num != "" {
			n := util.MustAtoi(num)
			seeds = append(seeds, n)
		}
	}

	// Blank line
	s.Scan()
	var maps []*node
	// Map label
	s.Scan()
	// Seed to soil
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Soil to fertilizer
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Fertilizer to water
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Water to light
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Light to temperature
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Temperature to humidity
	maps = append(maps, populateMapBST(s))
	// Map label
	s.Scan()
	// Humidity to location
	maps = append(maps, populateMapBST(s))

	minLocation := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			s := seed
			for _, m := range maps {
				s = m.valFromBST(s)
			}

			minLocation = min(minLocation, s)
		}
	}

	fmt.Printf("Answer Day Five, Part Two %v\n", minLocation)
}

type node struct {
	destStart, sourceStart, length int
	left, right                    *node
}

func populateMapBST(s *util.Scanner) *node {
	var n *node

	// Start
	s.Scan()
	for strings.TrimSpace(s.Text()) != "" {
		rnge := strings.Split(strings.TrimSpace(s.Text()), " ")
		sourceStart := util.MustAtoi(rnge[0])
		destStart := util.MustAtoi(rnge[1])
		length := util.MustAtoi(rnge[2])
		this := &node{
			destStart:   destStart,
			sourceStart: sourceStart,
			length:      length,
		}

		n = n.insert(this)

		s.Scan()
	}

	return n
}

func (n *node) insert(newNode *node) *node {
	if n == nil {
		return newNode
	}

	if newNode.destStart < n.destStart {
		n.left = n.left.insert(newNode)
	} else {
		n.right = n.right.insert(newNode)
	}

	return n
}

func (n *node) valFromBST(val int) int {
	if n == nil {
		return val
	}

	if val < n.destStart {
		return n.left.valFromBST(val)
	}
	if val >= n.destStart+n.length {
		return n.right.valFromBST(val)
	}

	return n.sourceStart + val - n.destStart
}
