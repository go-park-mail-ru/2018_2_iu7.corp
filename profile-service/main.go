package main

import (
	"2018_2_iu7.corp/profile-service/repositories"
	"2018_2_iu7.corp/profile-service/service/rest"
	"2018_2_iu7.corp/profile-service/service/rpc"
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
	DefaultAddress      = "0.0.0.0:8090"
	DefaultShutdownTime = 10
)

func main() {
	addressPtr := flag.String("addr", DefaultAddress, "service address")
	shutdownTimePtr := flag.Int("shutdown", DefaultShutdownTime, "service shutdown time [seconds]")

	dbHostPtr := flag.String("dbhost", repositories.DefaultHost, "database host")
	dbPortPtr := flag.String("dbport", repositories.DefaultPort, "database port")
	dbUserPtr := flag.String("dbuser", repositories.DefaultUser, "database user")
	dbPasswordPtr := flag.String("dbpassword", repositories.DefaultPassword, "database password")
	dbNamePtr := flag.String("dbname", repositories.DefaultDB, "database name")

	flag.Parse()

	if len(flag.Args()) != 0 {
		log.Fatal("unknown command-line arguments")
	}

	connParams := &repositories.ConnectionParams{
		Host:     *dbHostPtr,
		Port:     *dbPortPtr,
		User:     *dbUserPtr,
		Password: *dbPasswordPtr,
		Database: *dbNamePtr,
	}

	r := repositories.NewDBProfileRepository(connParams)
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
		log.Fatal("rest service not created")
	}

	rpcSrv, err := rpc.CreateService(r)
	if err != nil {
		log.Fatal("rpc service not started")
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

	shutdownTime := time.Duration(*shutdownTimePtr) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime)
	defer cancel()

	if err := restSrv.Shutdown(ctx); err != nil {
		log.Fatal("service shutdown failed")
	}
}
