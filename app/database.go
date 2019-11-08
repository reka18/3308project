package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

const DBNAME = "socialmediasite"
const PGNAME = "postgres"

// THESE ARE THE TABLES
var (
	userAccount = "user_account"
)

/*
PRIVATE METHODS
 */
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

/*
PUBLIC METHODS
 */

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
			_ = createDatabase(db)
			db, _ = Database(DBNAME)
			_ = createEnums(db)
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

func Encrypt(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	log.Printf("%s", hash)
	if err != nil {
		log.Printf("Unable to hash: %s", err)
	}
	return string(hash)
}

func VerifyPW(dbPasswordHash string, password string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	hash := []byte(dbPasswordHash)
	log.Printf("%s", hash)
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
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
	if e != nil {
		log.Println("Unable to execute query:", e)
	} else {
		log.Printf("Successfully added User <%s> to Database.", email)
	}
	return e

}

func LoginUserAccount(inputEmail string, inputPassword string, db *sql.DB) (bool, error) {

	/*
	FAST FAIL IF EMAIL OR PASSWORD ARE BLANK
	 */
	if len(inputEmail) == 0 || len(inputPassword) == 0 {
		log.Println("Email and/or password blank.")
		return false, &EmptyStringError{}
	}

	/*
	DEBUGGING
	 */
	hashpw := Encrypt(inputPassword)
	log.Println(hashpw)

	query := fmt.Sprintf("SELECT * FROM user_account WHERE email='%s';", inputEmail)

	r := db.QueryRow(query)

	var (
		password string
	)

	e := r.Scan(&password)
	if e != nil {
		log.Println("Email not found:")
		return false, e
	}
	if VerifyPW(password, inputPassword) {
		log.Println("Password verified.")
		return true, e
	}
	return false, e

}