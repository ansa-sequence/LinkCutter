package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 20)),
	)
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		encrypt, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = encrypt
	}
	return nil
}

func encryptString(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
