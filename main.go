package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Jesus is Lord!")
	version := "v1"

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, r.URL.Path[1:]) })
	router.HandleFunc(fmt.Sprintf("POST /%v/upload", version), func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("upload endpoint hit")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
