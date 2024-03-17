package users

import (
	"basic-mongo-crud/helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	user Users
	cpf  UserCpf
	un   UserName
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	users, _ := getUsers()

	jsonData, _ := json.Marshal(&users)
	_, err := w.Write(jsonData)

	if err != nil {
		http.Error(w, "error writing response", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u, err := helpers.DecoderJSON(r, &user)
	if err != nil {
		panic(err)
	}
	newUser, _ := u.(*Users)

	err = BodyValidator(newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	CreateUsers(*newUser)

	w.WriteHeader(http.StatusCreated)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	c, err := helpers.DecoderJSON(r, &cpf)

	if err != nil {
		http.Error(w, "error decoding body", http.StatusBadRequest)
	}

	cpfUser := c.(*UserCpf)

	_, e := DeleteUser(*cpfUser)
	if e != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vr := mux.Vars(r)
	id := vr["id"]

	_, err := helpers.DecoderJSON(r, &un)

	if err != nil {
		http.Error(w, "error decoding json", http.StatusInternalServerError)
	}

	uu, err := UpdateUsersName(id, un.Name)
	if err != nil {
		http.Error(w, "error updating user", http.StatusInternalServerError)
	}

	jd, err := json.Marshal(&uu)

	if err != nil {
		http.Error(w, "error when marshaling json", http.StatusInternalServerError)
	}

	_, err = w.Write(jd)

	if err != nil {
		http.Error(w, "error writing response", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
