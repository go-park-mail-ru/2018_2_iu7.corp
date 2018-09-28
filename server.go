package main

import (
	"log"
	"net/http"
)

func CreateServer(addr, staticPath, uploadsPath string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: createHandlers(staticPath, uploadsPath),
	}
}

func createHandlers(staticPath, uploadsPath string) http.Handler {
	reqHandlers := make(map[string]http.Handler)
	reqHandlers["register"] = POST(RegisterRequestHandler())
	reqHandlers["login"] = POST(LoginRequestHandler())
	reqHandlers["logout"] = POST(LogoutRequestHandler())
	reqHandlers["profile"] = GETorPATCH(ProfileRequestHandler(uploadsPath))
	reqHandlers["leaderboard"] = GET(LeaderBoardRequestHandler())

	for k, v := range reqHandlers {
		reqHandlers[k] = withLogging(v)
	}

	fileHandlers := make(map[string]http.Handler)
	fileHandlers["static"] = http.FileServer(http.Dir(staticPath))
	fileHandlers["uploads"] = http.FileServer(http.Dir(uploadsPath))

	mux := http.NewServeMux()

	mux.Handle("/register", reqHandlers["register"])
	mux.Handle("/login", reqHandlers["login"])
	mux.Handle("/logout", reqHandlers["logout"])
	mux.Handle("/profile", reqHandlers["profile"])
	mux.Handle("/leaderboard", reqHandlers["leaderboard"])

	mux.Handle("/static/", http.StripPrefix("/static", fileHandlers["static"]))
	mux.Handle("/uploads/", http.StripPrefix("/uploads", fileHandlers["uploads"]))

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
