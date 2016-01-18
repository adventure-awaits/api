package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adventure-awaits/adventure-server/models"
	"github.com/julienschmidt/httprouter"
)

// UserController represents the controller for operating on the User resource
type UserController struct{}

// NewUserController constructs a new UserController
func NewUserController() *UserController {
	return &UserController{}
}

// Get retrieves an individual user resource
func (uc UserController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an example user
	pID, _ := strconv.Atoi(p.ByName("id"))
	u := models.User{
		ID:     pID,
		Name:   "Jennifer Petersen",
		Gender: "female",
		Age:    31,
	}

	// Marshal provided interface into JSON structure
	// uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/vnd.api+json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

// Create creates a new user resource
func (uc UserController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.ID = 12

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/vnd.api+json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

// Remove removes an existing user resource
func (uc UserController) Remove(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write status for now
	w.WriteHeader(http.StatusOK)
}
