package main

import (
	"fmt"
	"io/fs"
	"os"

	"google.golang.org/genai"
)

func generate_education_yaml(pdf []byte, site_path string) {
	response := ask_gemini(pdf, genai.NewPartFromText(`
		Can you please map the education section of the uploaded resume to the following yaml format:

        # section information
        section:
            name: Education
            id: education
            template: sections/education.html # Use "sections/education-alt.html for alternate template.
            enable: true
            weight: 4
            showOnNavbar: true

        degrees:
        -   name: Ph.D in Quantum Cryptography
            icon: fa-microscope
            timeframe: 2016-2020
            institution:
                name: ABC University of Technology
                url: "#"
            grade: #(optional)
                scale: CGPA
                achieved: 3.6
                outOf: 4
            publications: #(optional)
            -   title: Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                url: "#"
            -   title: Fusce eu augue ut odio porttitor pulvinar.
                url: "#"
            -   title: Nullam vitae orci tincidunt purus viverra pulvinar.
                url: "#"
        -   name: B.Sc. in Computer Science & Engineering
            icon: fa-graduation-cap
            timeframe: 2012-2016
            institution:
                name: University of XYZ
                url: "#"
            grade: #(optional)
                scale: CGPA
                achieved: 3.5
                outOf: 4
            takenCourses: #(optional)
                # if true, the courses will be rendered as a table otherwise it will render as a list and the grades will be hidden.
                showGrades: true
                collapseAfter: 3
                courses:
                -   name: Data Structures and Algorithm
                    achieved: 3.75
                    outOf: 4
                -   name: Network Security
                    achieved: 3.80
                    outOf: 4
                -   name: Operating System
                    achieved: 3.5
                    outOf: 4
                -   name: Artificial Intelligent
                    achieved: 3.75
                    outOf: 4
            publications: #(optional)
            -   title: Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                url: "#"
            -   title: Fusce eu augue ut odio porttitor pulvinar.
                url: "#"
            extracurricularActivities: #(optional)
            -   In vitae facilisis est, eget porta sem.
            -   Vestibulum consectetur lorem justo, at laoreet lorem feugiat et.
            -   Duis sed massa feugiat, ornare justo et, aliquam est.
            -   Pellentesque ut fringilla magna.
            customSections: #(optional)
            -   name: Thesis
                content: Lorem ipsum dolor sit amet, consectetur adipiscing elit.
            -   name: Supervisor
                content: Fusce eu augue ut odio porttitor pulvinar.
        -   name: Higher Secondary School Certificate
            icon: fa-university
            timeframe: 2010-2012
            institution:
                name: MST College of Science
                url: "#"
            grade: #(optional)
                scale: GPA
                achieved: 5
                outOf: 5
            extracurricularActivities: #(optional)
            -   In vitae facilisis est, eget porta sem.
            -   Vestibulum consectetur lorem justo, at laoreet lorem feugiat et.
            -   Duis sed massa feugiat, ornare justo et, aliquam est.
            -   Pellentesque ut fringilla magna.
        -   name: Secondary School Certificate
            icon: fa-school
            timeframe: 2005-2010
            institution:
                name: JK School of Science
            grade: #(optional)
                scale: GPA
                achieved: 4.5
                outOf: 5

        Here are some further instructions:
        - No other sections should be added to this yaml
        - For any icons, only font awesome icons should be used
        - All lines should contain valid markdown.
	`))

	cleaned := clean_response(response)

	file_path := fmt.Sprintf("%v/data/en/sections/education.yaml", site_path)
	err := os.WriteFile(file_path, []byte(cleaned), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
