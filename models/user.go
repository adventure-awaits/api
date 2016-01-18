package models

import "time"

type (
	// User represents the structure of our resource
	User struct {
		ID       int       `jsonapi:"primary,users"`
		Name     string    `jsonapi:"attr,name"`
		Gender   string    `jsonapi:"attr,gender"`
		Age      int       `jsonapi:"attr,age"`
		JoinedAt time.Time `jsonapi:"attr,joined-at"`
	}
)
