package handlers

import (
	"2018_2_iu7.corp/profile-service/profiles/models"
	"2018_2_iu7.corp/profile-service/profiles/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProfile(r repositories.ProfileRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p models.NewProfile

		if err := getRequestEntity(c, &p); err != nil {
			writeError(c, err)
			return
		}

		if err := r.SaveNew(p.Get()); err != nil {
			writeError(c, err)
			return
		}

		writeSuccess(c, http.StatusOK)
	}
}
