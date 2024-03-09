package users

import "github.com/gorilla/mux"

func UserRouter(r *mux.Router) {
	usersRouter := r.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("/", GetUsers).Methods("GET")
}
