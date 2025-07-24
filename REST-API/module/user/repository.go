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

func (rep Repository) get(email string) string {
	sql := `
	SELECT password FROM users WHERE username = ?
	`

	stmt, err := rep.sto.GetDB().Prepare(sql)
	if err != nil {
		slog.Error("Could not prepare get user statement", "error", err)
		return ""
	}

	defer stmt.Close()

	var password string
	err = stmt.QueryRow(email).Scan(&password)
	if err != nil {
		slog.Error("Could not execute get user statement", "error", err)
		return ""
	}

	return password
}
