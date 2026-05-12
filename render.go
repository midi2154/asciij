package main

import "strings"

func RenderLine(text string, banner map[rune][]string) []string {
	ouput := []string{}

	for row := 0; row < 8; row++ {
		var result strings.Builder
		for _, char := range text {

			line := banner[char][row]
			result.WriteString(line)
		}
		ouput = append(ouput, result.String())
	}
	return ouput
}
