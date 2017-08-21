package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	// log.Debug("HTTP server listening on :9009")
	err := http.ListenAndServe(":9009", nil)
	if err != nil {
		// log.Error(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
