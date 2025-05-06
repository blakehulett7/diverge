package main

var section_getter_map = map[string]func([]byte, string){
	"about":           generate_about_yaml,
	"accomplishments": generate_accomplishments_yaml,
	"achievements":    generate_achievements_yaml,
	"author":          generate_author_yaml,
	"education":       generate_education_yaml,
	"experiences":     generate_experiences_yaml,
	"projects":        generate_projects_yaml,
	"publications":    generate_publications_yaml,
	"skills":          generate_skills_yaml,
}
