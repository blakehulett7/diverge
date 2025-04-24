package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/genai"
)

func prompting(pdf []byte) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	prompt := genai.NewContentFromParts(
		[]*genai.Part{
			genai.NewPartFromBytes(pdf, "application/pdf"),
			genai.NewPartFromText(`
                I have upload a resume that contains sections. Can you please list the sections?
            `),
		},
		"",
	)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		[]*genai.Content{prompt},
		nil,
	)

	fmt.Println(result.Text())
}
