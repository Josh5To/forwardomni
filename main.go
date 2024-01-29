package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/josh5to/forwardomni/bones"
)

type DocumentData struct {
	Head bones.Head
}

type Page struct {
	Data         DocumentData
	PageTemplate *template.Template
}

func main() {
	page, err := createHomepage()
	if err != nil {
		log.Fatalf("Error parsing templates: %v\n", err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := page.PageTemplate.ExecuteTemplate(w, "base", page.Data); err != nil {
			fmt.Printf("unable to execute: %v\n", err)
		}
	})

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func createHead() *bones.Head {
	return &bones.Head{
		Meta: bones.Meta{
			Charset:  "utf-8",
			Viewport: "width=device-width, initial-scale=1",
		},
		Title:      "Testing a Site",
		Stylesheet: []string{"/static/stylesheets/main.css", "/static/stylesheets/mobile.css"},
	}
}

func createHomepage() (*Page, error) {
	htmldoc := DocumentData{
		Head: *createHead(),
	}

	// docBase := filepath.Join("./bones", "base.tmpl")
	// docHead := filepath.Join("/Users/jsmith/repos/tt-fn/web/bones", "head.tmpl")
	// docBody := filepath.Join("/Users/jsmith/repos/tt-fn/web/bones", "body.html")

	//Parse base template
	base, err := template.ParseFiles("base.tmpl")
	if err != nil {
		return nil, err
	}

	//Parse the rest of the templates
	page, err := base.ParseGlob(filepath.Join("/Users/jsmith/repos/forwardomni-go/bones", "*"))
	if err != nil {
		return nil, err
	}

	return &Page{
		Data:         htmldoc,
		PageTemplate: page,
	}, nil
}
