package main

import (
	"fmt"
	"io/fs"
	"os"

	"google.golang.org/genai"
)

func generate_publications_yaml(pdf []byte) {
	response := ask_gemini(pdf, genai.NewPartFromText(`
		Can you please map the publications section of the uploaded resume to the following yaml format:

        # section information
        section:
            name: Publications
            id: publications
            enable: true
            weight: 6
            showOnNavbar: true

        # filter buttons
        buttons:
        -   name: All
            filter: "all"
        -   name: "Machine Learning"
            filter: "machinelearning"
        -   name: "Image Processing"
            filter: "image-processing"
        -   name: Security
            filter: "security"

        # your publications
        publications:
        -   title: An Example Paper on Machine Learning
            publishedIn:
                name: 2020 IEEE Region Symposium (TENSYMP)
                date: 7 June 2020
                url: https://example.com
            authors:
            -   name: Dr. Madman
                url: https://example.com
            -   name: Dr. Lessmad
                url: https://example.com
            -   name: Dr. Moremad
                url: https://example.com
            -   name: Dr. Goodman
                url: https://example.com
			paper:
				summary: Voluptate in id id voluptate laboris. Minim mollit aliquip sit aliqua ut exercitation voluptate eiusmod consequat pariatur sunt enim veniam. Velit esse tempor laboris anim tempor officia. Magna non labore duis do esse sit do ipsum culpa. Officia consequat id non duis culpa dolor. Excepteur magna non nostrud cupidatat aute aliqua aliquip.
				url: https://example.com
			categories: ["machinelearning","image-processing"]
			tags: ["Machine Learning", "Autonomous Driving", "Computer Vision"]

        -   title: An Sample Paper on Image Processing
			publishedIn:
				name: 2020 IEEE Region Symposium (TENSYMP)
				date: 7 June 2020
				url: https://example.com
        	authors:
			-   name: Dr. Madman
				url: https://example.com
			-   name: Dr. Lessmad
				url: https://example.com
			-   name: Dr. Moremad
				url: https://example.com
			-   name: Dr. Goodman
				url: https://example.com
			paper:
				summary: Ullamco magna minim cupidatat Lorem ea ex aliqua fugiat et. Dolor quis cillum ea duis irure et commodo aliquip consectetur ullamco labore ut anim nisi. Commodo reprehenderit est consectetur tempor adipisicing occaecat exercitation amet do aliquip dolor do irure. Labore officia ut magna pariatur reprehenderit et ex sit sunt. Magna proident ullamco adipisicing sit.
				url: https://example.com
			categories: ["image-processing"]
			tags: ["Image Processing", "Computer Vision"]

        -   title: An Example Paper on Security
			publishedIn:
				name: 2020 IEEE Region Symposium (TENSYMP)
				date: 7 June 2020
				url: https://example.com
			authors:
			-   name: Dr. Madman
				url: https://example.com
			-   name: Dr. Lessmad
				url: https://example.com
			-   name: Dr. Moremad
				url: https://example.com
			-   name: Dr. Goodman
				url: https://example.com
			paper:
				summary: Dolor ad cupidatat pariatur nulla ipsum ex ullamco nisi anim ullamco. Dolore elit esse in exercitation minim aliqua. Amet velit incididunt magna laboris. Proident reprehenderit deserunt ad officia duis dolor. Tempor nostrud ullamco ullamco sint deserunt cupidatat irure minim consectetur nulla adipisicing sunt.
				url: https://example.com
			categories: ["security"]
			tags: ["Machine Learning", "Autonomous Driving", "Machine Vision"]

        Here are some further instructions:
        - No other sections should be added to this yaml
        - For any icons, only font awesome icons should be used
        - All lines should contain valid markdown.
	`))

	cleaned := clean_response(response)

	err := os.WriteFile("output.yaml", []byte(cleaned), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
