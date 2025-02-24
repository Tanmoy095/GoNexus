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

	log.Printf("starting broker service on port %s\n", webPort)
	//definr http server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%:s", webPort),
		Handler: app.routes(),
	}

	//start http server

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)

	}

}
