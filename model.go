package main

import (
	"database/sql"
	"errors"
)

// Pet -
type Pet struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Species float64 `json:"price"`
	Owner   string  `json:"owner"`
}

// Owner -
type Owner struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
}

func (p *Owner) getOwner(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Owner) createOwner(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getOwners(db *sql.DB, start, count int) ([]Owner, error) {
	return nil, errors.New("Not implemented")
}

func (p *Pet) getPet(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Pet) createPet(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getPets(db *sql.DB, start, count int) ([]Pet, error) {
	return nil, errors.New("Not implemented")
}

func getOwnersPets(db *sql.DB, owner string, start, count int) ([]Pet, error) {
	return nil, errors.New("Not implemented")
}
