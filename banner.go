package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("Error reading file")
	}

	if len(file) == 0 {
		return nil, errors.New("Empty file")
	}

	content := strings.ReplaceAll(string(file), "\r\n", "\n")

	splitcontent := strings.Split(content, "\n")

	mapline := make(map[rune][]string)

	for j := 32; j < 127; j++ {
		char := rune(j)
		start := (j - 32) * 9

		if start+9 > len(splitcontent) {
			return nil, errors.New("Invalid banner format")
		}

		completeline := splitcontent[start+1 : start+9]
		if len(completeline) != 8 {
			return nil, errors.New("INVALID CHARACTER BLOCK")
		}

		mapline[char] = completeline
	}
	return mapline, nil
}
