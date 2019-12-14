package main

import (
	"database/sql"
	"log"
)

func GetUserByName(username string, db *sql.DB) User {

	var (
		uid			int
		uname		string
		firstname	string
		lastname	string
		age			int
		email		string
		gender		string
		public		bool
		joindate	string
		active		bool
	)

	e := db.QueryRow("SELECT * FROM users WHERE username=$1;",
		username).Scan(&uid, &uname, &firstname, &lastname, &age, &email,
			&gender, &public, &joindate, &active)
	if e != nil {
		log.Println(Warn("Unable to fetch user from database."))
		return User{}
	}

	return User{
		Id:        uid,
		Username:  uname,
		Firstname: firstname,
		Lastname:  lastname,
		Age:       age,
		Email:     email,
		Gender:    gender,
		Public:    public,
		Joindate:  joindate,
		Active:    active,
	}

}

func GetUserById(userid int, db *sql.DB) User {

	var (
		uid			int
		age			int
		firstname	string
		lastname	string
		email		string
		username	string
		public		bool
		joindate	string
		active		bool
		password	string	// we ignore this
		gender		string
	)

	e := db.QueryRow("SELECT * FROM users WHERE id=$1;",
		userid).Scan(&uid, &age, &firstname, &lastname, &email,
			&username, &public, &joindate, &active, &password, &gender)
	if e != nil {
		log.Println(Warn("Unable to fetch user from database."))
		return User{}
	}

	return User{
		Id:        uid,
		Username:  username,
		Firstname: firstname,
		Lastname:  lastname,
		Age:       age,
		Email:     email,
		Gender:    gender,
		Public:    public,
		Joindate:  joindate,
		Active:    active,
	}

}