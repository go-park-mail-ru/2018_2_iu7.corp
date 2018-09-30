package profiles

import (
	"2018_2_iu7.corp/errs"
	"fmt"
	"regexp"
)

type Profile struct {
	ID         uint64
	Email      string
	Username   string
	Password   string
	AvatarPath string
	Score      uint16
}

func (p *Profile) ParseOnRegister(m map[string]interface{}) error {
	return p.parseOnEdit(m, 3)
}

func (p *Profile) ParseOnLogin(m map[string]interface{}) error {
	return p.parseOnLogin(m)
}

func (p *Profile) ParseOnEdit(m map[string]interface{}) error {
	return p.parseOnEdit(m, -1)
}

func (p *Profile) GetPublicAttributes() map[string]interface{} {
	return map[string]interface{}{
		"id":       p.ID,
		"username": p.Username,
		"avatar":   p.AvatarPath,
		"score":    p.Score,
	}
}

func (p *Profile) GetPrivateAttributes() map[string]interface{} {
	m := p.GetPublicAttributes()
	m["email"] = p.Email
	return m
}

func (p *Profile) parseOnEdit(m map[string]interface{}, n int) error {
	if n != -1 && len(m) != n {
		return errs.NewInvalidFormatError("wrong number of attributes")
	}

	var err error = nil
	for k, v := range m {
		switch k {
		case "username":
			p.Username, err = parseUsername(v)
		case "email":
			p.Email, err = parseEmail(v)
		case "password":
			p.Password, err = parsePassword(v)
		default:
			return errs.NewInvalidFormatError(fmt.Sprintf("unknown attribute: %s", k))
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Profile) parseOnLogin(m map[string]interface{}) error {
	if len(m) != 2 {
		return errs.NewInvalidFormatError("wrong number of attributes")
	}

	var err error = nil
	for k, v := range m {
		switch k {
		case "username":
			p.Username, err = parseUsername(v)
		case "password":
			p.Password, err = parsePassword(v)
		default:
			return errs.NewInvalidFormatError(fmt.Sprintf("unknown attribute: %s", k))
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func parseUsername(v interface{}) (string, error) {
	username, ok := v.(string)
	if !ok {
		return "", errs.NewInvalidFormatError("invalid username: wrong type")
	}

	if m, err := regexp.MatchString("^\\w+$", username); err != nil {
		panic(err)
	} else if !m {
		return "", errs.NewInvalidFormatError("invalid username: pattern mismatch")
	}

	return username, nil
}

func parseEmail(v interface{}) (string, error) {
	email, ok := v.(string)
	if !ok {
		return "", errs.NewInvalidFormatError("invalid email: wrong type")
	}

	if m, err := regexp.MatchString("^.+@.+$", email); err != nil {
		panic(err)
	} else if !m {
		return "", errs.NewInvalidFormatError("invalid email: pattern mismatch")
	}

	return email, nil
}

func parsePassword(v interface{}) (string, error) {
	password, ok := v.(string)
	if !ok {
		return "", errs.NewInvalidFormatError("invalid password: wrong type")
	}

	if password == "" {
		return "", errs.NewInvalidFormatError("invalid password: empty")
	}

	return password, nil
}
