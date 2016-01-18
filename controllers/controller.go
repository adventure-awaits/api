package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Controller is an interface for all controllers
type Controller interface {
	NewController() Controller
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Get(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Remove(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
