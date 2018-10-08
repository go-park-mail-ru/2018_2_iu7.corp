package handlers

import (
	"2018_2_iu7.corp/service-registry/services/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func UpdateService(r repositories.ServiceRepository) context.Handler {
	return func(c iris.Context) {
		name := c.Params().Get("serviceName")
		addr := c.RemoteAddr()

		if err := r.UpdateService(name, addr); err != nil {
			writeErrorJSON(c, err)
			return
		}

		writeResponseOK(c)
	}
}
