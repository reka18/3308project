package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

func postsGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST ENDPOINT (GET)")

	username, ok := CompareTokens(w, r)
	if !ok {
		return
	}

	RefreshCookie(username) /* This updates cookie to restart clock. */

	limit := ParseLimitQuery(r, 5)

	db, _ := Database(DBNAME)
	defer db.Close()

	posts := GetPosts(username, db, limit)
	_, _ = w.Write(posts)


}

func postsPOST(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "POST ENDPOINT (POST)")

	username, ok := CompareTokens(w, r)
	if !ok {
		http.Redirect(w, r, "login", http.StatusSeeOther)
		return
	}

	RefreshCookie(username)

	var postContent = r.FormValue("content")

	db, _ := Database(DBNAME)
	defer db.Close()

	MakePost(username, postContent, db)

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
		id int
		ids []int
		response []Post
		post Post
	)

	r, e := db.Query("SELECT followid FROM follow WHERE follow.userid=(SELECT id FROM users WHERE username=$1) UNION SELECT id FROM users WHERE username=$1;", username)
	if e != nil {
		log.Println(Warn("Unable to retrieve relevant ids."))
		return []byte{}
	} else if r != nil {
		for r.Next() {
			_ = r.Scan(&id)
			ids = append(ids, id)
		}
	} else {
		log.Println(Warn("Unable to parse query response."))
		return []byte{}
	}

	for _, id := range ids {
		r, e := db.Query("SELECT * FROM posts WHERE userid=$1 ORDER BY date LIMIT $2;", id, pagelimit)
		if e != nil {
			return nil
		}
		for r.Next() {
			e = r.Scan(&post.Id, &post.UserId, &post.Content, &post.UpVotes, &post.DownVotes, &post.Deleted, &post.Date)
			if e != nil {
				log.Println(Warn("Error scanning post."))
			}

			timestamp := strings.Split(post.Date.String(), " ")
			date := timestamp[0]
			clock := strings.Split(timestamp[1], ".")[0][:5]
			post.FriendlyDate = fmt.Sprintf("%s @ %s", date, clock)


			e = db.QueryRow("SELECT username FROM users WHERE id=(SELECT userid FROM posts WHERE posts.id=$1);", post.Id).Scan(&post.UserName)
			if e != nil {
				log.Println(Warn("Unable to fetch post owner from database."))
			}

			response = append(response, post)
		}
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Date.Before(response[j].Date)
	})

	js, e := json.Marshal(response)
	if e != nil {
		log.Println(Warn("Error making posts query."))
	}

	return js
}

func MakePost(username string, post string, db *sql.DB) {

	r, e := db.Exec("INSERT INTO posts (userid, content, upvotes, downvotes, deleted, date) VALUES ((SELECT id FROM users WHERE username=$1), $2, 0, 0, false, $3);",
		username, post, time.Now())
	if e != nil {
		log.Println(Warn("Unable to make post in database."))
		log.Println(Warn(e))
	}

	c, _ := r.RowsAffected()
	if c != 0 {
		log.Println(Success("Successfully added post to database."))
	}

}
