package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const DBNAME = "socialmediasite"
const PGNAME = "postgres"

func createDatabase(db *sql.DB) error {

	query := fmt.Sprintf("CREATE DATABASE %s;", DBNAME)
	_, e := db.Exec(query)
	if e != nil {
		log.Println(Warn("Database not created."))
	} else {
		log.Println(Success("Database created."))
	}
	return e
}

func dropDatabase(db *sql.DB) error {

	query := fmt.Sprintf("DROP DATABASE %s;", DBNAME)
	_, e := db.Exec(query)
	if e != nil {
		log.Println(Warn("Database not dropped."))
	} else {
		log.Println(Success("Database dropped."))
	}
	return e

}

func generateTables(db *sql.DB) {
	/*
	THIS PROCEDURALLY GENERATES ALL THE NECESSARY TABLES
	 */
	e := createTable(db, genderEnum, "'gender' enum")
	if e != nil {
		log.Println(Fatal("Error"), e)
	} else {
		log.Println(Success("Success."))
	}
	e = createTable(db, usersTable, "'users' table")
	if e != nil {
		log.Println(Fatal("Error"), e)
	} else {
		log.Println(Success("Success."))
	}
	e = createTable(db, postTable, "'posts' table")
	if e != nil {
		log.Println(Fatal("Error"), e)
	} else {
		log.Println(Success("Success."))
	}
	e = createTable(db, avatarTable, "'avatar' table")
	if e != nil {
		log.Println(Fatal("Error"), e)
	} else {
		log.Println(Success("Success."))
	}

}

func DatabaseArgHandler() {

	if len(os.Args) > 1 {
		db, e := Database(PGNAME)
		if e != nil {
			log.Fatal(Fatal(e))
		}
		defer db.Close()

		arg := os.Args[1]

		if arg == "--naked" {
			log.Println(Info("Attempting to build naked database."))
			e = dropDatabase(db)
			if e != nil {
				log.Println(Warn(e))
			}
			e = createDatabase(db)
			if e != nil {
				log.Println(Fatal(e))
			} else {
				log.Println(Success("Naked database created."))
			}
		} else if arg == "--reset" {
			log.Println(Info("Attempting to reinitializing database..."))
			e := dropDatabase(db)
			if e != nil {
				log.Println(Warn(e))
			}
			e = createDatabase(db)
			if e != nil {
				log.Println(Warn(e))
			}
			db, e = Database(DBNAME)
			if e != nil {
				log.Println(Fatal("Unable to reset. Aborting."))
				log.Println(Fatal("Please recreate database."))
				log.Fatal(Fatal(e))
			}
			generateTables(db)
		} else if arg == "--init" {
			log.Println(Info("Attempting to create database with tables..."))
			e = createDatabase(db)
			if e != nil {
				log.Println(Warn(e))
			}
			db, e = Database(DBNAME)
			if e != nil {
				log.Println(Fatal("Tables not generated. Aborting."))
				log.Fatal(Fatal(e))
			}
			generateTables(db)
		} else if arg == "--drop" {
			log.Println(Info("Attempting to drop database..."))
			e = dropDatabase(db)
			if e != nil {
				log.Fatal(Fatal(e))
			}
		} else {
			log.Fatal(Fatal("Unknown database command... exiting."))
		}
		os.Exit(0)
	}

}

func Database(dbname string) (*sql.DB, error) {
	/*
	THIS OPENS THE DATABASE CONNECTION. NOTE THAT
	THE DATABASE IS BASICALLY IN WAIT MODE, THE
	CONNECTION ONLY ACTUALLY OPENS WHEN A QUERY IS
	MADE.
	*/
	dbInfo := fmt.Sprintf("dbname='%v' sslmode=disable", dbname)
	db, e := sql.Open("postgres", dbInfo)
	if e != nil {
		log.Println(Warn("Database not connected."))
	} else {
		log.Println(Detail("Database connected."))
	}
	return db, e

}

