package main

import (
	"database/sql"
	json2 "encoding/json"
	"log"
	"net/http"
)

func followGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "FOLLOW ENDPOINT (GET)")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username)

	limit := ParseLimit(r, 5)

	db, _ := Database(DBNAME)
	defer db.Close()
	code, _ := w.Write(GetFollowed(username, db, limit))
	log.Println(Info("Write-back response: ", code))

}

func followPOST(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "FOLLOW ENDPOINT (POST)")

	_ = r.ParseForm()

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username)

	targetUsername := r.FormValue("target")

	db, _ := Database(DBNAME)
	defer db.Close()

	e := FollowUser(username, targetUsername, db)
	if e != nil {
		log.Println(Warn("Failed to follow user."))
		log.Println(Warn(e))
	}

}

func FollowHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		followGET(w, r)
	case "POST":
		followPOST(w, r)
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

func GetFollowed(username string, db *sql.DB, limit int) []byte {

	var followid int

	r, e := db.Query("SELECT followid FROM follow WHERE userid=(SELECT id FROM users WHERE username=$1) ORDER BY date LIMIT $2;",
		username, limit)

	if e != nil {
		return nil
	}
	
	var response []User
	
	for r.Next() {
		_ = r.Scan(&followid)

		response = append(response, GetUser(followid, "", db))
	}
	json, e := json2.Marshal(response)
	if e != nil {
		log.Println(Warn("Error getting followed users."))
	}
	log.Println(Info("Followed users: ", json))

	return json

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