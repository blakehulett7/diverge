package main

import (
	"fmt"
	"io/fs"
	"os"

	"google.golang.org/genai"
)

func generate_experiences_yaml(pdf []byte) {
	response := ask_gemini(pdf, genai.NewPartFromText(`
		Can you please map the experiences section of the uploaded resume to the following yaml format:

        # section information
        section:
            name: Experiences
            id: experiences
            enable: true
            weight: 3
            showOnNavbar: true

        # Your experiences
        experiences:
        -   company:
                name: Example Co.
                url: "https://www.example.com"
                location: Dhaka Branch
                # company overview
                overview: Example Co. is a widely recognized company for cloud-native development. It builds tools for Kubernetes.
            positions:
            -   designation: Senior Software Engineer
                start: Nov 2019
                # don't provide end date if you are currently working there. It will be replaced by "Present"
                # end: Dec 2020
                # give some points about what was your responsibilities at the company.
                responsibilities:
                -   Design and develop XYZ tool for ABC task
                -   Design, develop and manage disaster recovery tool [Xtool](https://www.example.com) that backup Kubernetes volumes, databases, and cluster's resource definition.
                -   Lead backend team.

            -   designation: Junior Software Engineer
                start: Nov 2017
                end: Oct 2019
                responsibilities:
                -   Implement and test xyz feature for abc tool.
                -   Support client for abc tool.
                -   Learn k,d,w technology for xyz.

        -   company:
                name: PreExample Co.
                url: "https://www.example.com"
                location: Nowhere
                overview: PreExample Co. is a gateway company to enter into Example co. So, nothing special here.
            positions:
            -   designation: Software Engineer
                start: March 2016
                end: May 2017
                responsibilities:
                -   Write lots of example codes.
                -   Read lots of examples.
                -   See lots of example videos.

        -   company:
                name: Intern Counting Company (ICC).
                url: "https://www.example.com"
                location: Intern Land
                overview: Intern counting Company (ICC) is responsible for counting worldwide intern Engineers.
            positions:
            -   designation: Intern
                start: Jun 2015
                end: Jan 2016
                responsibilities:
                -   Count lost of interns.
                -   Count more interns.
                -   Count me as an intern.

        Here are some further instructions:
        - No other sections should be added to this yaml
        - Your response should contain no more than 80 lines
        - All lines should contain valid markdown.
	`))

	cleaned := clean_response(response)

	err := os.WriteFile("output.yaml", []byte(cleaned), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
