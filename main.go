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

	// Send HTTP dump to RealtimeLog.
	logIDs := r.URL.Query()["realtimelogid"]
	if len(logIDs) == 0 {
		return
	}
	logID := logIDs[0]
	if err := postToRealtimeLog(logID, string(dump)); err != nil {
		log.Printf("failed to post to realtimelog, error: %#v\n", err)
	}
}

// postToRealtimeLog sends dump to RealtimeLog.
func postToRealtimeLog(logID string, dump string) error {
	data := map[string]interface{}{"dump": dump}
	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data, error: %w", err)
	}

	url := realtimeLogBasePath + logID
	_, err = http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to send request, error: %w", err)
	}

	return nil
}
