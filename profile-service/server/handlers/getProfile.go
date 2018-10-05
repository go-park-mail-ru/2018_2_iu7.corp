package handlers

import (
	"2018_2_iu7.corp/profile-service/profiles/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func GetProfile(r repositories.ProfileRepository) context.Handler {
	return func(c iris.Context) {
		//TODO
	}
}
