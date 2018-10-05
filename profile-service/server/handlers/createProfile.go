package handlers

import (
	"2018_2_iu7.corp/profile-service/profiles/models"
	"2018_2_iu7.corp/profile-service/profiles/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func CreateProfile(r repositories.ProfileRepository) context.Handler {
	return func(c iris.Context) {
		var p models.NewProfile

		if err := getRequestEntity(c, &p); err != nil {
			writeError(c, err)
			return
		}

		if err := r.SaveNew(p.Get()); err != nil {
			writeError(c, err)
			return
		}

		writeSuccess(c)
	}
}
