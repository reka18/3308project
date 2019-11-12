package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func userLogoutGET(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("web/logout_success.html"))
	_ = t.Execute(w, "")

	log.Println("Logout page arrival cookies: ", r.Cookies())

	cookie, _ := r.Cookie("socialmediasite")
	values := strings.Split(cookie.Value, ":::")
	username := values[0]

	DeleteCookie(w, username)
}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	userLogoutGET(w, r)

}