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

func UserBuilder(id int, firstname string, lastname string, email string,
	gender string, public bool, joindate string, active bool) User {
	/*
		Returns a user.
	*/
	return User{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Gender:    gender,
		Public:    public,
		Joindate:  joindate,
		Active:    active,
	}
}