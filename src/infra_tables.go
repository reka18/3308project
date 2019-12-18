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
	"gender GENDER NOT NULL" +
	");"

var postTable = "CREATE TABLE posts (" +
	"id SERIAL PRIMARY KEY," +
	"userid INT REFERENCES users(id) NOT NULL," +
	"content VARCHAR(240) NOT NULL," +
	"upvotes INT NOT NULL," +
	"downvotes INT NOT NULL," +
	"deleted BOOLEAN NOT NULL," +
	"date TIMESTAMP NOT NULL" +
	");"

var avatarTable = "CREATE TABLE avatars (" +
	"id SERIAL PRIMARY KEY," +
	"userid INT REFERENCES users(id) NOT NULL," +
	"avatar BYTEA NOT NULL" +
	");"

var followTable = "CREATE TABLE follow (" +
	"PRIMARY KEY (userid, followid)," + // element id's are never referenced so we can do without ids
	"userid INT REFERENCES users(id) NOT NULL," +
	"followid INT REFERENCES users(id) NOT NULL," +
	"date TIMESTAMP NOT NULL" +
	");"

var messageTable = "CREATE TABLE messages (" +
	"id SERIAL PRIMARY KEY," +
	"fromid INT REFERENCES users(id) NOT NULL," +
	"toid INT REFERENCES users(id) NOT NULL," +
	"content VARCHAR(240) NOT NULL," +
	"date TIMESTAMP NOT NULL" +
	");"

var votesTable = "CREATE TABLE votes (" +
	"id SERIAL PRIMARY KEY," +
	"postid INT REFERENCES posts(id) NOT NULL," +
	"userid INT REFERENCES users(id) NOT NULL" +
	");"

func createTable(db *sql.DB, table string, label string) error {

	log.Println(Info("Creating ", label))
	_, e := db.Exec(table)
	return e

}