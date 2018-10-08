package main

import (
	"2018_2_iu7.corp/service-registry/server"
	"2018_2_iu7.corp/service-registry/services/repositories"
	"context"
	"flag"
	"github.com/kataras/iris"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	DefaultAddress      = ":8765"
	DefaultShutdownTime = 10
)

func main() {
	addressPtr := flag.String("addr", DefaultAddress, "server address")
	shutdownTimePtr := flag.Int("shutdown", DefaultShutdownTime, "server shutdown time [seconds]")

	flag.Parse()

	if len(flag.Args()) != 0 {
		log.Fatal("unknown command-line arguments")
	}

	r := repositories.NewInMemoryServiceRepository(repositories.DefaultExpireTime)
	if r == nil {
		log.Fatal("service repository not created")
	}
	r.StartMonitor()

	srv, err := server.CreateServer(r)
	if err != nil {
		log.Fatal("server not created")
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT|syscall.SIGTERM)

	go func() {
		err := srv.Run(iris.Addr(*addressPtr))
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-ch

	shutdownTime := time.Duration(*shutdownTimePtr) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown failed")
	}
}
