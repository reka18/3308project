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

	log.Println("Authenticated login page arrival cookies: ", r.Cookies())

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
			Name:	"login_cookie",
		}
		log.Printf("Cookie: %v", &c)
		http.SetCookie(w, &c)

		http.Redirect(w, r, "/", 303)

		log.Println("Create account page arrival cookies: ", r.Cookies())
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

func LoginUserAccount(inputUsername string, inputPassword string, db *sql.DB) (User, bool, error) {

	var (
		user User
	)
	/*
		FAST FAIL IF EMAIL OR PASSWORD ARE BLANK
	*/
	if len(inputUsername) == 0 || len(inputPassword) == 0 {
		log.Println("Email and/or password blank.")
		return user, false, &EmptyStringError{}
	}

	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s';", inputUsername)

	r := db.QueryRow(query)

	var (
		password	string
	)

	e := r.Scan(&password)

	// user = UserBuilder(id, firstname, lastname, email, gender, public, joindate, active)

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