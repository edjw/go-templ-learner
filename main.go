package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/edjw/go-templ-learner/components"
	"github.com/edjw/go-templ-learner/macServer"
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

	fmt.Print(templ.Classes()...)
	// Serve the app.
	index := components.Index()
	http.Handle("/", templ.Handler(index))

	about := components.About()
	http.Handle("/about", templ.Handler(about))

	macServer.Server()
}
