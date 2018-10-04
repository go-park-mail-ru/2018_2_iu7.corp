package profiles

//easyjson:json
type NewProfile struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (p NewProfile) Get() Profile {
	return Profile{
		Username: p.Username,
		Password: p.Password,
		Email:    p.Email,
	}
}
