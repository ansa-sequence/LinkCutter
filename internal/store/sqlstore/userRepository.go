package sqlstore

import "LinkCutter/internal/model"

type UserRepository struct {
	store *Store
}

func (repository *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if err := user.BeforeCreate(); err != nil {
		return err
	}

	return repository.store.db.QueryRow(
		"INSERT INTO information (email, encryptedpassword) VALUES ($1, $2) RETURNING id",
		user.Email,
		user.EncryptedPassword,
	).Scan(&user.Id)
}

func (repository *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := repository.store.db.QueryRow(
		"SELECT id, email, encryptedpassword FROM information WHERE email = $1",
		email,
	).Scan(&u.Id, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}

func (repository *UserRepository) RemoveByEmail(email string) error {
	u, err := repository.FindByEmail(email)
	if err != nil {
		return err
	}
	repository.store.db.QueryRow("delete from information where email = $1", u.Email)
	return nil
}
