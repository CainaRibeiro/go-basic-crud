package main

import (
	"basic-mongo-crud/users"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	users.UserRouter(r)

	err := http.ListenAndServe("localhost:5000", r)
	if err != nil {
		panic(err)
	}
}
