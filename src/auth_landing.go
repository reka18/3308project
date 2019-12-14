package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

func usrLandingGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "LANDING PAGE GET")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username) /* This updates cookie to restart clock. */

	// userInfo := loadUserInfo(username)
	t := template.Must(template.ParseFiles("web/auth_landing.html"))
	_ = t.Execute(w, username)
}

func UserLandingHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	usrLandingGET(w,r)

}

func loadUserInfo(username string) User {
	/* 
	GET USER INFORMATION
	*/
	db, _ := Database(DBNAME)
	defer db.Close()

	var user User

	row := db.QueryRow("SELECT firstname, lastname, username, age, gender, public, joindate, " +
		"active FROM users WHERE username=$1;", username)
	e := row.Scan(&user.Firstname, &user.Lastname, &user.Username, &user.Age, &user.Gender,
		&user.Public, &user.Joindate, &user.Active)
	
	if e != nil {
		if e == sql.ErrNoRows {
			log.Printf(Warn("No user information found for %s", username))
		} else {
			log.Printf(Warn(e))
		}
	}
	return user

}