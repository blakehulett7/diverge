package main

import (
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	pdf, _ := os.ReadFile("tmp/test.pdf")
	parse_r(pdf)
}
