package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func voteGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST ENDPOINT (GET)")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username) /* This updates cookie to restart clock. */

	db, _ := Database(DBNAME)
	defer db.Close()

	vote, postid := ParseVoteQuery(r)

	votes := CastVote(vote, postid, username, db)
	log.Println(Fail(string(votes)))

	http.Redirect(w, r, fmt.Sprintf("/%s", username), http.StatusSeeOther)

}

func VoteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		voteGET(w, r)
	case "POST":
		log.Println(Warn("Vote endpoint has no post method."))
	}

}

func CastVote(vote string, postid int, username string, db *sql.DB) []byte {

	var votes Votes

	if CheckHasVoted(postid, username, db) {
		log.Println(Warn(username, " has already voted for postid=", postid))
		votes = Votes{}
	}

	_, e := db.Exec("INSERT INTO votes (postid, userid) VALUES ($1, (SELECT id FROM users WHERE username=$2));",
		postid, username)
	if e != nil {
		votes = Votes{}
	}

	if vote == "up" {
		_, e := db.Exec("UPDATE posts SET upvotes=upvotes+1 WHERE id=$1;", postid)
		if e != nil {
			log.Println(Warn("Error incrementing upvote field in database."))
		}
		 votes = countVotes(postid, db)
	}
	if vote == "down" {
		_, e := db.Exec("UPDATE posts SET downvotes=downvotes+1 WHERE id=$1;", postid)
		if e != nil {
			log.Println(Warn("Error incrementing downvote field in database."))
		}
		votes = countVotes(postid, db)
	}

	js, e := json.Marshal(votes)
	if e != nil {
		log.Println(Warn("Unable to parse votes into json object."))
	}
	return js

}

func countVotes(postid int, db *sql.DB) Votes {

	var votes Votes

	r := db.QueryRow("SELECT upvotes, downvotes FROM posts WHERE id=$1;", postid).Scan(&votes.UpVotes, &votes.DownVotes)
	if r != nil {
		log.Println(Warn("Unable to fetch count for postid=", postid))
	}
	votes.PostId = postid

	return votes

}

func CheckHasVoted(postid int, username string, db *sql.DB) bool {

	var result = 0

	_ = db.QueryRow("SELECT count(*) FROM votes WHERE postid=$1 AND userid=(SELECT id FROM users WHERE username=$2);",
		postid, username).Scan(&result)

	if result != 0 {
		return true
	}
	return false

}