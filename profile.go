package main

type Profile struct {
	ID         uint64 `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	AvatarPath string `json:"avatar"`
	Score      uint16 `json:"score"`
}
