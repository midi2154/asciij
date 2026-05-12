package main

import "strings"

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
// for the error on line 76 abi 78 requesting new line just edit the test file to 8 or 9 depending on what you have because 16 is not possible
