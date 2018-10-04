package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func CreateServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: createHandlers(),
	}
}

func createHandlers() http.Handler {
	type RequestHandlerInfo struct {
		Methods    []string
		Handler    http.Handler
		Middleware []mux.MiddlewareFunc
	}

	handlers := make(map[string]RequestHandlerInfo)

	handlers["/new"] = RequestHandlerInfo{
		Methods: []string{http.MethodPost, http.MethodOptions},
		//TODO Handler,
		//TODO Middleware
	}
	handlers["/{profileID:[0-9]+}"] = RequestHandlerInfo{
		Methods: []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodOptions},
		//TODO Handler,
		//TODO Middleware
	}
	handlers["/leaders"] = RequestHandlerInfo{
		Methods: []string{http.MethodGet, http.MethodOptions},
		//TODO Handler,
		//TODO Middleware
	}

	for p, h := range handlers {
		for _, m := range h.Middleware {
			h.Handler = m(h.Handler)
		}
		handlers[p] = h
	}

	r := mux.NewRouter()
	for p, h := range handlers {
		r.Handle(p, h.Handler).Methods(h.Methods...)
	}

	return r
}
