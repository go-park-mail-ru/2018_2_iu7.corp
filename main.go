package main

import (
	"log"
	"os"
	"sync"
)

const (
	DefaultAddress = ":8080"
)

func main() {
	addr := os.Getenv("SERVER_ADDRESS")
	if addr == "" {
		addr = DefaultAddress
	}

	srv := CreateServer(addr)
	if srv == nil {
		log.Fatal("Server startup failed")
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
