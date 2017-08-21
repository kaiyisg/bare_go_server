package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	initDB()
	runHTTPServerLoop()
}

func initDB() {
	// psqlDB, err := postgres.NewPostgresDB(&postgres.PostgresDBOpts{
	// 	Host:     webscandbHost,
	// 	User:     webscandbUser,
	// 	Password: webscandbPassword,
	// 	Database: webscandbDatabase,
	// })
}

func runHTTPServerLoop() {
	http.HandleFunc("/", handler)
	log.Println("HTTP server listening on :9009")
	err := http.ListenAndServe(":9009", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
