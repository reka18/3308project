package main

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func postsGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST")

	username := CompareTokens(w, r)
	RefreshCookie(w, username) /* This updates cookie to restart clock. */

	db, _ := Database(DBNAME)
	defer db.Close()
	code, _ := w.Write(GetPosts(username, db))
	log.Println(Info("Post Response: ", code))

}

func postsPOST(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST")

	username := CompareTokens(w, r)
	RefreshCookie(w, username)

	var postContent = r.FormValue("content")

	db, _ := Database(DBNAME)
	defer db.Close()

	makePost(username, postContent, db)

	userPage := fmt.Sprintf("/%s", username)
	http.Redirect(w, r, userPage, http.StatusSeeOther)
}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		postsGET(w, r)
	case "POST":
		postsPOST(w, r)
	}
}

func makePost(username string, post string, db *sql.DB) {

	userid := GetUserId(username, db)

	_, e := db.Exec("INSERT INTO posts (userid, content, upvotes, downvotes, deleted, date) "+
		"VALUES ($1, $2, 0, 0, false, $3);", userid, post, time.Now())
	if e != nil {
		log.Println(Warn("Unable to execute user query."))
		log.Println(Warn(e))
	}

}

func GetPosts(username string, db *sql.DB) []byte {

	var (
		postid    int
		userid    int
		content   string
		upvotes   int
		downvotes int
		deleted   bool
		date      string
	)

	r, _ := db.Query("SELECT * FROM posts WHERE userid=(SELECT id FROM users WHERE username=$1) ORDER BY date LIMIT 35;", username)

	var response []Post

	if !r.Next() {
		return nil
	}

	for r.Next() {
		_ = r.Scan(&postid, &userid, &content, &upvotes, &downvotes, &deleted, &date)

		p := Post{
			Id:        postid,
			UserId:    userid,
			Content:   content,
			UpVotes:   upvotes,
			DownVotes: downvotes,
			Date:      date,
		}
		response = append(response, p)
	}
	json, e := json2.Marshal(response)
	if e != nil {
		log.Println(Warn("Error making posts query."))
	}

	posts := string(json)

	log.Println(Info(posts))
	return json
}
