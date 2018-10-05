package handlers

import (
	"2018_2_iu7.corp/profile-service/profiles/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func GetProfile(r repositories.ProfileRepository) context.Handler {
	return func(c iris.Context) {
		id, err := c.Params().GetInt64("profileID")
		if err != nil {
			panic(err)
		}

		p, err := r.FindByID(id)
		if err != nil {
			writeError(c, err)
			return
		}

		writeResponse(c, &p)
	}
}
