package main

import (
	"os"
	"os/signal"
	"time"
)

func main() {

	// RESETS THE DATABASE TO AN EMPTY STATE
	DatabaseArgHandler()

	serverConfig := Config {
		Host:			"localhost:8000",
		ReadTimeout:	5 * time.Second,
		WriteTimeout:	5 * time.Second,
	}

	OpenRedisConnection()
	server := Start(serverConfig)
	defer server.Stop()
	defer CloseRedisConnection()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	<-channel

}