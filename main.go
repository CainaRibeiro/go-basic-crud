package main

import (
	"basic-mongo-crud/db"
	"basic-mongo-crud/server"
	"fmt"
)

func main() {
	fmt.Println("Initializing server")
	server.StartServer()
	fmt.Println("Server running in port 8081")
	db.Connect()
}
