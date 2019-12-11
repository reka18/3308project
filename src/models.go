package main

import (
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
	Id        int
	Username  string
	Firstname string
	Lastname  string
	Age       int
	Email     string
	Gender    string
	Public    bool
	Joindate  string
	Active    bool
}

type Post struct {
	Id			int
	UserId		int
	Content		string
	UpVotes		int
	DownVotes	int
	Date		string
}