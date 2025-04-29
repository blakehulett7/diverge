package main

import (
	"fmt"
	"io/fs"
	"os"

	"google.golang.org/genai"
)

func generate_about_yaml(pdf []byte) {
	response := ask_gemini(pdf, genai.NewPartFromText(`
		Can you please map the uploaded resume to the following yaml format:

        # section information
        section:
            name: About
            id: about
            enable: true
            weight: 1
            showOnNavbar: true
            template: sections/about.html

        # your designation
        designation: Software Engineer
        # your company information
        company:
            name: Example Co.
            url: "https://www.example.com"

        # a summary about you
        summary: 'I am a passionate software engineer with x years of working experience. I built OSS tools for [Kubernetes](https://kubernetes.io/) using GO. My tools help people to deploy their workloads in Kubernetes. Sometimes, I work on some fun projects such as writing a theme, etc.'

        # your social links
        # give as many as you want. use font-awesome for the icons.
        socialLinks:
        -   name: Email
            icon: "fas fa-envelope"
            url: "example@gmail.com"

        -   name: Github
            icon: "fab fa-github"
            url: "https://www.github.com/example"

        -   name: Stackoverflow
            icon: "fab fa-stack-overflow"
            url: "#"

        -   name: LinkedIn
            icon: "fab fa-linkedin"
            url: "#"

        -   name: Twitter
            icon: "fab fa-twitter"
            url: "#"

        -   name: Facebook
            icon: "fab fa-facebook"
            url: "#"

        -   name: Mastadon
            icon: "fab fa-mastodon"
            url: "#"
            rel: "me noopener"

        -   name: ResearchGate
            icon: "fab fa-researchgate"
            url: "https://www.researchgate.net/profile/john-doe"

        # Show your badges
        # You can show your verifiable certificates from https://www.credly.com.
        # You can also show a circular bar indicating the level of expertise on a certain skill
        badges:
        -   type: soft-skill-indicator
            name: Leadership
            percentage: 85
            color: blue

        -   type: soft-skill-indicator
            name: Team Work
            percentage: 90
            color: yellow

        -   type: soft-skill-indicator
            name: Hard Working
            percentage: 85
            color: orange

        Here are some further instructions:
        - No other sections should be added to this yaml
        - Your response should contain no more than 75 lines
        - All lines should contain valid markdown.
	`))

	cleaned := clean_response(response)

	err := os.WriteFile("tmp/about.yaml", []byte(cleaned), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
