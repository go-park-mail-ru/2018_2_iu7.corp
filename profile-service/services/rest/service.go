package rest

import (
	"2018_2_iu7.corp/profile-service/repositories"
	"2018_2_iu7.corp/profile-service/services/rest/handlers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func CreateService(r repositories.ProfileRepository) (*iris.Application, error) {
	server := iris.Default()

	rLog := logger.New(logger.Config{
		Status: true,
		Method: true,
		Path:   true,
		Query:  true,
	})
	server.Use(rLog)

	server.Post("/new", handlers.CreateProfile(r))
	server.Get("/{profileID:int}", handlers.GetProfile(r))
	server.Put("/{profileID:int}", handlers.UpdateProfile(r))
	server.Delete("/{profileID:int}", handlers.DeleteProfile(r))
	server.Get("/leaders", handlers.GetLeaders(r))

	return server, nil
}
