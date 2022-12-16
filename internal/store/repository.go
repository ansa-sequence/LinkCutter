package store

import "LinkCutter/internal/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	RemoveByEmail(string) error
}
