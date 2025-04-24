package main

import (
	"context"
	"log"

	"google.golang.org/genai"
)

func ask_gemini(pdf []byte, prompt_text *genai.Part) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	prompt := genai.NewContentFromParts(
		[]*genai.Part{
			genai.NewPartFromBytes(pdf, "application/pdf"),
			prompt_text,
		},
		"",
	)

	response, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		[]*genai.Content{prompt},
		nil,
	)

	return response.Text()
}
