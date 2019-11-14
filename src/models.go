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
	Firstname string
	Lastname  string
	Email     string
	Gender    string
	Public    bool
	Joindate  string
	Active    bool
}