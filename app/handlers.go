package main

import (
	"html/template"
	"net/http"
)

func userLogin(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("css"))
	t := template.Must(template.ParseFiles("web/login.html"))
	_ = t.Execute(w, "")
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "css/main.css")
	push(w, "css/util.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		userLogin(w, r)
	}
}