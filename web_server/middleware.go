package main

import (
	"net/http"
)

func CheckAuth() Middleware {
    return func(f http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, req *http.Request) {
            flag := true
            if flag {
                f(w, req)
            } else {
                LoggerHandlerError(w, req, http.StatusUnauthorized, "Unauthorized!")
                return
            }
        }
    }
}

func Logger() Middleware {
    return func(f http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, req *http.Request) {
            LoggerHandler(f, w, req)
        }
    }
}
