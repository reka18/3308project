package main

import (
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

	r.ParseForm()

	var (
		email = strings.Join(r.Form["email"], "")
		password = strings.Join(r.Form["pass"], "")
	)

	db, _ := Database(DBNAME)
	defer db.Close()

	verified, e := LoginUserAccount(email, password, db)
	if e != nil || !verified {
		log.Println("User login failed.")
		t := template.Must(template.ParseFiles("web/login.html"))
		_ = t.Execute(w, "Incorrect email/password combination")
	} else {
		c := http.Cookie {
			Name:       Encrypt(email + "thisisasaltstring"),
			Value:      Encrypt(password + "thisisasaltstring"),
		}
		log.Printf("Cookie: %v", &c)
		http.SetCookie(w, &c)

		t := template.Must(template.ParseFiles("web/landing.html"))
		t.Execute(w, "")
	}
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {

	push(w, "css/main.css")
	push(w, "css/util.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		userLoginGET(w, r)
	case "POST":
		userLoginPOST(w, r)
	}

}