package model_tests

import (
	"LinkCutter/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserValidate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{{
		name: "valid",
		user: func() *model.User {
			return model.TestUser(t)
		},
		isValid: true,
	},
		{
			name: "empty email",
			user: func() *model.User {
				return &model.User{}
			},
			isValid: false,
		},
		{
			name: "invalid email",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = "short"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "encrypted password",
			user: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "fdafasd"
				return u
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.user().Validate())
			} else {
				assert.Error(t, tc.user().Validate())
			}
		})
	}
}

func TestUserBeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
