package models

type User struct {
	UserId   string
	Ip       string
	Name     string
	Email    string
	Password string `json:"-"`
}
