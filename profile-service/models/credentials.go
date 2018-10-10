package models

//easyjson:json
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *Credentials) AsProfile() (*Profile, error) {
	return &Profile{
		Username: p.Username,
		Password: p.Password,
	}, nil
}
