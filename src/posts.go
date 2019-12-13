package main

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func postsGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST")

	username := CompareTokens(w, r)
	RefreshCookie(w, username) /* This updates cookie to restart clock. */

	limit := 5

	value, ok := r.URL.Query()["limit"]
	if !ok {
		log.Println(Warn("No limit value passed in request."))
	} else {
		log.Println(Success(fmt.Sprintf("Limit value of %v detected.", value)))
		parsedLimit, e := strconv.Atoi(value[0])

		if e != nil {
			log.Println(Warn("Unable to parse page limit."))
		} else {
			log.Println(Success("Successfully parsed page limit."))
			limit = parsedLimit // only set this if no errors
		}
	}


	db, _ := Database(DBNAME)
	defer db.Close()
	code, _ := w.Write(GetPosts(username, db, limit))
	log.Println(Info("Post Response: ", code))

}

func postsPOST(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST")

	username := CompareTokens(w, r)
	RefreshCookie(w, username)

	var postContent = r.FormValue("content")

	db, _ := Database(DBNAME)
	defer db.Close()

	MakePost(username, postContent, db)

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

func GetPosts(username string, db *sql.DB, pagelimit int) []byte {

	var (
		postid    int
		userid    int
		content   string
		upvotes   int
		downvotes int
		deleted   bool
		date      string
	)

	r, e := db.Query("SELECT * FROM posts WHERE userid=(SELECT id FROM users WHERE username=$1) ORDER BY date LIMIT $2;",
		username, pagelimit)

	if e != nil {
		return nil
	}

	var response []Post

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

	posts := json

	log.Println(Info(posts))
	return json
}

func MakePost(username string, post string, db *sql.DB) {

	userid := GetUserId(username, db)

	_, e := db.Exec("INSERT INTO posts (userid, content, upvotes, downvotes, deleted, date) "+
		"VALUES ($1, $2, 0, 0, false, $3);", userid, post, time.Now())
	if e != nil {
		log.Println(Warn("Unable to execute post query."))
		log.Println(Warn(e))
	}

}
