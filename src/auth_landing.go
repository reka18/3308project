package main

import (
	"html/template"
	"log"
	"net/http"
)

func usrLandingGET(w http.ResponseWriter, r *http.Request) {

	log.Println("Landing page get cookies: ", r.Cookies())

	if CompareTokens(w, r) {
		t := template.Must(template.ParseFiles("web/auth_landing.html"))
		_ = t.Execute(w, "")
	}

}

func UserLandingHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	usrLandingGET(w,r)

}