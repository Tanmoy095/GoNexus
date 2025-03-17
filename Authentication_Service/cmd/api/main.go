package main

import (
	"Authentication_Service/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const webPort = "80"

var counts int64

type Config struct {
	// Add any necessary configuration settings here.
	// For example, database connection strings or API keys.
	Db     *sql.DB
	Models data.Models
}

func main() {
	fmt.Printf("Starting broker service on port %s\n", webPort)

	// Connect to the database
	connection := connectToDB()
	if connection == nil {
		log.Fatal("Failed to connect to the postgres database.")

	}

	//set up the config /.....
	app := Config{
		Db:     connection,
		Models: data.NewModels(connection),
	}
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
func OpenDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		conn, err := OpenDb(dsn)
		if err != nil {
			log.Println("postgres not yet ready...")
			counts++

		} else {
			log.Println("connected to postgres .......")
			return conn

		}
		if counts > 10 {
			log.Println("ERROR:")
			return nil
		}
		log.Println("Backing off for 2 seconds...")
		time.Sleep(time.Second * 2)
		continue

	}

}
