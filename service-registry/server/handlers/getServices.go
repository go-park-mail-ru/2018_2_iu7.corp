package handlers

import (
	"2018_2_iu7.corp/service-registry/services/models"
	"2018_2_iu7.corp/service-registry/services/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func GetServices(r repositories.ServiceRepository) context.Handler {
	return func(c iris.Context) {
		info, err := r.GetAllServicesInfo()
		if err != nil {
			writeResponseError(c, err)
			return
		}

		sInfo := models.Services{
			Services: info,
		}

		writeResponseJSON(c, sInfo)
	}
}
