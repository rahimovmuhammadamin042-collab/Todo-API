package main

import (
	"log"
	"net/http"
	"todo-api/config"
	"todo-api/router"
)

func main() {

	db := config.ConnectDB()

	r := router.SetupRouter(db)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
