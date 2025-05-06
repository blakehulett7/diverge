package main

import (
	"context"
	"encoding/json"
	"log"

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
                I have upload a resume that contains sections. The possible sections are as follows:
                [
                    "about",
                    "accomplishments",
                    "achievements",
                    "education",
                    "experiences",
                    "featured-posts",
                    "projects",
                    "publications",
                    "recent-posts",
                    "skills"
                ]
                Can you please list the sections that are present in the resume in a json array?
				Only use terms that are on the possible sections list and do not use any terms not present on the possible sections list.
                If no skills are listed on the resume, do not include "skills" on the return list

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
