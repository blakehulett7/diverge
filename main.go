package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Jesus is Lord!")

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "../app/index.html") })

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
