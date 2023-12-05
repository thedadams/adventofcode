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

	var ans int64
	for s.Scan() {
		var first, last int32
		for _, r := range s.Text() {
			if r >= '0' && r <= '9' {
				if first == 0 {
					first = r - '0'
				}
				last = r - '0'
			}
		}

		ans += int64(10*first + last)
	}

	fmt.Printf("Answer Day One, Part One: %v\n", ans)
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		return
	}

	var ans int64
	for s.Scan() {
		var first, last byte
		t := s.Text()
		for len(t) > 0 {
			if r := t[0]; r >= '1' && r <= '9' {
				if first == 0 {
					first = r - '0'
				}
				last = r - '0'
			}

			for key, val := range map[string]byte{
				"one":   1,
				"two":   2,
				"three": 3,
				"four":  4,
				"five":  5,
				"six":   6,
				"seven": 7,
				"eight": 8,
				"nine":  9,
			} {
				if strings.HasPrefix(t, key) {
					if first == 0 {
						first = val
					}
					last = val
					break
				}
			}

			t = t[1:]
		}

		ans += int64(10*first + last)
	}

	fmt.Printf("Answer Day One, Part Two: %v\n", ans)
}
