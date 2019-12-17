package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func userGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST ENDPOINT (GET)")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username) /* This updates cookie to restart clock. */

	db, _ := Database(DBNAME)
	defer db.Close()

	userQuery := ParseUserQuery(r)

	user := FetchUser(userQuery, db)
	_, _ = w.Write(user)

}

func UserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		userGET(w, r)
	case "POST":
		log.Println(Warn("User endpoint has no post method."))
	}

}

func FetchUser(username string, db *sql.DB) []byte {

	var (
		user User
		ignore string
	)

	r := db.QueryRow("SELECT * FROM users WHERE username=$1;", username)
	e := r.Scan(&user.Id, &user.Age, &user.Firstname, &user.Lastname, &user.Email, &user.Username, &user.Public,
		&user.Joindate, &user.Active, &ignore, &user.Gender)
	if e != nil {
		log.Printf(Warn("Unable to find '%s'"), username)
	}

	timestamp := strings.Split(user.Joindate.String(), " ")
	date := timestamp[0]
	clock := strings.Split(timestamp[1], ".")[0][:5]
	user.FriendlyJoinDate = fmt.Sprintf("%s @ %s", date, clock)

	js, e := json.Marshal(user)
	if e != nil {
		log.Println(Warn("Error querying user."))
	}
	log.Println(Info("User content: ", string(js)))

	return js

}