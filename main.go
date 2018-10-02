package main

import (
	"2018_2_iu7.corp/profiles"
	"2018_2_iu7.corp/server"
	"2018_2_iu7.corp/sessions"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	DefaultAddress      = ":8080"
	DefaultStaticPath   = "./static/"
	DefaultUploadsPath  = "./upload/"
	DefaultShutdownTime = 10
)

func main() {
	addressPtr := flag.String("addr", DefaultAddress, "server address")
	staticPathPtr := flag.String("static", DefaultStaticPath, "static files path")
	uploadsPathPtr := flag.String("uploads", DefaultUploadsPath, "uploaded files path")
	sessionKeyPtr := flag.String("key", "", "session key")
	shutdownTimePtr := flag.Int("shutdown", DefaultShutdownTime, "server shutdown time [seconds]")

	flag.Parse()

	if len(flag.Args()) != 0 {
		log.Fatal("unknown command-line arguments")
	}

	config := &server.Config{
		Address:     *addressPtr,
		StaticPath:  *staticPathPtr,
		UploadsPath: *uploadsPathPtr,
	}

	config.SessionStorage = sessions.NewCookieSessionStorage(*sessionKeyPtr)
	if config.SessionStorage == nil {
		log.Fatal("session storage not created")
	}

	config.ProfileRepository = profiles.NewInMemoryProfileRepository()
	if config.ProfileRepository == nil {
		log.Fatal("profile repository not created")
	}

	srv := server.CreateServer(config)
	if srv == nil {
		log.Fatal("server not created")
	}

	log.Printf("server is configured to start on %s", config.Address)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT|syscall.SIGTERM)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-ch

	shutdownTime := time.Duration(*shutdownTimePtr) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), shutdownTime)

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
