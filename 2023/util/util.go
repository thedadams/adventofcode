package util

import (
	"bufio"
	"embed"
	"fmt"
)

func ReadInputFile(f embed.FS) (*bufio.Scanner, error) {
	file, err := f.Open("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}

	return bufio.NewScanner(file), nil
}
