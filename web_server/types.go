package main

import "net/http"

type HTTPMethod string

const (
	HTTPMethodCONNECT HTTPMethod = http.MethodConnect
	HTTPMethodDELETE  HTTPMethod = http.MethodDelete
	HTTPMethodGET     HTTPMethod = http.MethodGet
	HTTPMethodHEAD    HTTPMethod = http.MethodHead
	HTTPMethodOPTIONS HTTPMethod = http.MethodOptions
	HTTPMethodPATCH   HTTPMethod = http.MethodPatch
	HTTPMethodPOST    HTTPMethod = http.MethodPost
	HTTPMethodPUT     HTTPMethod = http.MethodPut
	HTTPMethodTRACE   HTTPMethod = http.MethodTrace
)

type EndpointPath string

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Name  string
	Email string
	Phone string
}
