package main

import (
	"embed"
	"fmt"
	"regexp"
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

	fields := map[string]struct{}{
		"byr": {},
		"iyr": {},
		"eyr": {},
		"hgt": {},
		"hcl": {},
		"ecl": {},
		"pid": {},
	}

	var valid int
	var thisFields []string
	for s.Scan() {
		if s.Text() == "" {
			if len(thisFields) == len(fields) {
				valid++
			}
			thisFields = nil
			continue
		}

		for _, field := range strings.Split(s.Text(), " ") {
			f, _, _ := strings.Cut(field, ":")
			if _, ok := fields[f]; ok {
				thisFields = append(thisFields, f)
			}
		}
	}

	if len(thisFields) == len(fields) {
		valid++
	}

	fmt.Printf("Answer Day Four, Part One: %v\n", valid)
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	fields := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile(`^(19[2-9][0-9]|200[0-2])$`),
		"iyr": regexp.MustCompile(`^(201[0-9]|2020)$`),
		"eyr": regexp.MustCompile(`^(202[0-9]|2030)$`),
		"hgt": regexp.MustCompile(`^((1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)$`),
		"hcl": regexp.MustCompile(`^(#[0-9a-f]{6})$`),
		"ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
		"pid": regexp.MustCompile(`^([0-9]{9})$`),
	}

	var valid int
	var thisFields []string
	for s.Scan() {
		if s.Text() == "" {
			if len(thisFields) == len(fields) {
				valid++
			}
			thisFields = nil
			continue
		}

		for _, field := range strings.Split(s.Text(), " ") {
			f, val, _ := strings.Cut(field, ":")
			if reg := fields[f]; reg != nil && reg.MatchString(val) {
				thisFields = append(thisFields, f)
			}
		}
	}

	if len(thisFields) == len(fields) {
		valid++
	}

	fmt.Printf("Answer Day Four, Part Two: %v\n", valid)
}
