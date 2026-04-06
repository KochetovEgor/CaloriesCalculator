package models

type User struct {
	Username string `json:"username"`
}

type UserPWD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}
