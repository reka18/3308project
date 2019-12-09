package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func createUserAccountGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "CREATE ACCOUNT")

	t := template.Must(template.ParseFiles("web/create_account.html"))
	_ = t.Execute(w, "")

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
		log.Println(Warn("Passwords do not match."))
		t := template.Must(template.ParseFiles("web/create_account.html"))
		_ = t.Execute(w, "Passwords do not match.")
	} else {
		db, _ := Database(DBNAME)
		defer db.Close()

		e := AddNewUserAccount(age, firstname, lastname, email, username, true, GenerateKey(password), gender, db)
		if e != nil {
			log.Printf(Warn("User creation failed."))
			log.Println(Warn(e))
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
	img, _ := os.Open("web/images/default_avatar.png")
	defer img.Close()
	imgInfo, e := img.Stat()
	if e != nil {
		log.Println(Warn("Unable to load image."))
		log.Println(Warn(e))
		return e
	}
	size := imgInfo.Size()
	bytes := make([]byte, size)
	_, e = db.Exec("INSERT INTO avatars (userid, avatar) VALUES ($1, $2);",
		id, bytes)
	if e != nil {
		log.Println(Warn("Unable to execute image query."))
		log.Println(Warn(e))
		return e
	}
	log.Println(id)
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