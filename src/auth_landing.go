package main

import (
	"html/template"
	"log"
	"net/http"
)

func usrLandingGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "LANDING")

	ok, username := CompareTokens(w, r)
	if ok {
		RefreshCookie(w, username) /* This updates cookie to restart clock. */
		t := template.Must(template.ParseFiles("web/auth_landing.html"))
		_ = t.Execute(w, username)
	}

}

func UserLandingHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf(Info("Request to UserLandingHandler from: %s"), GetIP(r))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	usrLandingGET(w,r)

}