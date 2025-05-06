package main

import (
	"fmt"
	"io/fs"
	"os"

	"google.golang.org/genai"
)

func generate_achievements_yaml(pdf []byte, site_path string) {
	response := ask_gemini(pdf, genai.NewPartFromText(`
		Can you please map the achievements section of the uploaded resume to the following yaml format:

		# section information
		section:
			name: Achievements
			id: achievements
			enable: true
			weight: 10
			showOnNavbar: true

		# Your achievements achievements
		achievements:
		-   title: Best Presenter
			image: /images/sections/achievements/presenter.jpg
			summary: Best presenter in the 2020 XYZ conference.
		-   title: Champion
			image: /images/sections/achievements/sport.jpg
			summary: Champion in cycling inter-city cycling championship 2020.
		-   title: Graduation
			image: /images/sections/achievements/graduation-cap.jpg
			summary: Received Bachelor of Science (B.Sc.) in Computer Science and Engineer from XYZ University.
		-   title: Award Winner
			image: /images/sections/achievements/woman-winner.jpg
			summary: Wined best paper award at IEE Conference 2020.

        Here are some further instructions:
        - No other sections should be added to this yaml
        - For any icons, only font awesome icons should be used
        - All lines should contain valid markdown.
	`))

	cleaned := clean_response(response)

	file_path := fmt.Sprintf("%v/data/en/sections/achievements.yaml", site_path)
	err := os.WriteFile(file_path, []byte(cleaned), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
