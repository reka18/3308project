package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
)

const DBNAME = "socialmediasite"
const PGNAME = "postgres"

// THESE ARE THE TABLES
var (
	userAccount = "user_account"
)

func DatabaseArgHandler() {

	if len(os.Args) > 1 {
		db, _ := Database(PGNAME)

		arg := os.Args[1]

		if arg == "--reset" {
			log.Println("Manually dropping tables.")
			_ = dropTables(db)
			_ = createTables(db)
		}

		if arg == "--create" {
			log.Println("Manually creating database and initializing tables.")
			_ = createDatabase(db)
			db, _ = Database(DBNAME)
			_ = createEnums(db)
			_ = createTables(db)
		}

		if arg == "--drop" {
			log.Println("Manually dropping database.")
			_ = dropDatabase(db)
		}

		defer FailError(db.Close(), "Failed to close database.")
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

func createDatabase(db *sql.DB) error {

	q := fmt.Sprintf("CREATE DATABASE %v;", DBNAME)
	_, e := db.Query(q)
	if e == nil {
		log.Printf("Successfully created database.")
	} else {
		log.Println("Failed creating database:", e)
	}
	return e

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

func createEnums(db *sql.DB) error {

	q := fmt.Sprintf("CREATE TYPE gender AS ENUM ('M', 'F', 'O');")
	_, e := db.Query(q)
	if e == nil {
		log.Println("'gender' enum created successfully.")
	} else {
		log.Println("Unable to create enum:", e)
	}
	return e

}

func createTables(db *sql.DB) error {

	q := fmt.Sprintf("CREATE TABLE user_account (" +
		"id SERIAL PRIMARY KEY," +
		"age INT," +
		"firstName TEXT," +
		"lastName TEXT," +
		"email TEXT UNIQUE NOT NULL," +
		"gender gender NOT NULL," +
		"public BOOLEAN," +
		"joinDate DATE," +
		"active BOOLEAN," +
		"password TEXT" +
		");")

	_, e := db.Query(q)
	if e == nil {
		log.Printf("'user_account table' created successfully.")
	} else {
		log.Println("Unable to create tables:", e)
	}
	return e

}

func dropTables(db *sql.DB) error {

	q := fmt.Sprintf("DROP TABLE IF EXISTS %v;", userAccount)
	_, e := db.Query(q)
	if e == nil {
		log.Println("Successfully dropped tables:", userAccount)
	} else {
		log.Println("Unable to drop table: ", userAccount, e)
	}
	return e

}

func Encrypt(password string) string {

	h := sha256.New()
	_, e := io.WriteString(h, password)
	if e == nil {
		log.Println("Hashing successful.")
	} else {
		log.Fatal("Unknown hashing error:", e)
	}
	return hex.EncodeToString([]byte(fmt.Sprint(h)))

}

func AddNewUserAccount(age int, firstname string, lastname string,
	email string, gender string, public bool, password string, db *sql.DB) error {

	/*
	THIS CONNECTS TO THE DATABASE AND ADDS A USER
	*/
	q := fmt.Sprintf("INSERT INTO user_account(age, firstname, lastname, email, "+
		"gender, public, joindate, active, password)"+
		"VALUES (%d, '%s', '%s', '%s', '%s', '%t', now(), true, '%s');",
		age, firstname, lastname, email, gender, public, Encrypt(password))
	_, e := db.Query(q)
	if e == nil {
		log.Printf("Successfully added User <%s> to Database.", email)
	} else {
		log.Println("Unable to execute query:", e)
	}

	return e

}

func LoginUserAccount(inputEmail string, inputPassword string, db *sql.DB) User {

	query := fmt.Sprintf("SELECT * FROM user_account WHERE email='%s' AND password='%v';",
		inputEmail, Encrypt(inputPassword))
	r := db.QueryRow(query)

	var (
		id        int
		age       int
		firstname string
		lastname  string
		email     string
		gender    string
		public    bool
		joindate  string
		active    bool
		password  string
	)

	e := r.Scan(&id, &age, &firstname, &lastname, &email, &gender, &public, &joindate, &active, &password)
	if e == nil {
		log.Println("Login successful.")
	} else {
		log.Println("Incorrect username or password:", e)
	}

	return UserBuilder(id, firstname, lastname, email, gender, public, joindate, active)

}

func PrintUser(u User) {

	/*
	THIS IS A DEBUGGING TOOL
	*/
	log.Printf("\n\n\tUSER {\n" +
		"\tId: %v\n" +
		"\tFirst Name: %v\n" +
		"\tLast Name: %v\n" +
		"\tEmail: %v\n" +
		"\tGender: %v\n" +
		"\tPublic: %v\n" +
		"\tJoin Date: %v\n" +
		"\tActive: %v\n" +
		"\t}\n\n",
		u.Id, u.Firstname, u.Lastname,
		u.Email, u.Gender, u.Public,
		u.Joindate, u.Active)

}