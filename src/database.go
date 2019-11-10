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

func createDatabase(db *sql.DB) {

	q := fmt.Sprintf("CREATE DATABASE %v;", DBNAME)
	_, e := db.Query(q)
	if e == nil {
		log.Printf("Successfully created database.")
	} else {
		log.Println("Failed creating database:", e)
	}

}

func dropDatabase(db *sql.DB) error {

	q := fmt.Sprintf("DROP DATABASE %v;", DBNAME)
	_, e := db.Query(q)
	if e == nil {
		log.Printf("Successfully dropped database.")
	} else {
		log.Println("Failed dropping database:", e)
	}
	return e

}

func generateTables(db *sql.DB) {
	/*
	THIS PROCEDURALLY GENERATES ALL THE NECESSARY TABLES
	 */
	createTable(db, genderEnum, "gender enum")
	createTable(db, usersTable, "users table")
	createTable(db, postTable, "posts table")

}

func DatabaseArgHandler() {

	if len(os.Args) > 1 {
		db, _ := Database(PGNAME)

		arg := os.Args[1]

		if arg == "--reset" {
			log.Println("Manually dropping tables.")
			e := dropDatabase(db)
			if e != nil {
				log.Fatalf("Unable to drop database. Aborting.")
			}
			createDatabase(db)
			db, _ = Database(DBNAME)
			generateTables(db)
		}

		if arg == "--create" {
			log.Println("Manually creating database and initializing tables.")
			createDatabase(db)
			db, _ = Database(DBNAME)
			generateTables(db)
		}

		if arg == "--drop" {
			log.Println("Manually dropping database.")
			e := dropDatabase(db)
			if e != nil {
				log.Fatalf("Unable to drop database. Aborting.")
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
	if e == nil {
		log.Println("Database connection established.")
	} else {
		log.Println("Database connection failed:", e)
	}
	return db, e

}

