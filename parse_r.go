package main

func parse_r(pdf []byte) {
	sections := get_sections(pdf)

	for _, section := range sections {
		generator := section_getter_map[section]
		generator(pdf)
	}
}
