package main

import (
	"RESTAPI-MySQL-mics/app"
	"RESTAPI-MySQL-mics/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
