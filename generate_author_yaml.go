package main

import (
	"fmt"
	"os"

	"google.golang.org/genai"
)

func generate_author_yaml(pdf []byte) {
	response := ask_gemini(pdf, genai.NewPartFromText(`
		Can you please map the uploaded resume to the following yaml format:

        # some information about you
        name: "John Doe"
        nickname: "John"
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

        Please ensure that no other sections are added to this yaml and that all lines contain valid markdown.
	`))

	fmt.Println(response)

	cleaned := clean_response(response)
	if cleaned == "err" {
		fmt.Println("this is where I will try 4 more times")
	}

	os.WriteFile("output.yaml", []byte(cleaned), 777)
}
