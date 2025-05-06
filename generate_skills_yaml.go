package main

import (
	"fmt"
	"io/fs"
	"os"

	"google.golang.org/genai"
)

func generate_skills_yaml(pdf []byte, site_path string) {
	response := ask_gemini(pdf, genai.NewPartFromText(`
		Can you please map the skills section of the uploaded resume to the following yaml format:

        # section information
        section:
            name: Skills
            id: skills
            enable: true
            weight: 2
            showOnNavbar: true
            filter: true

        # filter buttons
        buttons:
        -   name: All
            filter: "all"
        -   name: Basic
            filter: "basic"
        -   name: Language
            filter: "language"
        -   name: Container
            filter: "container"
        -   name: Others
            filter: "others"

        # Your Skills.
        # Give a summary of you each skill in the summary section.
        skills:
        -   name: Kubernetes
            logo: /images/sections/skills/kubernetes.png
            summary: "Capable of deploying, managing application on Kubernetes. Experienced in writing Kubernetes controllers for CRDs."
            categories: ["container"]
            url: "https://kubernetes.io/"

        -   name: Go Development
            logo: /images/sections/skills/go.png
            summary: "Using as the main language for professional development. Capable of writing scalable, testable, and maintainable program."
            categories: ["basic", "language"]
            url: "https://golang.org/"

        -   name: Cloud Computing
            logo: /images/sections/skills/cloud.png
            summary: "Worked with most of the major clouds such as GCP, AWS, Azure etc."
            categories: ["others"]

        -   name: Docker
            logo: /images/sections/skills/docker.svg
            summary: "Write most of the programs as dockerized container. Experienced with multi-stage, multi-arch build process."
            categories: ["container"]
            url: "https://www.docker.com/"

        -   name: Prometheus
            logo: /images/sections/skills/prometheus.png
            summary: "Capable of setup, configure Prometheus metrics. Experienced with PromQL, AlertManager. Also, experienced with writing metric exporters."
            categories: ["basic"]
            url: "https://prometheus.io/"

        -   name: Linux
            logo: /images/sections/skills/linux.png
            summary: "Using as the main operating system. Capable of writing bash/shell scripts."
            categories: ["others"]

        -   name: Git
            logo: /images/sections/skills/git.png
            summary: "Experienced with git-based development. Mostly, use Github. Also, have experience in working with GitLab."
            categories: ["basic"]
            url: "https://git-scm.com/"

        -   name: C++
            logo: /images/sections/skills/c++.png
            summary: "Know basic C/C++ programming. Used for contest programming and problem solving."
            categories: ["basic", "language"]

        Here are some further instructions:
        - No other sections should be added to this yaml
        - In the skills section, please omit the logo unless it can be found in the following list:
            1. c++.png
            2. cloud.png
            3. docker.svg
            4. git.png
            5. go.png
            6. kubernetes.png
            7. linux.png
            8. prometheus.png
        - All lines should contain valid markdown.
	`))

	cleaned := clean_response(response)

	file_path := fmt.Sprintf("%v/data/en/sections/skills.yaml", site_path)
	err := os.WriteFile(file_path, []byte(cleaned), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
