package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	RegisterRoutes(r)

	http.ListenAndServe(":8080", r)
}
