package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Jesus is Lord!")
	const version = "v1"

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, r.URL.Path[1:]) })
	router.HandleFunc(fmt.Sprintf("/%v/upload"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("upload endpoint hit")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
