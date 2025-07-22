package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() *Storage {
	var sto Storage
	sto.db = sto.initDB()

	sto.db.SetMaxOpenConns(10)
	sto.db.SetMaxIdleConns(5)

	sto.createTable()

	return &sto
}

func (sto *Storage) GetDB() *sql.DB {
	if sto.db == nil {
		slog.Error("Database not initialized.")
		panic("Database not initialized.")
	}

	return sto.db
}

func (sto *Storage) initDB() *sql.DB {
	var err error
	db, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	return db
}

func (sto *Storage) createTable() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	username TEXT NOT NULL UNIQUE,
    	password TEXT NOT NULL
	)`
	_, err := sto.db.Exec(createUsersTable)
	if err != nil {
		slog.Error("Could not create table: users")
		panic(err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		create_time DATETIME current_timestamp,
		update_time DATETIME current_timestamp,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	_, err = sto.db.Exec(createEventsTable)
	if err != nil {
		slog.Error("Could not create table: events")
		panic(err)
	}
}
