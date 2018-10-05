package server

import (
	"2018_2_iu7.corp/profile-service/profiles/repositories"
	"2018_2_iu7.corp/profile-service/server/handlers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func CreateServer(r repositories.ProfileRepository) (*iris.Application, error) {
	server := iris.Default()

	rLog := logger.New(logger.Config{
		Status: true,
		Method: true,
		Path:   true,
		Query:  true,
	})
	server.Use(rLog)

	server.Post("/new", handlers.CreateProfile(r))
	server.Get("/{profileID:long}", handlers.GetProfile(r))
	server.Put("/{profileID:long}", handlers.UpdateProfile(r))
	server.Delete("/{profileID:long}", handlers.DeleteProfile(r))
	server.Get("/leaders", handlers.GetLeaders(r))

	return server, nil
}
