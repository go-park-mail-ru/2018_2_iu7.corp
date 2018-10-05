package handlers

import (
	"2018_2_iu7.corp/profile-service/profiles/repositories"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func GetLeaders(r repositories.ProfileRepository) context.Handler {
	return func(c iris.Context) {
		//TODO
	}
}
