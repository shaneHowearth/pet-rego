package main

import (
	"log"
	"os"
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
(id UUID    NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
 firstname       TEXT    NOT NULL,
 surname         TEXT    NOT NULL
)`

const petTableCreationQuery = `CREATE TABLE IF NOT EXISTS pet
(id UUID    NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
 name    TEXT    NOT NULL,
 species TEXT    NOT NULL,
 owner   UUID    NOT NULL REFERENCES owner(id)
)`
