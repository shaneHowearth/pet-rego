package main

import (
	"database/sql"
	"strings"
)

// Pet -
type Pet struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Owner   string `json:"owner"`
	Food    string `json:"food,omitempty"`
}

// Owner -
type Owner struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
}

func (o *Owner) getOwner(db *sql.DB) error {
	return db.QueryRow("SELECT firstname, surname FROM owner WHERE id=$1",
		o.ID).Scan(&o.Firstname, &o.Surname)
}

func (o *Owner) createOwner(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO owner(firstname, surname) VALUES($1, $2) RETURNING id",
		o.Firstname, o.Surname).Scan(&o.ID)

	if err != nil {
		return err
	}

	return nil
}

func getOwners(db *sql.DB) ([]Owner, error) {
	rows, err := db.Query("SELECT id, firstname, surname FROM owner")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	owners := []Owner{}

	for rows.Next() {
		var o Owner
		if err := rows.Scan(&o.ID, &o.Firstname, &o.Surname); err != nil {
			return nil, err
		}
		owners = append(owners, o)
	}

	return owners, nil
}

// Pets

func (p *Pet) getPet(db *sql.DB) error {
	var f sql.NullString
	err := db.QueryRow("SELECT name, species, food, owner FROM pet WHERE id=$1", p.ID).Scan(&p.Name, &p.Species, &f, &p.Owner)
	if f.Valid != false {
		p.Food = f.String
	}
	return err
}

func (p *Pet) createPet(db *sql.DB) error {
	if p.Food == "" {
		switch strings.ToLower(p.Species) {
		case "dog":
			p.Food = "bones"
		case "cat":
			p.Food = "fish"
		case "chicken":
			p.Food = "corn"
		case "snake":
			p.Food = "mice"
		}
	}
	err := db.QueryRow(
		"INSERT INTO pet(name, species, owner, food) VALUES($1, $2, $3, $4) RETURNING id",
		p.Name, p.Species, p.Owner, p.Food).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getPets(db *sql.DB) ([]Pet, error) {
	rows, err := db.Query("SELECT id, name, species, food, owner FROM pet")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pets := []Pet{}

	for rows.Next() {
		var p Pet
		var f sql.NullString
		if err := rows.Scan(&p.ID, &p.Name, &p.Species, &f, &p.Owner); err != nil {
			return nil, err
		}
		if f.Valid != false {
			p.Food = f.String
		}
		pets = append(pets, p)
	}

	return pets, nil
}

func (o *Owner) getOwnersPets(db *sql.DB) ([]Pet, error) {
	rows, err := db.Query("SELECT id, name, species, food, owner FROM pet WHERE owner = $1", o.ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Pets := []Pet{}

	for rows.Next() {
		var p Pet
		var f sql.NullString
		if err := rows.Scan(&p.ID, &p.Name, &p.Species, &f, &p.Owner); err != nil {
			return nil, err
		}
		if f.Valid != false {
			p.Food = f.String
		}
		Pets = append(Pets, p)
	}

	return Pets, nil
}
