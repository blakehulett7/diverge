package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func compare_sections() {
	fileBytes, _ := os.ReadFile("sections.json")
	sections := []string{}
	json.Unmarshal(fileBytes, &sections)

	fmt.Println(sections)
}
