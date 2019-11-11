package main

import (
	"database/sql"
	"fmt"
	"log"
)

var genderEnum = "CREATE TYPE gender " +
	"AS ENUM ('M', 'F', 'O');"

var usersTable = "CREATE TABLE users (" +
	"id SERIAL PRIMARY KEY," +
	"age INT NOT NULL," +
	"firstName TEXT NOT NULL," +
	"lastName TEXT NOT NULL," +
	"email TEXT UNIQUE NOT NULL," +
	"gender gender NOT NULL," +
	"public BOOLEAN NOT NULL," +
	"joinDate DATE NOT NULL," +
	"active BOOLEAN NOT NULL," +
	"password TEXT NOT NULL" +
	");"

var postTable = "CREATE TABLE posts (" +
	"id SERIAL PRIMARY KEY," +
	"user_id INT REFERENCES users(id)," +
	"content VARCHAR(240)," +
	"upvotes INT," +
	"downvotes INT," +
	"deleted BOOLEAN" +
	");"

func createTable(db *sql.DB, table string, label string) {

	query := fmt.Sprintf(table)
	_, e := db.Query(query)
	if e == nil {
		log.Printf("%s created successfully.", label)
	} else {
		log.Println("Unable to create table:", e)
	}

}