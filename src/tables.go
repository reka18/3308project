package main

import (
	"database/sql"
	"log"
)

var genderEnum = "CREATE TYPE gender " +
	"AS ENUM ('M', 'F', 'O');"

var usersTable = "CREATE TABLE users (" +
	"id SERIAL PRIMARY KEY," +
	"age INT NOT NULL," +
	"firstname TEXT NOT NULL," +
	"lastname TEXT NOT NULL," +
	"email TEXT UNIQUE NOT NULL," +
	"username TEXT UNIQUE NOT NULL," +
	"public BOOLEAN NOT NULL," +
	"joindate TIMESTAMP NOT NULL," +
	"active BOOLEAN NOT NULL," +
	"password TEXT NOT NULL," +
	"gender GENDER NOT NULL);"

var postTable = "CREATE TABLE posts (" +
	"id SERIAL PRIMARY KEY," +
	"userid INT REFERENCES users(id)," +
	"content VARCHAR(240)," +
	"upvotes INT," +
	"downvotes INT," +
	"deleted BOOLEAN);"

var avatarTable = "CREATE TABLE avatars (" +
	"id SERIAL PRIMARY KEY," +
	"userid INT REFERENCES users(id)," +
	"avatar BYTEA);"

func createTable(db *sql.DB, table string, label string) error {

	log.Println(Info("Attempting to create"), label)
	_, e := db.Exec(table)
	return e

}