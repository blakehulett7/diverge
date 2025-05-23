package main

import (
	"fmt"
	"io/fs"
	"os"

	"google.golang.org/genai"
)

func generate_projects_yaml(pdf []byte, site_path string) {
	response := ask_gemini(pdf, genai.NewPartFromText(`
		Can you please map the projects section of the uploaded resume to the following yaml format:

        # section information
        section:
            name: Projects
            id: projects
            enable: true
            weight: 5
            showOnNavbar: true

        # filter buttons
        buttons:
        -   name: All
            filter: "all"
        -   name: Professional
            filter: "professional"
        -   name: Academic
            filter: "academic"
        -   name: Hobby
            filter: "hobby"

        # your projects
        projects:
        -   name: Kubernetes
            logo: /images/sections/projects/kubernetes.png
            role: Contributor
            timeline: "March 2018 - Present"
            repo: https://github.com/kubernetes/kubernetes # If your project is a public repo on GitHub, then provide this link. it will show star count.
            #url: ""  # If your project is not a public repo but it has a website or any external details url then provide it here. don't provide "repo" and "url" simultaneously.
            summary: Production-Grade Container Scheduling and Management.
            tags: ["professional", "kubernetes", "cloud"]

        -   name: Tensorflow
            logo: /images/sections/projects/tensorflow.png
            role: Developer
            timeline: "Jun 2018 - Present"
            repo: https://github.com/tensorflow/tensorflow
            #url: ""
            summary: An Open Source Machine Learning Framework for Everyone.
            tags: ["professional", "machine-learning"]

        -   name: A sample academic paper
            role: Team Lead
            timeline: "Jan 2017 - Nov 2017"
            url: "https://www.example.com"
            summary: Lorem ipsum dolor sit amet consectetur adipisicing elit. Sapiente eius reprehenderit animi suscipit autem eligendi esse amet aliquid error eum. Accusantium distinctio soluta aliquid quas placeat modi suscipit eligendi nisi.
            tags: ["academic","iot"]

        -   name: Nocode
            logo: /images/sections/projects/no-code.png
            role: Nothing
            timeline: "Oct 2019 - Dec 2019"
            repo: https://github.com/kelseyhightower/nocode
            #url: ""
            summary: The best way to write secure and reliable applications. Write nothing; deploy nowhere.
            tags: ["hobby", "fun"]

        -   name: Toha
            logo: /images/sections/projects/toha.png
            image: /images/sections/projects/toha_website.png
            role: Owner
            timeline: "Jun 2019 - Present"
            repo: https://github.com/hossainemruz/toha
            summary: A Hugo theme for personal portfolio.
            tags: ["hobby","hugo","theme","professional"]

        Here are some further instructions:
        - No other sections should be added to this yaml
        - All lines should contain valid markdown.
	`))

	cleaned := clean_response(response)

	file_path := fmt.Sprintf("%v/data/en/sections/projects.yaml", site_path)
	err := os.WriteFile(file_path, []byte(cleaned), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
