package main

import (
	"database/sql"
	"errors"
)

// Owner -
type Owner struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
}

func (p *Owner) getOwner(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Owner) updateOwner(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Owner) deleteOwner(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Owner) createOwner(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getOwners(db *sql.DB, start, count int) ([]Owner, error) {
	return nil, errors.New("Not implemented")
}
