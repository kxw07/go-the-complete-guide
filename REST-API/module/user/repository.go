package user

import (
	"github.com/kxw07/REST-API/storage"
	"log/slog"
)

type Repository struct {
	sto *storage.Storage
}

func NewRepository(sto *storage.Storage) *Repository {
	return &Repository{sto: sto}
}

func (rep Repository) save(user User) (User, error) {
	sql := `
	INSERT INTO users (username, password)
	VALUES (?, ?)
	`

	stmt, err := rep.sto.GetDB().Prepare(sql)
	if err != nil {
		slog.Error("Could not save user", "error", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		slog.Error("Could not execute save user statement", "error", err)
		return User{}, err
	}

	user.ID, _ = result.LastInsertId()
	return user, nil
}

func (rep Repository) get(user User) (User, error) {
	sql := `
	SELECT id, password FROM users WHERE username = ?
	`

	stmt, err := rep.sto.GetDB().Prepare(sql)
	if err != nil {
		slog.Error("Could not prepare get user statement", "error", err)
		return User{}, nil
	}

	defer stmt.Close()

	err = stmt.QueryRow(user.Email).Scan(&user.ID, &user.Password)
	if err != nil {
		slog.Error("Could not execute get user statement", "error", err)
		return User{}, nil
	}

	return user, nil
}
