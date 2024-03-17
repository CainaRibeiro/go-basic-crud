package server

import (
	"basic-mongo-crud/users"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func middlewareHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, X-HTTP-Method-Override, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	r := mux.NewRouter()
	r.Use(middlewareHeaders)
	users.UserRouter(r)

	fmt.Println("Running server in port 4000...")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		fmt.Println(err)
	}
}
