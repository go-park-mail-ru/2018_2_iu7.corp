package main

import (
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
	sessionStorage    SessionStorage
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

	srv := CreateServer(addr, staticPath, uploadPath)
	if srv == nil {
		log.Fatal("Server not started")
	}

	sessionStorage = NewInMemorySessionStorage()
	if sessionStorage == nil {
		log.Fatal("Session storage not created")
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
