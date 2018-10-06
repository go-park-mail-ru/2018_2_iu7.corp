package handlers

import (
	"2018_2_iu7.corp/profile-service/errors"
	"2018_2_iu7.corp/profile-service/profiles/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func GetLeaders(r repositories.ProfileRepository) context.Handler {
	return func(c iris.Context) {
		p, err := c.URLParamInt("page")
		if err != nil {
			writeError(c, errors.NewInvalidFormatError("invalid query: no page"))
			return
		}

		s, err := c.URLParamInt("pageSize")
		if err != nil {
			writeError(c, errors.NewInvalidFormatError("invalid query: no page"))
			return
		}

		pfs, err := r.GetSeveralOrderByScorePaginated(p, s)
		if err != nil {
			writeError(c, err)
			return
		}

		writeResponse(c, pfs)
	}
}
