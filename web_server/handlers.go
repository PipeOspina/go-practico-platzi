package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hellowis!")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "This is the API Endpoint")
}
