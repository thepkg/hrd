// Package main represents main entry point for HRD application.
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

const (
	// header1 holds ENV header delimiter, just to separate output data.
	header1 = "===============[     ENV      ]==============="
	// header2 holds HTTP header delimiter, just to separate output data.
	header2 = "===============[ HTTP REQUEST ]==============="
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
	// Get ENV dump for current environment.
	envDump := dumpEnv()

	// Get HTTP dump for current request.
	httpDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	dump := fmt.Sprintf("%s \n\n %s \n\n %s \n\n %s", header1, envDump, header2, httpDump)

	// Log HTTP dump just for local environment.
	if AppEnv == "local" {
		log.Print(dump)
	}

	// Print HTTP dump to response.
	if _, err := fmt.Fprint(w, dump); err != nil {
		log.Printf("failed to print header, error: %#v\n", err)
	}
}

// dumpEnv returns ENV dump.
func dumpEnv() string {
	dump := ""

	for _, e := range os.Environ() {
		dump += e + "\n"
	}

	return dump
}
