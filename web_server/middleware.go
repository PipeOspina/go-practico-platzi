package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	start := time.Now()
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			flag := true
			if flag {
				f(w, req)
			} else {
				LoggerHandlerError(w, req, http.StatusUnauthorized, "Unauthorized!", start)
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

func JsonDecoder() Middleware {
	start := time.Now()
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			contentType := req.Header.Get("Content-Type")
			if req.Header.Get("Content-Type") != "application/json" {
				err := fmt.Sprintf(
					"Bad \"Content-Type\" header, expected \"application/json\", but found \"%s\"",
					contentType,
				)
				LoggerHandlerError(w, req, http.StatusBadRequest, err, start)
				return
			}

			buf, _ := ioutil.ReadAll(req.Body)

			body := ioutil.NopCloser(bytes.NewBuffer(buf))
			req.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

			decoder := json.NewDecoder(body)

			var metadata Metadata
			err := decoder.Decode(&metadata)
			if err != nil {
				errValue := fmt.Sprintf("Bad request! - %v", err.Error())
				LoggerHandlerError(w, req, http.StatusBadRequest, errValue, start)
				return
			}

			body.Close()

			f(w, req)
		}
	}
}
