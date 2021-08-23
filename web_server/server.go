package main

import (
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

func (s *Server) Handle(path string, handler http.HandlerFunc, middlewares ...Middleware) {
    allMiddlewares := append(s.middlewares, middlewares...)
    s.router.rules[path] = s.AddMiddleware(handler, allMiddlewares...)
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, middleware := range middlewares {
        f = middleware(f)
    }
    return f
}

func (s *Server) Listen() error {
    http.Handle("/", s.router)

    err := http.ListenAndServe(s.port, nil)
    if (err != nil) {
        return err
    }
    return nil
}
