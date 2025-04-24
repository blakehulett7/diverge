package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func compare_sections(pdf []byte) {
	parsed_sections := get_sections(pdf)

	fileBytes, _ := os.ReadFile("sections.json")
	possible_sections := []string{}
	json.Unmarshal(fileBytes, &possible_sections)

	fmt.Printf("possible: %v\n", possible_sections)
	fmt.Printf("parsed: %v\n", parsed_sections)
}
