package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func GetUserById(userid int, db *sql.DB) User {

	var password string // we don't do anything with this

	var user User

	e := db.QueryRow("SELECT * FROM users WHERE id=$1;",
		userid).Scan(&user.Id, &user.Age, &user.Firstname, &user.Lastname, &user.Email,
		&user.Username, &user.Public, &user.Joindate, &user.Active, &password, &user.Gender)
	if e != nil {
		log.Println(Warn("Unable to fetch user from database."))
		return User{}
	}

	timestamp := strings.Split(user.Joindate.String(), " ")
	date := timestamp[0]
	clock := strings.Split(timestamp[1], ".")[0][:5]
	user.FriendlyJoinDate = fmt.Sprintf("%s @ %s", date, clock)

	return user

}

func GetUserByName(username string, db *sql.DB) User {

	var password string // we don't do anything with this

	var user User

	e := db.QueryRow("SELECT * FROM users WHERE id=(SELECT id FROM users WHERE username=$1);",
		username).Scan(&user.Id, &user.Age, &user.Firstname, &user.Lastname, &user.Email,
		&user.Username, &user.Public, &user.Joindate, &user.Active, &password, &user.Gender)
	if e != nil {
		log.Println(Warn("Unable to fetch user from database."))
		return User{}
	}

	timestamp := strings.Split(user.Joindate.String(), " ")
	date := timestamp[0]
	clock := strings.Split(timestamp[1], ".")[0][:5]
	user.FriendlyJoinDate = fmt.Sprintf("%s @ %s", date, clock)

	return user

}