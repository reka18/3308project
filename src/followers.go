package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
)

func followPost(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()

	username := CompareTokens(w, r)

	RefreshCookie(w, username)

	targetUsername := strings.Join(r.Form["target_username"], "")

	db, _ := Database(DBNAME)
	defer db.Close()

	e := FollowUser(username, targetUsername, db)
	if e != nil {
		log.Println(Warn("Failed to follow user."))
		log.Println(Warn(e))
	}

}

func FollowUser(username string, targetUsername string, db *sql.DB) error {

	userid := GetUserId(username, db)
	target := GetUserId(targetUsername, db)

	_, e := db.Exec("INSERT INTO follow (userid, followid) VALUES ($1, $2);", userid, target)
	if e != nil {
		log.Println(Warn("Unable to execute follow query."))
	}
	return e

}

func IsFollower(username string, targetUsername string, db *sql.DB) bool {

	userid := GetUserId(username, db)
	targetid := GetUserId(targetUsername, db)

	var count int
	r := db.QueryRow("SELECT count(*) FROM follow WHERE userid=$1 AND followid=$2;", targetid, userid)
	_ = r.Scan(&count)
	if count != 0 {
		return true
	}
	return false
}

func FollowHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	followPost(w, r)
}