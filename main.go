package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

const (
	// header holds header delimiter, just to separate output data.
	header = "===============[ HTTP REQUEST ]==============="
)

var (
	// AppEnv holds application environment name.
	AppEnv = ""
	// AppPort holds port which use to run application.
	AppPort = "8080"
)

func main() {
	runApp()
}

// runApp starts application.
func runApp() {
	env := os.Getenv("APP_ENV")
	if env != "" {
		AppEnv = env
	}
	port := os.Getenv("APP_PORT")
	if port != "" {
		AppPort = port
	}

	// Init the only one HTTP handler.
	http.HandleFunc("/", dumpHandler)

	addr := fmt.Sprintf(":%s", AppPort)
	log.Printf("starting server on: %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("failed to ListenAndServe, error: %#v\n", err)
	}
}

// dumpHandler represents main HTTP handler.
func dumpHandler(w http.ResponseWriter, r *http.Request) {
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
