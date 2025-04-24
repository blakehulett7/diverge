package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func get_sections(pdf []byte) []string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	prompt := genai.NewContentFromParts(
		[]*genai.Part{
			genai.NewPartFromBytes(pdf, "application/pdf"),
			genai.NewPartFromText(`
                I have upload a resume that contains sections. Can you please list the sections in a json array?
                return Array<section>
            `),
		},
		"",
	)

	config := genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
	}

	response, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		[]*genai.Content{prompt},
		&config,
	)

	section_list := []string{}
	json.Unmarshal([]byte(response.Text()), &section_list)

	return section_list
}

	fileBytes, _ := os.ReadFile("sections.json")
	sections := []string{}
	json.Unmarshal(fileBytes, &sections)

	fmt.Println(sections)
}
