package main

import (
	"html/template"
	"os"
	"os/signal"
	"time"
)

var (
	loginHTML string
	loginTMPL *template.Template
)

func init() {
	loginHTML = "web/login.html"
	loginTMPL = template.Must(template.New("login").Parse(loginHTML))
}

func main() {

	// RESETS THE DATABASE TO AN EMPTY STATE
	DatabaseArgHandler()

	serverConfig := Config {
		Host:			"localhost:3000",
		ReadTimeout:	5 * time.Second,
		WriteTimeout:	5 * time.Second,
	}

	server := Start(serverConfig)
	defer server.Stop()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	<-channel

}