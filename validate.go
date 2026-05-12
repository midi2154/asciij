package main

import "fmt"

func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char < 32 || char > 126 {
			return char, fmt.Errorf("Invalid character: %q", char)
		}
	}
	return 0, nil
}
