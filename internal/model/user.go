package model

type User struct {
	Id       int
	Email    string
	Password string
}

func (u *User) NewUser() *User {
	return &User{}
}
