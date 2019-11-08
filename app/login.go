package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func userLoginGET(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("css"))
	t := template.Must(template.ParseFiles("web/login.html"))
	_ = t.Execute(w, "")
}

func userLoginPOST(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	log.Printf("%s", r.FormValue("username"))
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "css/main.css")
	push(w, "css/util.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		userLoginGET(w, r)
	case "POST":
		userLoginGET(w, r)
		userLoginPOST(w, r)
	}

}