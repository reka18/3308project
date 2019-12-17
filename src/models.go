package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Config struct {
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type HTMLServer struct {
	server *http.Server
	wg     sync.WaitGroup
}

type User struct {
	Id        			int
	Age       			int
	Firstname 			string
	Lastname  			string
	Email     			string
	Username  			string
	Public    			bool
	Joindate  			time.Time
	FriendlyJoinDate	string
	Active    			bool
	Gender   			string
}

type FollowedUser struct {
	User	User
	Mutual	bool
}

type SearchResult struct {
	User	User
	Count	int
}

type Post struct {
	Id				int
	UserId			int
	UserName		string // this is not stored in the post table
	Content			string
	UpVotes			int
	DownVotes		int
	Deleted			bool
	Date			time.Time
	FriendlyDate	string
}

type Message struct {
	Id				int
	FromId			int
	ToId			int
	Content			string
	Date			time.Time
	FriendlyDate	string
}

type Votes struct {
	PostId		int
	UpVotes		int
	DownVotes	int
}

type Avatar struct {
	Id		int
	UserId	int
	Avatar	[]byte
}

// ERRORS
type EmptyStringError struct {}

func (e *EmptyStringError) Error() string {
	return fmt.Sprint("empty string error")
}