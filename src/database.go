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
		log.Println("Database not created.")
	} else {
		log.Println("Database created.")
	}
	return e
}

func dropDatabase(db *sql.DB) error {

	query := fmt.Sprintf("DROP DATABASE %s;", DBNAME)
	_, e := db.Exec(query)
	if e != nil {
		log.Println("Database not dropped.")
	} else {
		log.Println("Database dropped.")
	}
	return e

}

func generateTables(db *sql.DB) {
	/*
	THIS PROCEDURALLY GENERATES ALL THE NECESSARY TABLES
	 */
	e := createTable(db, genderEnum, "'gender' enum")
	if e != nil {
		log.Println("Error", e)
	} else {
		log.Println("Success.")
	}
	e = createTable(db, usersTable, "'users' table")
	if e != nil {
		log.Println("Error", e)
	} else {
		log.Println("Success.")
	}
	e = createTable(db, postTable, "'posts' table")
	if e != nil {
		log.Println("Error", e)
	} else {
		log.Println("Success.")
	}

}

func DatabaseArgHandler() {

	if len(os.Args) > 1 {
		db, e := Database(PGNAME)
		if e != nil {
			log.Fatal(e)
		}

		arg := os.Args[1]

		if arg == "--naked" {
			log.Println("Attempting to build naked database.")
			e = dropDatabase(db)
			if e != nil {
				log.Println(e)
			}
			e = createDatabase(db)
			if e != nil {
				log.Println(e)
			} else {
				log.Println("Naked database created.")
			}
		}

		if arg == "--reset" {
			log.Println("Attempting to reinitializing database...")
			e := dropDatabase(db)
			if e != nil {
				log.Println(e)
			}
			e = createDatabase(db)
			if e != nil {
				log.Println(e)
			}
			db, e = Database(DBNAME)
			if e != nil {
				log.Println("Unable to reset. Aborting.")
				log.Println("Please recreate database.")
				log.Fatal(e)
			}
			generateTables(db)
		}

		if arg == "--init" {
			log.Println("Attempting to create database with tables...")
			e = createDatabase(db)
			if e != nil {
				log.Println(e)
			}
			db, e = Database(DBNAME)
			if e != nil {
				log.Println("Tables not generated. Aborting.")
				log.Fatal(e)
			}
			generateTables(db)
		}

		if arg == "--drop" {
			log.Println("Attempting to drop database...")
			e = dropDatabase(db)
			if e != nil {
				log.Println(e)
			}
		}

		defer db.Close()
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
		log.Println("Database not connected.")
	} else {
		log.Println("Database connected.")
	}
	return db, e

}

