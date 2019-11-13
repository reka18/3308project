package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func createUserAccountGET(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("web/create_account.html"))
	_ = t.Execute(w, "")

	log.Println("Create account page arrival cookies: ", r.Cookies())
}

func createUserAccountPOST(w http.ResponseWriter, r *http.Request) {
	/*
	THIS CREATES A NEW USER IN THE DATABASE
	 */
	_ = r.ParseForm()

	var (
		age, _          = strconv.Atoi(strings.Join(r.Form["age"], ""))
		firstname       = strings.Join(r.Form["firstname"], "")
		lastname        = strings.Join(r.Form["lastname"], "")
		email           = strings.Join(r.Form["email"], "")
		username        = strings.Join(r.Form["username"], "")
		password        = strings.Join(r.Form["pass"], "")
		confirmPassword = strings.Join(r.Form["confirm_pass"], "")
		gender			= strings.Join(r.Form["gender"], "")
	)

	if password != confirmPassword {
		log.Println("Passwords do not match.")
		t := template.Must(template.ParseFiles("web/create_account.html"))
		_ = t.Execute(w, "Passwords do not match.")
	} else {
		db, _ := Database(DBNAME)
		defer db.Close()

		e := AddNewUserAccount(age, firstname, lastname, email, username, true, password, gender, db)
		if e != nil {
			log.Printf("User creation failed with error: %s", e)
			t := template.Must(template.ParseFiles("web/create_account.html"))
			_ = t.Execute(w, "Please fill out all fields")
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
	_, e := db.Exec("INSERT INTO users (" +
		"age, firstname, lastname, email, username, public, active, password, gender, joindate)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);",
		age, firstname, lastname, email, username, public, true, password, gender, time.Now())
	if e != nil {
		log.Println("Unable to execute query:", e)
		return e
	} else {
		log.Printf("Successfully added User <%s> to Database.", email)
	}

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