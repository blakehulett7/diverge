package main

import "fmt"

func parse_r(pdf []byte, site_path string) {
	sections := get_sections(pdf)
	sections = append(sections, []string{"about", "author"}...)

	for _, section := range sections {
		fmt.Printf("generating %v yaml...\n", section)
		generator := section_getter_map[section]
		generator(pdf, site_path)
	}
}
