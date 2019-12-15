package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"
)

func messageGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "MESSAGE ENDPOINT (GET)")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username) /* This updates cookie to restart clock. */

	limit := ParseLimitQuery(r, 5)

	db, _ := Database(DBNAME)
	defer db.Close()
	code, _ := w.Write(GetMessages(username, db, limit))
	log.Println(Info("Write-back response: ", code))

}

func messagePOST(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "MESSAGE ENDPOINT (POST)")

	username, ok := CompareTokens(w, r)
	if !ok {
		http.Redirect(w, r, "login", http.StatusSeeOther)
		return
	}

	RefreshCookie(username)

	var toName = r.FormValue("send_to")
	var messageContent = r.FormValue("message")

	db, _ := Database(DBNAME)
	defer db.Close()

	if IsFollowerNameToName(toName, username, db) && IsFollowerNameToName(username, toName, db) {
		PostMessage(username, toName, messageContent, db)
	} else {
		log.Println(Warn("Unable to message user. Not mutual follower."))
	}

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

func GetMessages(username string, db *sql.DB, limit int) []byte {

	var (
		messages []Message
		msg Message
	)

	friends := FetchMutualFollowers(username, db)

	for _, id := range friends {
		r, e := db.Query("SELECT * FROM messages WHERE fromid=$1 LIMIT $2;", id, limit)
		if e != nil {
			log.Println(Warn("Unable to fetch messages for id=", id))
		} else if r != nil {
			for r.Next() {
				_ = r.Scan(&msg.Id, &msg.FromId, &msg.ToId, &msg.Content, &msg.Date, &msg.FriendlyDate)
				messages = append(messages, msg)
			}
		}
	}

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Date.Before(messages[j].Date)
	})

	js, e := json.Marshal(messages)
	if e != nil {
		log.Println(Warn("Error making posts query."))
	}
	log.Println(Info("Post content: ", string(js)))

	return js

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

