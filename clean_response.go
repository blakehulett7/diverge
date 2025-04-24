package main

import (
	"fmt"
	"strings"
)

func clean_response(s string) string {
	lines := strings.Split(s, "\n")
	if len(lines) != 16 {
		fmt.Printf("invalid response... not 16 lines, instead %v\n", len(lines))
		fmt.Println(s)
		return "err"
	}

	cleaned_lines := []string{}
	for idx, line := range lines {
		if idx == 0 {
			continue
		}
		if idx == 16 {
			continue
		}

		cleaned_lines = append(cleaned_lines, line)
	}

	return strings.Join(cleaned_lines, "\n")
}
