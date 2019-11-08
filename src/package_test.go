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
	e := CreateDatabase(db)
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

func TestInitializeDatabase(t *testing.T) {
	db, _ := Database(DBNAME)
	_ = DropTables(db)
	e := InitializeDatabase(db)
	defer db.Close()
	if e != nil {
		t.Log("Database initialization test failed: ", e)
		t.Fail()
	}
	t.Log("TestInitializeDatabase pass!")
}

func TestEncrypt(t *testing.T) {
	v := "267b5b31373739303333373033203331343431333432373720313031333" +
		"930343234322032373733343830373632203133353938393331313920323" +
		"630303832323932342035323837333436333520313534313435393232355" +
		"d205b3439203530203531203532203533203020302030203020302030203" +
		"020302030203020302030203020302030203020302030203020302030203" +
		"020302030203020302030203020302030203020302030203020302030203" +
		"020302030203020302030203020302030203020302030203020302030203" +
		"020302030203020302030203020305d203520352066616c73657d"
	c := Encrypt("12345")
	if c != v {
		t.Fail()
	}
	t.Log("TestEncrypt pass!")
}

func TestAddNewUserAccount(t *testing.T) {
	db, _ := Database(DBNAME)

	e := AddNewUserAccount(40, "Rodrigo", "Garcia", "rigo.garcia@colorado.edu",
		"M", true, "iamtheverymodelofthemodernmajorgeneral", db)
	defer db.Close()
	if e != nil {
		t.Log("Unable to add user:", e)
		t.Fail()
	}

	e = AddNewUserAccount(36, "Reagan", "Karnes", "reagan.karnes@colorado.edu",
		"M", true, "abcdefghijklmnopqrstuvwxyz1234567890", db)
	if e != nil {
		t.Log("Unable to add user:", e)
		t.Fail()
	}

	t.Log("TestAddNewUserAccount pass!")

}

func TestLoginUserAccount(t *testing.T) {
	db, _ := Database(DBNAME)

	u := LoginUserAccount("rigo.garcia@colorado.edu", "iamtheverymodelofthemodernmajorgeneral", db)
	u.joindate = "" // we do this because it is difficult to assert
	defer db.Close()

	v := user {
		id:			1,
		firstname:	"Rodrigo",
		lastname:	"Garcia",
		email:		"rigo.garcia@colorado.edu",
		gender:		"M",
		public:		true,
		joindate:	"",
		active:		true,
	}

	if u != v {
		t.Log("Login data mismatch.")
		t.Fail()
	}

	t.Log("TestLoginUserAccount pass!")
}