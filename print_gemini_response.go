package main

import (
	"fmt"

	"google.golang.org/genai"
)

func print_gemini_response(response *genai.GenerateContentResponse) {
	for _, cand := range response.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part.Text)
			}
		}
	}
}
