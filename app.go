// app.go

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App -
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialise -
func (a *App) Initialise(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initialiseRoutes()
}

// Run -
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) initialiseRoutes() {
	a.Router.HandleFunc("/pet", a.getPets).Methods("GET")
	a.Router.HandleFunc("/pet", a.createPet).Methods("POST")
	a.Router.HandleFunc("/pet/{id:[0-9]+}", a.getPet).Methods("GET")
	a.Router.HandleFunc("/owner", a.getOwners).Methods("GET")
	a.Router.HandleFunc("/owner", a.createOwner).Methods("POST")
	a.Router.HandleFunc("/owner/{id:[0-9]+}", a.getOwner).Methods("GET")
	a.Router.HandleFunc("/owner/{id:[0-9]+}/pets", a.getOwnersPets).Methods("GET")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getPet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid pet ID")
		return
	}

	p := Pet{ID: vars["id"]}
	if err := p.getPet(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Pet not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) getPets(w http.ResponseWriter, r *http.Request) {

	pets, err := getPets(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, pets)
}

func (a *App) createPet(w http.ResponseWriter, r *http.Request) {
	var p Pet
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.createPet(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) getOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid pet ID")
		return
	}

	o := Owner{ID: vars["id"]}
	if err := o.getOwner(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Owner not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, o)
}

func (a *App) getOwners(w http.ResponseWriter, r *http.Request) {

	owners, err := getOwners(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, owners)
}

func (a *App) createOwner(w http.ResponseWriter, r *http.Request) {
	var o Owner
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&o); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := o.createOwner(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, o)
}

func (a *App) getOwnersPets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid owner ID")
		return
	}

	o := Owner{ID: vars["id"]}
	owners, err := o.getOwnersPets(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, owners)
}
