// This is a simple server to help development on a Mac. It also works in production on Fly and Render.
// It works around the firewall issue on Macs where you have to manually allow incoming connections to your Go app on every rebuild.
package macServer

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Server() {
	port, portExists := os.LookupEnv("PORT")
	env, envExists := os.LookupEnv("GOENVIRONMENT")

	var ip string = "0.0.0.0" // Standard IP address on Fly and Render

	if envExists && env == "development" {
		ip = "127.0.0.1"
	}
	// For development, this relies on setting the GOENVIRONMENT variable
	// to "development" in your .zshrc or .bashrc file with
	// export GOENVIRONMENT=development

	if !portExists {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ip + ":" + port,
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("\n\n😀😀👏👏👏\nListening on:\nhttp://%s:%v\n\n", ip, port)

	// Except it'll actually be http://localhost:3000 if we're using browser-sync

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
