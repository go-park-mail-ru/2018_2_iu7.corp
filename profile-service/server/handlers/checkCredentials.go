package handlers

import (
	"2018_2_iu7.corp/profile-service/profiles/models"
	"2018_2_iu7.corp/profile-service/profiles/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func CheckCredentials(r repositories.ProfileRepository) context.Handler {
	return func(c iris.Context) {
		var cr models.Credentials
		if err := getRequestEntity(c, &cr); err != nil {
			writeError(c, err)
			return
		}

		p, err := r.FindByCredentials(cr)
		if err != nil {
			writeError(c, err)
			return
		}

		id := &models.ProfileID{
			Value: p.ProfileID,
		}

		writeResponse(c, id)
	}
}
