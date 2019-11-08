package main

import (
	"html/template"
	"log"
	"net/http"
)

func userLoginGET(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("web/login.html"))
	_ = t.Execute(w, "")

}

func userLoginPOST(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	log.Println("username:", r.Form["username"])
	log.Println("password:", r.Form["pass"])

	t := template.Must(template.ParseFiles("web/login.html"))
	t.Execute(w, "")

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