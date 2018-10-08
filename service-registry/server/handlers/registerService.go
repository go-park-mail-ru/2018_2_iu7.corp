package handlers

import (
	"2018_2_iu7.corp/common/regclient"
	"2018_2_iu7.corp/service-registry/services/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func RegisterService(r repositories.ServiceRepository) context.Handler {
	return func(c iris.Context) {
		var serviceInfo regclient.ServiceInfo
		if err := parseRequestBodyJSON(c, &serviceInfo); err != nil {
			writeResponseError(c, err)
			return
		}

		if err := r.RegisterService(serviceInfo.Name, serviceInfo.Address); err != nil {
			writeResponseError(c, err)
			return
		}

		writeResponseOK(c)
	}
}
