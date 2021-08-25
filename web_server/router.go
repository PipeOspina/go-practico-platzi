package main

import (
	"net/http"
)

type Router struct {
	rules map[EndpointPath]map[HTTPMethod]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[EndpointPath]map[HTTPMethod]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, pathExists, methodExists := r.FindHandler(EndpointPath(req.URL.Path), HTTPMethod(req.Method))

	if !pathExists {
		LoggerHandlerError(w, req, http.StatusNotFound, "Not found!")
		return
	}

	if !methodExists {
		LoggerHandlerError(w, req, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	handler(w, req)
}

func (r *Router) FindHandler(path EndpointPath, method HTTPMethod) (handler http.HandlerFunc, pathExists bool, methodExists bool) {
	methodsAllowed, pathExists := r.rules[path]

	if !pathExists {
		return nil, false, false
	}

	handler, methodExists = methodsAllowed[method]
	return
}
