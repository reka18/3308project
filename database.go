package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
)

const DBNAME = "socialmediasite"
const PGNAME = "postgres"

// THESE ARE THE TABLES
var (
	userAccount = "user_account"
)

type user struct {
	id        int
	firstname string
	lastname  string
	email     string
	gender    string
	public    bool
	joindate  string
	active    bool
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
		log.Println("Database connection failed.")
	}
	return db, e
}

func CreateDatabase(db *sql.DB) error {
	q := fmt.Sprintf("CREATE DATABASE %v;", DBNAME)
	_, e := db.Query(q)
	if e == nil {
		log.Printf("Successfully created database.")
	} else {
		log.Println("Failed creating database.")
	}
	return e
}

func InitializeDatabase(db *sql.DB) error {
	q := fmt.Sprintf("CREATE TYPE gender AS ENUM ('M', 'F', 'O');")
	_, e := db.Query(q)
	if e != nil {
		log.Println("Unable to create enum:", e)
	} else {
		log.Println("'gender' enum created successfully.")
	}

	q = fmt.Sprintf("CREATE TABLE user_account (" +
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
	_, e = db.Query(q)
	if e != nil {
		log.Println("Unable to create tables.")
		return e
	}
	log.Printf("'user_account table' created successfully.")
	return e
}

func DropTables(db *sql.DB) error {
	q := fmt.Sprintf("DROP TABLE IF EXISTS %v;", userAccount)
	_, e := db.Query(q)
	if e != nil {
		log.Println("Unable to drop table: ", userAccount)
		return e
	}
	log.Println("Successfully dropped tables:", userAccount)
	return nil
}

func Encrypt(password string) string {
	h := sha256.New()
	_, err := io.WriteString(h, password)
	if err != nil {
		log.Fatal(err, "Unknown hashing error.")
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
	if e != nil {
		log.Println("Unable to execute query:", e)
		return e
	}
	log.Printf("Successfully added user <%s> to Database.", email)
	return nil
}

func LoginUserAccount(inputEmail string, inputPassword string, db *sql.DB) user {
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
	if e != nil {
		log.Fatal("Incorrect username or password.")
	}

	return user{
		id:        id,
		firstname: firstname,
		lastname:  lastname,
		email:     email,
		gender:    gender,
		public:    public,
		joindate:  joindate,
		active:    active,
	}
}

func PrintUser(u user) {
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
		u.id, u.firstname, u.lastname,
		u.email, u.gender, u.public,
		u.joindate, u.active)
}