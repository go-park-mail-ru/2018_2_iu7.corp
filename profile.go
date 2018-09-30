package main

import (
	"fmt"
	"regexp"
)

type Profile struct {
	ID         uint64 `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AvatarPath string `json:"avatar"`
	Score      uint16 `json:"score"`
}

func (p *Profile) ValidateNew() error {
	if err := p.Validate(); err != nil {
		return err
	}

	p.ID = 0
	p.AvatarPath = ""
	p.Score = 0

	return nil
}

func (p Profile) Validate() error {
	if m, err := regexp.MatchString("^\\w+$", p.Username); err != nil {
		panic(err)
	} else if !m {
		return NewInvalidFormatError("invalid username")
	}

	if m, err := regexp.MatchString("^.+@.+$", p.Email); err != nil {
		panic(err)
	} else if !m {
		return NewInvalidFormatError("invalid email")
	}

	if p.Password == "" {
		return NewInvalidFormatError("invalid password")
	}

	return nil
}

func ParseProfileOnLogin(m map[string]interface{}) (*Profile, error) {
	p := &Profile{}

	if len(m) != 2 {
		return nil, NewInvalidFormatError("wrong number of attributes")
	}

	var ok bool
	for k, v := range m {
		switch k {
		case "username":
			p.Username, ok = v.(string)
			if !ok {
				return nil, NewInvalidFormatError("invalid username: wrong type")
			}
		case "password":
			p.Password, ok = v.(string)
			if !ok {
				return nil, NewInvalidFormatError("invalid password: wrong type")
			}
		default:
			return nil, NewInvalidFormatError(fmt.Sprintf("unknown attribute: %s", k))
		}
	}

	return p, nil
}
