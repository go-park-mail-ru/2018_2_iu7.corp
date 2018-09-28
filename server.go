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
	handlers := make(map[string]http.Handler)

	handlers["/register"] = POST(RegisterRequestHandler())
	handlers["/login"] = POST(LoginRequestHandler())
	handlers["/logout"] = POST(LogoutRequestHandler())
	handlers["/profile"] = GETorPATCH(ProfileRequestHandler(uploadsPath))
	handlers["/leaderboard"] = GET(LeaderBoardRequestHandler())

	handlers["/static/"] = http.StripPrefix("/static", http.FileServer(http.Dir(staticPath)))
	handlers["/uploads/"] = http.StripPrefix("/uploads", http.FileServer(http.Dir(uploadsPath)))

	for k, v := range handlers {
		handlers[k] = withLogging(v)
	}

	mux := http.NewServeMux()
	for k, v := range handlers {
		mux.Handle(k, v)
	}

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
