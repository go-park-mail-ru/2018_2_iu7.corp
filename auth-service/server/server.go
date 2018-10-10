package server

import (
	"2018_2_iu7.corp/auth-service/server/handlers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func CreateServer() (*iris.Application, error) {
	server := iris.Default()

	rLog := logger.New(logger.Config{
		Status: true,
		Method: true,
		Path:   true,
		Query:  true,
	})
	server.Use(rLog)

	server.Post("/session", handlers.CreateSession())
	server.Delete("/session", handlers.DeleteSession())

	return server, nil
}
