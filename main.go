package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ledongthuc/pdf"
)

func main() {
	fmt.Println("Jesus is Lord!")
	version := "v1"

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, r.URL.Path[1:]) })
	router.HandleFunc(fmt.Sprintf("POST /%v/upload", version), func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("upload endpoint hit")

		file, header, err := r.FormFile("resume")
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("filename: %v\n", header.Filename)

		echo, err := os.Create("echo/test.pdf")

		_, err = io.Copy(echo, file)
		echo.Close()

		file, reader, err := pdf.Open("echo/test.pdf")
		defer file.Close()
		var buffer bytes.Buffer
		b, err := reader.GetPlainText()

		buffer.ReadFrom(b)
		fmt.Println(buffer.String())

		http.Redirect(w, r, "/redirect.html", http.StatusSeeOther)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
