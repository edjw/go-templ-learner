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
	component := components.Layout()

	http.Handle("/", templ.Handler(component))

	macServer.Server()
}
