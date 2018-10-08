package server

import (
	"2018_2_iu7.corp/service-registry/server/handlers"
	"2018_2_iu7.corp/service-registry/services/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func CreateServer(r repositories.ServiceRepository) (*iris.Application, error) {
	server := iris.Default()

	rLog := logger.New(logger.Config{
		Status: true,
		Method: true,
		Path:   true,
		Query:  true,
	})
	server.Use(rLog)

	server.Get("/", handlers.GetServices(r))
	server.Get("/{serviceName}", handlers.GetService(r))
	server.Post("/{serviceName}", handlers.RegisterService(r))
	server.Put("/{serviceName}", handlers.UpdateService(r))
	server.Delete("/{serviceName}", handlers.UnregisterService(r))

	return server, nil
}
