package users

import "go.mongodb.org/mongo-driver/mongo"

type Users struct {
	Name               string `json:"name"`
	Age                int    `json:"age"`
	Address            string `json:"address"`
	Cpf                string `json:"cpf"`
	ActiveTransactions int    `json:"activeTransactions"`
}

type UserCpf struct {
	Cpf string `json:"cpf"`
}

type UserName struct {
	Name string `json:"name"`
}

type UpdateStruct struct {
	filter mongo.UpdateOneModel
	update mongo.UpdateOneModel
}
