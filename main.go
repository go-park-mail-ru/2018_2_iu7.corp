package main

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"os"
	"sync"
)

const (
	DefaultAddress     = ":8080"
	DefaultStaticPath  = "./static/"
	DefaultUploadsPath = "./upload/"
)

var (
	sessionStore      sessions.Store
	profileRepository ProfileRepository
)

func main() {
	addr := os.Getenv("SERVER_ADDRESS")
	if addr == "" {
		addr = DefaultAddress
	}

	staticPath := os.Getenv("SERVER_STATIC_PATH")
	if staticPath == "" {
		staticPath = DefaultStaticPath
	}

	uploadPath := os.Getenv("SERVER_UPLOAD_PATH")
	if uploadPath == "" {
		uploadPath = DefaultUploadsPath
	}

	var sk []byte

	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		sk = securecookie.GenerateRandomKey(32)
	} else {
		sk = []byte(sessionKey)
	}

	srv := CreateServer(addr, staticPath, uploadPath)
	if srv == nil {
		log.Fatal("Server not started")
	}

	sessionStore = sessions.NewCookieStore(sk)
	if sessionStore == nil {
		log.Fatal("Session store not created")
	}

	profileRepository = NewInMemoryProfileRepository()
	if profileRepository == nil {
		log.Fatal("Profile repository not created")
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	var err error = nil
	go func() {
		err = srv.ListenAndServe()
	}()

	if err != nil {
		wg.Done()
		log.Fatal(err.Error())
	} else {
		log.Printf("Server started at %s", addr)
	}

	wg.Wait()
}
