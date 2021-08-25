package main

import (
	"log"
	"net/http"
)

type Server struct {
    port string
    router *Router
    middlewares []Middleware
}

func NewServer(port string, middlewares ...Middleware) *Server {
    return &Server{
        port: port,
        router: NewRouter(),
        middlewares: middlewares,
    }
}

func (s *Server) Handle(path EndpointPath, method HTTPMethod, handler http.HandlerFunc, middlewares ...Middleware) {
    _, exists := s.router.rules[path]

    if !exists {
        s.router.rules[path] = make(map[HTTPMethod]http.HandlerFunc)
    }
    allMiddlewares := append(s.middlewares, middlewares...)
    s.router.rules[path][method] = s.AddMiddleware(handler, allMiddlewares...)
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, middleware := range middlewares {
        f = middleware(f)
    }
    return f
}

func (s *Server) Listen() error {
    http.Handle("/", s.router)

    log.Printf("\033[0;32m%s%s\033[0m", "Server listening on http://localhost", s.port)
    err := http.ListenAndServe(s.port, nil)
    if (err != nil) {
        log.Fatalf("\033[0;31m%s", err)
        return err
    }
    return nil
}
