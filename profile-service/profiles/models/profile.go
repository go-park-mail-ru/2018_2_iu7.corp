package models

import (
	"2018_2_iu7.corp/profile-service/errors"
	"regexp"
)

//easyjson:json
type Profile struct {
	ID       int64  `json:"id"       gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Username string `json:"username" gorm:"type:varchar(20);unique"`
	Password string `json:"-"        gorm:"type:varchar(50)"`
	Email    string `json:"email"    gorm:"type:varchar(100);unique"`
	Score    int32  `json:"score"    gorm:"DEFAULT:0"`
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

//easyjson:json
type Profiles []Profile
