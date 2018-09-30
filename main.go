package main

import (
	"2018_2_iu7.corp/handlers"
	"2018_2_iu7.corp/profiles"
	"2018_2_iu7.corp/sessions"
	"log"
	"os"
)

const (
	DefaultAddress     = ":8080"
	DefaultStaticPath  = "./static/"
	DefaultUploadsPath = "./upload/"
)

func main() {
	config := handlers.ServerConfig{}

	config.Address = os.Getenv("SERVER_ADDRESS")
	if config.Address == "" {
		config.Address = DefaultAddress
	}

	config.StaticPath = os.Getenv("SERVER_STATIC_PATH")
	if config.StaticPath == "" {
		config.StaticPath = DefaultStaticPath
	}

	config.UploadsPath = os.Getenv("SERVER_UPLOAD_PATH")
	if config.UploadsPath == "" {
		config.UploadsPath = DefaultUploadsPath
	}

	config.SessionStorage = sessions.NewInMemorySessionStorage()
	if config.SessionStorage == nil {
		log.Fatal("Session storage not created")
	}

	config.ProfileRepository = profiles.NewInMemoryProfileRepository()
	if config.ProfileRepository == nil {
		log.Fatal("Profile repository not created")
	}

	srv := handlers.CreateServer(config)
	if srv == nil {
		log.Fatal("Server not started")
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
