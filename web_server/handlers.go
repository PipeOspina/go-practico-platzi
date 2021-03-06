package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func LoggerHandler(f http.HandlerFunc, w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	defer func() {
		log.Printf(
			"\033[0;32m%d\033[0m - %s %s %s",
			200,
			req.Method,
			req.URL.Path,
			time.Since(start),
		)
	}()
	f(w, req)
}

func LoggerHandlerError(
	w http.ResponseWriter,
	req *http.Request,
	statusCode int,
	err string,
	start time.Time,
) (http.ResponseWriter, *http.Request) {
	defer func() {
		log.Printf(
			"\033[0;31m%d - %s\033[0m - %s %s %s",
			statusCode,
			err,
			req.Method,
			req.URL.Path,
			time.Since(start),
		)
	}()
	w.WriteHeader(statusCode)
	fmt.Fprint(w, err)
	return w, req
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hellowis!")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the API Endpoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata Metadata
	err := decoder.Decode(&metadata)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}
