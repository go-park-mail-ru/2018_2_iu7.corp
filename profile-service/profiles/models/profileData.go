package models

//easyjson:json
type ProfileData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (p ProfileData) AsProfile() Profile {
	return Profile{
		Username: p.Username,
		Password: p.Password,
		Email:    p.Email,
	}
}
