package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestEmptyOwnerTable(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/owner", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentOwner(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/owner/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Owner not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Owner not found'. Got '%s'", m["error"])
	}
}

func TestCreateOwner(t *testing.T) {
	clearTables()

	payload := []byte(`{"firstname":"test","surname":"owner"}`)

	req, _ := http.NewRequest("POST", "/owner", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["firstname"] != "test" {
		t.Errorf("Expected owner firstname to be 'test'. Got '%v'", m["firstname"])
	}

	if m["surname"] != "owner" {
		t.Errorf("Expected owner surname to be 'owner'. Got '%v'", m["surname"])
	}

	if m["id"] != "1" {
		t.Errorf("Expected owner ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetOwner(t *testing.T) {
	clearTables()
	addOwnersAndPets(1)

	req, _ := http.NewRequest("GET", "/owner/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestMultipleGetOwner(t *testing.T) {
	clearTables()
	addOwnersAndPets(5)

	req, _ := http.NewRequest("GET", "/owner", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetOwnerPets(t *testing.T) {
	clearTables()
	addOwnersAndPets(5)

	req, _ := http.NewRequest("GET", "/owner/4/pets", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
