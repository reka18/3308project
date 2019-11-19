package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"fmt"
)

func usrLandingGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "LANDING")

	ok, username := CompareTokens(w, r)
	if ok {
		userInfo, e := loadUserInfo(username)
		if e == nil {
			RefreshCookie(w, username) /* This updates cookie to restart clock. */
			t := template.Must(template.ParseFiles("web/auth_landing.html"))
			_ = t.Execute(w, userInfo)
		} else {
			http.Redirect(w, r, "/logout", http.StatusUnauthorized)
		}
	} 
}

func UserLandingHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	usrLandingGET(w,r)

}

func loadUserInfo(username string) (User, error) {
	/* 
	GET USER INFORMATION FROM AN EXISTING DATABASE CONNECTION 
	*/
	db, _ := Database(DBNAME)
		defer db.Close()

	var user User
	var sqlStatement = fmt.Sprintf(`SELECT firstname, lastname, username, age, gender, public, joindate, 
	active FROM users WHERE username='%s';`, username) 

	row := db.QueryRow(sqlStatement)
	e := row.Scan(&user.Firstname, &user.Lastname, &user.Username, &user.Age, &user.Gender, &user.Public, &user.Joindate, &user.Active)
	
	if e != nil {
		if e == sql.ErrNoRows {
			log.Printf(Warn("No user information found for %s", username))
			return user, e
		} else {
			log.Printf(Warn(e))
			return user, e
		}
	}
	return user, nil

}