package main

import (
	"strings"
)

func clean_response(s string) string {
	lines := strings.Split(s, "\n")

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
