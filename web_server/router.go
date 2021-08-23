package main

import (
	"fmt"
	"net/http"
)

type Router struct {
    rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
    return &Router{
        rules: make(map[string]http.HandlerFunc),
    }
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    handler, exists := r.FindHandler(req.URL.Path)

    if !exists {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "Not found!")
        return
    }

    handler(w, req)
}

func (r *Router) FindHandler(path string) (handler http.HandlerFunc, exists bool) {
    handler, exists = r.rules[path]
    return
}
