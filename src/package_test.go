package main

import (
	"testing"
)

func TestDatabase(t *testing.T) {
	db, e := Database(PGNAME)
	defer db.Close()
	if e != nil {
		t.Info("Unable to open connection.")
		t.Fail()
	}
	t.Info("TestDatabase pass!")
}

func TestCreateDatabase(t *testing.T) {
	db, _ := Database(PGNAME)
	e := createDatabase(db)
	defer db.Close()
	if e != nil {
		t.Info("Warning: Problem creating database:", e)
		t.Info("Continuing...")
		if e.Error() != "pq: database \"socialmediasite\" already exists" {
			t.Fail()
		}
	}
	t.Info("TestCreateDatabase pass!")
}

func TestAddNewUserAccount(t *testing.T) {
	db, _ := Database(DBNAME)

	e := AddNewUserAccount(40, "Rodrigo", "Garcia", "rigo.garcia@colorado.edu",
		"M", true, "iamtheverymodelofthemodernmajorgeneral", db)
	defer db.Close()
	if e != nil {
		t.Info("Unable to add User:", e)
		t.Fail()
	}

	e = AddNewUserAccount(36, "Reagan", "Karnes", "reagan.karnes@colorado.edu",
		"M", true, "abcdefghijklmnopqrstuvwxyz1234567890", db)
	if e != nil {
		t.Info("Unable to add User:", e)
		t.Fail()
	}

	t.Info("TestAddNewUserAccount pass!")

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
		t.Info("Login data mismatch.")
		t.Fail()
	}

	t.Info("TestLoginUserAccount pass!")
}