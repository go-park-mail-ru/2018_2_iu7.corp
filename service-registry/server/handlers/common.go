package handlers

import (
	"2018_2_iu7.corp/common/errors"
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
)

func parseRequestBodyJSON(c iris.Context, e RequestEntity) error {
	rb, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return errors.NewServiceError()
	}
	c.Request().Body.Close()

	err = e.UnmarshalJSON(rb)
	if err != nil {
		return errors.NewInvalidFormatError("invalid request body: format mismatch")
	}

	return nil
}

func writeResponseOK(c iris.Context) {
	c.ResponseWriter().WriteHeader(http.StatusOK)
}

func writeErrorJSON(c iris.Context, err error) {
	switch err.(type) {
	default:
		c.StatusCode(http.StatusInternalServerError)
	}
	c.JSON(iris.Map{"error": err.Error()})
}

func writeResponseJSON(c iris.Context, v ResponseEntity) {
	resp, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}
	c.ResponseWriter().Header().Set("Content-Type", "application/json")
	c.ResponseWriter().WriteHeader(http.StatusOK)
	c.ResponseWriter().Write(resp)
}
