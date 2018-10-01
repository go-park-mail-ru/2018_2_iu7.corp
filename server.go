package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func CreateServer(addr, staticPath, uploadsPath string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: createHandlers(staticPath, uploadsPath),
	}
}

func createHandlers(staticPath, uploadsPath string) http.Handler {
	type RequestHandlerInfo struct {
		Methods    []string
		Handler    http.Handler
		Middleware []mux.MiddlewareFunc
	}

	handlers := make(map[string]RequestHandlerInfo)

	handlers["/auth/register"] = RequestHandlerInfo{
		Methods:    []string{http.MethodPost},
		Handler:    RegisterRequestHandler(),
		Middleware: []mux.MiddlewareFunc{NotAuthenticatedMiddleware, LoggingMiddleware},
	}
	handlers["/auth/login"] = RequestHandlerInfo{
		Methods:    []string{http.MethodPost},
		Handler:    LoginRequestHandler(),
		Middleware: []mux.MiddlewareFunc{NotAuthenticatedMiddleware, LoggingMiddleware},
	}
	handlers["/auth/logout"] = RequestHandlerInfo{
		Methods:    []string{http.MethodPost},
		Handler:    LogoutRequestHandler(),
		Middleware: []mux.MiddlewareFunc{AuthenticatedMiddleware, LoggingMiddleware},
	}
	handlers["/profiles/{id:[0-9]+}"] = RequestHandlerInfo{
		Methods:    []string{http.MethodGet},
		Handler:    ProfileRequestHandler(),
		Middleware: []mux.MiddlewareFunc{LoggingMiddleware},
	}
	handlers["/profiles/current"] = RequestHandlerInfo{
		Methods:    []string{http.MethodGet, http.MethodPut},
		Handler:    CurrentProfileRequestHandler(),
		Middleware: []mux.MiddlewareFunc{AuthenticatedMiddleware, LoggingMiddleware},
	}
	handlers["/profiles/leaderboard/pages/{page:[0-9]+}"] = RequestHandlerInfo{
		Methods:    []string{http.MethodGet},
		Handler:    LeaderBoardRequestHandler(),
		Middleware: []mux.MiddlewareFunc{LoggingMiddleware},
	}

	handlers["/static/"] = RequestHandlerInfo{
		Methods:    []string{http.MethodGet},
		Handler:    http.StripPrefix("/static", http.FileServer(http.Dir(staticPath))),
		Middleware: []mux.MiddlewareFunc{LoggingMiddleware},
	}
	handlers["/uploads/"] = RequestHandlerInfo{
		Methods:    []string{http.MethodGet},
		Handler:    http.StripPrefix("/uploads", http.FileServer(http.Dir(uploadsPath))),
		Middleware: []mux.MiddlewareFunc{LoggingMiddleware},
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
