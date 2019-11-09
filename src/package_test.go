package main

import (
	"testing"
)

func TestDatabase(t *testing.T) {
	db, e := Database(PGNAME)
	defer db.Close()
	if e != nil {
		t.Log("Unable to open connection.")
		t.Fail()
	}
	t.Log("TestDatabase pass!")
}

func TestCreateDatabase(t *testing.T) {
	db, _ := Database(PGNAME)
	e := createDatabase(db)
	defer db.Close()
	if e != nil {
		t.Log("Warning: Problem creating database:", e)
		t.Log("Continuing...")
		if e.Error() != "pq: database \"socialmediasite\" already exists" {
			t.Fail()
		}
	}
	t.Log("TestCreateDatabase pass!")
}

func TestAddNewUserAccount(t *testing.T) {
	db, _ := Database(DBNAME)

	e := AddNewUserAccount(40, "Rodrigo", "Garcia", "rigo.garcia@colorado.edu",
		"M", true, "iamtheverymodelofthemodernmajorgeneral", db)
	defer db.Close()
	if e != nil {
		t.Log("Unable to add User:", e)
		t.Fail()
	}

	e = AddNewUserAccount(36, "Reagan", "Karnes", "reagan.karnes@colorado.edu",
		"M", true, "abcdefghijklmnopqrstuvwxyz1234567890", db)
	if e != nil {
		t.Log("Unable to add User:", e)
		t.Fail()
	}

	t.Log("TestAddNewUserAccount pass!")

}

func TestLoginUserAccount(t *testing.T) {

	db, _ := Database(DBNAME)

	u, _, _ := LoginUserAccount("rigo.garcia@colorado.edu", "iamtheverymodelofthemodernmajorgeneral", db)



	defer db.Close()

	v := User{
		Id:			1,
		Firstname:	"Rodrigo",
		Lastname:	"Garcia",
		Email:		"rigo.garcia@colorado.edu",
		Gender:		"M",
		Public:		true,
		Joindate:	"",
		Active:		true,
	}

	if u != v {
		t.Log("Login data mismatch.")
		t.Fail()
	}

	t.Log("TestLoginUserAccount pass!")
}