package util

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
)

func ReadInputFile(f embed.FS) (*Scanner, error) {
	file, err := f.Open("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}

	return &Scanner{
		Scanner: bufio.NewScanner(file),
		file:    file,
	}, nil
}

type Scanner struct {
	*bufio.Scanner
	file fs.File
}

func (s *Scanner) Close() error {
	return s.file.Close()
}
