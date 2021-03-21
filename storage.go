package auth

import (
	"time"

	"github.com/joaosoft/dbr"
)

type storagePostgres struct {
	config *AuthConfig
	db     *dbr.Dbr
}

func newStoragePostgres(config *AuthConfig) (*storagePostgres, error) {
	dbr, err := dbr.New(dbr.WithConfiguration(config.Dbr))
	if err != nil {
		return nil, err
	}

	return &storagePostgres{
		config: config,
		db:     dbr,
	}, nil
}

func (storage *storagePostgres) getUserByIdUserAndRefreshToken(idUser, refreshToken string) (*User, error) {
	user := &User{}
	count, err := storage.db.
		Select("*").
		From(authTableUser).
		Where("id_user = ?", idUser).
		Where("refresh_token = ?", refreshToken).
		Where("active").
		Load(user)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrorNotFound
	}

	return user, nil
}

func (storage *storagePostgres) getUserByEmailAndPassword(email, password string) (*User, error) {
	user := &User{}
	count, err := storage.db.
		Select("*").
		From(authTableUser).
		Where("email = ?", email).
		Where("password_hash = ?", password).
		Where("active").
		Load(user)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrorNotFound
	}

	return user, nil
}

func (storage *storagePostgres) updateUserRefreshToken(idUser, refreshToken string) error {
	result, err := storage.db.
		Update(authTableUser).
		Set("refresh_token", refreshToken).
		Set("updated_at", time.Now()).
		Where("id_user = ?", idUser).
		Where("active").
		Exec()

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrorNotFound
	}

	return nil
}

func (storage *storagePostgres) signUp(user *User) error {
	result, err := storage.db.
		Insert().Into(authTableUser).
		Record(user).
		Exec()

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrorNotFound
	}

	return nil
}

func (storage *storagePostgres) updateUserStatus(idUser string, isActive bool) error {
	result, err := storage.db.
		Update(authTableUser).
		Set("active", isActive).
		Set("updated_at", time.Now()).
		Where("id_user = ?", idUser).
		Exec()

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrorNotFound
	}

	return nil
}
