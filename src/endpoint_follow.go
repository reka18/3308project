package main

import (
	"database/sql"
	json2 "encoding/json"
	"log"
	"net/http"
	"time"
)

func followGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "FOLLOW ENDPOINT (GET)")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username)

	limit := ParseLimitQuery(r, 5)

	db, _ := Database(DBNAME)
	defer db.Close()

	user := ParseUserQuery(r)
	if user != "" {
		e := FollowUser(username, user, db)
		if e != nil {
			log.Println(Warn("Unable to follow ", user))
		}
		w.WriteHeader(http.StatusOK)
	}

	// TODO display followed
	followList := FetchFollowed(username, db, limit)
	_, _ = w.Write(followList)

}

func FollowHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		followGET(w, r)
	case "POST":
		log.Println(Warn("Follow endpoint has no POST method."))
	}

}

func FollowUser(username string, targetname string, db *sql.DB) error {

	_, e := db.Exec("INSERT INTO follow (userid, followid, date) VALUES ((SELECT id FROM users WHERE username=$1), (SELECT id FROM users WHERE username=$2), $3);",
		username, targetname, time.Now())
	if e != nil {
		log.Println(Warn("Already followed."))
	}
	return e

}

func FetchFollowed(username string, db *sql.DB, limit int) []byte {

	var followid int

	r, e := db.Query("SELECT followid FROM follow WHERE userid=(SELECT id FROM users WHERE username=$1) ORDER BY date LIMIT $2;",
		username, limit)

	if e != nil {
		return nil
	}
	
	var response []FollowedUser

	for r.Next() {
		_ = r.Scan(&followid)

		var f FollowedUser
		f.User = GetUserById(followid, db)
		f.Mutual = IsFollowerIdToName(followid, username, db)

		response = append(response, f)
	}
	json, e := json2.Marshal(response)
	if e != nil {
		log.Println(Warn("Error getting followed users."))
	}
	log.Println(Info("Followed users: ", string(json)))

	return json

}

func FetchFollowedIds(username string, db *sql.DB) map[int]bool {

	var followid int

	r, e := db.Query("SELECT followid FROM follow WHERE userid=(SELECT id FROM users WHERE username=$1) ORDER BY date;",
		username)

	if e != nil {
		return nil
	}

	idMap := make(map[int]bool)

	for r.Next() {
		_ = r.Scan(&followid)

		idMap[followid] = true
	}
	return idMap

}

func FetchMutualFollowers(username string, db *sql.DB) []int {

	var id int
	var ids []int

	r, e := db.Query("SELECT followid FROM follow WHERE userid=(SELECT id FROM users WHERE username=$1);",
		username)
	if e != nil {
		log.Println(Warn("Unable to retrieve relevant ids."))
	}
	if r != nil {
		for r.Next() {
			_ = r.Scan(&id)
			if IsFollowerIdToName(id, username, db) {
				ids = append(ids, id)
			}
		}
	}
	return ids
}

func IsFollowerIdToName(follower int, followed string, db *sql.DB) bool {

	r, e := db.Exec("SELECT * FROM follow WHERE userid=$1 AND followid=(SELECT id FROM users WHERE username=$2);",
		follower, followed)
	if e != nil {
		log.Println(Warn("Error making follow status database query."))
	}
	if r != nil {
		count, e := r.RowsAffected()
		if e != nil {
			log.Println(Warn("Error getting follow status count from database response."))
		}
		if count != 0 {
			return true
		}
	}
	return false

}

func IsFollowerNameToName(follower string, followed string, db *sql.DB) bool {

	r, e := db.Exec("SELECT * FROM follow WHERE userid=(SELECT id FROM users WHERE username=$1) AND followid=(SELECT id FROM users WHERE username=$2);",
		follower, followed)
	if e != nil {
		log.Println(Warn("Error making follow status database query."))
	}
	if r != nil {
		count, e := r.RowsAffected()
		if e != nil {
			log.Println(Warn("Error getting follow status count from database response."))
		}
		if count != 0 {
			return true
		}
	}
	return false

}