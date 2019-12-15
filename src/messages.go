package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

func messageGET(w http.ResponseWriter, r *http.Request) {



}

func messagePOST(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "MESSAGE ENDPOINT (POST)")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username)

}

func MessageHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		messageGET(w, r)
	case "POST":
		messagePOST(w, r)
	}

}

func GetMessages() {



}

func PostMessage(username string, toname string, content string, db *sql.DB) {

	r, e := db.Exec("INSERT INTO messages (fromid, toid, content, date) VALUES ((SELECT id FROM users WHERE username=$1), (SELECT id FROM users WHERE username=$2), $3, $4);",
		username, toname, content, time.Now())
	if e != nil {
		log.Println(Warn("Unable to make message in database."))
		log.Println(Warn(e))
	}

	c, _ := r.RowsAffected()
	if c != 0 {
		log.Println(Success("Successfully added message top database."))
	}

}