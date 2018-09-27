package main

import "net/http"

func CreateServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: createHandler(),
	}
}

func createHandler() http.Handler {
	return nil //TODO
}
