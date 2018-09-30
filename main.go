package main

import (
	"2018_2_iu7.corp/handlers"
	"2018_2_iu7.corp/profiles"
	"2018_2_iu7.corp/sessions"
	"log"
	"os"
	"sync"
)

const (
	DefaultAddress     = ":8080"
	DefaultStaticPath  = "./static/"
	DefaultUploadsPath = "./upload/"
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

	srv := handlers.CreateServer(addr, staticPath, uploadPath)
	if srv == nil {
		log.Fatal("Server not started")
	}

	sessionStorage := sessions.NewInMemorySessionStorage()
	if sessionStorage == nil {
		log.Fatal("Session storage not created")
	}

	profileRepository := profiles.NewInMemoryProfileRepository()
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
