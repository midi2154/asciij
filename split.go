package main

import "strings"

func SplitInput(input string) []string {
	splitinput := strings.Split(input, "\\n")
	return splitinput
}
