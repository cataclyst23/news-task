//go:build ignore

package main

import (
	"log"
	"net/http"
	"text/template"

	client "github.com/cataclyst23/homework/pkg/client"
	"github.com/cataclyst23/homework/pkg/utils"
)

const templateDir = "tmpl"

func init() {
	utils.Templates = template.Must(template.ParseFiles(templateDir + "/display.html"))
}

func main() {
	// Serve CSS files
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", client.MakeHandler(client.ResponseHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
