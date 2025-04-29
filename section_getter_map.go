package main

var section_getter_map = map[string]func([]byte){
	"about":           generate_about_yaml,
	"accomplishments": nil,
	"achievements":    nil,
	"author":          generate_author_yaml,
	"education":       generate_education_yaml,
	"experiences":     generate_experiences_yaml,
	"featured-posts":  nil,
	"projects":        generate_projects_yaml,
	"publications":    generate_publications_yaml,
	"recent-posts":    nil,
	"skills":          generate_skills_yaml,
}
