package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func createUserAccountGET(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("web/new_account.html"))
	_ = t.Execute(w, "")

}

func createUserAccountPOST(w http.ResponseWriter, r *http.Request) {
	/*
	THIS CREATES A NEW USER IN THE DATABASE
	TODO: Certain fields need to have constraints added. Age cannot be negative,
	TODO: names cannot be blank or too long, email must be a valid email, gender
	TODO: should be restricted to M, F, or O (it already is in the database but
	TODO: lets catch it earlier, and password should have certain constraints.
	 */
	_ = r.ParseForm()

	var (
		age, _ = strconv.Atoi(strings.Join(r.Form["age"], ""))
		firstname = strings.Join(r.Form["firstname"], "")
		lastname = strings.Join(r.Form["lastname"], "")
		email = strings.Join(r.Form["email"], "")
		gender = strings.Join(r.Form["gender"], "")
		password = strings.Join(r.Form["pass"], "")
	)

	db, _ := Database(DBNAME)
	defer db.Close()

	e := AddNewUserAccount(age, firstname, lastname, email, gender, true, password, db)
	if e != nil {
		log.Printf("User creation failed with error: %s", e)
		t := template.Must(template.ParseFiles("web/new_account.html"))
		_ = t.Execute(w, "Please fill out all fields")
	} else {
		t := template.Must(template.ParseFiles("web/account_creation_success.html"))
		_ = t.Execute(w, "")
	}

}

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {

	push(w, "css/main.css")
	push(w, "css/util.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		createUserAccountGET(w, r)
	case "POST":
		createUserAccountPOST(w, r)
	}

}

func AddNewUserAccount(age int, firstname string, lastname string,
	email string, gender string, public bool, password string, db *sql.DB) error {
	/*
		THIS CONNECTS TO THE DATABASE AND ADDS A USER
	*/
	q := fmt.Sprintf("INSERT INTO user_account(age, firstname, lastname, email, "+
		"gender, public, joindate, active, password)"+
		"VALUES (%d, '%s', '%s', '%s', '%s', '%t', now(), true, '%s');",
		age, firstname, lastname, email, gender, public, Encrypt(password))
	_, e := db.Query(q)
	if e != nil {
		log.Println("Unable to execute query:", e)
	} else {
		log.Printf("Successfully added User <%s> to Database.", email)
	}
	return e

}