package main

import (
	"html/template"
	"net/http"
)

func usrLandingGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "LANDING PAGE (GET)")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username) /* This updates cookie to restart clock. */

	db, _ := Database(DBNAME)
	defer db.Close()

	user := GetUserByNameJson(username, db)
	_, _ = w.Write(user)

	t := template.Must(template.ParseFiles("web/auth_landing.html"))
	_ = t.Execute(w, username)
}

func UserLandingHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	usrLandingGET(w,r)

}