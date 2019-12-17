package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var (
	Success = Green
	Info    = White
	Warn    = Yellow
	Fail    = Red
	Detail  = Teal
)
var (
	Black   = Color("\033[1;90m%s\033[0m")
	Red     = Color("\033[1;91m%s\033[0m")
	Green   = Color("\033[1;92m%s\033[0m")
	Yellow  = Color("\033[1;93m%s\033[0m")
	Purple  = Color("\033[1;94m%s\033[0m")
	Magenta = Color("\033[1;95m%s\033[0m")
	Teal    = Color("\033[1;96m%s\033[0m")
	White   = Color("\033[1;97m%s\033[0m")
)
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

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