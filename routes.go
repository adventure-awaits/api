package main

import (
	"github.com/adventure-awaits/adventure-server/controllers"
	"github.com/julienschmidt/httprouter"
)

// RegisterRoutes initializes the routes for the service
func RegisterRoutes(r *httprouter.Router) {
	UC := controllers.NewUserController()
	AC := controllers.NewAdventureController()

	r.GET("/users/:id", UC.Get)
	r.POST("/users", UC.Create)
	r.DELETE("/users/:id", UC.Remove)

	r.GET("/adventures", AC.GetAll)
	r.GET("/adventures/:id", AC.Get)
	r.POST("/adventures", AC.Create)
	r.DELETE("/adventures/:id", AC.Remove)
}
