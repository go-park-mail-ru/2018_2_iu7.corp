package handlers

import (
	"2018_2_iu7.corp/profile-service/profiles/models"
	"2018_2_iu7.corp/profile-service/profiles/repositories"
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

		var np models.ProfileData
		if err := getRequestEntity(c, &np); err != nil {
			writeError(c, err)
			return
		}

		p := np.AsProfile()
		p.ProfileID = id

		if err = r.SaveExisting(p); err != nil {
			writeError(c, err)
			return
		}

		writeSuccess(c)
	}
}
