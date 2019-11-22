package routes

import (
	"github.com/gorilla/mux"
)

// Router initialize a new routes
func Router() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	// User routes
	userRoutes(r)

	return r
}
