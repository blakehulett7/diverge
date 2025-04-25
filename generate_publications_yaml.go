package main

import (
	"fmt"
	"io/fs"
	"os"

	"google.golang.org/genai"
)

func generate_publications_yaml(pdf []byte) {
	response := ask_gemini(pdf, genai.NewPartFromText(`
		Can you please map the publications section of the uploaded resume to the following yaml format:


        Here are some further instructions:
        - No other sections should be added to this yaml
        - For any icons, only font awesome icons should be used
        - Your response should contain no more than 80 lines
        - All lines should contain valid markdown.
	`))

	cleaned := clean_response(response)

	err := os.WriteFile("output.yaml", []byte(cleaned), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
