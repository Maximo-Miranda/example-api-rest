package routes

import (
	"github.com/gorilla/mux"

	hdl "github.com/Maximo-Miranda/example-api-rest/internal/user"
)

// userRoutes ...
func userRoutes(r *mux.Router) {

	// Prefix API version
	r = r.PathPrefix("/api/v1").Subrouter()

	// Create user
	r.HandleFunc("/user", hdl.StoreUser).Methods("POST").Name("storeUser")

}
