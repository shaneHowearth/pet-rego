package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialise(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	ensureTablesExist()

	code := m.Run()

	clearTables()

	os.Exit(code)
}

func ensureTablesExist() {
	if _, err := a.DB.Exec(ownerTableCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(petTableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTables() {
	a.DB.Exec("DELETE FROM pet")
	a.DB.Exec("ALTER SEQUENCE pet_id_seq RESTART WITH 1")
	a.DB.Exec("DELETE FROM owner")
	a.DB.Exec("ALTER SEQUENCE owner_id_seq RESTART WITH 1")
}

const ownerTableCreationQuery = `CREATE TABLE IF NOT EXISTS owner
(id SERIAL PRIMARY KEY,
 firstname       TEXT    NOT NULL,
 surname         TEXT    NOT NULL
)`

const petTableCreationQuery = `CREATE TABLE IF NOT EXISTS pet
(id SERIAL PRIMARY KEY,
 name    TEXT    NOT NULL,
 species TEXT    NOT NULL,
 owner   INT    NOT NULL REFERENCES owner(id)
)`

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func addOwnersAndPets(count int) {
	if count < 1 {
		count = 1
	}

	// Pets must have an owner
	// We'll create 1 to n-1 owners so that one owner will have multiple pets
	ownerCount := count
	if count != 1 {
		ownerCount--
	}

	for j := 0; j < ownerCount; j++ {
		jCount := strconv.Itoa(j)
		_, err := a.DB.Exec("INSERT INTO owner(firstname, surname) VALUES($1, $2)", "Name "+jCount, "lName "+jCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	for i := 0; i < count; i++ {
		_, err := a.DB.Exec("INSERT INTO pet(name, species, owner) VALUES($1, $2, $3)", "Pet "+strconv.Itoa(i), "Rat "+strconv.Itoa(i), strconv.Itoa(((i+1)%ownerCount)+1))
		if err != nil {
			log.Fatal(err)
		}
	}
}
