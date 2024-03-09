package users

import (
	"fmt"
	"net/http"
)

func GetUsers(h http.ResponseWriter, r *http.Request) {

	h.WriteHeader(200)
	fmt.Println("GETTING USERS...")
}
