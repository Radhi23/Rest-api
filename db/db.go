package db

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var DB *sql.DB

func InitDB() {

	db, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to the db ..")
	}
	DB = db
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()

}

func createTable() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL )`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("could not create events table .." + err.Error())
	}
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("could not create events table .." + err.Error())
	}
	_, err = DB.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		panic("could not create events table .." + err.Error())
	}
}
