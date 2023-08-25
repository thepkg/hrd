// Package main represents main entry point for HRD application.
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

const (
	// headerENV holds ENV header delimiter, just to separate output data.
	headerENV = "===============[     ENV      ]==============="
	// headerHTTP holds HTTP header delimiter, just to separate output data.
	headerHTTP = "===============[ HTTP REQUEST ]==============="
)

var (
	// AppEnv holds application environment name.
	AppEnv = ""
	// AppPort holds port which use to run application.
	AppPort = "8080"
	// AppWithEnv holds flag whether to dump environment variables.
	AppWithEnv = false
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
	withEnv := os.Getenv("APP_WITH_ENV")
	if strings.ToLower(withEnv) == "true" {
		AppWithEnv = true
	}

	// Init the only one HTTP handler.
	http.Handle("/", mainMiddleware(http.HandlerFunc(dumpHandler)))

	addr := fmt.Sprintf(":%s", AppPort)
	log.Printf("Starting server on: %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("Failed to ListenAndServe, error: %#v\n", err)
	}
}

// mainMiddleware represents main HTTP middleware.
func mainMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
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

	dump := ""
	// Add ENV data into dump.
	if AppWithEnv {
		dump += fmt.Sprintf("%s\n\n%s\n\n", headerENV, envDump)
	}
	// Add HTTP data into dump.
	dump += fmt.Sprintf("%s\n\n%s\n\n", headerHTTP, httpDump)

	// Log HTTP dump just for local environment.
	if AppEnv == "local" {
		log.Print(dump)
	}

	// Print dump to response.
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
