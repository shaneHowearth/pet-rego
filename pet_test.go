package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestEmptyPetTable(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/pet", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentPet(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/pet/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Pet not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Pet not found'. Got '%s'", m["error"])
	}
}

func TestCreatePet(t *testing.T) {
	clearTables()
	addOwnersAndPets(1)

	payload := []byte(`{"name":"test pet","species":"dog","owner":"1"}`)

	req, _ := http.NewRequest("POST", "/pet", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test pet" {
		t.Errorf("Expected pet name to be 'test pet'. Got '%v'", m["name"])
	}

	if m["species"] != "dog" {
		t.Errorf("Expected pet species to be 'dog'. Got '%v'", m["species"])
	}

	if m["owner"] != "1" {
		t.Errorf("Expected pet owner to be '1'. Got '%v'", m["owner"])
	}

	// the setup creates a pet, so this is the second one to be created
	if m["id"] != "2" {
		t.Errorf("Expected owner ID to be '2'. Got '%v'", m["id"])
	}
}

func TestGetPet(t *testing.T) {
	clearTables()
	addOwnersAndPets(1)

	req, _ := http.NewRequest("GET", "/pet/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetMulitplePets(t *testing.T) {
	clearTables()
	addOwnersAndPets(5)

	req, _ := http.NewRequest("GET", "/pet", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
