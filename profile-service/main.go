package main

import (
	"2018_2_iu7.corp/profile-service/profiles/repositories"
	"2018_2_iu7.corp/profile-service/server"
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
	DefaultAddress      = ":8080"
	DefaultShutdownTime = 10
)

func main() {
	addressPtr := flag.String("addr", DefaultAddress, "server address")
	shutdownTimePtr := flag.Int("shutdown", DefaultShutdownTime, "server shutdown time [seconds]")

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
