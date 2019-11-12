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

	log.Println("Login page get cookies: ", r.Cookies())

}

func userLoginPOST(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()

	var (
		login = strings.Join(r.Form["login"], "")
		password = strings.Join(r.Form["pass"], "")
	)

	db, _ := Database(DBNAME)
	defer db.Close()

	username, v, e := LoginUserAccount(login, password, db)
	if e != nil || !v {
		log.Println("User login failed.")
		t := template.Must(template.ParseFiles("web/login.html"))
		_ = t.Execute(w, "Incorrect email/password combination")
	} else {
		AddCookie(w, username)

		userPage := fmt.Sprintf("/%s", username)
		http.Redirect(w, r, userPage, 303)

		log.Println("Login page post cookies: ", r.Cookies())
	}
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	pushAllResources(w) /* This is the only place we do this since it is the landing page */
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		userLoginGET(w, r)
	case "POST":
		userLoginPOST(w, r)
	}

}

func LoginUserAccount(inputUsernameOrEmail string, inputPassword string, db *sql.DB) (string, bool, error) {

	/*
	FAST FAIL IF USERNAME/EMAIL OR PASSWORD ARE BLANK
	*/
	if len(inputUsernameOrEmail) == 0 || len(inputPassword) == 0 {
		log.Println("Username and/or password blank.")
		return "", false, &EmptyStringError{}
	}

	/* Login with either username OR email. */
	query := fmt.Sprintf("SELECT password, username, email FROM users WHERE email = '%s' OR username = '%s';",
		inputUsernameOrEmail, inputUsernameOrEmail)

	r := db.QueryRow(query)

	var (
		password	string
		username	string
		email		string
	)

	e := r.Scan(&password, &username, &email)

	if e != nil {
		log.Println("Account not found: ", e)
		return "", false, e
	}
	if VerifyKey(password, inputPassword) {
		log.Println("Password verified.")
		/* Here we actually want the username. */
		return username, true, e
	}
	/* Here we just want to return whatever the user passed for error logging */
	return inputUsernameOrEmail, false, e

}