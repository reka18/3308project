package main

import (
	"html/template"
	"net/http"
	"strings"
)

func userLogoutGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "LOGOUT PAGE (GET)")

	t := template.Must(template.ParseFiles("web/logout_success.html"))
	_ = t.Execute(w, "")

	cookie, _ := r.Cookie("screebit")
	values := strings.Split(cookie.Value, ":")
	username := values[0]

	DeleteCookie(username)
}

func UserLogoutHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	userLogoutGET(w, r)

}