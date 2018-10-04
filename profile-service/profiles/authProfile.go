package profiles

//easyjson:json
type AuthProfile struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *AuthProfile) Get() (*Profile, error) {
	return &Profile{
		Username: p.Username,
		Password: p.Password,
	}, nil
}
