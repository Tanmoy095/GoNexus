package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// Define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// Start http server
	log.Printf("Server configured to listen on %s\n", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panicf("Error starting server: %v\n", err)
	}
}
