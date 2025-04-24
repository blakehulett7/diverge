package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGemini(t *testing.T) {
	pdf, _ := os.ReadFile("tmp/test.pdf")
	fmt.Println(generate_author_yaml(pdf))
}
