package profiles

import (
	"2018_2_iu7.corp/profile-service/errors"
	"regexp"
)

//easyjson:json
type Profile struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Score    int32  `json:"score"`
}

func (p Profile) Validate() error {
	if m, err := regexp.MatchString("^\\w+$", p.Username); err != nil {
		panic(err)
	} else if !m {
		return errors.NewInvalidFormatError("invalid username: pattern mismatch")
	}

	if m, err := regexp.MatchString("^.+@.+$", p.Email); err != nil {
		panic(err)
	} else if !m {
		return errors.NewInvalidFormatError("invalid email: pattern mismatch")
	}

	if p.Password == "" {
		return errors.NewInvalidFormatError("invalid password: empty")
	}

	return nil
}
