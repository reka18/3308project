package main

import (
	"html/template"
	"net/http"
	"strings"
)

func userLogoutGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "LOGOUT")

	t := template.Must(template.ParseFiles("web/logout_success.html"))
	_ = t.Execute(w, "")

	cookie, _ := r.Cookie("socialmediasite")
	values := strings.Split(cookie.Value, ":")
	username := values[0]

	DeleteCookie(w, username)
}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	userLogoutGET(w, r)

}