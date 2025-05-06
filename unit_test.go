package main

import (
	"os"
	"testing"
)

func TestSiteGenerator(t *testing.T) {
	pdf, _ := os.ReadFile("tmp/test.pdf")
	create_site(pdf)
}
