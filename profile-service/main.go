package main

import (
	"2018_2_iu7.corp/profile-service/repositories"
	"2018_2_iu7.corp/profile-service/services/rest"
	"2018_2_iu7.corp/profile-service/services/rpc"
	"context"
	"flag"
	"github.com/kataras/iris"
	_ "github.com/micro/go-micro"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	DefaultAddress      = "127.0.0.1:8090"
	DefaultShutdownTime = 5 * time.Second
)

func main() {
	addressPtr := flag.String("-addr", DefaultAddress, "services address")
	flag.Parse()

	r := repositories.NewDBProfileRepository()
	if r == nil {
		log.Fatal("profile repository not created")
	}

	err := r.Open()
	if err != nil {
		log.Fatal("profile repository not available")
	}
	defer r.Close()

	restSrv, err := rest.CreateService(r)
	if err != nil {
		log.Fatal("rest services not created")
	}

	rpcSrv, err := rpc.CreateService(r)
	if err != nil {
		log.Fatal("rpc services not started")
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT|syscall.SIGTERM)

	go func() {
		if err := restSrv.Run(iris.Addr(*addressPtr)); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := (*rpcSrv).Run(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), DefaultShutdownTime)
	defer cancel()

	if err := restSrv.Shutdown(ctx); err != nil {
		log.Fatal("services shutdown failed")
	}
}
