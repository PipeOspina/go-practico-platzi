package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
    return func(f http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, req *http.Request) {
            flag := true
            fmt.Println("Checking Authentication")
            if flag {
                f(w, req)
            } else {
                return
            }
        }
    }
}

func Logger() Middleware {
    return func(f http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, req *http.Request) {
            start := time.Now()
            defer func() {
                log.Println(req.Method, req.URL.Path, time.Since(start))
            }()
            f(w, req)
        }
    }
}
