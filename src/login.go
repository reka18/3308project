package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func userLoginGET(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("web/login.html"))
	_ = t.Execute(w, "")

}

func userLoginPOST(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()

	var (
		email = strings.Join(r.Form["email"], "")
		password = strings.Join(r.Form["pass"], "")
	)

	db, _ := Database(DBNAME)
	defer db.Close()

	_, v, e := LoginUserAccount(email, password, db)
	if e != nil || !v {
		log.Println("User login failed.")
		t := template.Must(template.ParseFiles("web/login.html"))
		_ = t.Execute(w, "Incorrect email/password combination")
	} else {
		c := http.Cookie {
			Name:	email,
		}
		log.Printf("Cookie: %v", &c)
		http.SetCookie(w, &c)

		t := template.Must(template.ParseFiles("web/auth_landing.html"))
		_ = t.Execute(w, "")
	}
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	pushAllResources(w)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		userLoginGET(w, r)
	case "POST":
		userLoginPOST(w, r)
	}

}

func LoginUserAccount(inputEmail string, inputPassword string, db *sql.DB) (User, bool, error) {

	var (
		user User
	)
	/*
		FAST FAIL IF EMAIL OR PASSWORD ARE BLANK
	*/
	if len(inputEmail) == 0 || len(inputPassword) == 0 {
		log.Println("Email and/or password blank.")
		return user, false, &EmptyStringError{}
	}

	query := fmt.Sprintf("SELECT * FROM user_account WHERE email='%s';", inputEmail)

	r := db.QueryRow(query)

	var (
		id        int
		age       int
		firstname string
		lastname  string
		email     string
		gender    string
		public    bool
		joindate  string
		active    bool
		password  string
	)

	e := r.Scan(&id, &age, &firstname, &lastname, &email, &gender, &public, &joindate, &active, &password)

	user = UserBuilder(id, firstname, lastname, email, gender, public, joindate, active)

	if e != nil {
		log.Println("Email not found: ", e)
		return user, false, e
	}
	if VerifyPW(password, inputPassword) {
		log.Println("Password verified.")
		return user, true, e
	}
	return user, false, e

}