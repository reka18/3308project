package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func createUserAccountGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "CREATE ACCOUNT (GET)")

	w.WriteHeader(http.StatusOK)
	t := template.Must(template.ParseFiles("web/create_account.html"))
	_ = t.Execute(w, "")

}

func createUserAccountPOST(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "CREATE ACCOUNT (POST)")

	_ = r.ParseForm()

	var (
		age, _          = strconv.Atoi(r.FormValue("age"))
		firstname       = r.FormValue("firstname")
		lastname        = r.FormValue("lastname")
		email           = r.FormValue("email")
		username        = r.FormValue("username")
		password        = r.FormValue("pass")
		confirmPassword = r.FormValue("confirm_pass")
		gender			= r.FormValue("gender")
	)

	if password != confirmPassword {
		log.Println(Warn("Passwords do not match."))
		t := template.Must(template.ParseFiles("web/create_account.html"))
		_ = t.Execute(w, "Passwords do not match")
	} else {
		db, _ := Database(DBNAME)
		defer db.Close()

		e := AddNewUserAccount(age, firstname, lastname, email, username, true, GenerateKey(password), gender, db)
		if e != nil {
			log.Printf(Warn("User creation failed."))
			log.Println(Warn(e))
			t := template.Must(template.ParseFiles("web/create_account.html"))

			err := e.Error()

			if err == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
				_ = t.Execute(w, "Username already exists")
			} else if err == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
				_ = t.Execute(w, "Email already exists")
			} else {
				_ = t.Execute(w, "Please fill out all fields")
			}
		} else {
			t := template.Must(template.ParseFiles("web/account_created.html"))
			_ = t.Execute(w, "")
		}
	}
}

func AddNewUserAccount(age int, firstname string, lastname string, email string, username string,
	public bool, password string, gender string, db *sql.DB) error {
	/*
	THIS CONNECTS TO THE DATABASE AND ADDS A USER
	*/
	var id int
	e := db.QueryRow("INSERT INTO users (" +
		"age, firstname, lastname, email, username, public, active, password, gender, joindate)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;",
		age, firstname, lastname, email, username, public, true, password, gender, time.Now()).Scan(&id)
	if e != nil {
		log.Println(Warn("Unable to execute user query."))
		log.Println(Warn(e))
		return e
	}
	NewUserAvatar(username, db)
	return e

}

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		createUserAccountGET(w, r)
	case "POST":
		createUserAccountPOST(w, r)
	}

}