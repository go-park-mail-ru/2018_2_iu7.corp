package handlers

import (
	"2018_2_iu7.corp/service-registry/services/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func GetService(r repositories.ServiceRepository) context.Handler {
	return func(c iris.Context) {
		name := c.Params().Get("serviceName")

		info, err := r.GetServiceInfo(name)
		if err != nil {
			writeResponseError(c, err)
			return
		}

		writeResponseJSON(c, info)
	}
}
