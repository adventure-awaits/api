package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adventure-awaits/adventure-server/models"
	"github.com/julienschmidt/httprouter"
	"github.com/shwoodard/jsonapi"
)

// AdventureController represents the controller for operating on the Adventure resource
type AdventureController struct{}

// NewAdventureController constructs a new AdventureController
func NewAdventureController() *AdventureController {
	return &AdventureController{}
}

// Get retrieves an individual adventure resource
func (ac AdventureController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("adventures.get")

	// Stub an example adventure
	aID, _ := strconv.Atoi(p.ByName("id"))
	a := models.Adventure{}

	for _, adventure := range models.AdventureTests {
		if aID == adventure.ID {
			a = *adventure
			break
		}
	}

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)

	if err := jsonapiRuntime.MarshalOnePayload(w, &a); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetAll retrieves all adventure resources
func (ac AdventureController) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("adventures.get-all")

	as := models.AdventureTests
	asInterface := make([]interface{}, len(as))

	for i, a := range as {
		asInterface[i] = a
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)

	if err := jsonapiRuntime.MarshalManyPayload(w, asInterface); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Create creates a new user resource
func (ac AdventureController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Stub an user to be populated from the body
	a := models.Adventure{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&a)

	// Add an Id
	a.ID = 12

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(a); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Remove removes an existing user resource
func (ac AdventureController) Remove(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write status for now
	w.WriteHeader(http.StatusOK)
}
