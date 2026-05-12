package main

import "strings"

func SplitInput(text string) []string {
	return strings.Split(text, "\n")
}

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	if strings.Trim(input, `\n`) == `` {
		strings.ReplaceAll(input, `\n`, "\n")
	}

	Splitinput := SplitInput(input)

	var result strings.Builder

	for i, lines := range Splitinput {
		if lines == "" {
			if i == len(Splitinput)-1 {
				continue
			}
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
