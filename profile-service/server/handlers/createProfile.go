package handlers

import (
	"2018_2_iu7.corp/profile-service/errors"
	"2018_2_iu7.corp/profile-service/profiles"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func CreateProfile(r profiles.ProfileRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		rb, err := getRequestBody(c)
		if err != nil {
			writeError(c, http.StatusInternalServerError, errors.NewServiceError())
			return
		}

		var p profiles.NewProfile
		if err = p.UnmarshalJSON(rb); err != nil {
			writeError(c, http.StatusBadRequest, err)
			return
		}

		if err = r.SaveNew(p.Get()); err != nil {
			writeError(c, http.StatusInternalServerError, errors.NewServiceError())
			return
		}

		writeSuccess(c, http.StatusOK)
	}
}

func getRequestBody(c *gin.Context) ([]byte, error) {
	rb, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	c.Request.Body.Close()
	return rb, nil
}

func writeSuccess(c *gin.Context, status int) {
	c.Writer.WriteHeader(status)
}

func writeError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{"error": err.Error()})
}

func writeProfile(c *gin.Context, p profiles.Profile) {
	resp, err := p.MarshalJSON()
	if err != nil {
		panic(err)
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(resp)
}

func writeProfiles(c *gin.Context, p []profiles.Profile) {
	c.JSON(http.StatusOK, gin.H{"profiles": p})
}
