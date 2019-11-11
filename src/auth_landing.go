package main

import (
	"html/template"
	"log"
	"net/http"
)

func usrLandingGET(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("web/auth_landing.html"))
	_ = t.Execute(w, "")

	log.Println("Authenticated landing page arrival cookies: ", r.Cookies())
}

func UserLandingHandler(w http.ResponseWriter, r *http.Request) {

	pushAllResources(w)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	usrLandingGET(w,r)

}