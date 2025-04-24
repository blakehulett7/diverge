package main

import "google.golang.org/genai"

func generate_author_yaml(pdf []byte) string {
	return ask_gemini(pdf, genai.NewPartFromText(`
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

        Please ensure that all lines contain valid markdown.
	`))
}
