package main

import (
	"log"
	"net/http"
)

func CreateServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: createHandlers(),
	}
}

func createHandlers() http.Handler {
	handlers := make(map[string]http.Handler)

	handlers["register"] = POST(RegisterRequestHandler())
	handlers["login"] = POST(LoginRequestHandler())
	handlers["logout"] = POST(LogoutRequestHandler())

	for k, v := range handlers {
		handlers[k] = withLogging(v)
	}

	mux := http.NewServeMux()
	mux.Handle("/register", handlers["register"])
	mux.Handle("/login", handlers["login"])
	mux.Handle("/logout", handlers["logout"])

	return mux
}

func GET(h http.Handler) http.Handler {
	return checkMethod(h, []string{http.MethodGet})
}

func POST(h http.Handler) http.Handler {
	return checkMethod(h, []string{http.MethodPost})
}

func checkMethod(h http.Handler, allowed []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, method := range allowed {
			if r.Method == method {
				h.ServeHTTP(w, r)
				return
			}
		}
		http.NotFound(w, r)
	})
}

func withLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
