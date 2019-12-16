package main

import (
	"encoding/json"
	"html/template"
	"log"
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

	user := GetUserByName(username, db)

	js, e := json.Marshal(user)
	if e != nil {
		log.Println(Warn("Error making user query."))
	}
	log.Println(Info("User content: ", string(js)))

	//code, _ := w.Write(js)
	//log.Println(Info("Write-back response: ", code))

	t := template.Must(template.ParseFiles("web/auth_landing.html"))
	_ = t.Execute(w, user)
}

func UserLandingHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	usrLandingGET(w,r)

}