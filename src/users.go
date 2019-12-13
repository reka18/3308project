package main

import (
	"database/sql"
	"log"
)

func GetUser(userid int, username string, db *sql.DB) User {

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

	e := db.QueryRow("SELECT * FROM users WHERE id=$1 OR username=$2;",
		userid, username).Scan(&uid, &uname, &firstname, &lastname, &age, &email,
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