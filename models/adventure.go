package models

import "time"

// Adventure represents the structure of an adventure
type Adventure struct {
	ID            int       `jsonapi:"primary,adventures"`
	Title         string    `jsonapi:"attr,title"`
	Body          string    `jsonapi:"attr,body"`
	Description   string    `jsonapi:"attr,description"`
	Images        []string  `jsonapi:"attr,images"`
	PlannedDate   time.Time `jsonapi:"attr,planned-date"`
	CompletedDate time.Time `jsonapi:"attr,completed-date"`
	Completed     bool      `jsonapi:"attr,completed"`
	Author        *User     `jsonapi:"relation,author"`
}

// Adventures is a collection of individual adventures
type Adventures []*Adventure
