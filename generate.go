package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	Splitinput := SplitInput(input)

	allEmpty := true
	for _, r := range Splitinput {
		if r != "" {
			allEmpty = false
			break
		}
	}
	if allEmpty {
		return strings.Repeat("\n", len(Splitinput)-1)
	}

	var result strings.Builder

	for i, lines := range Splitinput {
		if lines == "" && i < len(Splitinput)-1 { // handles intermediate newline (A\n\nB)
			result.WriteString("\n")
			continue
		}
		genlines := RenderLine(lines, banner)
		for _, line := range genlines {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}
	return result.String()
}
