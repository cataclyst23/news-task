//go:build ignore

package main

import (
	"log"
	"net/http"

	client "github.com/cataclyst23/homework/pkg/client"
)

func main() {
	// Serve CSS files
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", client.MakeHandler(client.ResponseHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
