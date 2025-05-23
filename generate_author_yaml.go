package main

import (
	"fmt"
	"io/fs"
	"os"

	"google.golang.org/genai"
)

func generate_author_yaml(pdf []byte, site_path string) {
	response := ask_gemini(pdf, genai.NewPartFromText(`
		Can you please map the uploaded resume to the following yaml format:

        # some information about you
        name: "John Doe"
        nickname: "John"
        # greeting message before your name. it will default to "Hi! I am" if not provided
        greeting: "Hi, I am"
        image: "images/author/john.png"
        # give your some contact information. they will be used in the footer
        contactInfo:
            email: "johndoe@example.com"
            phone: "+0123456789"
            github: johndoe
            linkedin: johndoe
            researchgate: john-doe

        # some summary about what you do
        summary:
            - I am a Developer
            - I am a Devops
            - I love servers
            - I work on open-source projects
            - I love to work with some fun projects

        Here are some further instructions:
        - No other sections should be added to this yaml
        - The image line should not change. It should remain as follows, image: "images/author/john.png"
        - Make sure to include the greeting line with the greeting "Hi, I am"
        - Each summary line should be in a brief statement that begins with I...
        - All lines contain valid markdown.
	`))

	cleaned := clean_response(response)

	file_path := fmt.Sprintf("%v/data/en/author.yaml", site_path)
	err := os.WriteFile(file_path, []byte(cleaned), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
