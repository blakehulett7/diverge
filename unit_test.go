package main

import (
	"os"
	"testing"
)

func TestGemini(t *testing.T) {
	pdf, _ := os.ReadFile("tmp/test.pdf")
	parse_r(pdf)
}
