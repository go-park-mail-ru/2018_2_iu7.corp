package handlers

import (
	"2018_2_iu7.corp/profile-service/errors"
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
	"strconv"
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
	case *errors.NotFoundError:
		c.StatusCode(http.StatusNotFound)
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
	c.ResponseWriter().Header().Set("Content-Type", "application/json")
	c.ResponseWriter().WriteHeader(http.StatusOK)
	c.ResponseWriter().Write(resp)
}

func getProfileID(c iris.Context) (uint32, error) {
	id, err := strconv.ParseUint(c.Params().Get("profileID"), 0, 32)
	if err != nil {
		return 0, errors.NewInvalidFormatError("invalid id: type mismatch")
	}
	return uint32(id), nil
}
