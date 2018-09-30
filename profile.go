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

func ParseProfileOnRegister(m map[string]interface{}) (p *Profile, err error) {
	p = &Profile{}

	if len(m) != 3 {
		return nil, NewInvalidFormatError("wrong number of attributes")
	}

	err = nil
	for k, v := range m {
		switch k {
		case "username":
			p.Username, err = parseUsername(v)
		case "email":
			p.Password, err = parseEmail(v)
		case "password":
			p.Password, err = parsePassword(v)
		default:
			return nil, NewInvalidFormatError(fmt.Sprintf("unknown attribute: %s", k))
		}

		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func ParseProfileOnLogin(m map[string]interface{}) (p *Profile, err error) {
	p = &Profile{}

	if len(m) != 2 {
		return nil, NewInvalidFormatError("wrong number of attributes")
	}

	err = nil
	for k, v := range m {
		switch k {
		case "username":
			p.Username, err = parseUsername(v)
		case "password":
			p.Password, err = parsePassword(v)
		default:
			return nil, NewInvalidFormatError(fmt.Sprintf("unknown attribute: %s", k))
		}

		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func parseUsername(v interface{}) (string, error) {
	username, ok := v.(string)
	if !ok {
		return "", NewInvalidFormatError("invalid username: wrong type")
	}

	if m, err := regexp.MatchString("^\\w+$", username); err != nil {
		panic(err)
	} else if !m {
		return "", NewInvalidFormatError("invalid username: pattern mismatch")
	}

	return username, nil
}

func parseEmail(v interface{}) (string, error) {
	email, ok := v.(string)
	if !ok {
		return "", NewInvalidFormatError("invalid email: wrong type")
	}

	if m, err := regexp.MatchString("^.+@.+$", email); err != nil {
		panic(err)
	} else if !m {
		return "", NewInvalidFormatError("invalid email: pattern mismatch")
	}

	return email, nil
}

func parsePassword(v interface{}) (string, error) {
	password, ok := v.(string)
	if !ok {
		return "", NewInvalidFormatError("invalid password: wrong type")
	}

	if password == "" {
		return "", NewInvalidFormatError("invalid password: empty")
	}

	return password, nil
}
