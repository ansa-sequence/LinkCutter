package store

import "LinkCutter/internal/model"

type UserRepository struct {
	store *Store
}

func (repository *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := repository.store.db.QueryRow(
		"insert into information (email, encryptedpassword) values ($1,$2) RETURNING id",
		user.Email,
		user.Password,
	).Scan(&user.Id); err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := repository.store.db.QueryRow(
		"SELECT id, email, encryptedpassword FROM information WHERE email = $1",
		email,
	).Scan(&u.Id, &u.Email, &u.Password); err != nil {
		return nil, err
	}

	return u, nil
}
