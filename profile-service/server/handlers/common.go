package handlers

import (
	"2018_2_iu7.corp/profile-service/errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func getRequestEntity(c *gin.Context, entity requestEntity) error {
	rb, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	c.Request.Body.Close()

	err = entity.UnmarshalJSON(rb)
	if err != nil {
		return errors.NewInvalidFormatError("invalid request body: format mismatch")
	}

	return nil
}

func writeSuccess(c *gin.Context, status int) {
	c.Writer.WriteHeader(status)
}

func writeError(c *gin.Context, err error) {
	switch err.(type) {
	case *errors.InvalidFormatError:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func writeResponse(c *gin.Context, v responseEntity) {
	resp, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(resp)
}
