package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

const (
	realtimeLogBasePath = "https://realtimelog.herokuapp.com:443/"
	header              = "===============[ HTTP REQUEST DUMP ]==============="
)

var (
	AppEnv  = ""
	AppPort = "8080"
)

func main() {
	cli()
}

// cli starts application from CLI environment.
func cli() {
	env := os.Getenv("APP_ENV")
	if env != "" {
		AppEnv = env
	}
	port := os.Getenv("APP_PORT")
	if port != "" {
		AppPort = port
	}

	// Init the only one HTTP handler.
	http.HandleFunc("/", Dump)

	addr := fmt.Sprintf(":%s", AppPort)
	log.Printf("starting server on: %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("failed to ListenAndServe, error: %#v\n", err)
	}
}

// Dump represents main HTTP handler.
func Dump(w http.ResponseWriter, r *http.Request) {
	// Get HTTP dump for current request.
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	// Log HTTP dump just for local environment.
	if AppEnv == "local" {
		log.Printf("[%s] \n\n%s", header, dump)
	}

	// Print HTTP dump to response.
	if _, err := fmt.Fprintf(w, "%s \n\n%s", header, dump); err != nil {
		log.Printf("failed to print header, error: %#v\n", err)
	}
}
