package main

import "fmt"

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
	if p.Username == "" {
		return NewInvalidFormatError("invalid username: empty")
	}

	for i, c := range p.Username {
		if c < 'a' || c > 'z' {
			return NewInvalidFormatError(fmt.Sprintf("invalid username: position %d", i))
		}
	}

	flag := false
	for i, c := range p.Email {
		if c == '@' {
			if !flag {
				flag = true
			} else {
				return NewInvalidFormatError(fmt.Sprintf("invalid email: position %d", i))
			}
		}
	}

	if p.Password == "" {
		return NewInvalidFormatError("invalid password: empty")
	}

	return nil
}
