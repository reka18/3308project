package main

import (
	"html/template"
	"log"
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
	databaseArgHandler()

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

func databaseArgHandler() {
	if len(os.Args) > 1 {
		db, _ := Database(PGNAME)

		if os.Args[1] == "--reset" {
			log.Println("Manually dropping tables.")
			_ = DropTables(db)
		}

		if os.Args[1] == "--create" {
			log.Println("Manually creating database and initializing tables.")
			_ = CreateDatabase(db)
			_ = InitializeDatabase(db)
		}

		defer FailError(db.Close(), "Failed to close database.")
	}
}