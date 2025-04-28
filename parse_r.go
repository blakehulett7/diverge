package main

import "fmt"

func parse_r(pdf []byte) {
	sections := get_sections(pdf)

	for _, section := range sections {
		fmt.Printf("generating %v yaml...", section)
		generator := section_getter_map[section]
		generator(pdf)
	}
}
