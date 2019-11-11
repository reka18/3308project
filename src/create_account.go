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
		age, _ = strconv.Atoi(strings.Join(r.Form["age"], ""))
		firstname = strings.Join(r.Form["firstname"], "")
		lastname = strings.Join(r.Form["lastname"], "")
		email = strings.Join(r.Form["email"], "")
		username = strings.Join(r.Form["username"], "")
		gender = strings.Join(r.Form["gender"], "")
		password = strings.Join(r.Form["pass"], "")
	)

	log.Println(gender)
	log.Println(username)

	db, _ := Database(DBNAME)
	defer db.Close()

	e := AddNewUserAccount(age, firstname, lastname, email, username, gender, true, password, db)
	if e != nil {
		log.Printf("User creation failed with error: %s", e)
		t := template.Must(template.ParseFiles("web/create_account.html"))
		_ = t.Execute(w, "Please fill out all fields")
	} else {
		SecureCookieController(w, email)
		t := template.Must(template.ParseFiles("web/account_created.html"))
		_ = t.Execute(w, "")

		log.Println("Account created page arrival cookies: ", r.Cookies())
	}

}

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {

	pushAllResources(w)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		createUserAccountGET(w, r)
	case "POST":
		createUserAccountPOST(w, r)
	}

}

func AddNewUserAccount(age int, firstname string, lastname string, email string, username string, gender string,
	public bool, password string, db *sql.DB) error {
	/*
		THIS CONNECTS TO THE DATABASE AND ADDS A USER
	*/

	query := fmt.Sprintf("INSERT INTO users (" +
		"age, firstname, lastname, email, username, gender, public, active, password)"+
		"VALUES (%d, '%s', '%s', '%s', '%s' ,'%s', '%t', '%t', '%s');",
		age, firstname, lastname, email, username, gender, public, true, password)
	_, e := db.Query(query)
	if e != nil {
		log.Println("Unable to execute query:", e)
		return e
	} else {
		log.Printf("Successfully added User <%s> to Database.", email)
	}

	return e
}