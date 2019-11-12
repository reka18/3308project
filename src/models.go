package main

import (
	"log"
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

func PrintUser(u User) {
	/*
		THIS IS A DEBUGGING TOOL
	*/
	log.Printf("\n\nUSER\t{\n" +
		"\t  Id: %v\n" +
		"\t  First Name: %v\n" +
		"\t  Last Name: %v\n" +
		"\t  Email: %v\n" +
		"\t  Gender: %v\n" +
		"\t  Public: %v\n" +
		"\t  Join Date: %v\n" +
		"\t  Active: %v\n" +
		"\t}\n\n",
		u.Id, u.Firstname, u.Lastname,
		u.Email, u.Gender, u.Public,
		u.Joindate, u.Active)

}