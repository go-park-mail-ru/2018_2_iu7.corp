package main

import (
	"log"
	"net/http"
)

type Server struct {
	http.Server

	StaticPath string
	UploadPath string
}

func CreateServer(addr, staticPath, uploadPath string) *Server {
	server := &Server{
		StaticPath: staticPath,
		UploadPath: uploadPath,
	}
	server.Addr = addr
	server.Handler = createHandlers()
	return server
}

func createHandlers() http.Handler {
	handlers := make(map[string]http.Handler)

	handlers["register"] = POST(RegisterRequestHandler())
	handlers["login"] = POST(LoginRequestHandler())
	handlers["logout"] = POST(LogoutRequestHandler())
	handlers["profile"] = GETorPATCH(ProfileRequestHandler())
	handlers["leaderboard"] = GET(LeaderBoardRequestHandler())

	for k, v := range handlers {
		handlers[k] = withLogging(v)
	}

	mux := http.NewServeMux()
	mux.Handle("/register", handlers["register"])
	mux.Handle("/login", handlers["login"])
	mux.Handle("/logout", handlers["logout"])
	mux.Handle("/profile", handlers["profile"])
	mux.Handle("/leaderboard", handlers["leaderboard"])

	return mux
}

func GET(h http.Handler) http.Handler {
	return checkMethod(h, []string{http.MethodGet})
}

func POST(h http.Handler) http.Handler {
	return checkMethod(h, []string{http.MethodPost})
}

func GETorPATCH(h http.Handler) http.Handler {
	return checkMethod(h, []string{http.MethodGet, http.MethodPatch})
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
