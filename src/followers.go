package main

import (
	"database/sql"
	"log"
	"net/http"
)

func followGET(w http.ResponseWriter, r *http.Request) {

	

}

func followPOST(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(w, r, username)

	targetUsername := r.FormValue("target")

	db, _ := Database(DBNAME)
	defer db.Close()

	e := FollowUser(username, targetUsername, db)
	if e != nil {
		log.Println(Warn("Failed to follow user."))
		log.Println(Warn(e))
	}

}

func FollowUser(username string, targetUsername string, db *sql.DB) error {

	_, e := db.Exec("INSERT INTO follow (userid, followid) VALUES ((SELECT id FROM users WHERE username=$1), (SELECT id FROM users WHERE username=$2));",
		username, targetUsername)
	if e != nil {
		log.Println(Warn("Unable to execute follow query."))
	}
	return e

}

func IsFollower(username string, targetUsername string, db *sql.DB) bool {


	var count int
	r := db.QueryRow("SELECT count(*) FROM follow WHERE userid=(SELECT id FROM users WHERE username=$1) AND followid=(SELECT id FROM users WHERE username=$2);",
		username, targetUsername)
	_ = r.Scan(&count)
	if count != 0 {
		return true
	}
	return false
}

func FollowHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	followPOST(w, r)
}