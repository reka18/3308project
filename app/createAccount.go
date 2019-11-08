package main

import (
	"html/template"
	"log"
	"net/http"
)

func createUserAccountGET(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("web/new_account.html"))
	_ = t.Execute(w, "")

}

func createUserAccountPOST(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	log.Println("firstname:", r.Form["firstname"])
	log.Println("lastname:", r.Form["lastname"])
	log.Println("email:", r.Form["email"])
	log.Println("password:", r.Form["pass"])

	t := template.Must(template.ParseFiles("web/new_account.html"))
	_ = t.Execute(w, "")

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