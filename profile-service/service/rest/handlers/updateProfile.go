package handlers

import (
	"2018_2_iu7.corp/profile-service/models"
	"2018_2_iu7.corp/profile-service/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func UpdateProfile(r repositories.ProfileRepository) context.Handler {
	return func(c iris.Context) {
		id, err := getProfileID(c)
		if err != nil {
			writeError(c, err)
			return
		}

		var u models.ProfileDataUpdate
		if err := getRequestEntity(c, &u); err != nil {
			writeError(c, err)
			return
		}

		if err = r.SaveExisting(id, u); err != nil {
			writeError(c, err)
			return
		}

		writeSuccess(c)
	}
}
