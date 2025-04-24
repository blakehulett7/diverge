package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handle_file_upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload endpoint hit")

	file, header, err := r.FormFile("resume")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("filename: %v\n", header.Filename)

	echo, err := os.Create(fmt.Sprintf("tmp/%v.pdf", header.Filename))
	defer echo.Close()
	_, err = io.Copy(echo, file)

	http.Redirect(w, r, "/redirect.html", http.StatusSeeOther)
}
