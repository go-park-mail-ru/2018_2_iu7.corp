package server

import (
	"2018_2_iu7.corp/profile-service/profiles/repositories"
	"2018_2_iu7.corp/profile-service/server/handlers"
	"github.com/gin-gonic/gin"
)

func CreateServer(r repositories.ProfileRepository) (*gin.Engine, error) {
	router := gin.Default()

	router.POST("/new", handlers.CreateProfile(r))
	router.GET("/:profileID", handlers.GetProfile)
	router.PUT("/:profileID", handlers.UpdateProfile)
	router.DELETE("/:profileID", handlers.DeleteProfile)
	router.GET("/leaders", handlers.GetLeaders)

	return router, nil
}
