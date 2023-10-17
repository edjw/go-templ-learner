package main

import (
	"embed"
	"go-templ-learner/components"
	"go-templ-learner/macServer"
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

//go:embed public/*
var public embed.FS

func main() {

	// Serve the public folder.
	publicFS, err := fs.Sub(public, "public")
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.FS(publicFS))
	http.Handle("/public/", http.StripPrefix("/public", fileServer))

	// Serve the app.
	index := components.Index()
	http.Handle("/", templ.Handler(index))

	about := components.About()
	http.Handle("/about", templ.Handler(about))

	macServer.Server()
}
