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
		for _, s := range strings.Split(s.Text(), ",") {
			ans += hash(s)
		}
	}

	fmt.Printf("Answer Day Fourteen, Part One: %v\n", ans)
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	boxes := make([][]lens, 256)
	for s.Scan() {
		for _, s := range strings.Split(s.Text(), ",") {
			if strings.HasSuffix(s, "-") {
				h := hash(s[:len(s)-1])
				for i, l := range boxes[h] {
					if l.label == s[:len(s)-1] {
						boxes[h] = append(boxes[h][:i], boxes[h][i+1:]...)
						break
					}
				}
			} else if label, focalLength, ok := strings.Cut(s, "="); ok {
				var found bool
				h := hash(label)
				for i, l := range boxes[h] {
					if l.label == label {
						boxes[h][i].focalLength = util.MustAtoi(focalLength)
						found = true
						break
					}
				}
				if !found {
					boxes[h] = append(boxes[h], lens{label: label, focalLength: util.MustAtoi(focalLength)})
				}
			}
		}
	}

	var ans int
	for b, box := range boxes {
		for p, l := range box {
			ans += (b + 1) * (p + 1) * l.focalLength
		}
	}

	fmt.Printf("Answer Day Fourteen, Part One: %v\n", ans)
}

func hash(s string) int {
	var current int
	for _, r := range s {
		current += int(r)
		current *= 17
		current %= 256
	}
	return current
}

type lens struct {
	label       string
	focalLength int
}
