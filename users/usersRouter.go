package users

import "github.com/gorilla/mux"

func UserRouter(r *mux.Router) {
	usersRouter := r.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("", GetUsers).Methods("GET")
	usersRouter.HandleFunc("/create", CreateUser).Methods("POST")
	usersRouter.HandleFunc("/delete", RemoveUser).Methods("DELETE")
	usersRouter.HandleFunc("/{id}", UpdateUser).Methods("PUT")

}
