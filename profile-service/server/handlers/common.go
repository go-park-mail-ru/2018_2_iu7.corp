package handlers

import (
	"2018_2_iu7.corp/profile-service/errors"
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
)

func getRequestEntity(c iris.Context, e requestEntity) error {
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

func writeSuccess(c iris.Context) {
	c.ResponseWriter().WriteHeader(http.StatusOK)
}

func writeError(c iris.Context, err error) {
	switch err.(type) {
	case *errors.InvalidFormatError:
		c.StatusCode(http.StatusBadRequest)
	case *errors.ConstraintViolationError:
		c.StatusCode(http.StatusConflict)
	default:
		c.StatusCode(http.StatusInternalServerError)
	}
	c.JSON(iris.Map{"error": err.Error()})
}

func writeResponse(c iris.Context, v responseEntity) {
	resp, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}
	c.ResponseWriter().WriteHeader(http.StatusOK)
	c.ResponseWriter().Write(resp)
}
